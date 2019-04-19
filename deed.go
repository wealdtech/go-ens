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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v2/contracts/deed"
)

// Deed is the structure for the deed
type Deed struct {
	contract *deed.Contract
}

// NewDeed obtains the deed contract for a given domain
func NewDeed(client *ethclient.Client, domain string) (*Deed, error) {

	// Obtain the auction registrar for the deed
	auctionRegistrar, err := NewAuctionRegistrar(client, domain)
	if err != nil {
		return nil, err
	}

	entry, err := auctionRegistrar.Entry(domain)
	if err != nil {
		return nil, err
	}

	return NewDeedAt(client, entry.Deed)
}

// NewDeedAt creates a deed contract at a given address
func NewDeedAt(client *ethclient.Client, address common.Address) (*Deed, error) {
	contract, err := deed.NewContract(address, client)
	if err != nil {
		return nil, err
	}

	return &Deed{
		contract: contract,
	}, nil
}

// Owner obtains the owner of the deed
func (c *Deed) Owner() (common.Address, error) {
	return c.contract.Owner(nil)
}

// PreviousOwner obtains the previous owner of the deed
func (c *Deed) PreviousOwner() (common.Address, error) {
	return c.contract.PreviousOwner(nil)
}

// SetOwner sets the owner of the deed
func (c *Deed) SetOwner(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error) {
	return c.contract.SetOwner(opts, address)
}
