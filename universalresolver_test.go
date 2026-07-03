// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens_test

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	ens "github.com/wealdtech/go-ens/v4"
)

// chainIDBackend embeds a nil ContractBackend and overrides ChainID, so the
// constructor's type-assertion fires but no method outside ChainID can be
// called (which is fine because the guardrail returns before any other call).
type chainIDBackend struct {
	bind.ContractBackend
	chainID *big.Int
}

func (b *chainIDBackend) ChainID(_ context.Context) (*big.Int, error) {
	return b.chainID, nil
}

func TestNewUniversalResolver_AcceptsMainnet(t *testing.T) {
	ur, err := ens.NewUniversalResolver(context.Background(), &chainIDBackend{chainID: big.NewInt(1)})
	require.NoError(t, err)
	require.NotNil(t, ur)
}

func TestNewUniversalResolver_AcceptsSepolia(t *testing.T) {
	ur, err := ens.NewUniversalResolver(context.Background(), &chainIDBackend{chainID: big.NewInt(11155111)})
	require.NoError(t, err)
	require.NotNil(t, ur)
}

func TestNewUniversalResolver_AcceptsHolesky(t *testing.T) {
	ur, err := ens.NewUniversalResolver(context.Background(), &chainIDBackend{chainID: big.NewInt(17000)})
	require.NoError(t, err)
	require.NotNil(t, ur)
}

func TestNewUniversalResolver_RejectsUnknownChain(t *testing.T) {
	_, err := ens.NewUniversalResolver(context.Background(), &chainIDBackend{chainID: big.NewInt(8453)}) // Base
	require.Error(t, err)
	var typed *ens.UnknownChainError
	require.ErrorAs(t, err, &typed)
	require.Equal(t, int64(8453), typed.ChainID.Int64())
	require.Contains(t, err.Error(), "WithAddress")
}

// NewUniversalResolverAt skips the chain-ID check by design — useful for
// devnets / forks / alt-L1s where the UR is deployed at a non-canonical
// address.
func TestNewUniversalResolverAt_AcceptsAnyChain(t *testing.T) {
	custom := common.HexToAddress("0x000000000000000000000000000000000000bEEF")
	ur, err := ens.NewUniversalResolver(context.Background(), &chainIDBackend{chainID: big.NewInt(8453)}, ens.WithAddress(custom))
	require.NoError(t, err)
	require.Equal(t, custom, ur.Address())
}

// slowChainIDBackend blocks on ctx until released, then returns its chain ID.
// Used to assert that NewUniversalResolverContext propagates ctx cancellation
// into the chain-ID probe instead of relying on context.Background().
type slowChainIDBackend struct {
	bind.ContractBackend
	chainID *big.Int
	release chan struct{}
}

func (b *slowChainIDBackend) ChainID(ctx context.Context) (*big.Int, error) {
	select {
	case <-b.release:
		return b.chainID, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// TestNewUniversalResolverContext_PropagatesCancellation asserts that a
// cancelled ctx surfaces as ctx.Err() from the constructor's chain-ID probe,
// rather than the probe running to completion on an uncancellable background
// context.
func TestNewUniversalResolverContext_PropagatesCancellation(t *testing.T) {
	backend := &slowChainIDBackend{
		chainID: big.NewInt(1),
		release: make(chan struct{}),
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := ens.NewUniversalResolver(ctx, backend)
	require.ErrorIs(t, err, context.Canceled,
		"a cancelled ctx must surface as context.Canceled from the chain-ID probe")
}

// TestNewUniversalResolverContext_HonoursDeadline mirrors the cancellation
// test but with a context deadline that expires during the probe.
func TestNewUniversalResolverContext_HonoursDeadline(t *testing.T) {
	backend := &slowChainIDBackend{
		chainID: big.NewInt(1),
		release: make(chan struct{}),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	_, err := ens.NewUniversalResolver(ctx, backend)
	require.ErrorIs(t, err, context.DeadlineExceeded,
		"an expired deadline must surface as context.DeadlineExceeded from the chain-ID probe")
}

// TestNewUniversalResolverContext_AcceptsKnownChain ensures the context-aware
// constructor still returns a usable resolver on the happy path.
func TestNewUniversalResolverContext_AcceptsKnownChain(t *testing.T) {
	ur, err := ens.NewUniversalResolver(context.Background(), &chainIDBackend{chainID: big.NewInt(1)})
	require.NoError(t, err)
	require.NotNil(t, ur)
}

// mainnetClient connects to a public Ethereum mainnet RPC. The endpoint can
// be overridden with GO_ENS_TEST_RPC; without an override the test falls
// back to a public endpoint so that `go test` works out of the box.
//
// Live-RPC tests skip themselves under `go test -short`, so CI runs that
// can't reach the network (or that are rate-limited) keep working.
func mainnetClient(t *testing.T) *ethclient.Client {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping live-RPC test under -short")
	}
	url := os.Getenv("GO_ENS_TEST_RPC")
	if url == "" {
		// Public mainnet RPC. Surfaces revert data via the JSON-RPC `data`
		// field, which is required for CCIP-Read.
		url = "https://ethereum.publicnode.com"
	}
	client, err := ethclient.Dial(url)
	require.NoError(t, err, "failed to dial %s", url)
	return client
}

// TestUniversalResolverIntegrationName is the canonical ENSv2 readiness
// check from https://docs.ens.domains/web/ensv2-readiness/. A library that
// resolves through the UniversalResolver returns the v2 sentinel address;
// a library that still uses legacy registry-walking returns the v1 sentinel.
func TestUniversalResolverIntegrationName(t *testing.T) {
	const (
		name           = "ur.integration-tests.eth"
		expectedV2     = "0x2222222222222222222222222222222222222222"
		legacyV1Marker = "0x1111111111111111111111111111111111111111"
	)
	client := mainnetClient(t)

	addr, err := ens.Resolve(context.Background(), client, name)
	require.NoError(t, err, "Resolve(%s) failed", name)

	got := strings.ToLower(addr.Hex())
	require.NotEqual(t, strings.ToLower(legacyV1Marker), got,
		"got the legacy v1 sentinel %s — resolution did not route through the UniversalResolver", legacyV1Marker)
	require.Equal(t, strings.ToLower(expectedV2), got,
		"expected ENSv2 sentinel %s, got %s", expectedV2, got)
}

// TestUniversalResolverCCIPRead exercises the OffchainLookup path: this
// name resolves only after the client follows an ERC-3668 revert and queries
// a CCIP-Read gateway.
func TestUniversalResolverCCIPRead(t *testing.T) {
	const (
		name     = "test.offchaindemo.eth"
		expected = "0x779981590E7Ccc0CFAe8040Ce7151324747cDb97"
	)
	client := mainnetClient(t)

	addr, err := ens.Resolve(context.Background(), client, name)
	require.NoError(t, err, "Resolve(%s) failed", name)
	require.Equal(t, strings.ToLower(expected), strings.ToLower(addr.Hex()))
}

// TestUniversalResolverDirectResolve checks the lower-level helper used by
// Resolve, asserting that calling UniversalResolver.ResolveAddress directly
// produces the same result as the top-level Resolve entrypoint.
func TestUniversalResolverDirectResolve(t *testing.T) {
	client := mainnetClient(t)
	ur, err := ens.NewUniversalResolver(context.Background(), client)
	require.NoError(t, err)

	addr, err := ur.ResolveAddress(context.Background(), "ur.integration-tests.eth")
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0x2222222222222222222222222222222222222222"), addr)
}
