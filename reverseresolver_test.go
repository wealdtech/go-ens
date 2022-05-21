// Copyright 2022 Weald Technology Trading
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

package ens_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	ens "github.com/wealdtech/go-ens/v3"
)

// TestReverseResolve tests the reverse resolution functionality.
func TestReverseResolve(t *testing.T) {

	tests := []struct {
		name    string
		address common.Address
		res     string
		err     string
	}{
		{
			name:    "NoResolver",
			address: common.Address{},
			err:     "not a resolver",
		},
		{
			name:    "NoReverseRecord",
			address: common.HexToAddress("0x388ea662ef2c223ec0b047d41bf3c0f362142ad5"),
			err:     "no resolution",
		},
		{
			name:    "Exists",
			address: common.HexToAddress("0x809FA673fe2ab515FaA168259cB14E2BeDeBF68e"),
			res:     "avsa.eth",
		},
	}

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ens.ReverseResolve(client, test.address)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.res, res)
			}
		})
	}
}
