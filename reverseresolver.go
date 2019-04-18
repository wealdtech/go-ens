// Copyright 2017-2019 Weald Technology Trading
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
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/contracts/reverseresolver"
)

// ReverseResolver is the structure for the reverse resolver contract
type ReverseResolver struct {
	contract *reverseresolver.Contract
}

// NewReverseResolver obtains the reverse resolver
func NewReverseResolver(client *ethclient.Client) (*ReverseResolver, error) {
	reverseRegistrar, err := NewReverseRegistrar(client)
	if err != nil {
		return nil, err
	}

	// Now fetch the default resolver
	address, err := reverseRegistrar.DefaultResolverAddress()
	if err != nil {
		return nil, err
	}

	return NewReverseResolverAt(client, address)
}

// NewReverseResolverAt obtains the reverse resolver at a given address
func NewReverseResolverAt(client *ethclient.Client, address common.Address) (*ReverseResolver, error) {
	// Instantiate the reverse registrar contract
	contract, err := reverseresolver.NewContract(address, client)
	if err != nil {
		return nil, err
	}

	// Ensure the contract is a resolver
	_, err = contract.Name(nil, NameHash("0.addr.reverse"))
	if err != nil && err.Error() == "no contract code at given address" {
		return nil, fmt.Errorf("not a resolver")
	}

	return &ReverseResolver{
		contract: contract,
	}, nil
}

// Name obtains the name for an address
func (r *ReverseResolver) Name(address common.Address) (string, error) {
	return r.contract.Name(nil, NameHash(fmt.Sprintf("%s.addr.reverse", address.Hex()[2:])))
}

// Format provides a string version of an address, reverse resolving it if possible
func Format(client *ethclient.Client, address common.Address) string {
	result, err := ReverseResolve(client, address)
	if err != nil {
		result = fmt.Sprintf("%s", address.Hex())
	}
	return result
}

// ReverseResolve resolves an address in to an ENS name
// This will return an error if the name is not found or otherwise 0
func ReverseResolve(client *ethclient.Client, address common.Address) (string, error) {
	resolver, err := NewReverseResolver(client)
	if err != nil {
		return "", err
	}

	// Resolve the name
	name, err := resolver.Name(address)
	if name == "" {
		err = errors.New("no resolution")
	}

	return name, err
}
