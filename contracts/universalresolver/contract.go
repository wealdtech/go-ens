// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalresolver

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AbstractUniversalResolverResolverInfo is an auto generated low-level Go binding around an user-defined struct.
type AbstractUniversalResolverResolverInfo struct {
	Name     []byte
	Offset   *big.Int
	Node     [32]byte
	Resolver common.Address
	Extended bool
}

// CCIPBatcherBatch is an auto generated low-level Go binding around an user-defined struct.
type CCIPBatcherBatch struct {
	Lookups  []CCIPBatcherLookup
	Gateways []string
}

// CCIPBatcherLookup is an auto generated low-level Go binding around an user-defined struct.
type CCIPBatcherLookup struct {
	Target common.Address
	Call   []byte
	Data   []byte
	Flags  *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractENS\",\"name\":\"ens\",\"type\":\"address\"},{\"internalType\":\"contractIGatewayProvider\",\"name\":\"batchGatewayProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dns\",\"type\":\"bytes\"}],\"name\":\"DNSDecodingFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ens\",\"type\":\"string\"}],\"name\":\"DNSEncodingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"status\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"HttpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBatchGatewayResponse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"OffsetOutOfBoundsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"ResolverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"ResolverNotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"ResolverNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"primaryAddress\",\"type\":\"bytes\"}],\"name\":\"ReverseAddressMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"UnsupportedResolverProfile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"batchGatewayProvider\",\"outputs\":[{\"internalType\":\"contractIGatewayProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"ccipBatch\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ccipBatchCallback\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structCCIPBatcher.Lookup[]\",\"name\":\"lookups\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"internalType\":\"structCCIPBatcher.Batch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ccipReadCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"findResolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractENS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"requireResolver\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extended\",\"type\":\"bool\"}],\"internalType\":\"structAbstractUniversalResolver.ResolverInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveBatchCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveCallback\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveDirectCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"resolveDirectCallbackError\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"name\":\"resolveWithGateways\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"name\":\"resolveWithResolver\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"lookupAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"reverse\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"reverseAddressCallback\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reverseResolver\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"reverseNameCallback\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"lookupAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"gateways\",\"type\":\"string[]\"}],\"name\":\"reverseWithGateways\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"primary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reverseResolver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BatchGatewayProvider is a free data retrieval call binding the contract method 0x02cf2578.
//
// Solidity: function batchGatewayProvider() view returns(address)
func (_Contract *ContractCaller) BatchGatewayProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchGatewayProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchGatewayProvider is a free data retrieval call binding the contract method 0x02cf2578.
//
// Solidity: function batchGatewayProvider() view returns(address)
func (_Contract *ContractSession) BatchGatewayProvider() (common.Address, error) {
	return _Contract.Contract.BatchGatewayProvider(&_Contract.CallOpts)
}

// BatchGatewayProvider is a free data retrieval call binding the contract method 0x02cf2578.
//
// Solidity: function batchGatewayProvider() view returns(address)
func (_Contract *ContractCallerSession) BatchGatewayProvider() (common.Address, error) {
	return _Contract.Contract.BatchGatewayProvider(&_Contract.CallOpts)
}

// CcipBatch is a free data retrieval call binding the contract method 0x9f28e99d.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (_Contract *ContractCaller) CcipBatch(opts *bind.CallOpts, batch CCIPBatcherBatch) (CCIPBatcherBatch, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ccipBatch", batch)

	if err != nil {
		return *new(CCIPBatcherBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPBatcherBatch)).(*CCIPBatcherBatch)

	return out0, err

}

// CcipBatch is a free data retrieval call binding the contract method 0x9f28e99d.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (_Contract *ContractSession) CcipBatch(batch CCIPBatcherBatch) (CCIPBatcherBatch, error) {
	return _Contract.Contract.CcipBatch(&_Contract.CallOpts, batch)
}

// CcipBatch is a free data retrieval call binding the contract method 0x9f28e99d.
//
// Solidity: function ccipBatch(((address,bytes,bytes,uint256)[],string[]) batch) view returns(((address,bytes,bytes,uint256)[],string[]))
func (_Contract *ContractCallerSession) CcipBatch(batch CCIPBatcherBatch) (CCIPBatcherBatch, error) {
	return _Contract.Contract.CcipBatch(&_Contract.CallOpts, batch)
}

// CcipBatchCallback is a free data retrieval call binding the contract method 0xb536af76.
//
// Solidity: function ccipBatchCallback(bytes response, bytes extraData) view returns(((address,bytes,bytes,uint256)[],string[]) batch)
func (_Contract *ContractCaller) CcipBatchCallback(opts *bind.CallOpts, response []byte, extraData []byte) (CCIPBatcherBatch, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ccipBatchCallback", response, extraData)

	if err != nil {
		return *new(CCIPBatcherBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(CCIPBatcherBatch)).(*CCIPBatcherBatch)

	return out0, err

}

// CcipBatchCallback is a free data retrieval call binding the contract method 0xb536af76.
//
// Solidity: function ccipBatchCallback(bytes response, bytes extraData) view returns(((address,bytes,bytes,uint256)[],string[]) batch)
func (_Contract *ContractSession) CcipBatchCallback(response []byte, extraData []byte) (CCIPBatcherBatch, error) {
	return _Contract.Contract.CcipBatchCallback(&_Contract.CallOpts, response, extraData)
}

// CcipBatchCallback is a free data retrieval call binding the contract method 0xb536af76.
//
// Solidity: function ccipBatchCallback(bytes response, bytes extraData) view returns(((address,bytes,bytes,uint256)[],string[]) batch)
func (_Contract *ContractCallerSession) CcipBatchCallback(response []byte, extraData []byte) (CCIPBatcherBatch, error) {
	return _Contract.Contract.CcipBatchCallback(&_Contract.CallOpts, response, extraData)
}

// CcipReadCallback is a free data retrieval call binding the contract method 0xef46c0b8.
//
// Solidity: function ccipReadCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractCaller) CcipReadCallback(opts *bind.CallOpts, response []byte, extraData []byte) error {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ccipReadCallback", response, extraData)

	if err != nil {
		return err
	}

	return err

}

// CcipReadCallback is a free data retrieval call binding the contract method 0xef46c0b8.
//
// Solidity: function ccipReadCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractSession) CcipReadCallback(response []byte, extraData []byte) error {
	return _Contract.Contract.CcipReadCallback(&_Contract.CallOpts, response, extraData)
}

// CcipReadCallback is a free data retrieval call binding the contract method 0xef46c0b8.
//
// Solidity: function ccipReadCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractCallerSession) CcipReadCallback(response []byte, extraData []byte) error {
	return _Contract.Contract.CcipReadCallback(&_Contract.CallOpts, response, extraData)
}

// FindResolver is a free data retrieval call binding the contract method 0xa1cbcbaf.
//
// Solidity: function findResolver(bytes name) view returns(address, bytes32, uint256)
func (_Contract *ContractCaller) FindResolver(opts *bind.CallOpts, name []byte) (common.Address, [32]byte, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "findResolver", name)

	if err != nil {
		return *new(common.Address), *new([32]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FindResolver is a free data retrieval call binding the contract method 0xa1cbcbaf.
//
// Solidity: function findResolver(bytes name) view returns(address, bytes32, uint256)
func (_Contract *ContractSession) FindResolver(name []byte) (common.Address, [32]byte, *big.Int, error) {
	return _Contract.Contract.FindResolver(&_Contract.CallOpts, name)
}

// FindResolver is a free data retrieval call binding the contract method 0xa1cbcbaf.
//
// Solidity: function findResolver(bytes name) view returns(address, bytes32, uint256)
func (_Contract *ContractCallerSession) FindResolver(name []byte) (common.Address, [32]byte, *big.Int, error) {
	return _Contract.Contract.FindResolver(&_Contract.CallOpts, name)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contract *ContractCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contract *ContractSession) Registry() (common.Address, error) {
	return _Contract.Contract.Registry(&_Contract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contract *ContractCallerSession) Registry() (common.Address, error) {
	return _Contract.Contract.Registry(&_Contract.CallOpts)
}

// RequireResolver is a free data retrieval call binding the contract method 0xc285238a.
//
// Solidity: function requireResolver(bytes name) view returns((bytes,uint256,bytes32,address,bool) info)
func (_Contract *ContractCaller) RequireResolver(opts *bind.CallOpts, name []byte) (AbstractUniversalResolverResolverInfo, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "requireResolver", name)

	if err != nil {
		return *new(AbstractUniversalResolverResolverInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AbstractUniversalResolverResolverInfo)).(*AbstractUniversalResolverResolverInfo)

	return out0, err

}

// RequireResolver is a free data retrieval call binding the contract method 0xc285238a.
//
// Solidity: function requireResolver(bytes name) view returns((bytes,uint256,bytes32,address,bool) info)
func (_Contract *ContractSession) RequireResolver(name []byte) (AbstractUniversalResolverResolverInfo, error) {
	return _Contract.Contract.RequireResolver(&_Contract.CallOpts, name)
}

// RequireResolver is a free data retrieval call binding the contract method 0xc285238a.
//
// Solidity: function requireResolver(bytes name) view returns((bytes,uint256,bytes32,address,bool) info)
func (_Contract *ContractCallerSession) RequireResolver(name []byte) (AbstractUniversalResolverResolverInfo, error) {
	return _Contract.Contract.RequireResolver(&_Contract.CallOpts, name)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (_Contract *ContractCaller) Resolve(opts *bind.CallOpts, name []byte, data []byte) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolve", name, data)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (_Contract *ContractSession) Resolve(name []byte, data []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.Resolve(&_Contract.CallOpts, name, data)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (_Contract *ContractCallerSession) Resolve(name []byte, data []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.Resolve(&_Contract.CallOpts, name, data)
}

// ResolveBatchCallback is a free data retrieval call binding the contract method 0x491fc4f9.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractCaller) ResolveBatchCallback(opts *bind.CallOpts, response []byte, extraData []byte) error {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveBatchCallback", response, extraData)

	if err != nil {
		return err
	}

	return err

}

// ResolveBatchCallback is a free data retrieval call binding the contract method 0x491fc4f9.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractSession) ResolveBatchCallback(response []byte, extraData []byte) error {
	return _Contract.Contract.ResolveBatchCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveBatchCallback is a free data retrieval call binding the contract method 0x491fc4f9.
//
// Solidity: function resolveBatchCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractCallerSession) ResolveBatchCallback(response []byte, extraData []byte) error {
	return _Contract.Contract.ResolveBatchCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveCallback is a free data retrieval call binding the contract method 0xb4a85801.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) pure returns(bytes, address)
func (_Contract *ContractCaller) ResolveCallback(opts *bind.CallOpts, response []byte, extraData []byte) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveCallback", response, extraData)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// ResolveCallback is a free data retrieval call binding the contract method 0xb4a85801.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) pure returns(bytes, address)
func (_Contract *ContractSession) ResolveCallback(response []byte, extraData []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.ResolveCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveCallback is a free data retrieval call binding the contract method 0xb4a85801.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) pure returns(bytes, address)
func (_Contract *ContractCallerSession) ResolveCallback(response []byte, extraData []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.ResolveCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveDirectCallback is a free data retrieval call binding the contract method 0x55391bb8.
//
// Solidity: function resolveDirectCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractCaller) ResolveDirectCallback(opts *bind.CallOpts, response []byte, extraData []byte) error {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveDirectCallback", response, extraData)

	if err != nil {
		return err
	}

	return err

}

// ResolveDirectCallback is a free data retrieval call binding the contract method 0x55391bb8.
//
// Solidity: function resolveDirectCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractSession) ResolveDirectCallback(response []byte, extraData []byte) error {
	return _Contract.Contract.ResolveDirectCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveDirectCallback is a free data retrieval call binding the contract method 0x55391bb8.
//
// Solidity: function resolveDirectCallback(bytes response, bytes extraData) view returns()
func (_Contract *ContractCallerSession) ResolveDirectCallback(response []byte, extraData []byte) error {
	return _Contract.Contract.ResolveDirectCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveDirectCallbackError is a free data retrieval call binding the contract method 0x3c6cbda8.
//
// Solidity: function resolveDirectCallbackError(bytes response, bytes ) pure returns()
func (_Contract *ContractCaller) ResolveDirectCallbackError(opts *bind.CallOpts, response []byte, arg1 []byte) error {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveDirectCallbackError", response, arg1)

	if err != nil {
		return err
	}

	return err

}

// ResolveDirectCallbackError is a free data retrieval call binding the contract method 0x3c6cbda8.
//
// Solidity: function resolveDirectCallbackError(bytes response, bytes ) pure returns()
func (_Contract *ContractSession) ResolveDirectCallbackError(response []byte, arg1 []byte) error {
	return _Contract.Contract.ResolveDirectCallbackError(&_Contract.CallOpts, response, arg1)
}

// ResolveDirectCallbackError is a free data retrieval call binding the contract method 0x3c6cbda8.
//
// Solidity: function resolveDirectCallbackError(bytes response, bytes ) pure returns()
func (_Contract *ContractCallerSession) ResolveDirectCallbackError(response []byte, arg1 []byte) error {
	return _Contract.Contract.ResolveDirectCallbackError(&_Contract.CallOpts, response, arg1)
}

// ResolveWithGateways is a free data retrieval call binding the contract method 0xa1472844.
//
// Solidity: function resolveWithGateways(bytes name, bytes data, string[] gateways) view returns(bytes result, address resolver)
func (_Contract *ContractCaller) ResolveWithGateways(opts *bind.CallOpts, name []byte, data []byte, gateways []string) (struct {
	Result   []byte
	Resolver common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveWithGateways", name, data, gateways)

	outstruct := new(struct {
		Result   []byte
		Resolver common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Resolver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ResolveWithGateways is a free data retrieval call binding the contract method 0xa1472844.
//
// Solidity: function resolveWithGateways(bytes name, bytes data, string[] gateways) view returns(bytes result, address resolver)
func (_Contract *ContractSession) ResolveWithGateways(name []byte, data []byte, gateways []string) (struct {
	Result   []byte
	Resolver common.Address
}, error) {
	return _Contract.Contract.ResolveWithGateways(&_Contract.CallOpts, name, data, gateways)
}

// ResolveWithGateways is a free data retrieval call binding the contract method 0xa1472844.
//
// Solidity: function resolveWithGateways(bytes name, bytes data, string[] gateways) view returns(bytes result, address resolver)
func (_Contract *ContractCallerSession) ResolveWithGateways(name []byte, data []byte, gateways []string) (struct {
	Result   []byte
	Resolver common.Address
}, error) {
	return _Contract.Contract.ResolveWithGateways(&_Contract.CallOpts, name, data, gateways)
}

// ResolveWithResolver is a free data retrieval call binding the contract method 0x4a3e3994.
//
// Solidity: function resolveWithResolver(address resolver, bytes name, bytes data, string[] gateways) view returns(bytes)
func (_Contract *ContractCaller) ResolveWithResolver(opts *bind.CallOpts, resolver common.Address, name []byte, data []byte, gateways []string) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveWithResolver", resolver, name, data, gateways)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ResolveWithResolver is a free data retrieval call binding the contract method 0x4a3e3994.
//
// Solidity: function resolveWithResolver(address resolver, bytes name, bytes data, string[] gateways) view returns(bytes)
func (_Contract *ContractSession) ResolveWithResolver(resolver common.Address, name []byte, data []byte, gateways []string) ([]byte, error) {
	return _Contract.Contract.ResolveWithResolver(&_Contract.CallOpts, resolver, name, data, gateways)
}

// ResolveWithResolver is a free data retrieval call binding the contract method 0x4a3e3994.
//
// Solidity: function resolveWithResolver(address resolver, bytes name, bytes data, string[] gateways) view returns(bytes)
func (_Contract *ContractCallerSession) ResolveWithResolver(resolver common.Address, name []byte, data []byte, gateways []string) ([]byte, error) {
	return _Contract.Contract.ResolveWithResolver(&_Contract.CallOpts, resolver, name, data, gateways)
}

// Reverse is a free data retrieval call binding the contract method 0x5d78a217.
//
// Solidity: function reverse(bytes lookupAddress, uint256 coinType) view returns(string, address, address)
func (_Contract *ContractCaller) Reverse(opts *bind.CallOpts, lookupAddress []byte, coinType *big.Int) (string, common.Address, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverse", lookupAddress, coinType)

	if err != nil {
		return *new(string), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// Reverse is a free data retrieval call binding the contract method 0x5d78a217.
//
// Solidity: function reverse(bytes lookupAddress, uint256 coinType) view returns(string, address, address)
func (_Contract *ContractSession) Reverse(lookupAddress []byte, coinType *big.Int) (string, common.Address, common.Address, error) {
	return _Contract.Contract.Reverse(&_Contract.CallOpts, lookupAddress, coinType)
}

// Reverse is a free data retrieval call binding the contract method 0x5d78a217.
//
// Solidity: function reverse(bytes lookupAddress, uint256 coinType) view returns(string, address, address)
func (_Contract *ContractCallerSession) Reverse(lookupAddress []byte, coinType *big.Int) (string, common.Address, common.Address, error) {
	return _Contract.Contract.Reverse(&_Contract.CallOpts, lookupAddress, coinType)
}

// ReverseAddressCallback is a free data retrieval call binding the contract method 0x94fbfa87.
//
// Solidity: function reverseAddressCallback(bytes response, bytes extraData) pure returns(string primary, address resolver, address reverseResolver)
func (_Contract *ContractCaller) ReverseAddressCallback(opts *bind.CallOpts, response []byte, extraData []byte) (struct {
	Primary         string
	Resolver        common.Address
	ReverseResolver common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverseAddressCallback", response, extraData)

	outstruct := new(struct {
		Primary         string
		Resolver        common.Address
		ReverseResolver common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Primary = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Resolver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReverseResolver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ReverseAddressCallback is a free data retrieval call binding the contract method 0x94fbfa87.
//
// Solidity: function reverseAddressCallback(bytes response, bytes extraData) pure returns(string primary, address resolver, address reverseResolver)
func (_Contract *ContractSession) ReverseAddressCallback(response []byte, extraData []byte) (struct {
	Primary         string
	Resolver        common.Address
	ReverseResolver common.Address
}, error) {
	return _Contract.Contract.ReverseAddressCallback(&_Contract.CallOpts, response, extraData)
}

// ReverseAddressCallback is a free data retrieval call binding the contract method 0x94fbfa87.
//
// Solidity: function reverseAddressCallback(bytes response, bytes extraData) pure returns(string primary, address resolver, address reverseResolver)
func (_Contract *ContractCallerSession) ReverseAddressCallback(response []byte, extraData []byte) (struct {
	Primary         string
	Resolver        common.Address
	ReverseResolver common.Address
}, error) {
	return _Contract.Contract.ReverseAddressCallback(&_Contract.CallOpts, response, extraData)
}

// ReverseNameCallback is a free data retrieval call binding the contract method 0x575de750.
//
// Solidity: function reverseNameCallback(bytes response, bytes extraData) view returns(string primary, address, address)
func (_Contract *ContractCaller) ReverseNameCallback(opts *bind.CallOpts, response []byte, extraData []byte) (string, common.Address, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverseNameCallback", response, extraData)

	if err != nil {
		return *new(string), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// ReverseNameCallback is a free data retrieval call binding the contract method 0x575de750.
//
// Solidity: function reverseNameCallback(bytes response, bytes extraData) view returns(string primary, address, address)
func (_Contract *ContractSession) ReverseNameCallback(response []byte, extraData []byte) (string, common.Address, common.Address, error) {
	return _Contract.Contract.ReverseNameCallback(&_Contract.CallOpts, response, extraData)
}

// ReverseNameCallback is a free data retrieval call binding the contract method 0x575de750.
//
// Solidity: function reverseNameCallback(bytes response, bytes extraData) view returns(string primary, address, address)
func (_Contract *ContractCallerSession) ReverseNameCallback(response []byte, extraData []byte) (string, common.Address, common.Address, error) {
	return _Contract.Contract.ReverseNameCallback(&_Contract.CallOpts, response, extraData)
}

// ReverseWithGateways is a free data retrieval call binding the contract method 0xb7d6ca64.
//
// Solidity: function reverseWithGateways(bytes lookupAddress, uint256 coinType, string[] gateways) view returns(string primary, address resolver, address reverseResolver)
func (_Contract *ContractCaller) ReverseWithGateways(opts *bind.CallOpts, lookupAddress []byte, coinType *big.Int, gateways []string) (struct {
	Primary         string
	Resolver        common.Address
	ReverseResolver common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverseWithGateways", lookupAddress, coinType, gateways)

	outstruct := new(struct {
		Primary         string
		Resolver        common.Address
		ReverseResolver common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Primary = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Resolver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReverseResolver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ReverseWithGateways is a free data retrieval call binding the contract method 0xb7d6ca64.
//
// Solidity: function reverseWithGateways(bytes lookupAddress, uint256 coinType, string[] gateways) view returns(string primary, address resolver, address reverseResolver)
func (_Contract *ContractSession) ReverseWithGateways(lookupAddress []byte, coinType *big.Int, gateways []string) (struct {
	Primary         string
	Resolver        common.Address
	ReverseResolver common.Address
}, error) {
	return _Contract.Contract.ReverseWithGateways(&_Contract.CallOpts, lookupAddress, coinType, gateways)
}

// ReverseWithGateways is a free data retrieval call binding the contract method 0xb7d6ca64.
//
// Solidity: function reverseWithGateways(bytes lookupAddress, uint256 coinType, string[] gateways) view returns(string primary, address resolver, address reverseResolver)
func (_Contract *ContractCallerSession) ReverseWithGateways(lookupAddress []byte, coinType *big.Int, gateways []string) (struct {
	Primary         string
	Resolver        common.Address
	ReverseResolver common.Address
}, error) {
	return _Contract.Contract.ReverseWithGateways(&_Contract.CallOpts, lookupAddress, coinType, gateways)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}
