// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnsresolver

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"setDNSRecords\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameEntriesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_key\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"hasDNSRecords\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"records\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_contentType\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_name\",\"type\":\"bytes32\"},{\"name\":\"_resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"clearDNSZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"versions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"Deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"Cleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"}]"

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
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) constant returns(uint256 contentType, bytes data)
func (_Contract *ContractCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	ret := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "ABI", node, contentTypes)
	return *ret, err
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) constant returns(uint256 contentType, bytes data)
func (_Contract *ContractSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _Contract.Contract.ABI(&_Contract.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) constant returns(uint256 contentType, bytes data)
func (_Contract *ContractCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _Contract.Contract.ABI(&_Contract.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) constant returns(address ret)
func (_Contract *ContractCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) constant returns(address ret)
func (_Contract *ContractSession) Addr(node [32]byte) (common.Address, error) {
	return _Contract.Contract.Addr(&_Contract.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) constant returns(address ret)
func (_Contract *ContractCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Contract.Contract.Addr(&_Contract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(bytes32 node) constant returns(bytes32 ret)
func (_Contract *ContractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(bytes32 node) constant returns(bytes32 ret)
func (_Contract *ContractSession) Content(node [32]byte) ([32]byte, error) {
	return _Contract.Contract.Content(&_Contract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(bytes32 node) constant returns(bytes32 ret)
func (_Contract *ContractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _Contract.Contract.Content(&_Contract.CallOpts, node)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 _node, bytes32 _name, uint16 _resource) constant returns(bytes)
func (_Contract *ContractCaller) DnsRecord(opts *bind.CallOpts, _node [32]byte, _name [32]byte, _resource uint16) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "dnsRecord", _node, _name, _resource)
	return *ret0, err
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 _node, bytes32 _name, uint16 _resource) constant returns(bytes)
func (_Contract *ContractSession) DnsRecord(_node [32]byte, _name [32]byte, _resource uint16) ([]byte, error) {
	return _Contract.Contract.DnsRecord(&_Contract.CallOpts, _node, _name, _resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(bytes32 _node, bytes32 _name, uint16 _resource) constant returns(bytes)
func (_Contract *ContractCallerSession) DnsRecord(_node [32]byte, _name [32]byte, _resource uint16) ([]byte, error) {
	return _Contract.Contract.DnsRecord(&_Contract.CallOpts, _node, _name, _resource)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 _node, bytes32 _name) constant returns(bool)
func (_Contract *ContractCaller) HasDNSRecords(opts *bind.CallOpts, _node [32]byte, _name [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "hasDNSRecords", _node, _name)
	return *ret0, err
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 _node, bytes32 _name) constant returns(bool)
func (_Contract *ContractSession) HasDNSRecords(_node [32]byte, _name [32]byte) (bool, error) {
	return _Contract.Contract.HasDNSRecords(&_Contract.CallOpts, _node, _name)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(bytes32 _node, bytes32 _name) constant returns(bool)
func (_Contract *ContractCallerSession) HasDNSRecords(_node [32]byte, _name [32]byte) (bool, error) {
	return _Contract.Contract.HasDNSRecords(&_Contract.CallOpts, _node, _name)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) constant returns(string ret)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "name", node)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) constant returns(string ret)
func (_Contract *ContractSession) Name(node [32]byte) (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) constant returns(string ret)
func (_Contract *ContractCallerSession) Name(node [32]byte) (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts, node)
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x0dee3863.
//
// Solidity: function nameEntriesCount(bytes32 , uint16 , bytes32 ) constant returns(uint16)
func (_Contract *ContractCaller) NameEntriesCount(opts *bind.CallOpts, arg0 [32]byte, arg1 uint16, arg2 [32]byte) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "nameEntriesCount", arg0, arg1, arg2)
	return *ret0, err
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x0dee3863.
//
// Solidity: function nameEntriesCount(bytes32 , uint16 , bytes32 ) constant returns(uint16)
func (_Contract *ContractSession) NameEntriesCount(arg0 [32]byte, arg1 uint16, arg2 [32]byte) (uint16, error) {
	return _Contract.Contract.NameEntriesCount(&_Contract.CallOpts, arg0, arg1, arg2)
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x0dee3863.
//
// Solidity: function nameEntriesCount(bytes32 , uint16 , bytes32 ) constant returns(uint16)
func (_Contract *ContractCallerSession) NameEntriesCount(arg0 [32]byte, arg1 uint16, arg2 [32]byte) (uint16, error) {
	return _Contract.Contract.NameEntriesCount(&_Contract.CallOpts, arg0, arg1, arg2)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) constant returns(bytes32 x, bytes32 y)
func (_Contract *ContractCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	ret := new(struct {
		X [32]byte
		Y [32]byte
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "pubkey", node)
	return *ret, err
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) constant returns(bytes32 x, bytes32 y)
func (_Contract *ContractSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Contract.Contract.Pubkey(&_Contract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) constant returns(bytes32 x, bytes32 y)
func (_Contract *ContractCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Contract.Contract.Pubkey(&_Contract.CallOpts, node)
}

// Records is a free data retrieval call binding the contract method 0x4f96bc00.
//
// Solidity: function records(bytes32 , uint16 , bytes32 , uint16 ) constant returns(bytes)
func (_Contract *ContractCaller) Records(opts *bind.CallOpts, arg0 [32]byte, arg1 uint16, arg2 [32]byte, arg3 uint16) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "records", arg0, arg1, arg2, arg3)
	return *ret0, err
}

// Records is a free data retrieval call binding the contract method 0x4f96bc00.
//
// Solidity: function records(bytes32 , uint16 , bytes32 , uint16 ) constant returns(bytes)
func (_Contract *ContractSession) Records(arg0 [32]byte, arg1 uint16, arg2 [32]byte, arg3 uint16) ([]byte, error) {
	return _Contract.Contract.Records(&_Contract.CallOpts, arg0, arg1, arg2, arg3)
}

// Records is a free data retrieval call binding the contract method 0x4f96bc00.
//
// Solidity: function records(bytes32 , uint16 , bytes32 , uint16 ) constant returns(bytes)
func (_Contract *ContractCallerSession) Records(arg0 [32]byte, arg1 uint16, arg2 [32]byte, arg3 uint16) ([]byte, error) {
	return _Contract.Contract.Records(&_Contract.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) constant returns(string ret)
func (_Contract *ContractCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "text", node, key)
	return *ret0, err
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) constant returns(string ret)
func (_Contract *ContractSession) Text(node [32]byte, key string) (string, error) {
	return _Contract.Contract.Text(&_Contract.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) constant returns(string ret)
func (_Contract *ContractCallerSession) Text(node [32]byte, key string) (string, error) {
	return _Contract.Contract.Text(&_Contract.CallOpts, node, key)
}

// Versions is a free data retrieval call binding the contract method 0xc7cec7f8.
//
// Solidity: function versions(bytes32 ) constant returns(uint16)
func (_Contract *ContractCaller) Versions(opts *bind.CallOpts, arg0 [32]byte) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "versions", arg0)
	return *ret0, err
}

// Versions is a free data retrieval call binding the contract method 0xc7cec7f8.
//
// Solidity: function versions(bytes32 ) constant returns(uint16)
func (_Contract *ContractSession) Versions(arg0 [32]byte) (uint16, error) {
	return _Contract.Contract.Versions(&_Contract.CallOpts, arg0)
}

// Versions is a free data retrieval call binding the contract method 0xc7cec7f8.
//
// Solidity: function versions(bytes32 ) constant returns(uint16)
func (_Contract *ContractCallerSession) Versions(arg0 [32]byte) (uint16, error) {
	return _Contract.Contract.Versions(&_Contract.CallOpts, arg0)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(bytes32 _node) returns()
func (_Contract *ContractTransactor) ClearDNSZone(opts *bind.TransactOpts, _node [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "clearDNSZone", _node)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(bytes32 _node) returns()
func (_Contract *ContractSession) ClearDNSZone(_node [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClearDNSZone(&_Contract.TransactOpts, _node)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(bytes32 _node) returns()
func (_Contract *ContractTransactorSession) ClearDNSZone(_node [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClearDNSZone(&_Contract.TransactOpts, _node)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 _node, uint256 _contentType, bytes _data) returns()
func (_Contract *ContractTransactor) SetABI(opts *bind.TransactOpts, _node [32]byte, _contentType *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setABI", _node, _contentType, _data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 _node, uint256 _contentType, bytes _data) returns()
func (_Contract *ContractSession) SetABI(_node [32]byte, _contentType *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetABI(&_Contract.TransactOpts, _node, _contentType, _data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 _node, uint256 _contentType, bytes _data) returns()
func (_Contract *ContractTransactorSession) SetABI(_node [32]byte, _contentType *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetABI(&_Contract.TransactOpts, _node, _contentType, _data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 _node, address _addr) returns()
func (_Contract *ContractTransactor) SetAddr(opts *bind.TransactOpts, _node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAddr", _node, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 _node, address _addr) returns()
func (_Contract *ContractSession) SetAddr(_node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr(&_Contract.TransactOpts, _node, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 _node, address _addr) returns()
func (_Contract *ContractTransactorSession) SetAddr(_node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr(&_Contract.TransactOpts, _node, _addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(bytes32 _node, bytes32 _hash) returns()
func (_Contract *ContractTransactor) SetContent(opts *bind.TransactOpts, _node [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setContent", _node, _hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(bytes32 _node, bytes32 _hash) returns()
func (_Contract *ContractSession) SetContent(_node [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetContent(&_Contract.TransactOpts, _node, _hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(bytes32 _node, bytes32 _hash) returns()
func (_Contract *ContractTransactorSession) SetContent(_node [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetContent(&_Contract.TransactOpts, _node, _hash)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 _node, bytes _data) returns()
func (_Contract *ContractTransactor) SetDNSRecords(opts *bind.TransactOpts, _node [32]byte, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDNSRecords", _node, _data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 _node, bytes _data) returns()
func (_Contract *ContractSession) SetDNSRecords(_node [32]byte, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetDNSRecords(&_Contract.TransactOpts, _node, _data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(bytes32 _node, bytes _data) returns()
func (_Contract *ContractTransactorSession) SetDNSRecords(_node [32]byte, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetDNSRecords(&_Contract.TransactOpts, _node, _data)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 _node, string _name) returns()
func (_Contract *ContractTransactor) SetName(opts *bind.TransactOpts, _node [32]byte, _name string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setName", _node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 _node, string _name) returns()
func (_Contract *ContractSession) SetName(_node [32]byte, _name string) (*types.Transaction, error) {
	return _Contract.Contract.SetName(&_Contract.TransactOpts, _node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 _node, string _name) returns()
func (_Contract *ContractTransactorSession) SetName(_node [32]byte, _name string) (*types.Transaction, error) {
	return _Contract.Contract.SetName(&_Contract.TransactOpts, _node, _name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 _node, bytes32 _x, bytes32 _y) returns()
func (_Contract *ContractTransactor) SetPubkey(opts *bind.TransactOpts, _node [32]byte, _x [32]byte, _y [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPubkey", _node, _x, _y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 _node, bytes32 _x, bytes32 _y) returns()
func (_Contract *ContractSession) SetPubkey(_node [32]byte, _x [32]byte, _y [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetPubkey(&_Contract.TransactOpts, _node, _x, _y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 _node, bytes32 _x, bytes32 _y) returns()
func (_Contract *ContractTransactorSession) SetPubkey(_node [32]byte, _x [32]byte, _y [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetPubkey(&_Contract.TransactOpts, _node, _x, _y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 _node, string _key, string _value) returns()
func (_Contract *ContractTransactor) SetText(opts *bind.TransactOpts, _node [32]byte, _key string, _value string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setText", _node, _key, _value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 _node, string _key, string _value) returns()
func (_Contract *ContractSession) SetText(_node [32]byte, _key string, _value string) (*types.Transaction, error) {
	return _Contract.Contract.SetText(&_Contract.TransactOpts, _node, _key, _value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 _node, string _key, string _value) returns()
func (_Contract *ContractTransactorSession) SetText(_node [32]byte, _key string, _value string) (*types.Transaction, error) {
	return _Contract.Contract.SetText(&_Contract.TransactOpts, _node, _key, _value)
}

// ContractABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the Contract contract.
type ContractABIChangedIterator struct {
	Event *ContractABIChanged // Event containing the contract specifics and raw log

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
func (it *ContractABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractABIChanged)
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
		it.Event = new(ContractABIChanged)
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
func (it *ContractABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractABIChanged represents a ABIChanged event raised by the Contract contract.
type ContractABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Contract *ContractFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*ContractABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractABIChangedIterator{contract: _Contract.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Contract *ContractFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *ContractABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractABIChanged)
				if err := _Contract.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ContractAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the Contract contract.
type ContractAddrChangedIterator struct {
	Event *ContractAddrChanged // Event containing the contract specifics and raw log

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
func (it *ContractAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddrChanged)
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
		it.Event = new(ContractAddrChanged)
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
func (it *ContractAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddrChanged represents a AddrChanged event raised by the Contract contract.
type ContractAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Contract *ContractFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddrChangedIterator{contract: _Contract.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Contract *ContractFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ContractAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddrChanged)
				if err := _Contract.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ContractClearedIterator is returned from FilterCleared and is used to iterate over the raw logs and unpacked data for Cleared events raised by the Contract contract.
type ContractClearedIterator struct {
	Event *ContractCleared // Event containing the contract specifics and raw log

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
func (it *ContractClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCleared)
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
		it.Event = new(ContractCleared)
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
func (it *ContractClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCleared represents a Cleared event raised by the Contract contract.
type ContractCleared struct {
	Node [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCleared is a free log retrieval operation binding the contract event 0xf9c9793a5e1c345ae880b8c5e3b0a241f568d7242dcf6e21a2bea5dfb817d88e.
//
// Solidity: event Cleared(bytes32 node)
func (_Contract *ContractFilterer) FilterCleared(opts *bind.FilterOpts) (*ContractClearedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Cleared")
	if err != nil {
		return nil, err
	}
	return &ContractClearedIterator{contract: _Contract.contract, event: "Cleared", logs: logs, sub: sub}, nil
}

// WatchCleared is a free log subscription operation binding the contract event 0xf9c9793a5e1c345ae880b8c5e3b0a241f568d7242dcf6e21a2bea5dfb817d88e.
//
// Solidity: event Cleared(bytes32 node)
func (_Contract *ContractFilterer) WatchCleared(opts *bind.WatchOpts, sink chan<- *ContractCleared) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Cleared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCleared)
				if err := _Contract.contract.UnpackLog(event, "Cleared", log); err != nil {
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

// ContractContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the Contract contract.
type ContractContentChangedIterator struct {
	Event *ContractContentChanged // Event containing the contract specifics and raw log

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
func (it *ContractContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractContentChanged)
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
		it.Event = new(ContractContentChanged)
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
func (it *ContractContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractContentChanged represents a ContentChanged event raised by the Contract contract.
type ContractContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: event ContentChanged(bytes32 indexed node, bytes32 hash)
func (_Contract *ContractFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractContentChangedIterator{contract: _Contract.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: event ContentChanged(bytes32 indexed node, bytes32 hash)
func (_Contract *ContractFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *ContractContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractContentChanged)
				if err := _Contract.contract.UnpackLog(event, "ContentChanged", log); err != nil {
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

// ContractDeletedIterator is returned from FilterDeleted and is used to iterate over the raw logs and unpacked data for Deleted events raised by the Contract contract.
type ContractDeletedIterator struct {
	Event *ContractDeleted // Event containing the contract specifics and raw log

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
func (it *ContractDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeleted)
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
		it.Event = new(ContractDeleted)
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
func (it *ContractDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeleted represents a Deleted event raised by the Contract contract.
type ContractDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleted is a free log retrieval operation binding the contract event 0x133052c72ea386f24d31f74751f618e877370038e43ae5a1571abd4e7039a10b.
//
// Solidity: event Deleted(bytes32 node, bytes name, uint16 resource)
func (_Contract *ContractFilterer) FilterDeleted(opts *bind.FilterOpts) (*ContractDeletedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Deleted")
	if err != nil {
		return nil, err
	}
	return &ContractDeletedIterator{contract: _Contract.contract, event: "Deleted", logs: logs, sub: sub}, nil
}

// WatchDeleted is a free log subscription operation binding the contract event 0x133052c72ea386f24d31f74751f618e877370038e43ae5a1571abd4e7039a10b.
//
// Solidity: event Deleted(bytes32 node, bytes name, uint16 resource)
func (_Contract *ContractFilterer) WatchDeleted(opts *bind.WatchOpts, sink chan<- *ContractDeleted) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Deleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeleted)
				if err := _Contract.contract.UnpackLog(event, "Deleted", log); err != nil {
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

// ContractNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the Contract contract.
type ContractNameChangedIterator struct {
	Event *ContractNameChanged // Event containing the contract specifics and raw log

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
func (it *ContractNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNameChanged)
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
		it.Event = new(ContractNameChanged)
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
func (it *ContractNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNameChanged represents a NameChanged event raised by the Contract contract.
type ContractNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Contract *ContractFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractNameChangedIterator{contract: _Contract.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Contract *ContractFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *ContractNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNameChanged)
				if err := _Contract.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ContractPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the Contract contract.
type ContractPubkeyChangedIterator struct {
	Event *ContractPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *ContractPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPubkeyChanged)
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
		it.Event = new(ContractPubkeyChanged)
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
func (it *ContractPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPubkeyChanged represents a PubkeyChanged event raised by the Contract contract.
type ContractPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractPubkeyChangedIterator{contract: _Contract.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *ContractPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPubkeyChanged)
				if err := _Contract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ContractTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the Contract contract.
type ContractTextChangedIterator struct {
	Event *ContractTextChanged // Event containing the contract specifics and raw log

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
func (it *ContractTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTextChanged)
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
		it.Event = new(ContractTextChanged)
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
func (it *ContractTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTextChanged represents a TextChanged event raised by the Contract contract.
type ContractTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Contract *ContractFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ContractTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ContractTextChangedIterator{contract: _Contract.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Contract *ContractFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ContractTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTextChanged)
				if err := _Contract.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ContractUpdatedIterator is returned from FilterUpdated and is used to iterate over the raw logs and unpacked data for Updated events raised by the Contract contract.
type ContractUpdatedIterator struct {
	Event *ContractUpdated // Event containing the contract specifics and raw log

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
func (it *ContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdated)
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
		it.Event = new(ContractUpdated)
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
func (it *ContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdated represents a Updated event raised by the Contract contract.
type ContractUpdated struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdated is a free log retrieval operation binding the contract event 0xb09b445b0776998f69b77758912de74d641195b10c8f03062f124495b13ef9f6.
//
// Solidity: event Updated(bytes32 node, bytes name, uint16 resource)
func (_Contract *ContractFilterer) FilterUpdated(opts *bind.FilterOpts) (*ContractUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedIterator{contract: _Contract.contract, event: "Updated", logs: logs, sub: sub}, nil
}

// WatchUpdated is a free log subscription operation binding the contract event 0xb09b445b0776998f69b77758912de74d641195b10c8f03062f124495b13ef9f6.
//
// Solidity: event Updated(bytes32 node, bytes name, uint16 resource)
func (_Contract *ContractFilterer) WatchUpdated(opts *bind.WatchOpts, sink chan<- *ContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdated)
				if err := _Contract.contract.UnpackLog(event, "Updated", log); err != nil {
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
