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
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Name represents an ENS name, for example 'foo.bar.eth'.
type Name struct {
	client *ethclient.Client
	// Name is the full name of an ENS domain e.g. foo.bar.eth
	Name string
	// Domain is the domain of an ENS domain e.g. bar.eth
	Domain string
	// Label is the name part of an ENS domain e.g. foo
	Label string
}

// NewName creates an ENS name structure
// Note this does not create the name on-chain.
func NewName(client *ethclient.Client, name string) (*Name, error) {
	name = NormaliseDomain(name)
	domain := Domain(name)
	label, err := DomainPart(name, 1)
	if err != nil {
		return nil, err
	}

	controller, err := NewETHController(client, domain)
	if err != nil {
		return nil, err
	}
	isValid, err := controller.IsValid(name)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("name is not valid according to the rules of the registrar (too short, invalid characters, etc.)")
	}

	return &Name{
		client: client,
		Name:   name,
		Domain: domain,
		Label:  label,
	}, nil
}

// IsRegistered returns true if the name is registered in the registrar
func (n *Name) IsRegistered() (bool, error) {
	owner, err := n.Owner()
	if err != nil {
		return false, err
	}
	return owner != UnknownAddress, nil
}

// ExtendRegistration sends a transaction that extends the registration of the name.
func (n *Name) ExtendRegistration(opts *bind.TransactOpts) (*types.Transaction, error) {
	controller, err := NewETHController(n.client, n.Domain)
	if err != nil {
		return nil, err
	}

	isRegistered, err := n.IsRegistered()
	if err != nil {
		return nil, err
	}
	if !isRegistered {
		return nil, errors.New("name is not registered")
	}

	// TODO should compare to cost per second
	if opts.Value == nil || opts.Value.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("not enough funds to extend the registration")
	}

	return controller.Renew(opts, n.Name)
}

// RegistrationInterval obtains the minimum interval between commit and reveal
// when registering this name.
func (n *Name) RegistrationInterval() (time.Duration, error) {
	controller, err := NewETHController(n.client, n.Domain)
	if err != nil {
		return time.Duration(0), err
	}
	interval, err := controller.MinCommitmentInterval()
	if err != nil {
		return time.Duration(0), err
	}
	return time.Duration(interval.Int64()) * time.Second, nil
}

// RegisterStageOne sends a transaction that starts the registration process.
func (n *Name) RegisterStageOne(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, [32]byte, error) {
	var secret [32]byte
	_, err := rand.Read(secret[:])
	if err != nil {
		return nil, secret, err
	}

	isRegistered, err := n.IsRegistered()
	if err != nil {
		return nil, secret, err
	}
	if isRegistered {
		return nil, secret, errors.New("name is already registered")
	}

	controller, err := NewETHController(n.client, n.Domain)
	if err != nil {
		return nil, secret, err
	}

	signedTx, err := controller.Commit(opts, n.Label, owner, secret)
	return signedTx, secret, err
}

// RegisterStageTwo sends a transaction that completes the registration process.
// The owner must be the same as supplied in RegisterStageOne.
// The secret is that returned by RegisterStageOne.
// At least RegistrationInterval() time must have passed since the stage one
// transaction was mined for this to work.
func (n *Name) RegisterStageTwo(opts *bind.TransactOpts, owner common.Address, secret [32]byte) (*types.Transaction, error) {
	controller, err := NewETHController(n.client, n.Domain)
	if err != nil {
		return nil, err
	}
	commitTS, err := controller.CommitmentTime(n.Label, owner, secret)
	if err != nil {
		return nil, err
	}
	if commitTS.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("stage 2 attempted prior to successful stage 1 transaction")
	}
	commit := time.Unix(commitTS.Int64(), 0)

	minCommitIntervalTS, err := controller.MinCommitmentInterval()
	if err != nil {
		return nil, err
	}
	minCommitInterval := time.Duration(minCommitIntervalTS.Int64()) * time.Second
	minStageTwoTime := commit.Add(minCommitInterval)
	if time.Now().Before(minStageTwoTime) {
		return nil, errors.New("too early to send second transaction")
	}

	maxCommitIntervalTS, err := controller.MaxCommitmentInterval()
	if err != nil {
		return nil, err
	}
	maxCommitInterval := time.Duration(maxCommitIntervalTS.Int64()) * time.Second
	maxStageTwoTime := commit.Add(maxCommitInterval)
	if time.Now().After(maxStageTwoTime) {
		return nil, errors.New("too late to send second transaction")
	}

	return controller.Reveal(opts, n.Label, owner, secret)
}

// Expires obtain the time at which the registration for this name expires.
func (n *Name) Expires() (time.Time, error) {
	registrar, err := NewBaseRegistrar(n.client, n.Domain)
	if err != nil {
		return time.Unix(0, 0), err
	}

	expiryTS, err := registrar.Expiry(n.Label)
	if err != nil {
		return time.Unix(0, 0), err
	}

	if expiryTS.Int64() == 0 {
		return time.Unix(0, 0), errors.New("not registered")
	}

	return time.Unix(expiryTS.Int64(), 0), nil
}

// Administrator obtains the administrator for this name.
// The administrator can carry out operations on the name such as setting
// records, but is not the ultimate owner of the name.
func (n *Name) Administrator() (common.Address, error) {
	registry, err := NewRegistry(n.client)
	if err != nil {
		return UnknownAddress, err
	}
	return registry.Owner(n.Name)
}

// SetAdministrator sets the administrator for this name.
// The administrator can carry out operations on the name such as setting
// records, but is not the ultimate owner of the name.
func (n *Name) SetAdministrator(opts *bind.TransactOpts, administrator common.Address) (*types.Transaction, error) {
	registry, err := NewRegistry(n.client)
	if err != nil {
		return nil, err
	}
	// TODO ensure is either administrator or owner
	return registry.SetOwner(opts, n.Name, administrator)
}

// Owner obtains the owner for this name.
func (n *Name) Owner() (common.Address, error) {
	registrar, err := NewBaseRegistrar(n.client, n.Domain)
	if err != nil {
		return UnknownAddress, err
	}
	return registrar.Owner(n.Label)
}

// SetOwner sets the owner for this name.
func (n *Name) SetOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	registrar, err := NewBaseRegistrar(n.client, n.Domain)
	if err != nil {
		return nil, err
	}
	// TODO ensure is either owner
	return registrar.SetOwner(opts, n.Label, owner)
}
