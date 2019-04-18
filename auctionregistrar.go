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
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/contracts/auctionregistrar"
)

// AuctionRegistrar is the structure for the auction registrar contract
type AuctionRegistrar struct {
	client   *ethclient.Client
	domain   string
	contract *auctionregistrar.Contract
}

// AuctionEntry is an auction entry
type AuctionEntry struct {
	State        string
	Deed         common.Address
	Registration time.Time
	Value        *big.Int
	HighestBid   *big.Int
}

// NewAuctionRegistrar creates a new auction registrar for a given domain
func NewAuctionRegistrar(client *ethclient.Client, domain string) (*AuctionRegistrar, error) {
	address, err := RegistrarContractAddress(client, domain)
	if err != nil {
		return nil, err
	}

	return NewAuctionRegistrarContractAt(client, domain, address)
}

// NewAuctionRegistrarContractAt creates an auction registrar for a given domain at a given address
func NewAuctionRegistrarContractAt(client *ethclient.Client, domain string, address common.Address) (*AuctionRegistrar, error) {
	contract, err := auctionregistrar.NewContract(address, client)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrar{
		client:   client,
		domain:   domain,
		contract: contract,
	}, nil
}

// State returns the state of a nam
func (r *AuctionRegistrar) State(name string) (string, error) {
	entry, err := r.Entry(name)
	return entry.State, err
}

// Entry obtains a registrar entry for a name
func (r *AuctionRegistrar) Entry(name string) (*AuctionEntry, error) {
	domain, err := DomainPart(name, 1)
	if err != nil {
		return nil, fmt.Errorf("invalid name %s", name)
	}

	status, deedAddress, registration, value, highestBid, err := r.contract.Entries(nil, LabelHash(domain))
	if err != nil {
		return nil, err
	}

	entry := &AuctionEntry{
		Deed:       deedAddress,
		Value:      value,
		HighestBid: highestBid,
	}
	entry.Registration = time.Unix(registration.Int64(), 0)
	switch status {
	case 0:
		entry.State = "Available"
	case 1:
		entry.State = "Bidding"
	case 2:
		// Might be won or owned
		registry, err := NewRegistry(r.client)
		if err != nil {
			return nil, err
		}

		owner, err := registry.Owner(name)
		if err != nil {
			return nil, err
		}

		if owner == UnknownAddress {
			entry.State = "Won"
		} else {
			entry.State = "Owned"
		}
	case 3:
		entry.State = "Forbidden"
	case 4:
		entry.State = "Revealing"
	case 5:
		entry.State = "Unavailable"
	default:
		entry.State = "Unknown"
	}

	return entry, nil
}

// Migrate migrates a domain to the permanent registrar
func (r *AuctionRegistrar) Migrate(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	domain, err := DomainPart(name, 1)
	if err != nil {
		return nil, fmt.Errorf("invalid name %s", name)
	}

	return r.contract.TransferRegistrars(opts, LabelHash(domain))
}

// Owner obtains the owner of the deed that represents the name.
func (r *AuctionRegistrar) Owner(name string) (common.Address, error) {
	entry, err := r.Entry(name)
	if err != nil {
		return UnknownAddress, err
	}

	deed, err := NewDeedAt(r.client, entry.Deed)
	if err != nil {
		return UnknownAddress, err
	}
	return deed.Owner()
}
