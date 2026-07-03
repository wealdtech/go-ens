// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ErrUnregistered is returned when a name has no resolver in the registry.
// Callers may compare against it with errors.Is.
var ErrUnregistered = errors.New("unregistered name")

// ErrNoAddress is returned when a resolver is found but the address record
// is missing or zero.
var ErrNoAddress = errors.New("no address")

// ResolverNotFoundError corresponds to the UniversalResolver custom error
// `ResolverNotFound(bytes)`. It is reported when no resolver is set anywhere
// up the parent chain of the requested name.
type ResolverNotFoundError struct {
	// DNSName is the DNS-encoded name as passed to resolve().
	DNSName []byte
}

func (e *ResolverNotFoundError) Error() string  { return "unregistered name" }
func (e *ResolverNotFoundError) Unwrap() error  { return ErrUnregistered }

// ResolverNotContractError corresponds to `ResolverNotContract(bytes,address)`:
// the registry pointed at an address that holds no code.
type ResolverNotContractError struct {
	DNSName  []byte
	Resolver common.Address
}

func (e *ResolverNotContractError) Error() string {
	return fmt.Sprintf("resolver address has no code: %s", e.Resolver.Hex())
}

// UnsupportedResolverProfileError corresponds to
// `UnsupportedResolverProfile(bytes4)`: the resolver did not implement the
// requested function selector.
type UnsupportedResolverProfileError struct {
	Selector [4]byte
}

func (e *UnsupportedResolverProfileError) Error() string {
	return fmt.Sprintf("resolver does not implement profile 0x%s", hex.EncodeToString(e.Selector[:]))
}

// ResolverRevertError corresponds to `ResolverError(bytes)`: the downstream
// resolver itself reverted; the inner revert bytes are preserved verbatim.
type ResolverRevertError struct {
	Data []byte
}

func (e *ResolverRevertError) Error() string {
	return fmt.Sprintf("resolver reverted: 0x%s", hex.EncodeToString(e.Data))
}

// EmptyAddressError corresponds to `EmptyAddress()`: thrown by reverse() when
// the lookup address is empty.
type EmptyAddressError struct{}

func (e *EmptyAddressError) Error() string { return "no address" }
func (e *EmptyAddressError) Unwrap() error { return ErrNoAddress }

// HTTPGatewayError corresponds to `HttpError(uint16,string)`: a CCIP-Read
// gateway returned a non-success status whose body the UR forwarded.
type HTTPGatewayError struct {
	Status  uint16
	Message string
}

func (e *HTTPGatewayError) Error() string {
	return fmt.Sprintf("ccip-read gateway HTTP %d: %s", e.Status, e.Message)
}

// ReverseAddressMismatchError corresponds to `ReverseAddressMismatch(string,bytes)`:
// the reverse-resolved name's forward record does not point back at the
// queried address.
type ReverseAddressMismatchError struct {
	Name    string
	Address []byte
}

func (e *ReverseAddressMismatchError) Error() string {
	return fmt.Sprintf("reverse name %q does not resolve back to %s", e.Name, "0x"+hex.EncodeToString(e.Address))
}

// Custom-error selectors. Computed at init from the canonical signatures.
var (
	selResolverNotFound          [4]byte
	selResolverNotContract       [4]byte
	selUnsupportedResolverProfile [4]byte
	selResolverError             [4]byte
	selEmptyAddress              [4]byte
	selHTTPError                 [4]byte
	selReverseAddressMismatch    [4]byte
)

// Argument lists for decoding error payloads.
var (
	argsResolverNotFound          abi.Arguments
	argsResolverNotContract       abi.Arguments
	argsUnsupportedResolverProfile abi.Arguments
	argsResolverError             abi.Arguments
	argsHTTPError                 abi.Arguments
	argsReverseAddressMismatch    abi.Arguments
)

func init() {
	bytesT, _ := abi.NewType("bytes", "", nil)
	bytes4T, _ := abi.NewType("bytes4", "", nil)
	addressT, _ := abi.NewType("address", "", nil)
	uint16T, _ := abi.NewType("uint16", "", nil)
	stringT, _ := abi.NewType("string", "", nil)

	argsResolverNotFound = abi.Arguments{{Type: bytesT}}
	argsResolverNotContract = abi.Arguments{{Type: bytesT}, {Type: addressT}}
	argsUnsupportedResolverProfile = abi.Arguments{{Type: bytes4T}}
	argsResolverError = abi.Arguments{{Type: bytesT}}
	argsHTTPError = abi.Arguments{{Type: uint16T}, {Type: stringT}}
	argsReverseAddressMismatch = abi.Arguments{{Type: stringT}, {Type: bytesT}}

	selResolverNotFound = errorSelector("ResolverNotFound(bytes)")
	selResolverNotContract = errorSelector("ResolverNotContract(bytes,address)")
	selUnsupportedResolverProfile = errorSelector("UnsupportedResolverProfile(bytes4)")
	selResolverError = errorSelector("ResolverError(bytes)")
	selEmptyAddress = errorSelector("EmptyAddress()")
	selHTTPError = errorSelector("HttpError(uint16,string)")
	selReverseAddressMismatch = errorSelector("ReverseAddressMismatch(string,bytes)")
}

func errorSelector(sig string) [4]byte {
	h := crypto.Keccak256([]byte(sig))
	var s [4]byte
	copy(s[:], h[:4])
	return s
}

// translateURRevert tries to decode a CallContract error as one of the
// UniversalResolver's typed custom errors. If recognised, it returns a typed
// Go error with a human-friendly message; otherwise it returns err unchanged.
func translateURRevert(err error) error {
	if err == nil {
		return nil
	}
	revert, ok := extractRevertData(err)
	if !ok || len(revert) < 4 {
		return err
	}
	var sel [4]byte
	copy(sel[:], revert[:4])
	body := revert[4:]

	switch sel {
	case selResolverNotFound:
		out := &ResolverNotFoundError{}
		if vals, derr := argsResolverNotFound.Unpack(body); derr == nil && len(vals) == 1 {
			if name, ok := vals[0].([]byte); ok {
				out.DNSName = name
			}
		}
		return out
	case selResolverNotContract:
		out := &ResolverNotContractError{}
		if vals, derr := argsResolverNotContract.Unpack(body); derr == nil && len(vals) == 2 {
			if name, ok := vals[0].([]byte); ok {
				out.DNSName = name
			}
			if a, ok := vals[1].(common.Address); ok {
				out.Resolver = a
			}
		}
		return out
	case selUnsupportedResolverProfile:
		out := &UnsupportedResolverProfileError{}
		if vals, derr := argsUnsupportedResolverProfile.Unpack(body); derr == nil && len(vals) == 1 {
			if s, ok := vals[0].([4]byte); ok {
				out.Selector = s
			}
		}
		return out
	case selResolverError:
		out := &ResolverRevertError{}
		if vals, derr := argsResolverError.Unpack(body); derr == nil && len(vals) == 1 {
			if data, ok := vals[0].([]byte); ok {
				out.Data = data
			}
		}
		return out
	case selEmptyAddress:
		return &EmptyAddressError{}
	case selHTTPError:
		out := &HTTPGatewayError{}
		if vals, derr := argsHTTPError.Unpack(body); derr == nil && len(vals) == 2 {
			if s, ok := vals[0].(uint16); ok {
				out.Status = s
			}
			if m, ok := vals[1].(string); ok {
				out.Message = m
			}
		}
		return out
	case selReverseAddressMismatch:
		out := &ReverseAddressMismatchError{}
		if vals, derr := argsReverseAddressMismatch.Unpack(body); derr == nil && len(vals) == 2 {
			if n, ok := vals[0].(string); ok {
				out.Name = n
			}
			if a, ok := vals[1].([]byte); ok {
				out.Address = a
			}
		}
		return out
	}
	return err
}

// extractRevertData pulls the raw revert payload from an RPC error returned
// by ContractBackend.CallContract. Returns ok=false if the error is not an
// RPC error or carries no data field.
func extractRevertData(err error) ([]byte, bool) {
	type dataError interface{ ErrorData() interface{} }
	var de dataError
	if !errors.As(err, &de) {
		return nil, false
	}
	return normalizeErrorData(de.ErrorData())
}

// normalizeErrorData accepts the various concrete types that go-ethereum's
// rpc.DataError implementers return for ErrorData(): a 0x-prefixed string
// (the common case for net/http JSON-RPC), or raw bytes (some wrapped or
// mock backends). hexutil.Bytes is reported as []byte via its underlying
// type, so the []byte branch catches it too.
func normalizeErrorData(v interface{}) ([]byte, bool) {
	switch d := v.(type) {
	case string:
		raw, err := hex.DecodeString(strings.TrimPrefix(d, "0x"))
		if err != nil {
			return nil, false
		}
		return raw, true
	case []byte:
		// Already-decoded revert bytes — accept as-is.
		return d, true
	default:
		return nil, false
	}
}
