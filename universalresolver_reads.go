// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/wealdtech/go-ens/v4/contracts/resolver"
)

// resolverABI is the ABI of the legacy PublicResolver, used to pack the
// calldata for individual resolver profiles (text, contenthash, multicoin
// addr) that UniversalResolver.resolve forwards to the discovered resolver.
// Parsed once at init; the profile signatures are stable across resolver
// versions.
var resolverABI abi.ABI

func init() {
	parsed, err := resolver.ContractMetaData.GetAbi()
	if err != nil {
		// The binding ABI is a compile-time constant string; a parse failure
		// here means the generated binding is broken, which is a build-level
		// bug rather than a runtime condition.
		panic(fmt.Sprintf("ens: parse resolver abi: %v", err))
	}
	resolverABI = *parsed
}

// Text resolves a text record for a name through the UniversalResolver. Unlike
// Resolver.Text, this follows ENSIP-10 wildcard resolution and ERC-3668
// CCIP-Read, so offchain and L2 text records resolve correctly.
//
// This is the read half of the unified read path: it builds calldata for the
// resolver's text(bytes32,string) profile and hands it to the same resolve()
// choke point as Address and Contenthash, so wildcard/CCIP behaviour is
// identical across every profile.
func (u *UniversalResolver) Text(ctx context.Context, name, key string) (string, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return "", err
	}
	callData, err := resolverABI.Pack("text", nameHash, key)
	if err != nil {
		return "", fmt.Errorf("pack text: %w", err)
	}
	result, _, err := u.Resolve(ctx, name, callData)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", nil
	}
	values, err := resolverABI.Unpack("text", result)
	if err != nil {
		return "", fmt.Errorf("unpack text: %w", err)
	}
	if len(values) != 1 {
		return "", fmt.Errorf("universal resolver: unexpected text output arity")
	}
	out, ok := values[0].(string)
	if !ok {
		return "", fmt.Errorf("universal resolver: text result has unexpected type %T", values[0])
	}
	return out, nil
}

// Contenthash resolves the contenthash record for a name through the
// UniversalResolver, following ENSIP-10 wildcard resolution and ERC-3668
// CCIP-Read.
func (u *UniversalResolver) Contenthash(ctx context.Context, name string) ([]byte, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	callData, err := resolverABI.Pack("contenthash", nameHash)
	if err != nil {
		return nil, fmt.Errorf("pack contenthash: %w", err)
	}
	result, _, err := u.Resolve(ctx, name, callData)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	values, err := resolverABI.Unpack("contenthash", result)
	if err != nil {
		return nil, fmt.Errorf("unpack contenthash: %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("universal resolver: unexpected contenthash output arity")
	}
	out, ok := values[0].([]byte)
	if !ok {
		return nil, fmt.Errorf("universal resolver: contenthash result has unexpected type %T", values[0])
	}
	return out, nil
}

// MultiAddress resolves the address of a name for an arbitrary SLIP-0044 coin
// type through the UniversalResolver, following ENSIP-10 wildcard resolution
// and ERC-3668 CCIP-Read. It invokes the resolver's addr(bytes32,uint256)
// profile. For coin type 60 (ETH) prefer ResolveAddress, which returns a typed
// common.Address.
func (u *UniversalResolver) MultiAddress(ctx context.Context, name string, coinType uint64) ([]byte, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return nil, err
	}
	callData, err := resolverABI.Pack("addr0", nameHash, new(big.Int).SetUint64(coinType))
	if err != nil {
		return nil, fmt.Errorf("pack addr(bytes32,uint256): %w", err)
	}
	result, _, err := u.Resolve(ctx, name, callData)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, ErrNoAddress
	}
	values, err := resolverABI.Unpack("addr0", result)
	if err != nil {
		return nil, fmt.Errorf("unpack addr(bytes32,uint256): %w", err)
	}
	if len(values) != 1 {
		return nil, fmt.Errorf("universal resolver: unexpected addr output arity")
	}
	out, ok := values[0].([]byte)
	if !ok {
		return nil, fmt.Errorf("universal resolver: addr result has unexpected type %T", values[0])
	}
	return out, nil
}
