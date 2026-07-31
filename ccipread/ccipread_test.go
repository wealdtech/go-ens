// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ccipread_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wealdtech/go-ens/v4/ccipread"
)

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

// mockRevertErr satisfies the rpc.DataError interface used by
// decodeOffchainLookup: ErrorData() returns the raw revert bytes as a
// 0x-prefixed hex string.
type mockRevertErr struct {
	data string
}

func (e *mockRevertErr) Error() string             { return "execution reverted" }
func (e *mockRevertErr) ErrorData() interface{}    { return e.data }
func (e *mockRevertErr) Unwrap() error             { return nil }

// loopBackend is a bind.ContractCaller that always returns the same revert
// error. It records how many CallContract calls it received.
type loopBackend struct {
	revertHex string
	calls     atomic.Int32
}

func (b *loopBackend) CallContract(ctx context.Context, _ ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	b.calls.Add(1)
	return nil, &mockRevertErr{data: b.revertHex}
}

func (b *loopBackend) CodeAt(_ context.Context, _ common.Address, _ *big.Int) ([]byte, error) {
	return nil, nil
}

// buildOffchainLookupRevert hand-encodes a valid OffchainLookup revert payload.
// Mirrors the ABI used by the contract: OffchainLookup(address,string[],bytes,bytes4,bytes).
func buildOffchainLookupRevert(t *testing.T, sender common.Address, urls []string) string {
	t.Helper()
	addrT, _ := abi.NewType("address", "", nil)
	stringArrT, _ := abi.NewType("string[]", "", nil)
	bytesT, _ := abi.NewType("bytes", "", nil)
	bytes4T, _ := abi.NewType("bytes4", "", nil)
	args := abi.Arguments{
		{Type: addrT}, {Type: stringArrT}, {Type: bytesT}, {Type: bytes4T}, {Type: bytesT},
	}
	callback := [4]byte{0xaa, 0xbb, 0xcc, 0xdd}
	body, err := args.Pack(sender, urls, []byte{0x01, 0x02}, callback, []byte{0x03, 0x04})
	require.NoError(t, err)
	revert := append([]byte{0x55, 0x6f, 0x18, 0x30}, body...) // OffchainLookup selector
	return "0x" + hex.EncodeToString(revert)
}

// ---------------------------------------------------------------------------
// A1 — hop-counter bound: MaxRedirects=4 ⇒ 4 gateway calls, not 5.
// ---------------------------------------------------------------------------

func TestCall_HopCounterStopsAtMaxRedirects(t *testing.T) {
	var gatewayHits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		gatewayHits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":"0x"}`))
	}))
	defer srv.Close()

	target := common.HexToAddress("0x1111111111111111111111111111111111111111")
	revert := buildOffchainLookupRevert(t, target, []string{srv.URL})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "exceeded")
	// Before the fix the loop allowed 5 gateway queries with MaxRedirects=4.
	assert.Equal(t, int32(ccipread.MaxRedirects), gatewayHits.Load(),
		"gateway should be hit exactly MaxRedirects times before the loop errors out")
	assert.Equal(t, int32(ccipread.MaxRedirects), backend.calls.Load(),
		"backend should be called exactly MaxRedirects times")
}

// ---------------------------------------------------------------------------
// #1 — io.LimitReader caps the response body.
// ---------------------------------------------------------------------------

func TestCall_LimitReaderTruncatesOversizedBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Pad an otherwise-valid envelope with junk far beyond MaxGatewayBodyBytes.
		// LimitReader should cap reading, leaving the JSON unparseable, which we
		// surface as a clean error rather than an OOM.
		body := make([]byte, ccipread.MaxGatewayBodyBytes+(1<<20))
		for i := range body {
			body[i] = 'x'
		}
		_, _ = w.Write(body)
	}))
	defer srv.Close()

	target := common.HexToAddress("0x2222222222222222222222222222222222222222")
	revert := buildOffchainLookupRevert(t, target, []string{srv.URL})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "exceeded", "should fail on JSON, not on hop limit")
}

// ---------------------------------------------------------------------------
// #3 — JSON unmarshal is gated on Content-Type.
// ---------------------------------------------------------------------------

func TestCall_RejectsNonJSONContentType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html>0x1234</html>`))
	}))
	defer srv.Close()

	target := common.HexToAddress("0x3333333333333333333333333333333333333333")
	revert := buildOffchainLookupRevert(t, target, []string{srv.URL})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Content-Type")
}

func TestCall_AcceptsJSONWithCharsetParameter(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"data":"0xdeadbeef"}`))
	}))
	defer srv.Close()

	target := common.HexToAddress("0x4444444444444444444444444444444444444444")
	revert := buildOffchainLookupRevert(t, target, []string{srv.URL})
	// On the SECOND CallContract (after following the lookup), succeed so the
	// loop terminates cleanly and we can assert what Call returned.
	backend := &succeedAfterOneRevert{revertHex: revert, success: []byte{0x99}}

	out, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.NoError(t, err)
	assert.Equal(t, []byte{0x99}, out)
}

func TestCall_RejectsMissingContentType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// no Content-Type header
		_, _ = w.Write([]byte(`{"data":"0x"}`))
	}))
	defer srv.Close()

	target := common.HexToAddress("0x5555555555555555555555555555555555555555")
	revert := buildOffchainLookupRevert(t, target, []string{srv.URL})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Content-Type")
}

// ---------------------------------------------------------------------------
// A4 — gateway URL scheme allowlist.
// ---------------------------------------------------------------------------

func TestCall_RejectsNonHTTPScheme(t *testing.T) {
	target := common.HexToAddress("0x6666666666666666666666666666666666666666")
	revert := buildOffchainLookupRevert(t, target, []string{"file:///etc/passwd"})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported URL scheme")
}

func TestCall_RejectsGopherScheme(t *testing.T) {
	target := common.HexToAddress("0x7777777777777777777777777777777777777777")
	revert := buildOffchainLookupRevert(t, target, []string{"gopher://internal.svc/data"})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported URL scheme")
}

// ---------------------------------------------------------------------------
// A5 — HTTP redirects are not followed by default.
// ---------------------------------------------------------------------------

func TestCall_DoesNotFollowRedirects(t *testing.T) {
	// The "redirect target" server would record a hit if the client followed
	// the redirect — we assert it doesn't.
	var followed atomic.Int32
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		followed.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":"0x"}`))
	}))
	defer target.Close()

	// The "gateway" server returns a 302 pointing at the target.
	redirector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Location", target.URL)
		w.WriteHeader(http.StatusFound)
	}))
	defer redirector.Close()

	tgt := common.HexToAddress("0x8888888888888888888888888888888888888888")
	revert := buildOffchainLookupRevert(t, tgt, []string{redirector.URL})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, tgt, []byte{0x00}, nil)
	require.Error(t, err, "gateway returning a 3xx without data must error, not silently follow")
	assert.Equal(t, int32(0), followed.Load(), "redirect target must not be reached")
}

// ---------------------------------------------------------------------------
// A6 — typed GatewayHTTPError for non-2xx responses.
// ---------------------------------------------------------------------------

func TestCall_4xxReturnsTypedError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"bad request"}`))
	}))
	defer srv.Close()

	target := common.HexToAddress("0x9999999999999999999999999999999999999999")
	revert := buildOffchainLookupRevert(t, target, []string{srv.URL})
	backend := &loopBackend{revertHex: revert}

	_, err := ccipread.Call(context.Background(), backend, target, []byte{0x00}, nil)
	require.Error(t, err)

	var typed *ccipread.GatewayHTTPError
	require.ErrorAs(t, err, &typed, "expected GatewayHTTPError, got %T: %v", err, err)
	assert.Equal(t, 400, typed.Status)
	assert.Equal(t, srv.URL, typed.URL)
	assert.Contains(t, typed.Body, "bad request")
}

// ---------------------------------------------------------------------------
// Helper backend that reverts only on the first call, then succeeds.
// ---------------------------------------------------------------------------

type succeedAfterOneRevert struct {
	revertHex string
	success   []byte
	calls     atomic.Int32
}

func (b *succeedAfterOneRevert) CallContract(_ context.Context, _ ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	n := b.calls.Add(1)
	if n == 1 {
		return nil, &mockRevertErr{data: b.revertHex}
	}
	return b.success, nil
}

func (b *succeedAfterOneRevert) CodeAt(_ context.Context, _ common.Address, _ *big.Int) ([]byte, error) {
	return nil, nil
}

