// Copyright 2019 Weald Technology Trading
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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v2/contracts/baseregistrar"
	"golang.org/x/crypto/sha3"
)

// BaseRegistrar is the structure for the registrar
type BaseRegistrar struct {
	client   *ethclient.Client
	domain   string
	contract *baseregistrar.Contract
}

// NewBaseRegistrar obtains the registrar contract for a given domain
func NewBaseRegistrar(client *ethclient.Client, domain string) (*BaseRegistrar, error) {
	address, err := RegistrarContractAddress(client, domain)
	if err != nil {
		return nil, err
	}

	if address == UnknownAddress {
		return nil, fmt.Errorf("no registrar for domain %s", domain)
	}

	contract, err := baseregistrar.NewContract(address, client)
	if err != nil {
		return nil, err
	}

	// Ensure this really is a base registrar.  To do this confirm that it supports
	// the expected interfaces.
	supported, err := contract.SupportsInterface(nil, [4]byte{0x6c, 0xcb, 0x2d, 0xf4})
	if err != nil {
		return nil, err
	}
	if !supported {
		return nil, fmt.Errorf("purported registrar for domain %s does not support nametoken functionality", domain)
	}
	supported, err = contract.SupportsInterface(nil, [4]byte{0x2d, 0xab, 0xbe, 0xed})
	if err != nil {
		return nil, err
	}
	if !supported {
		return nil, fmt.Errorf("purported registrar for domain %s does not support reclaim functionality", domain)
	}

	return &BaseRegistrar{
		client:   client,
		domain:   domain,
		contract: contract,
	}, nil
}

// PriorAuctionContract obtains the previous (auction) registrar contract
func (r *BaseRegistrar) PriorAuctionContract() (*AuctionRegistrar, error) {
	address, err := r.contract.PreviousRegistrar(nil)
	if err != nil {
		return nil, errors.New("no prior auction contract")
	}
	auctionContract, err := NewAuctionRegistrarAt(r.client, r.domain, address)
	if err != nil {
		return nil, errors.New("failed to instantiate prior auction contract")
	}

	// Confirm ths really is an auction contract by trying to create (but not submit) a bid
	var shaBid [32]byte
	var emptyHash [32]byte
	sha := sha3.NewLegacyKeccak256()
	sha.Write(emptyHash[:])
	sha.Write(UnknownAddress.Bytes())
	var amountBytes [32]byte
	sha.Write(amountBytes[:])
	sha.Write(emptyHash[:])
	sha.Sum(shaBid[:0])

	contractShaBid, err := auctionContract.ShaBid(emptyHash, UnknownAddress, big.NewInt(0), emptyHash)
	if err != nil {
		return nil, errors.New("failed to confirm auction contract")
	}
	if bytes.Compare(contractShaBid[:], shaBid[:]) != 0 {
		return nil, errors.New("failed to confirm auction contract")
	}

	return auctionContract, nil
}

// RegisteredWith returns one of "temporary", "permanent" or "none" for the
// registrar on which this name is registered
func (r *BaseRegistrar) RegisteredWith(domain string) (string, error) {
	name, err := UnqualifiedName(domain, r.domain)
	if err != nil {
		return "", err
	}
	// See if we're registered at all - fetch the owner to find out
	registry, err := NewRegistry(r.client)
	if err != nil {
		return "", err
	}
	owner, err := registry.Owner(domain)
	if err != nil {
		return "", err
	}
	if owner == UnknownAddress {
		return "", fmt.Errorf("%s not registered", domain)
	}

	// Fetch the temporary registrar and see if we're registered there
	auctionRegistrar, err := r.PriorAuctionContract()
	if err != nil {
		return "", err
	}
	if auctionRegistrar != nil {
		state, err := auctionRegistrar.State(name)
		if err != nil {
			return "", err
		}
		if state == "Won" || state == "Owned" {
			return "temporary", nil
		}
	}

	// No temporary registrar or no entry in same
	if owner == UnknownAddress {
		return "none", nil
	}
	return "permanent", nil
}

// Owner obtains the owner of the underlying token that represents the name.
func (r *BaseRegistrar) Owner(domain string) (common.Address, error) {
	name, err := UnqualifiedName(domain, r.domain)
	if err != nil {
		return UnknownAddress, err
	}
	hash := LabelHash(name)
	owner, err := r.contract.OwnerOf(nil, new(big.Int).SetBytes(hash[:]))
	// Registrar reverts rather than provide a 0 owner, so...
	if err != nil && err.Error() == "abi: unmarshalling empty output" {
		return UnknownAddress, nil
	}
	return owner, err
}

// SetOwner sets the owner of the token holding the name
func (r *BaseRegistrar) SetOwner(opts *bind.TransactOpts, domain string, newOwner common.Address) (*types.Transaction, error) {
	name, err := UnqualifiedName(domain, r.domain)
	if err != nil {
		return nil, err
	}
	owner, err := r.Owner(name)
	if err != nil {
		return nil, err
	}
	hash := LabelHash(name)
	id := new(big.Int).SetBytes(hash[:])
	return r.contract.TransferFrom(opts, owner, newOwner, id)
}

// Expiry obtains the unix timestamp at which the registration expires.
func (r *BaseRegistrar) Expiry(domain string) (*big.Int, error) {
	name, err := UnqualifiedName(domain, r.domain)
	if err != nil {
		return nil, err
	}
	hash := LabelHash(name)
	id := new(big.Int).SetBytes(hash[:])
	return r.contract.NameExpires(nil, id)
}