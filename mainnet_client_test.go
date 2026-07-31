// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

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
