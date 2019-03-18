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
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/contracts/reverseregistrar"
	"github.com/wealdtech/go-ens/contracts/reverseresolver"
	"github.com/wealdtech/go-ens/util"
)

// Format provides a string version of an address, reverse resolving it if possible
func Format(client *ethclient.Client, input *common.Address) string {
	result, err := ReverseResolve(client, input)
	if err != nil {
		result = fmt.Sprintf("%s", input.Hex())
	}
	return result
}

// ReverseResolve resolves an address in to an ENS name
// This will return an error if the name is not found or otherwise 0
func ReverseResolve(client *ethclient.Client, input *common.Address) (string, error) {
	if input == nil {
		return "", errors.New("No address supplied")
	}

	nameHash := NameHash(input.Hex()[2:] + ".addr.reverse")

	contract, err := ReverseResolver(client)
	if err != nil {
		return "", err
	}

	// Resolve the name
	name, err := contract.Name(nil, nameHash)
	if name == "" {
		err = errors.New("No resolution")
	}

	return name, err
}

// ReverseResolver obtains the reverse resolver contract
func ReverseResolver(client *ethclient.Client) (*reverseresolver.ReverseResolver, error) {
	registryContract, err := RegistryContract(client)
	if err != nil {
		return nil, err
	}

	// Obtain the reverse registrar address
	reverseRegistrarAddress, err := registryContract.Owner(nil, NameHash("addr.reverse"))
	if err != nil {
		return nil, err
	}
	if bytes.Compare(reverseRegistrarAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		return nil, errors.New("unregistered name")
	}

	// Instantiate the reverse registrar contract
	reverseRegistrarContract, err := reverseregistrar.NewReverseRegistrarContract(reverseRegistrarAddress, client)
	if err != nil {
		return nil, err
	}

	// Now fetch the default resolver
	reverseResolverAddress, err := reverseRegistrarContract.DefaultResolver(nil)
	if err != nil {
		return nil, err
	}

	// Finally we can obtain the resolver itself
	return reverseresolver.NewReverseResolver(reverseResolverAddress, client)
}

// CreateReverseResolverSession creates a session suitable for multiple calls
func CreateReverseResolverSession(chainID *big.Int, wallet *accounts.Wallet, account *accounts.Account, passphrase string, contract *reverseresolver.ReverseResolver, gasPrice *big.Int) *reverseresolver.ReverseResolverSession {
	// Create a signer
	signer := util.AccountSigner(chainID, wallet, account, passphrase)

	// Return our session
	session := &reverseresolver.ReverseResolverSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     account.Address,
			Signer:   signer,
			GasPrice: gasPrice,
		},
	}

	return session
}
