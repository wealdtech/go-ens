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
	"bytes"
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v2/contracts/resolver"
)

var zeroHash = make([]byte, 32)

// UnknownAddress is the address to which unknown entries resolve
var UnknownAddress = common.HexToAddress("00")

// Resolver is the structure for the resolver contract
type Resolver struct {
	contract *resolver.Contract
	domain   string
}

// NewResolver obtains an ENS resolver for a given domain
func NewResolver(client *ethclient.Client, domain string) (*Resolver, error) {
	registry, err := NewRegistry(client)
	if err != nil {
		return nil, err
	}

	// Ensure the name is registered
	ownerAddress, err := registry.Owner(domain)
	if err != nil {
		return nil, err
	}
	if bytes.Compare(ownerAddress.Bytes(), UnknownAddress.Bytes()) == 0 {
		return nil, errors.New("unregistered name")
	}

	// Obtain the resolver address for this domain
	resolver, err := registry.ResolverAddress(domain)
	if err != nil {
		return nil, err
	}
	return NewResolverAt(client, domain, resolver)
}

// NewResolverAt obtains an ENS resolver at a given address
func NewResolverAt(client *ethclient.Client, domain string, address common.Address) (*Resolver, error) {
	contract, err := resolver.NewContract(address, client)
	if err != nil {
		return nil, err
	}

	// Ensure this really is a resolver contract
	_, err = contract.Addr(nil, NameHash("test.eth"))
	if err != nil {
		if err.Error() == "no contract code at given address" {
			return nil, errors.New("no resolver")
		}
		return nil, err
	}

	return &Resolver{
		contract: contract,
		domain:   domain,
	}, nil
}

// PublicResolverAddress obtains the address of the public resolver for a chain
func PublicResolverAddress(client *ethclient.Client) (common.Address, error) {
	return Resolve(client, "resolver.eth")
}

// Address returns the address of the domain
func (r *Resolver) Address(domain string) (common.Address, error) {
	return r.contract.Addr(nil, NameHash(domain))
}

// SetAddress sets the address of the domain
func (r *Resolver) SetAddress(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error) {
	return r.contract.SetAddr(opts, NameHash(r.domain), address)
}

// Contenthash returns the content hash of the domain
func (r *Resolver) Contenthash() ([]byte, error) {
	return r.contract.Contenthash(nil, NameHash(r.domain))
}

// SetContenthash sets the content hash of the domain
func (r *Resolver) SetContenthash(opts *bind.TransactOpts, contenthash []byte) (*types.Transaction, error) {
	return r.contract.SetContenthash(opts, NameHash(r.domain), contenthash)
}

// InterfaceImplementer returns the address of the contract that implements the given interface for the given domain
func (r *Resolver) InterfaceImplementer(interfaceID [4]byte) (common.Address, error) {
	return r.contract.InterfaceImplementer(nil, NameHash(r.domain), interfaceID)
}

// Resolve resolves an ENS name in to an Etheruem address
// This will return an error if the name is not found or otherwise 0
func Resolve(client *ethclient.Client, input string) (address common.Address, err error) {
	if strings.Contains(input, ".") {
		return resolveName(client, input)
	}
	if (strings.HasPrefix(input, "0x") && len(input) > 42) || (!strings.HasPrefix(input, "0x") && len(input) > 40) {
		err = errors.New("address too long")
	} else {
		address = common.HexToAddress(input)
		if address == UnknownAddress {
			err = errors.New("could not parse address")
		}
	}

	return
}

func resolveName(client *ethclient.Client, input string) (address common.Address, err error) {
	var nameHash [32]byte
	nameHash = NameHash(input)
	if bytes.Compare(nameHash[:], zeroHash) == 0 {
		err = errors.New("Bad name")
	} else {
		address, err = resolveHash(client, input)
	}
	return
}

func resolveHash(client *ethclient.Client, domain string) (address common.Address, err error) {
	resolver, err := NewResolver(client, domain)
	if err != nil {
		return UnknownAddress, err
	}

	// Resolve the domain
	address, err = resolver.Address(domain)
	if err != nil {
		return UnknownAddress, err
	}
	if bytes.Compare(address.Bytes(), UnknownAddress.Bytes()) == 0 {
		return UnknownAddress, errors.New("no address")
	}

	return
}

// SetAbi sets the ABI associated with a name
func SetAbi(session *resolver.ContractSession, name string, abi string, contentType *big.Int) (tx *types.Transaction, err error) {
	var data []byte
	if contentType.Cmp(big.NewInt(1)) == 0 {
		// Uncompressed JSON
		data = []byte(abi)
	} else if contentType.Cmp(big.NewInt(2)) == 0 {
		// Zlib-compressed JSON
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write([]byte(abi))
		w.Close()
		data = b.Bytes()
	} else {
		err = errors.New("Unsupported content type")
	}
	// Gas cost is around 50000 + 64 per byte; add 10000 headroom to be safe
	//session.TransactOpts.GasLimit = big.NewInt(int64(600000 + len(data)*64))
	tx, err = session.SetABI(NameHash(name), contentType, data)

	return
}

// Abi returns the ABI associated with a name
func Abi(resolver *resolver.Contract, name string) (abi string, err error) {
	contentTypes := big.NewInt(3)
	contentType, data, err := resolver.ABI(nil, NameHash(name), contentTypes)
	if err == nil {
		if contentType.Cmp(big.NewInt(1)) == 0 {
			// Uncompressed JSON
			abi = string(data)
		} else if contentType.Cmp(big.NewInt(2)) == 0 {
			// Zlib-compressed JSON
			b := bytes.NewReader(data)
			var z io.ReadCloser
			z, err = zlib.NewReader(b)
			if err != nil {
				return
			}
			defer z.Close()
			var uncompressed []byte
			uncompressed, err = ioutil.ReadAll(z)
			abi = string(uncompressed)
		}
	}
	return
}
