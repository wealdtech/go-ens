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
	assert.Equal(t, "unregistered name", err.Error(), "Unexpected error")
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

func TestOffchainResolverAddress(t *testing.T) {
	expected := "849151d7d0bf1f34b70d5cad5149d28cc2308bf1"
	actual, err := Resolve(client, "jesse.cb.id")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestInvalidSubdomainResolveAddress(t *testing.T) {
	_, err := Resolve(client, "xxxx.cb.id")
	require.NotNil(t, err, "Error resolving address")
	assert.Equal(t, "unregistered name", err.Error(), "Unexpected error")
}

func TestOnchainReadText(t *testing.T) {
	expected := "blockful_io"
	r, err := NewResolver(client, "blockful.eth")
	require.Nil(t, err, "Error creating resolver")
	actual, err := r.Text("com.twitter")
	require.Nil(t, err, "Error reading com.twitter")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}

func TestOffchainReadText(t *testing.T) {
	expected := "https://twitter.com/sid_coelho"
	r, err := NewResolver(client, "sid.cb.id")
	require.Nil(t, err, "Error creating resolver")
	actual, err := r.Text("com.twitter")
	require.Nil(t, err, "Error reading twitter")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}

func TestOffchainReadAddressTLD(t *testing.T) {
	expected := "179a862703a4adfb29896552df9e307980d19285"
	actual, err := Resolve(client, "gregskril.eth")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestOffchainReadAddressSubdomain(t *testing.T) {
	expected := "179a862703a4adfb29896552df9e307980d19285"
	actual, err := Resolve(client, "gregskril.uni.eth")
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestOffchainReadAddressSubdomainThroughResolver(t *testing.T) {
	expected := "179a862703a4adfb29896552df9e307980d19285"
	r, err := NewResolver(client, "gregskril.uni.eth")
	require.Nil(t, err, "Error resolving address")
	actual, err := r.Address()
	require.Nil(t, err, "Error resolving address")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestOffchainReadTextSubdomain(t *testing.T) {
	expected := "gregskril"
	r, err := NewResolver(client, "gregskril.uni.eth")
	require.Nil(t, err, "Error creating resolver")
	actual, err := r.Text("com.twitter")
	require.Nil(t, err, "Error reading twitter")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}

func TestBaseText(t *testing.T) {
	expected := "jessepollak"
	r, err := NewResolver(client, "jesse.base.eth")
	require.Nil(t, err, "Error creating resolver")
	actual, err := r.Text("com.github")
	require.Nil(t, err, "Error reading github")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}

func TestBaseAddress(t *testing.T) {
	expected := common.HexToAddress("0x849151d7D0bF1F34b70d5caD5149D28CC2308bf1")
	r, err := NewResolver(client, "jesse.base.eth")
	require.Nil(t, err, "Error creating resolver")
	actual, err := r.Address()
	require.Nil(t, err, "Error reading github")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}

// it has a custom CCIP-Read callback `addrCallback` instead of `resolveWithProof`
func TestLineaAddress(t *testing.T) {
	expected := common.HexToAddress("0x3C6aE1aEE2AE2D6f9cB79f1F9EBF466402EF1bb6")
	r, err := NewResolver(client, "james.linea.eth")
	require.Nil(t, err, "Error creating resolver")
	actual, err := r.Address()
	require.Nil(t, err, "Error reading github")
	assert.Equal(t, expected, actual, "Did not receive expected result")
}
