// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

// Package ccipread implements the client side of ERC-3668 (CCIP Read)
// so that calls into ENS contracts which revert with OffchainLookup are
// transparently followed: the gateway is queried, the callback is invoked,
// and the chain is repeated until a non-revert result is returned.
package ccipread

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	neturl "net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// MaxRedirects bounds the number of OffchainLookup hops a single call may chain.
// Four matches the limit used by ENS-aware libraries (viem, ethers).
const MaxRedirects = 4

// MaxGatewayBodyBytes caps the gateway response body that will be read into
// memory. CCIP-Read responses are small JSON envelopes (a 0x-prefixed hex
// payload), so anything beyond a few hundred kilobytes is either pathological
// or hostile. Cap conservatively at 16 MiB to keep a misbehaving gateway from
// exhausting the caller's heap.
const MaxGatewayBodyBytes = 16 * 1024 * 1024

// GatewayHTTPError is returned when a CCIP-Read gateway responds with a
// non-2xx HTTP status. Callers can use errors.As to recover the URL, status,
// and (truncated) body so they can decide whether to retry or surface the
// failure.
//
// This is distinct from the on-chain HTTPGatewayError typed-error in the
// parent ens package, which decodes the UR's HttpError(uint16,string) revert.
type GatewayHTTPError struct {
	URL    string
	Status int
	Body   string
}

func (e *GatewayHTTPError) Error() string {
	return fmt.Sprintf("ccipread: gateway %s: HTTP %d: %s", e.URL, e.Status, truncate(e.Body, 256))
}

// noRedirect is the CheckRedirect policy installed when the caller's HTTP
// client doesn't supply one of its own. ERC-3668 gateways are expected to
// return data directly, so a 3xx is suspicious; reflecting the response back
// instead of following protects against SSRF via Location: redirects to
// internal hosts (cloud metadata services, RFC1918 ranges, etc.).
func noRedirect(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}

// offchainLookupSelector is the first 4 bytes of
// keccak256("OffchainLookup(address,string[],bytes,bytes4,bytes)").
var offchainLookupSelector = []byte{0x55, 0x6f, 0x18, 0x30}

// offchainLookupArgs is the ABI for the OffchainLookup error parameters,
// used to decode the revert payload that follows the 4-byte selector.
var offchainLookupArgs abi.Arguments

// callbackArgs is the ABI for the callback signature
// callbackFunction(bytes response, bytes extraData).
var callbackArgs abi.Arguments

func init() {
	addrT, _ := abi.NewType("address", "", nil)
	stringArrT, _ := abi.NewType("string[]", "", nil)
	bytesT, _ := abi.NewType("bytes", "", nil)
	bytes4T, _ := abi.NewType("bytes4", "", nil)
	offchainLookupArgs = abi.Arguments{
		{Name: "sender", Type: addrT},
		{Name: "urls", Type: stringArrT},
		{Name: "callData", Type: bytesT},
		{Name: "callbackFunction", Type: bytes4T},
		{Name: "extraData", Type: bytesT},
	}
	callbackArgs = abi.Arguments{
		{Name: "response", Type: bytesT},
		{Name: "extraData", Type: bytesT},
	}
}

// OffchainLookup is a decoded OffchainLookup revert.
type OffchainLookup struct {
	Sender           common.Address
	URLs             []string
	CallData         []byte
	CallbackFunction [4]byte
	ExtraData        []byte
}

// Caller is the subset of bind.ContractBackend needed for CCIP-Read calls.
type Caller interface {
	bind.ContractCaller
}

// Options tune the CCIP-Read client. A zero value is valid and uses
// http.DefaultClient with the package default for MaxRedirects.
type Options struct {
	HTTPClient   *http.Client
	MaxRedirects int
	// BlockNumber pins the eth_call to a specific block. Nil means latest.
	BlockNumber *big.Int
}

// Call executes a contract call and follows any OffchainLookup reverts.
// On success it returns the raw bytes returned by the (possibly callback-)
// invoked function. The decoding of those bytes is the caller's job — for
// UniversalResolver.resolve they ABI-decode as (bytes, address); for
// UniversalResolver.reverse they decode as (string, address, address).
func Call(ctx context.Context, backend Caller, to common.Address, data []byte, opts *Options) ([]byte, error) {
	if opts == nil {
		opts = &Options{}
	}
	max := opts.MaxRedirects
	if max <= 0 {
		max = MaxRedirects
	}
	httpClient := opts.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if httpClient.CheckRedirect == nil {
		// Defensive copy so we don't mutate the caller's shared client.
		cp := *httpClient
		cp.CheckRedirect = noRedirect
		httpClient = &cp
	}

	target := to
	callData := data
	for hop := 0; hop < max; hop++ {
		out, callErr := backend.CallContract(ctx, ethereum.CallMsg{To: &target, Data: callData}, opts.BlockNumber)
		if callErr == nil {
			return out, nil
		}
		lookup, ok := decodeOffchainLookup(callErr)
		if !ok {
			return nil, callErr
		}
		// Per ERC-3668 the sender field MUST equal the contract that
		// reverted; otherwise the response could be misdirected.
		if lookup.Sender != target {
			return nil, fmt.Errorf("ccipread: OffchainLookup sender %s does not match call target %s", lookup.Sender.Hex(), target.Hex())
		}
		response, err := queryGateways(ctx, httpClient, lookup.URLs, lookup.Sender, lookup.CallData)
		if err != nil {
			return nil, err
		}
		nextData, err := encodeCallback(lookup.CallbackFunction, response, lookup.ExtraData)
		if err != nil {
			return nil, err
		}
		callData = nextData
		// target stays the same: the callback lives on `sender`, which we just
		// validated equals the previous target.
	}
	return nil, fmt.Errorf("ccipread: exceeded %d OffchainLookup redirects", max)
}

// DecodeOffchainLookup parses an OffchainLookup revert from raw revert bytes
// (selector + ABI-encoded args). It returns false if the data does not start
// with the OffchainLookup selector.
func DecodeOffchainLookup(revert []byte) (*OffchainLookup, bool) {
	if len(revert) < 4 || !bytes.Equal(revert[:4], offchainLookupSelector) {
		return nil, false
	}
	values, err := offchainLookupArgs.Unpack(revert[4:])
	if err != nil || len(values) != 5 {
		return nil, false
	}
	sender, ok := values[0].(common.Address)
	if !ok {
		return nil, false
	}
	urls, ok := values[1].([]string)
	if !ok {
		return nil, false
	}
	callData, ok := values[2].([]byte)
	if !ok {
		return nil, false
	}
	callbackFunction, ok := values[3].([4]byte)
	if !ok {
		return nil, false
	}
	extraData, ok := values[4].([]byte)
	if !ok {
		return nil, false
	}
	return &OffchainLookup{
		Sender:           sender,
		URLs:             urls,
		CallData:         callData,
		CallbackFunction: callbackFunction,
		ExtraData:        extraData,
	}, true
}

// decodeOffchainLookup attempts to extract an OffchainLookup payload from a
// CallContract error. Backends expose the revert data via the rpc.DataError
// interface; geth and ethclient both satisfy it.
func decodeOffchainLookup(err error) (*OffchainLookup, bool) {
	if err == nil {
		return nil, false
	}
	type dataError interface{ ErrorData() interface{} }
	var de dataError
	if !errors.As(err, &de) {
		return nil, false
	}
	revert, ok := normalizeErrorData(de.ErrorData())
	if !ok {
		return nil, false
	}
	return DecodeOffchainLookup(revert)
}

// normalizeErrorData accepts the various concrete types go-ethereum's
// rpc.DataError implementers return for ErrorData(): a 0x-prefixed hex
// string (the JSON-RPC case) or raw bytes (mocked / wrapped backends).
func normalizeErrorData(v interface{}) ([]byte, bool) {
	switch d := v.(type) {
	case string:
		raw, err := hex.DecodeString(strings.TrimPrefix(d, "0x"))
		if err != nil {
			return nil, false
		}
		return raw, true
	case []byte:
		return d, true
	default:
		return nil, false
	}
}

func encodeCallback(selector [4]byte, response, extraData []byte) ([]byte, error) {
	body, err := callbackArgs.Pack(response, extraData)
	if err != nil {
		return nil, fmt.Errorf("ccipread: encode callback args: %w", err)
	}
	out := make([]byte, 0, 4+len(body))
	out = append(out, selector[:]...)
	out = append(out, body...)
	return out, nil
}

// gatewayRequest is the JSON body the spec defines for POST queries.
type gatewayRequest struct {
	Data   string `json:"data"`
	Sender string `json:"sender"`
}

// gatewayResponse is the JSON envelope returned by gateways.
type gatewayResponse struct {
	Data string `json:"data"`
}

// queryGateways tries each URL in order and returns the first successful
// response. Per ERC-3668 a 4xx aborts the whole call; a 5xx falls through
// to the next URL.
func queryGateways(ctx context.Context, client *http.Client, urls []string, sender common.Address, data []byte) ([]byte, error) {
	if len(urls) == 0 {
		return nil, errors.New("ccipread: OffchainLookup returned no gateway URLs")
	}
	hexData := "0x" + hex.EncodeToString(data)
	hexSender := strings.ToLower(sender.Hex())
	var lastErr error
	for _, raw := range urls {
		resp, status, err := queryGateway(ctx, client, raw, hexSender, hexData)
		if err == nil {
			return resp, nil
		}
		// 4xx: per spec, abort immediately — request is malformed and other
		// gateways will reject it the same way.
		if status >= 400 && status < 500 {
			return nil, err
		}
		lastErr = err
	}
	if lastErr == nil {
		lastErr = errors.New("ccipread: no gateway URL succeeded")
	}
	return nil, lastErr
}

func queryGateway(ctx context.Context, client *http.Client, url, sender, data string) ([]byte, int, error) {
	// Reject obviously hostile gateway URLs before we even build a request.
	// Gateway URLs come from a smart contract (untrusted input), so a malicious
	// resolver could return file://, gopher://, etc. By default Go's net/http
	// transport rejects unknown schemes, but a caller-supplied custom transport
	// might support them; this is defense-in-depth.
	parsedURL, perr := neturl.Parse(url)
	if perr != nil {
		return nil, 0, fmt.Errorf("ccipread: gateway %q: invalid URL: %w", url, perr)
	}
	switch parsedURL.Scheme {
	case "http", "https":
		// ok
	default:
		return nil, 0, fmt.Errorf("ccipread: gateway %q: unsupported URL scheme %q", url, parsedURL.Scheme)
	}
	hasData := strings.Contains(url, "{data}")
	hasSender := strings.Contains(url, "{sender}")
	var req *http.Request
	var err error
	if hasData {
		// GET form: substitute placeholders directly into the URL.
		populated := strings.ReplaceAll(url, "{data}", data)
		if hasSender {
			populated = strings.ReplaceAll(populated, "{sender}", sender)
		}
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, populated, nil)
	} else {
		body, jerr := json.Marshal(gatewayRequest{Data: data, Sender: sender})
		if jerr != nil {
			return nil, 0, jerr
		}
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
		if err == nil {
			req.Header.Set("Content-Type", "application/json")
		}
	}
	if err != nil {
		return nil, 0, err
	}
	if client.Timeout == 0 {
		// Best-effort timeout when the caller didn't supply one. Without this
		// a wedged gateway would block resolution indefinitely. Use the smaller
		// of the caller's remaining deadline (if any) and 30s so a tight
		// per-request deadline isn't extended.
		timeout := 30 * time.Second
		if deadline, ok := ctx.Deadline(); ok {
			if remaining := time.Until(deadline); remaining < timeout {
				timeout = remaining
			}
		}
		ctxT, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()
		req = req.WithContext(ctxT)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("ccipread: gateway %s: %w", url, err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, MaxGatewayBodyBytes))
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, resp.StatusCode, &GatewayHTTPError{
			URL:    url,
			Status: resp.StatusCode,
			Body:   string(bodyBytes),
		}
	}
	// Gate the JSON decode on Content-Type so a misbehaving gateway returning
	// HTML or text/plain with a 200 status surfaces as a clean error instead of
	// a misleading "invalid JSON".
	if ct := resp.Header.Get("Content-Type"); !contentTypeIsJSON(ct) {
		return nil, resp.StatusCode, fmt.Errorf("ccipread: gateway %s: unexpected Content-Type %q", url, ct)
	}
	var parsed gatewayResponse
	if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
		return nil, resp.StatusCode, fmt.Errorf("ccipread: gateway %s: invalid JSON: %w", url, err)
	}
	out, err := hex.DecodeString(strings.TrimPrefix(parsed.Data, "0x"))
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("ccipread: gateway %s: invalid hex in response.data: %w", url, err)
	}
	return out, resp.StatusCode, nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// contentTypeIsJSON reports whether a Content-Type header indicates JSON,
// allowing for media-type parameters like `application/json; charset=utf-8`.
func contentTypeIsJSON(ct string) bool {
	if ct == "" {
		return false
	}
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = ct[:i]
	}
	ct = strings.TrimSpace(strings.ToLower(ct))
	return ct == "application/json"
}
