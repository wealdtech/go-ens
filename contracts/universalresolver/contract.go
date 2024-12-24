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

// Result is an auto generated low-level Go binding around an user-defined struct.
type Result struct {
	Success    bool
	ReturnData []byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_urls\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_resolveSingle\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gateways\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"callbackFunction\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"metaData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchGatewayURLs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"findResolver\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractResolver\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractENS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gateways\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structResult[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gateways\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structResult[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveCallback\",\"inputs\":[{\"name\":\"response\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structResult[]\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolveSingleCallback\",\"inputs\":[{\"name\":\"response\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reverse\",\"inputs\":[{\"name\":\"reverseName\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gateways\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reverse\",\"inputs\":[{\"name\":\"reverseName\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reverseCallback\",\"inputs\":[{\"name\":\"response\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setGatewayURLs\",\"inputs\":[{\"name\":\"_urls\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OffchainLookup\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"urls\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callbackFunction\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ResolverError\",\"inputs\":[{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ResolverNotContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ResolverNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ResolverWildcardNotSupported\",\"inputs\":[]}]",
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

// ResolveSingle is a free data retrieval call binding the contract method 0x8e25a0f3.
//
// Solidity: function _resolveSingle(bytes name, bytes data, string[] gateways, bytes4 callbackFunction, bytes metaData) view returns(bytes, address)
func (_Contract *ContractCaller) ResolveSingle(opts *bind.CallOpts, name []byte, data []byte, gateways []string, callbackFunction [4]byte, metaData []byte) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_resolveSingle", name, data, gateways, callbackFunction, metaData)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// ResolveSingle is a free data retrieval call binding the contract method 0x8e25a0f3.
//
// Solidity: function _resolveSingle(bytes name, bytes data, string[] gateways, bytes4 callbackFunction, bytes metaData) view returns(bytes, address)
func (_Contract *ContractSession) ResolveSingle(name []byte, data []byte, gateways []string, callbackFunction [4]byte, metaData []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.ResolveSingle(&_Contract.CallOpts, name, data, gateways, callbackFunction, metaData)
}

// ResolveSingle is a free data retrieval call binding the contract method 0x8e25a0f3.
//
// Solidity: function _resolveSingle(bytes name, bytes data, string[] gateways, bytes4 callbackFunction, bytes metaData) view returns(bytes, address)
func (_Contract *ContractCallerSession) ResolveSingle(name []byte, data []byte, gateways []string, callbackFunction [4]byte, metaData []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.ResolveSingle(&_Contract.CallOpts, name, data, gateways, callbackFunction, metaData)
}

// BatchGatewayURLs is a free data retrieval call binding the contract method 0xa6b16419.
//
// Solidity: function batchGatewayURLs(uint256 ) view returns(string)
func (_Contract *ContractCaller) BatchGatewayURLs(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchGatewayURLs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BatchGatewayURLs is a free data retrieval call binding the contract method 0xa6b16419.
//
// Solidity: function batchGatewayURLs(uint256 ) view returns(string)
func (_Contract *ContractSession) BatchGatewayURLs(arg0 *big.Int) (string, error) {
	return _Contract.Contract.BatchGatewayURLs(&_Contract.CallOpts, arg0)
}

// BatchGatewayURLs is a free data retrieval call binding the contract method 0xa6b16419.
//
// Solidity: function batchGatewayURLs(uint256 ) view returns(string)
func (_Contract *ContractCallerSession) BatchGatewayURLs(arg0 *big.Int) (string, error) {
	return _Contract.Contract.BatchGatewayURLs(&_Contract.CallOpts, arg0)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
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

// Resolve is a free data retrieval call binding the contract method 0x0667cfea.
//
// Solidity: function resolve(bytes name, bytes data, string[] gateways) view returns(bytes, address)
func (_Contract *ContractCaller) Resolve(opts *bind.CallOpts, name []byte, data []byte, gateways []string) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolve", name, data, gateways)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Resolve is a free data retrieval call binding the contract method 0x0667cfea.
//
// Solidity: function resolve(bytes name, bytes data, string[] gateways) view returns(bytes, address)
func (_Contract *ContractSession) Resolve(name []byte, data []byte, gateways []string) ([]byte, common.Address, error) {
	return _Contract.Contract.Resolve(&_Contract.CallOpts, name, data, gateways)
}

// Resolve is a free data retrieval call binding the contract method 0x0667cfea.
//
// Solidity: function resolve(bytes name, bytes data, string[] gateways) view returns(bytes, address)
func (_Contract *ContractCallerSession) Resolve(name []byte, data []byte, gateways []string) ([]byte, common.Address, error) {
	return _Contract.Contract.Resolve(&_Contract.CallOpts, name, data, gateways)
}

// Resolve0 is a free data retrieval call binding the contract method 0x206c74c9.
//
// Solidity: function resolve(bytes name, bytes[] data) view returns((bool,bytes)[], address)
func (_Contract *ContractCaller) Resolve0(opts *bind.CallOpts, name []byte, data [][]byte) ([]Result, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolve0", name, data)

	if err != nil {
		return *new([]Result), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]Result)).(*[]Result)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Resolve0 is a free data retrieval call binding the contract method 0x206c74c9.
//
// Solidity: function resolve(bytes name, bytes[] data) view returns((bool,bytes)[], address)
func (_Contract *ContractSession) Resolve0(name []byte, data [][]byte) ([]Result, common.Address, error) {
	return _Contract.Contract.Resolve0(&_Contract.CallOpts, name, data)
}

// Resolve0 is a free data retrieval call binding the contract method 0x206c74c9.
//
// Solidity: function resolve(bytes name, bytes[] data) view returns((bool,bytes)[], address)
func (_Contract *ContractCallerSession) Resolve0(name []byte, data [][]byte) ([]Result, common.Address, error) {
	return _Contract.Contract.Resolve0(&_Contract.CallOpts, name, data)
}

// Resolve1 is a free data retrieval call binding the contract method 0x76286c00.
//
// Solidity: function resolve(bytes name, bytes[] data, string[] gateways) view returns((bool,bytes)[], address)
func (_Contract *ContractCaller) Resolve1(opts *bind.CallOpts, name []byte, data [][]byte, gateways []string) ([]Result, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolve1", name, data, gateways)

	if err != nil {
		return *new([]Result), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]Result)).(*[]Result)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Resolve1 is a free data retrieval call binding the contract method 0x76286c00.
//
// Solidity: function resolve(bytes name, bytes[] data, string[] gateways) view returns((bool,bytes)[], address)
func (_Contract *ContractSession) Resolve1(name []byte, data [][]byte, gateways []string) ([]Result, common.Address, error) {
	return _Contract.Contract.Resolve1(&_Contract.CallOpts, name, data, gateways)
}

// Resolve1 is a free data retrieval call binding the contract method 0x76286c00.
//
// Solidity: function resolve(bytes name, bytes[] data, string[] gateways) view returns((bool,bytes)[], address)
func (_Contract *ContractCallerSession) Resolve1(name []byte, data [][]byte, gateways []string) ([]Result, common.Address, error) {
	return _Contract.Contract.Resolve1(&_Contract.CallOpts, name, data, gateways)
}

// Resolve2 is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (_Contract *ContractCaller) Resolve2(opts *bind.CallOpts, name []byte, data []byte) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolve2", name, data)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Resolve2 is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (_Contract *ContractSession) Resolve2(name []byte, data []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.Resolve2(&_Contract.CallOpts, name, data)
}

// Resolve2 is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes, address)
func (_Contract *ContractCallerSession) Resolve2(name []byte, data []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.Resolve2(&_Contract.CallOpts, name, data)
}

// ResolveCallback is a free data retrieval call binding the contract method 0xb4a85801.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) view returns((bool,bytes)[], address)
func (_Contract *ContractCaller) ResolveCallback(opts *bind.CallOpts, response []byte, extraData []byte) ([]Result, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveCallback", response, extraData)

	if err != nil {
		return *new([]Result), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]Result)).(*[]Result)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// ResolveCallback is a free data retrieval call binding the contract method 0xb4a85801.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) view returns((bool,bytes)[], address)
func (_Contract *ContractSession) ResolveCallback(response []byte, extraData []byte) ([]Result, common.Address, error) {
	return _Contract.Contract.ResolveCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveCallback is a free data retrieval call binding the contract method 0xb4a85801.
//
// Solidity: function resolveCallback(bytes response, bytes extraData) view returns((bool,bytes)[], address)
func (_Contract *ContractCallerSession) ResolveCallback(response []byte, extraData []byte) ([]Result, common.Address, error) {
	return _Contract.Contract.ResolveCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveSingleCallback is a free data retrieval call binding the contract method 0xe0a85412.
//
// Solidity: function resolveSingleCallback(bytes response, bytes extraData) view returns(bytes, address)
func (_Contract *ContractCaller) ResolveSingleCallback(opts *bind.CallOpts, response []byte, extraData []byte) ([]byte, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveSingleCallback", response, extraData)

	if err != nil {
		return *new([]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// ResolveSingleCallback is a free data retrieval call binding the contract method 0xe0a85412.
//
// Solidity: function resolveSingleCallback(bytes response, bytes extraData) view returns(bytes, address)
func (_Contract *ContractSession) ResolveSingleCallback(response []byte, extraData []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.ResolveSingleCallback(&_Contract.CallOpts, response, extraData)
}

// ResolveSingleCallback is a free data retrieval call binding the contract method 0xe0a85412.
//
// Solidity: function resolveSingleCallback(bytes response, bytes extraData) view returns(bytes, address)
func (_Contract *ContractCallerSession) ResolveSingleCallback(response []byte, extraData []byte) ([]byte, common.Address, error) {
	return _Contract.Contract.ResolveSingleCallback(&_Contract.CallOpts, response, extraData)
}

// Reverse is a free data retrieval call binding the contract method 0xb241d0d3.
//
// Solidity: function reverse(bytes reverseName, string[] gateways) view returns(string, address, address, address)
func (_Contract *ContractCaller) Reverse(opts *bind.CallOpts, reverseName []byte, gateways []string) (string, common.Address, common.Address, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverse", reverseName, gateways)

	if err != nil {
		return *new(string), *new(common.Address), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// Reverse is a free data retrieval call binding the contract method 0xb241d0d3.
//
// Solidity: function reverse(bytes reverseName, string[] gateways) view returns(string, address, address, address)
func (_Contract *ContractSession) Reverse(reverseName []byte, gateways []string) (string, common.Address, common.Address, common.Address, error) {
	return _Contract.Contract.Reverse(&_Contract.CallOpts, reverseName, gateways)
}

// Reverse is a free data retrieval call binding the contract method 0xb241d0d3.
//
// Solidity: function reverse(bytes reverseName, string[] gateways) view returns(string, address, address, address)
func (_Contract *ContractCallerSession) Reverse(reverseName []byte, gateways []string) (string, common.Address, common.Address, common.Address, error) {
	return _Contract.Contract.Reverse(&_Contract.CallOpts, reverseName, gateways)
}

// Reverse0 is a free data retrieval call binding the contract method 0xec11c823.
//
// Solidity: function reverse(bytes reverseName) view returns(string, address, address, address)
func (_Contract *ContractCaller) Reverse0(opts *bind.CallOpts, reverseName []byte) (string, common.Address, common.Address, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverse0", reverseName)

	if err != nil {
		return *new(string), *new(common.Address), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// Reverse0 is a free data retrieval call binding the contract method 0xec11c823.
//
// Solidity: function reverse(bytes reverseName) view returns(string, address, address, address)
func (_Contract *ContractSession) Reverse0(reverseName []byte) (string, common.Address, common.Address, common.Address, error) {
	return _Contract.Contract.Reverse0(&_Contract.CallOpts, reverseName)
}

// Reverse0 is a free data retrieval call binding the contract method 0xec11c823.
//
// Solidity: function reverse(bytes reverseName) view returns(string, address, address, address)
func (_Contract *ContractCallerSession) Reverse0(reverseName []byte) (string, common.Address, common.Address, common.Address, error) {
	return _Contract.Contract.Reverse0(&_Contract.CallOpts, reverseName)
}

// ReverseCallback is a free data retrieval call binding the contract method 0x6dc4fb73.
//
// Solidity: function reverseCallback(bytes response, bytes extraData) view returns(string, address, address, address)
func (_Contract *ContractCaller) ReverseCallback(opts *bind.CallOpts, response []byte, extraData []byte) (string, common.Address, common.Address, common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverseCallback", response, extraData)

	if err != nil {
		return *new(string), *new(common.Address), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// ReverseCallback is a free data retrieval call binding the contract method 0x6dc4fb73.
//
// Solidity: function reverseCallback(bytes response, bytes extraData) view returns(string, address, address, address)
func (_Contract *ContractSession) ReverseCallback(response []byte, extraData []byte) (string, common.Address, common.Address, common.Address, error) {
	return _Contract.Contract.ReverseCallback(&_Contract.CallOpts, response, extraData)
}

// ReverseCallback is a free data retrieval call binding the contract method 0x6dc4fb73.
//
// Solidity: function reverseCallback(bytes response, bytes extraData) view returns(string, address, address, address)
func (_Contract *ContractCallerSession) ReverseCallback(response []byte, extraData []byte) (string, common.Address, common.Address, common.Address, error) {
	return _Contract.Contract.ReverseCallback(&_Contract.CallOpts, response, extraData)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetGatewayURLs is a paid mutator transaction binding the contract method 0x8e5ea8df.
//
// Solidity: function setGatewayURLs(string[] _urls) returns()
func (_Contract *ContractTransactor) SetGatewayURLs(opts *bind.TransactOpts, _urls []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGatewayURLs", _urls)
}

// SetGatewayURLs is a paid mutator transaction binding the contract method 0x8e5ea8df.
//
// Solidity: function setGatewayURLs(string[] _urls) returns()
func (_Contract *ContractSession) SetGatewayURLs(_urls []string) (*types.Transaction, error) {
	return _Contract.Contract.SetGatewayURLs(&_Contract.TransactOpts, _urls)
}

// SetGatewayURLs is a paid mutator transaction binding the contract method 0x8e5ea8df.
//
// Solidity: function setGatewayURLs(string[] _urls) returns()
func (_Contract *ContractTransactorSession) SetGatewayURLs(_urls []string) (*types.Transaction, error) {
	return _Contract.Contract.SetGatewayURLs(&_Contract.TransactOpts, _urls)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
