package ens

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3/contracts/universalresolver"
)

type UniversalResolver struct {
	backend      bind.ContractBackend
	Contract     *universalresolver.Contract
	ContractAddr common.Address
}

// NewUniversalResolver obtains the ENS UniversalResolver.
func NewUniversalResolver(backend bind.ContractBackend) (*UniversalResolver, error) {
	address, err := UniversalResolverContractAddress(backend, "mainnet")
	if err != nil {
		return nil, err
	}

	return NewUniversalResolverAt(backend, address)
}

// NewUniversalResolverAt obtains the ENS UniversalResolver at a given address.
func NewUniversalResolverAt(backend bind.ContractBackend, address common.Address) (*UniversalResolver, error) {
	contract, err := universalresolver.NewContract(address, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalResolver{
		backend:      backend,
		Contract:     contract,
		ContractAddr: address,
	}, nil
}

// UniversalResolverContractAddress obtains the address of the UniversalResolver contract for a chain.
func UniversalResolverContractAddress(_ bind.ContractBackend, network string) (common.Address, error) {
	switch network {
	case "mainnet":
		return common.HexToAddress("0xce01f8eee7E479C928F8919abD53E553a36CeF67"), nil
	case "sepolia":
		return common.HexToAddress("0xc8af999e38273d658be1b921b88a9ddf005769cc"), nil
	}
	return common.Address{}, fmt.Errorf("universal resolver not found on %s", network)
}
