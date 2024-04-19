// Copyright 2017 Weald Technology Trading
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ens

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var client, _ = ethclient.Dial("https://mainnet.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

func TestResolveEmpty(t *testing.T) {
	_, err := Resolve(client, "")
	assert.NotNil(t, err, "Resolved empty name")
}

func TestResolveZero(t *testing.T) {
	_, err := Resolve(client, "0")
	assert.NotNil(t, err, "Resolved empty name")
}

func TestResolveNotPresent(t *testing.T) {
	_, err := Resolve(client, "sirnotappearinginthisregistry.eth")
	require.NotNil(t, err, "Resolved name that does not exist")
	assert.Equal(t, "unregistered name", err.Error(), "Unexpected error")
}

// func TestResolveNoResolver(t *testing.T) {
// 	_, err := Resolve(client, "noresolver.eth")
// 	require.NotNil(t, err, "Resolved name without a resolver")
// 	assert.Equal(t, "no resolver", err.Error(), "Unexpected error")
// }

func TestResolveBadResolver(t *testing.T) {
	_, err := Resolve(client, "resolvestozero.eth")
	require.NotNil(t, err, "Resolved name with a bad resolver")
	assert.Equal(t, "no address", err.Error(), "Unexpected error")
}

func TestResolveTestEnsTest(t *testing.T) {
	expected := "ed96dd3be847b387217ef9de5b20d8392a6cdf40"
	actual, err := Resolve(client, "test.enstest.eth")
	require.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveResolverEth(t *testing.T) {
	expected := "231b0ee14048e9dccd1d247744d114a4eb5e8e63"
	actual, err := Resolve(client, "resolver.eth")
	require.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveEthereum(t *testing.T) {
	expected := "de0b295669a9fd93d5f28d9ec85e40f4cb697bae"
	actual, err := Resolve(client, "ethereum.eth")
	require.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveAddress(t *testing.T) {
	expected := "b8c2c29ee19d8307cb7255e1cd9cbde883a267d5"
	actual, err := Resolve(client, "0xb8c2C29ee19D8307cb7255e1Cd9CbDE883A267d5")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveShortAddress(t *testing.T) {
	expected := "0000000000000000000000000000000000000001"
	actual, err := Resolve(client, "0x1")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveHexString(t *testing.T) {
	_, err := Resolve(client, "0xe32c6d1a964749b6de2130e20daed821a45b9e7261118801ff5319d0ffc6b54a")
	assert.NotNil(t, err, "Resolved too-long hex string")
}

func TestReverseResolveTestEnsTest(t *testing.T) {
	expected := "nick.eth"
	address := common.HexToAddress("b8c2C29ee19D8307cb7255e1Cd9CbDE883A267d5")
	actual, err := ReverseResolve(client, address)
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}

func TestSubdomainResolveAddress(t *testing.T) {
	expected := "51050ec063d393217b436747617ad1c2285aeeee"
	actual, err := Resolve(client, "331.moo.nft-owner.eth")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestExternalResolverAddress(t *testing.T) {
	_, err := Resolve(client, "jesse.cb.id")
	require.NotNil(t, err, "Error resolving address")
	assert.Equal(t, "external resolver", err.Error(), "Unexpected error")
}

func TestInvalidSubdomainResolveAddress(t *testing.T) {
	_, err := Resolve(client, "a.b.c.d.e.eth")
	require.NotNil(t, err, "Error resolving address")
	assert.Equal(t, "unregistered name", err.Error(), "Unexpected error")
}
