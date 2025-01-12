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
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/wealdtech/go-ens/v3/contracts/offchainresolver"
	"github.com/wealdtech/go-ens/v3/contracts/resolver"
)

var zeroHash = make([]byte, 32)
var offchainLookupSignature = "0x556f1830"

// UnknownAddress is the address to which unknown entries resolve.
var UnknownAddress = common.HexToAddress("00")

// Resolver is the structure for the resolver contract.
type Resolver struct {
	Contract     *resolver.Contract
	ContractAddr common.Address
	domain       string
	backend      bind.ContractBackend
	offResolver  *offchainresolver.Contract
}

// NewResolver obtains an ENS resolver for a given domain.
func NewResolver(backend bind.ContractBackend, domain string) (*Resolver, error) {
	ur, err := NewUniversalResolver(backend)
	if err != nil {
		return nil, err
	}
	lhash, err := DNSEncode(domain)
	if err != nil {
		return nil, err
	}
	rAddr, _, _, err := ur.Contract.FindResolver(nil, lhash)
	if err != nil || rAddr == common.Address(zeroHash) {
		return nil, errors.New("unregistered name")
	}
	return NewResolverAt(backend, domain, rAddr)
}

// NewResolverAt obtains an ENS resolver at a given address.
func NewResolverAt(backend bind.ContractBackend, domain string, address common.Address) (*Resolver, error) {
	contract, err := resolver.NewContract(address, backend)
	if err != nil {
		return nil, err
	}

	offR, err := offchainresolver.NewContract(address, backend)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		Contract:     contract,
		ContractAddr: address,
		domain:       domain,
		backend:      backend,
		offResolver:  offR,
	}, nil
}

// PublicResolverAddress obtains the address of the public resolver for a chain.
func PublicResolverAddress(backend bind.ContractBackend) (common.Address, error) {
	return Resolve(backend, "resolver.eth")
}

// Address returns the Ethereum address of the domain.
func (r *Resolver) Address() (common.Address, error) {
	nameHash, err := NameHash(r.domain)
	if err != nil {
		return UnknownAddress, err
	}
	addr, err := r.Contract.Addr(nil, nameHash)
	if err == nil && addr != common.Address(zeroHash) {
		return addr, err
	}
	ccipErr, errData := getCcipReadError(err)

	// try using the ENSIP-10 when the method is not found
	if !ccipErr {
		rAbi, _ := resolver.ContractMetaData.GetAbi()
		m := rAbi.Methods["addr"]
		args, _ := m.Inputs.Pack(nameHash)
		lhash, err := DNSEncode(r.domain)
		if err != nil {
			return common.Address{}, err
		}
		rawAddr, err := r.offResolver.Resolve(nil, lhash, append(m.ID, args...))
		if err == nil {
			return common.BytesToAddress(rawAddr), nil
		}
		ccipErr, errData = getCcipReadError(err)
		if errData == "" && addr == common.Address(zeroHash) {
			return UnknownAddress, errors.New("unregistered name")
		}
	}

	if !ccipErr {
		return UnknownAddress, errors.New("unregistered name")
	}

	rawAddr, err := ccipRead(r.backend, r.ContractAddr, errData)
	if err != nil || bytes.Equal(rawAddr, zeroHash) {
		return UnknownAddress, errors.New("unregistered name")
	}
	return common.BytesToAddress(rawAddr), nil
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

// Resolve resolves an ENS name in to an Etheruem address.
// This will return an error if the name is not found or otherwise 0.
func Resolve(backend bind.ContractBackend, input string) (common.Address, error) {
	if strings.Contains(input, ".") {
		return resolveName(backend, input)
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

func resolveName(backend bind.ContractBackend, input string) (common.Address, error) {
	nameHash, err := NameHash(input)
	if err != nil {
		return UnknownAddress, err
	}
	if bytes.Equal(nameHash[:], zeroHash) {
		return UnknownAddress, errors.New("bad name")
	}
	address, err := resolveHash(backend, input)
	if err != nil {
		return UnknownAddress, err
	}

	return address, nil
}

func resolveHash(backend bind.ContractBackend, domain string) (common.Address, error) {
	resolver, err := NewResolver(backend, domain)
	if err != nil {
		return UnknownAddress, err
	}
	return resolver.Address()
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
	text, err := r.Contract.Text(nil, nameHash, name)
	if err == nil && text != "" {
		return text, nil
	}
	ccipErr, errData := getCcipReadError(err)
	rAbi, _ := resolver.ContractMetaData.GetAbi()
	m := rAbi.Methods["text"]

	// try using the ENSIP-10 when the method is not found
	if !ccipErr {
		args, _ := m.Inputs.Pack(nameHash, name)
		lhash, err := DNSEncode(r.domain)
		if err != nil {
			return "", err
		}

		rawResp, err := r.offResolver.Resolve(nil, lhash, append(m.ID, args...))
		if err == nil {
			return string(rawResp), err
		}
		_, errData = getCcipReadError(err)
	}

	rawResp, err := ccipRead(r.backend, r.ContractAddr, errData)
	if err != nil {
		return "", err
	}
	outputs, err := m.Outputs.Unpack(rawResp)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(outputs[0]), nil
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
