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
	"context"
	"errors"
	"io"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/wealdtech/go-ens/v4/contracts/resolver"
)

// UnknownAddress is the address to which unknown entries resolve.
var UnknownAddress = common.HexToAddress("00")

// Resolver is the structure for the resolver contract.
type Resolver struct {
	Contract     *resolver.Contract
	ContractAddr common.Address
	domain       string
}

// NewResolver obtains an ENS resolver for a given domain.
func NewResolver(backend bind.ContractBackend, domain string) (*Resolver, error) {
	registry, err := NewRegistry(backend)
	if err != nil {
		return nil, err
	}

	// Ensure the name is registered.
	ownerAddress, err := registry.Owner(domain)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(ownerAddress.Bytes(), UnknownAddress.Bytes()) {
		return nil, errors.New("unregistered name")
	}

	// Obtain the resolver address for this domain.
	resolver, err := registry.ResolverAddress(domain)
	if err != nil {
		return nil, err
	}
	return NewResolverAt(backend, domain, resolver)
}

// NewResolverAt obtains an ENS resolver at a given address.
func NewResolverAt(backend bind.ContractBackend, domain string, address common.Address) (*Resolver, error) {
	contract, err := resolver.NewContract(address, backend)
	if err != nil {
		return nil, err
	}

	// Ensure this really is a resolver contract.
	nameHash, err := NameHash("test.eth")
	if err != nil {
		return nil, err
	}
	_, err = contract.Addr(nil, nameHash)
	if err != nil {
		if err.Error() == "no contract code at given address" {
			return nil, errors.New("no resolver")
		}
		return nil, err
	}

	return &Resolver{
		Contract:     contract,
		ContractAddr: address,
		domain:       domain,
	}, nil
}

// PublicResolverAddress obtains the address of the public resolver for a chain.
func PublicResolverAddress(ctx context.Context, backend bind.ContractBackend) (common.Address, error) {
	return Resolve(ctx, backend, "resolver.eth")
}

// Address returns the Ethereum address of the domain.
//
// This reads the addr(bytes32) record directly from the resolver contract this
// Resolver is bound to (located via the registry when the Resolver was
// created). It does NOT perform ENSIP-10 wildcard resolution or follow ERC-3668
// CCIP-Read, so names whose addresses live off-chain or on an L2 return the
// zero address or a raw revert error here. For wildcard- and CCIP-Read-aware
// resolution use Resolve or UniversalResolver.ResolveAddress instead.
func (r *Resolver) Address() (common.Address, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return UnknownAddress, err
	}
	return r.Contract.Addr(nil, nameHash)
}

// SetAddress sets the Ethereum address of the domain.
func (r *Resolver) SetAddress(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.SetAddr(opts, nameHash, address)
}

// MultiAddress returns the address of the domain for a given coin type.
// The coin type is as per https://github.com/satoshilabs/slips/blob/master/slip-0044.md
func (r *Resolver) MultiAddress(coinType uint64) ([]byte, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.Addr0(nil, nameHash, big.NewInt(int64(coinType)))
}

// SetMultiAddress sets the iaddress of the domain for a given coin type.
// The coin type is as per https://github.com/satoshilabs/slips/blob/master/slip-0044.md
func (r *Resolver) SetMultiAddress(opts *bind.TransactOpts, coinType uint64, address []byte) (*types.Transaction, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.SetAddr0(opts, nameHash, big.NewInt(int64(coinType)), address)
}

// PubKey returns the public key of the domain.
func (r *Resolver) PubKey() ([32]byte, [32]byte, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}
	res, err := r.Contract.Pubkey(nil, nameHash)
	return res.X, res.Y, err
}

// SetPubKey sets the public key of the domain.
func (r *Resolver) SetPubKey(opts *bind.TransactOpts, x [32]byte, y [32]byte) (*types.Transaction, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.SetPubkey(opts, nameHash, x, y)
}

// Contenthash returns the content hash of the domain.
func (r *Resolver) Contenthash() ([]byte, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.Contenthash(nil, nameHash)
}

// SetContenthash sets the content hash of the domain.
func (r *Resolver) SetContenthash(opts *bind.TransactOpts, contenthash []byte) (*types.Transaction, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.SetContenthash(opts, nameHash, contenthash)
}

// InterfaceImplementer returns the address of the contract that implements the given interface for the given domain.
func (r *Resolver) InterfaceImplementer(interfaceID [4]byte) (common.Address, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return UnknownAddress, err
	}
	return r.Contract.InterfaceImplementer(nil, nameHash, interfaceID)
}

// Resolve resolves an ENS name into an Ethereum address, or parses input
// directly if it is already a hex address.
//
// If input contains a dot it is treated as an ENS name and resolved through
// the ENS UniversalResolver, so ENSIP-10 wildcard resolution and ERC-3668
// CCIP-Read are followed transparently — names backed by offchain or L2 data
// resolve the same way as fully on-chain names. If input contains no dot it is
// treated as a literal hex address and returned after validation, without any
// on-chain lookup.
//
// ctx is honoured for cancellation and deadline propagation, flowing through
// the CCIP-Read hops, so a caller that sets a per-request deadline can cap the
// total resolution time end-to-end.
//
// It returns an error if the name is not found or otherwise resolves to the
// zero address.
func Resolve(ctx context.Context, backend bind.ContractBackend, input string) (common.Address, error) {
	if strings.Contains(input, ".") {
		return resolveName(ctx, backend, input)
	}
	if (strings.HasPrefix(input, "0x") && len(input) > 42) || (!strings.HasPrefix(input, "0x") && len(input) > 40) {
		return UnknownAddress, errors.New("address too long")
	}
	address := common.HexToAddress(input)
	if address == UnknownAddress {
		return UnknownAddress, errors.New("could not parse address")
	}

	return address, nil
}

func resolveName(ctx context.Context, backend bind.ContractBackend, input string) (common.Address, error) {
	if _, err := NameHash(input); err != nil {
		return UnknownAddress, err
	}
	ur, err := NewUniversalResolver(ctx, backend)
	if err != nil {
		return UnknownAddress, err
	}
	return ur.ResolveAddress(ctx, input)
}

// SetText sets the text associated with a name.
func (r *Resolver) SetText(opts *bind.TransactOpts, name string, value string) (*types.Transaction, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.SetText(opts, nameHash, name, value)
}

// Text obtains the text associated with a name.
func (r *Resolver) Text(name string) (string, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return "", err
	}
	return r.Contract.Text(nil, nameHash, name)
}

// SetABI sets the ABI associated with a name.
func (r *Resolver) SetABI(opts *bind.TransactOpts, _ string, abi string, contentType *big.Int) (*types.Transaction, error) {
	var data []byte
	switch contentType.Uint64() {
	case 1:
		// Uncompressed JSON.
		data = []byte(abi)
	case 2:
		// Zlib-compressed JSON.
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		if _, err := w.Write([]byte(abi)); err != nil {
			return nil, err
		}
		w.Close()
		data = b.Bytes()
	default:
		return nil, errors.New("unsupported content type")
	}

	nameHash, err := NameHash(r.domain)
	if err != nil {
		return nil, err
	}
	return r.Contract.SetABI(opts, nameHash, contentType, data)
}

// ABI returns the ABI associated with a name.
func (r *Resolver) ABI(name string) (string, error) {
	contentTypes := big.NewInt(3)
	nameHash, err := NameHash(name)
	if err != nil {
		return "", err
	}
	contentType, data, err := r.Contract.ABI(nil, nameHash, contentTypes)
	var abi string
	if err == nil {
		if contentType.Cmp(big.NewInt(1)) == 0 {
			// Uncompressed JSON.
			abi = string(data)
		} else if contentType.Cmp(big.NewInt(2)) == 0 {
			// Zlib-compressed JSON.
			b := bytes.NewReader(data)
			var z io.ReadCloser
			z, err = zlib.NewReader(b)
			if err != nil {
				return "", err
			}
			defer z.Close()
			var uncompressed []byte
			uncompressed, err = io.ReadAll(z)
			if err != nil {
				return "", err
			}
			abi = string(uncompressed)
		}
	}
	return abi, nil
}
