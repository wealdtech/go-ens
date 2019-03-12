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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/contracts/auctionregistrar"
	"github.com/wealdtech/go-ens/contracts/deed"
)

// DeedContract obtains the deed contract at a particular address
func DeedContract(client *ethclient.Client, address *common.Address) (*deed.DeedContract, error) {
	return deed.NewDeedContract(*address, client)
}

// DeedContractFor obtains the deed contract for a particular name
func DeedContractFor(client *ethclient.Client, registrar *auctionregistrar.AuctionRegistrarContract, name string) (*deed.DeedContract, error) {
	_, deedAddress, _, _, _, err := Entry(registrar, client, name)
	if err != nil {
		return nil, err
	}
	return DeedContract(client, &deedAddress)
}

// Owner obtains the owner of a deed
func Owner(contract *deed.DeedContract) (common.Address, error) {
	return contract.Owner(nil)
}

// PreviousOwner obtains the previous owner of a deed
func PreviousOwner(contract *deed.DeedContract) (common.Address, error) {
	return contract.PreviousOwner(nil)
}
