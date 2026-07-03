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
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v4/ccipread"
	"github.com/wealdtech/go-ens/v4/contracts/universalresolver"
)

// UniversalResolverContractAddress is the proxy address of the ENS
// UniversalResolver. The same vanity address is used on Ethereum mainnet
// and L1 testnets where the ENS DAO has deployed the proxy (Sepolia, Holesky).
const UniversalResolverContractAddress = "0xeEeEEEeE14D718C2B47D9923Deab1335E144EeEe"

// CoinTypeETH is SLIP-0044 coin type 60: native Ethereum addresses.
const CoinTypeETH = uint64(60)

// addrSelector is keccak256("addr(bytes32)")[:4]: the legacy ETH-address
// resolver profile. UniversalResolver.resolve forwards calldata starting with
// this selector to the discovered resolver and returns the abi-encoded result.
var addrSelector = [4]byte{0x3b, 0x3b, 0x57, 0xde}

// knownURChains lists the chain IDs on which the canonical UR proxy address
// is deployed by the ENS DAO. Callers on other chains must supply an explicit
// address via WithAddress.
var knownURChains = map[uint64]struct{}{
	1:        {}, // Ethereum mainnet
	11155111: {}, // Sepolia testnet
	17000:    {}, // Holesky testnet
}

// UnknownChainError is returned by NewUniversalResolver when the backend
// reports a chain ID without a known UR deployment at the canonical address.
type UnknownChainError struct {
	ChainID *big.Int
}

func (e *UnknownChainError) Error() string {
	return fmt.Sprintf("universal resolver: no canonical deployment on chain ID %s; supply an explicit address with WithAddress", e.ChainID)
}

// Parameter is a functional option for NewUniversalResolver.
type Parameter func(*parameters)

type parameters struct {
	address    *common.Address
	httpClient *http.Client
	maxHops    int
}

// WithAddress binds the resolver to an explicit UniversalResolver address
// instead of the canonical proxy. Use this on chains without a canonical
// deployment (devnets, forks, alt-L1s); it also skips the chain-ID guardrail.
func WithAddress(address common.Address) Parameter {
	return func(p *parameters) { p.address = &address }
}

// WithHTTPClient sets the HTTP client used for ERC-3668 CCIP-Read gateway
// requests. When unset, a default client is used.
func WithHTTPClient(client *http.Client) Parameter {
	return func(p *parameters) { p.httpClient = client }
}

// WithMaxHops caps the number of OffchainLookup redirects a single resolution
// may chain. When unset, the ccipread package default applies.
func WithMaxHops(hops int) Parameter {
	return func(p *parameters) { p.maxHops = hops }
}

// UniversalResolver wraps the ENS Universal Resolver and exposes resolution
// helpers that transparently follow ERC-3668 OffchainLookup reverts via the
// ccipread package.
type UniversalResolver struct {
	backend    bind.ContractBackend
	addr       common.Address
	abi        abi.ABI
	httpClient *http.Client
	maxHops    int
}

// NewUniversalResolver returns a UniversalResolver for the chain reached
// through backend. It honours ctx for the chain-ID probe, so a cancelled or
// expired ctx surfaces as ctx.Err().
//
// By default it binds to the canonical proxy address and, when the backend
// exposes a chain ID (ethclient.Client and most production backends do),
// verifies the chain has a known UR deployment — returning *UnknownChainError
// otherwise. Supply WithAddress to target a non-standard deployment (devnet,
// fork, alt-L1), which also skips the chain-ID guardrail.
func NewUniversalResolver(ctx context.Context, backend bind.ContractBackend, params ...Parameter) (*UniversalResolver, error) {
	var p parameters
	for _, param := range params {
		param(&p)
	}

	addr := common.HexToAddress(UniversalResolverContractAddress)
	if p.address != nil {
		addr = *p.address
	} else if cid, ok := backend.(interface {
		ChainID(context.Context) (*big.Int, error)
	}); ok {
		// Only guard the canonical address; an explicit WithAddress opts out.
		chainID, err := cid.ChainID(ctx)
		if err == nil && chainID != nil {
			if _, known := knownURChains[chainID.Uint64()]; !known {
				return nil, &UnknownChainError{ChainID: chainID}
			}
		} else if err != nil && ctx.Err() != nil {
			// Surface ctx cancellation/deadline from the probe itself.
			return nil, ctx.Err()
		}
	}

	parsed, err := universalresolver.ContractMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("parse universal resolver abi: %w", err)
	}
	return &UniversalResolver{
		backend:    backend,
		addr:       addr,
		abi:        *parsed,
		httpClient: p.httpClient,
		maxHops:    p.maxHops,
	}, nil
}

// callOptions builds the ccipread options from the resolver's configuration,
// returning nil when nothing has been customised so ccipread uses its defaults.
func (u *UniversalResolver) callOptions() *ccipread.Options {
	if u.httpClient == nil && u.maxHops == 0 {
		return nil
	}
	return &ccipread.Options{HTTPClient: u.httpClient, MaxRedirects: u.maxHops}
}

// Address returns the on-chain address of the bound UniversalResolver.
func (u *UniversalResolver) Address() common.Address {
	return u.addr
}

// Resolve performs UniversalResolver.resolve(name, callData) using the chain
// connected via the backend, transparently following any OffchainLookup
// reverts. The returned bytes are the abi-encoded return value of the
// downstream resolver call (e.g. for `addr(bytes32)` the bytes decode as
// `address`); resolverAddr is the resolver contract that produced it.
func (u *UniversalResolver) Resolve(ctx context.Context, name string, callData []byte) ([]byte, common.Address, error) {
	dnsName, err := DNSEncode(name)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("dns-encode %q: %w", name, err)
	}
	input, err := u.abi.Pack("resolve", dnsName, callData)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("pack resolve: %w", err)
	}
	out, err := ccipread.Call(ctx, u.backend, u.addr, input, u.callOptions())
	if err != nil {
		return nil, common.Address{}, translateURRevert(err)
	}
	values, err := u.abi.Unpack("resolve", out)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("unpack resolve: %w", err)
	}
	if len(values) != 2 {
		return nil, common.Address{}, errors.New("universal resolver: unexpected resolve output arity")
	}
	result, ok := values[0].([]byte)
	if !ok {
		return nil, common.Address{}, fmt.Errorf("universal resolver: resolve result has unexpected type %T", values[0])
	}
	resolver, ok := values[1].(common.Address)
	if !ok {
		return nil, common.Address{}, fmt.Errorf("universal resolver: resolve resolver has unexpected type %T", values[1])
	}
	return result, resolver, nil
}

// ResolveAddress resolves an ENS name to its primary Ethereum address by
// invoking the legacy `addr(bytes32)` resolver profile through the
// UniversalResolver.
func (u *UniversalResolver) ResolveAddress(ctx context.Context, name string) (common.Address, error) {
	nameHash, err := NameHash(name)
	if err != nil {
		return UnknownAddress, err
	}
	callData := make([]byte, 4+32)
	copy(callData[:4], addrSelector[:])
	copy(callData[4:], nameHash[:])
	result, _, err := u.Resolve(ctx, name, callData)
	if err != nil {
		return UnknownAddress, err
	}
	if len(result) == 0 {
		return UnknownAddress, ErrNoAddress
	}
	if len(result) != 32 {
		return UnknownAddress, fmt.Errorf("addr result has unexpected length %d", len(result))
	}
	addr := common.BytesToAddress(result)
	if addr == UnknownAddress {
		return UnknownAddress, ErrNoAddress
	}
	return addr, nil
}

// Reverse performs UniversalResolver.reverse(lookupAddress, coinType) and
// returns the primary name, the resolver address, and the reverse resolver
// address. CCIP-Read is followed transparently.
func (u *UniversalResolver) Reverse(ctx context.Context, address []byte, coinType uint64) (string, common.Address, common.Address, error) {
	input, err := u.abi.Pack("reverse", address, new(big.Int).SetUint64(coinType))
	if err != nil {
		return "", common.Address{}, common.Address{}, fmt.Errorf("pack reverse: %w", err)
	}
	out, err := ccipread.Call(ctx, u.backend, u.addr, input, u.callOptions())
	if err != nil {
		return "", common.Address{}, common.Address{}, translateURRevert(err)
	}
	values, err := u.abi.Unpack("reverse", out)
	if err != nil {
		return "", common.Address{}, common.Address{}, fmt.Errorf("unpack reverse: %w", err)
	}
	if len(values) != 3 {
		return "", common.Address{}, common.Address{}, errors.New("universal resolver: unexpected reverse output arity")
	}
	name, ok := values[0].(string)
	if !ok {
		return "", common.Address{}, common.Address{}, fmt.Errorf("universal resolver: reverse name has unexpected type %T", values[0])
	}
	resolver, ok := values[1].(common.Address)
	if !ok {
		return "", common.Address{}, common.Address{}, fmt.Errorf("universal resolver: reverse resolver has unexpected type %T", values[1])
	}
	reverseResolver, ok := values[2].(common.Address)
	if !ok {
		return "", common.Address{}, common.Address{}, fmt.Errorf("universal resolver: reverse reverseResolver has unexpected type %T", values[2])
	}
	return name, resolver, reverseResolver, nil
}
