// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auctionregistrar

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

// AuctionRegistrarContractABI is the input ABI used to generate the binding from.
const AuctionRegistrarContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"releaseDeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getAllowedTime\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unhashedName\",\"type\":\"string\"}],\"name\":\"invalidateName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"shaBid\",\"outputs\":[{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\"},{\"name\":\"seal\",\"type\":\"bytes32\"}],\"name\":\"cancelBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"entries\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"unsealBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"transferRegistrars\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sealedBids\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"finalizeAuction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registryStarted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"launchLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"name\":\"newBid\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"labels\",\"type\":\"bytes32[]\"}],\"name\":\"eraseNode\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"startAuctions\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"deed\",\"type\":\"address\"},{\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"acceptRegistrarTransfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"startAuction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashes\",\"type\":\"bytes32[]\"},{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"name\":\"startAuctionsAndBid\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_rootNode\",\"type\":\"bytes32\"},{\"name\":\"_startDate\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"HashRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"HashReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"HashInvalidated\",\"type\":\"event\"}]"

// AuctionRegistrarContract is an auto generated Go binding around an Ethereum contract.
type AuctionRegistrarContract struct {
	AuctionRegistrarContractCaller     // Read-only binding to the contract
	AuctionRegistrarContractTransactor // Write-only binding to the contract
	AuctionRegistrarContractFilterer   // Log filterer for contract events
}

// AuctionRegistrarContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionRegistrarContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionRegistrarContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionRegistrarContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionRegistrarContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionRegistrarContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionRegistrarContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionRegistrarContractSession struct {
	Contract     *AuctionRegistrarContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AuctionRegistrarContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionRegistrarContractCallerSession struct {
	Contract *AuctionRegistrarContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AuctionRegistrarContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionRegistrarContractTransactorSession struct {
	Contract     *AuctionRegistrarContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AuctionRegistrarContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRegistrarContractRaw struct {
	Contract *AuctionRegistrarContract // Generic contract binding to access the raw methods on
}

// AuctionRegistrarContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionRegistrarContractCallerRaw struct {
	Contract *AuctionRegistrarContractCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionRegistrarContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionRegistrarContractTransactorRaw struct {
	Contract *AuctionRegistrarContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionRegistrarContract creates a new instance of AuctionRegistrarContract, bound to a specific deployed contract.
func NewAuctionRegistrarContract(address common.Address, backend bind.ContractBackend) (*AuctionRegistrarContract, error) {
	contract, err := bindAuctionRegistrarContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContract{AuctionRegistrarContractCaller: AuctionRegistrarContractCaller{contract: contract}, AuctionRegistrarContractTransactor: AuctionRegistrarContractTransactor{contract: contract}, AuctionRegistrarContractFilterer: AuctionRegistrarContractFilterer{contract: contract}}, nil
}

// NewAuctionRegistrarContractCaller creates a new read-only instance of AuctionRegistrarContract, bound to a specific deployed contract.
func NewAuctionRegistrarContractCaller(address common.Address, caller bind.ContractCaller) (*AuctionRegistrarContractCaller, error) {
	contract, err := bindAuctionRegistrarContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractCaller{contract: contract}, nil
}

// NewAuctionRegistrarContractTransactor creates a new write-only instance of AuctionRegistrarContract, bound to a specific deployed contract.
func NewAuctionRegistrarContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionRegistrarContractTransactor, error) {
	contract, err := bindAuctionRegistrarContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractTransactor{contract: contract}, nil
}

// NewAuctionRegistrarContractFilterer creates a new log filterer instance of AuctionRegistrarContract, bound to a specific deployed contract.
func NewAuctionRegistrarContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionRegistrarContractFilterer, error) {
	contract, err := bindAuctionRegistrarContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractFilterer{contract: contract}, nil
}

// bindAuctionRegistrarContract binds a generic wrapper to an already deployed contract.
func bindAuctionRegistrarContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionRegistrarContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionRegistrarContract *AuctionRegistrarContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuctionRegistrarContract.Contract.AuctionRegistrarContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionRegistrarContract *AuctionRegistrarContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.AuctionRegistrarContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionRegistrarContract *AuctionRegistrarContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.AuctionRegistrarContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuctionRegistrarContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.contract.Transact(opts, method, params...)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) Ens() (common.Address, error) {
	return _AuctionRegistrarContract.Contract.Ens(&_AuctionRegistrarContract.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) Ens() (common.Address, error) {
	return _AuctionRegistrarContract.Contract.Ens(&_AuctionRegistrarContract.CallOpts)
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(bytes32 _hash) constant returns(uint8, address, uint256, uint256, uint256)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) Entries(opts *bind.CallOpts, _hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _AuctionRegistrarContract.contract.Call(opts, out, "entries", _hash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(bytes32 _hash) constant returns(uint8, address, uint256, uint256, uint256)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) Entries(_hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AuctionRegistrarContract.Contract.Entries(&_AuctionRegistrarContract.CallOpts, _hash)
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(bytes32 _hash) constant returns(uint8, address, uint256, uint256, uint256)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) Entries(_hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _AuctionRegistrarContract.Contract.Entries(&_AuctionRegistrarContract.CallOpts, _hash)
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(bytes32 _hash) constant returns(uint256 timestamp)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) GetAllowedTime(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "getAllowedTime", _hash)
	return *ret0, err
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(bytes32 _hash) constant returns(uint256 timestamp)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) GetAllowedTime(_hash [32]byte) (*big.Int, error) {
	return _AuctionRegistrarContract.Contract.GetAllowedTime(&_AuctionRegistrarContract.CallOpts, _hash)
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(bytes32 _hash) constant returns(uint256 timestamp)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) GetAllowedTime(_hash [32]byte) (*big.Int, error) {
	return _AuctionRegistrarContract.Contract.GetAllowedTime(&_AuctionRegistrarContract.CallOpts, _hash)
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(bytes32 _hash, uint256 _timestamp) constant returns(bool allowed)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) IsAllowed(opts *bind.CallOpts, _hash [32]byte, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "isAllowed", _hash, _timestamp)
	return *ret0, err
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(bytes32 _hash, uint256 _timestamp) constant returns(bool allowed)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) IsAllowed(_hash [32]byte, _timestamp *big.Int) (bool, error) {
	return _AuctionRegistrarContract.Contract.IsAllowed(&_AuctionRegistrarContract.CallOpts, _hash, _timestamp)
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(bytes32 _hash, uint256 _timestamp) constant returns(bool allowed)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) IsAllowed(_hash [32]byte, _timestamp *big.Int) (bool, error) {
	return _AuctionRegistrarContract.Contract.IsAllowed(&_AuctionRegistrarContract.CallOpts, _hash, _timestamp)
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) LaunchLength(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "launchLength")
	return *ret0, err
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) LaunchLength() (uint32, error) {
	return _AuctionRegistrarContract.Contract.LaunchLength(&_AuctionRegistrarContract.CallOpts)
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) LaunchLength() (uint32, error) {
	return _AuctionRegistrarContract.Contract.LaunchLength(&_AuctionRegistrarContract.CallOpts)
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) RegistryStarted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "registryStarted")
	return *ret0, err
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) RegistryStarted() (*big.Int, error) {
	return _AuctionRegistrarContract.Contract.RegistryStarted(&_AuctionRegistrarContract.CallOpts)
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) RegistryStarted() (*big.Int, error) {
	return _AuctionRegistrarContract.Contract.RegistryStarted(&_AuctionRegistrarContract.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "rootNode")
	return *ret0, err
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) RootNode() ([32]byte, error) {
	return _AuctionRegistrarContract.Contract.RootNode(&_AuctionRegistrarContract.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) RootNode() ([32]byte, error) {
	return _AuctionRegistrarContract.Contract.RootNode(&_AuctionRegistrarContract.CallOpts)
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids(address , bytes32 ) constant returns(address)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) SealedBids(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "sealedBids", arg0, arg1)
	return *ret0, err
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids(address , bytes32 ) constant returns(address)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) SealedBids(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _AuctionRegistrarContract.Contract.SealedBids(&_AuctionRegistrarContract.CallOpts, arg0, arg1)
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids(address , bytes32 ) constant returns(address)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) SealedBids(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _AuctionRegistrarContract.Contract.SealedBids(&_AuctionRegistrarContract.CallOpts, arg0, arg1)
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(bytes32 hash, address owner, uint256 value, bytes32 salt) constant returns(bytes32 sealedBid)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) ShaBid(opts *bind.CallOpts, hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "shaBid", hash, owner, value, salt)
	return *ret0, err
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(bytes32 hash, address owner, uint256 value, bytes32 salt) constant returns(bytes32 sealedBid)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) ShaBid(hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	return _AuctionRegistrarContract.Contract.ShaBid(&_AuctionRegistrarContract.CallOpts, hash, owner, value, salt)
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(bytes32 hash, address owner, uint256 value, bytes32 salt) constant returns(bytes32 sealedBid)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) ShaBid(hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	return _AuctionRegistrarContract.Contract.ShaBid(&_AuctionRegistrarContract.CallOpts, hash, owner, value, salt)
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(bytes32 _hash) constant returns(uint8)
func (_AuctionRegistrarContract *AuctionRegistrarContractCaller) State(opts *bind.CallOpts, _hash [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AuctionRegistrarContract.contract.Call(opts, out, "state", _hash)
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(bytes32 _hash) constant returns(uint8)
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) State(_hash [32]byte) (uint8, error) {
	return _AuctionRegistrarContract.Contract.State(&_AuctionRegistrarContract.CallOpts, _hash)
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(bytes32 _hash) constant returns(uint8)
func (_AuctionRegistrarContract *AuctionRegistrarContractCallerSession) State(_hash [32]byte) (uint8, error) {
	return _AuctionRegistrarContract.Contract.State(&_AuctionRegistrarContract.CallOpts, _hash)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(bytes32 hash, address deed, uint256 registrationDate) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) AcceptRegistrarTransfer(opts *bind.TransactOpts, hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "acceptRegistrarTransfer", hash, deed, registrationDate)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(bytes32 hash, address deed, uint256 registrationDate) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) AcceptRegistrarTransfer(hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.AcceptRegistrarTransfer(&_AuctionRegistrarContract.TransactOpts, hash, deed, registrationDate)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(bytes32 hash, address deed, uint256 registrationDate) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) AcceptRegistrarTransfer(hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.AcceptRegistrarTransfer(&_AuctionRegistrarContract.TransactOpts, hash, deed, registrationDate)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(address bidder, bytes32 seal) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) CancelBid(opts *bind.TransactOpts, bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "cancelBid", bidder, seal)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(address bidder, bytes32 seal) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) CancelBid(bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.CancelBid(&_AuctionRegistrarContract.TransactOpts, bidder, seal)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(address bidder, bytes32 seal) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) CancelBid(bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.CancelBid(&_AuctionRegistrarContract.TransactOpts, bidder, seal)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(bytes32[] labels) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) EraseNode(opts *bind.TransactOpts, labels [][32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "eraseNode", labels)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(bytes32[] labels) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) EraseNode(labels [][32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.EraseNode(&_AuctionRegistrarContract.TransactOpts, labels)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(bytes32[] labels) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) EraseNode(labels [][32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.EraseNode(&_AuctionRegistrarContract.TransactOpts, labels)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) FinalizeAuction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "finalizeAuction", _hash)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) FinalizeAuction(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.FinalizeAuction(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) FinalizeAuction(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.FinalizeAuction(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(string unhashedName) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) InvalidateName(opts *bind.TransactOpts, unhashedName string) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "invalidateName", unhashedName)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(string unhashedName) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) InvalidateName(unhashedName string) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.InvalidateName(&_AuctionRegistrarContract.TransactOpts, unhashedName)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(string unhashedName) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) InvalidateName(unhashedName string) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.InvalidateName(&_AuctionRegistrarContract.TransactOpts, unhashedName)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(bytes32 sealedBid) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) NewBid(opts *bind.TransactOpts, sealedBid [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "newBid", sealedBid)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(bytes32 sealedBid) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) NewBid(sealedBid [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.NewBid(&_AuctionRegistrarContract.TransactOpts, sealedBid)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(bytes32 sealedBid) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) NewBid(sealedBid [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.NewBid(&_AuctionRegistrarContract.TransactOpts, sealedBid)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) ReleaseDeed(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "releaseDeed", _hash)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) ReleaseDeed(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.ReleaseDeed(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) ReleaseDeed(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.ReleaseDeed(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) StartAuction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "startAuction", _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) StartAuction(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.StartAuction(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) StartAuction(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.StartAuction(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(bytes32[] _hashes) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) StartAuctions(opts *bind.TransactOpts, _hashes [][32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "startAuctions", _hashes)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(bytes32[] _hashes) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) StartAuctions(_hashes [][32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.StartAuctions(&_AuctionRegistrarContract.TransactOpts, _hashes)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(bytes32[] _hashes) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) StartAuctions(_hashes [][32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.StartAuctions(&_AuctionRegistrarContract.TransactOpts, _hashes)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(bytes32[] hashes, bytes32 sealedBid) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) StartAuctionsAndBid(opts *bind.TransactOpts, hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "startAuctionsAndBid", hashes, sealedBid)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(bytes32[] hashes, bytes32 sealedBid) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) StartAuctionsAndBid(hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.StartAuctionsAndBid(&_AuctionRegistrarContract.TransactOpts, hashes, sealedBid)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(bytes32[] hashes, bytes32 sealedBid) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) StartAuctionsAndBid(hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.StartAuctionsAndBid(&_AuctionRegistrarContract.TransactOpts, hashes, sealedBid)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 _hash, address newOwner) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) Transfer(opts *bind.TransactOpts, _hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "transfer", _hash, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 _hash, address newOwner) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) Transfer(_hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.Transfer(&_AuctionRegistrarContract.TransactOpts, _hash, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 _hash, address newOwner) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) Transfer(_hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.Transfer(&_AuctionRegistrarContract.TransactOpts, _hash, newOwner)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) TransferRegistrars(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "transferRegistrars", _hash)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) TransferRegistrars(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.TransferRegistrars(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(bytes32 _hash) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) TransferRegistrars(_hash [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.TransferRegistrars(&_AuctionRegistrarContract.TransactOpts, _hash)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(bytes32 _hash, uint256 _value, bytes32 _salt) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactor) UnsealBid(opts *bind.TransactOpts, _hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.contract.Transact(opts, "unsealBid", _hash, _value, _salt)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(bytes32 _hash, uint256 _value, bytes32 _salt) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractSession) UnsealBid(_hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.UnsealBid(&_AuctionRegistrarContract.TransactOpts, _hash, _value, _salt)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(bytes32 _hash, uint256 _value, bytes32 _salt) returns()
func (_AuctionRegistrarContract *AuctionRegistrarContractTransactorSession) UnsealBid(_hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _AuctionRegistrarContract.Contract.UnsealBid(&_AuctionRegistrarContract.TransactOpts, _hash, _value, _salt)
}

// AuctionRegistrarContractAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractAuctionStartedIterator struct {
	Event *AuctionRegistrarContractAuctionStarted // Event containing the contract specifics and raw log

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
func (it *AuctionRegistrarContractAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRegistrarContractAuctionStarted)
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
		it.Event = new(AuctionRegistrarContractAuctionStarted)
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
func (it *AuctionRegistrarContractAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRegistrarContractAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRegistrarContractAuctionStarted represents a AuctionStarted event raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractAuctionStarted struct {
	Hash             [32]byte
	RegistrationDate *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x87e97e825a1d1fa0c54e1d36c7506c1dea8b1efd451fe68b000cf96f7cf40003.
//
// Solidity: event AuctionStarted(bytes32 indexed hash, uint256 registrationDate)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) FilterAuctionStarted(opts *bind.FilterOpts, hash [][32]byte) (*AuctionRegistrarContractAuctionStartedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.FilterLogs(opts, "AuctionStarted", hashRule)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractAuctionStartedIterator{contract: _AuctionRegistrarContract.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x87e97e825a1d1fa0c54e1d36c7506c1dea8b1efd451fe68b000cf96f7cf40003.
//
// Solidity: event AuctionStarted(bytes32 indexed hash, uint256 registrationDate)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *AuctionRegistrarContractAuctionStarted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.WatchLogs(opts, "AuctionStarted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRegistrarContractAuctionStarted)
				if err := _AuctionRegistrarContract.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// AuctionRegistrarContractBidRevealedIterator is returned from FilterBidRevealed and is used to iterate over the raw logs and unpacked data for BidRevealed events raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractBidRevealedIterator struct {
	Event *AuctionRegistrarContractBidRevealed // Event containing the contract specifics and raw log

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
func (it *AuctionRegistrarContractBidRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRegistrarContractBidRevealed)
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
		it.Event = new(AuctionRegistrarContractBidRevealed)
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
func (it *AuctionRegistrarContractBidRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRegistrarContractBidRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRegistrarContractBidRevealed represents a BidRevealed event raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractBidRevealed struct {
	Hash   [32]byte
	Owner  common.Address
	Value  *big.Int
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidRevealed is a free log retrieval operation binding the contract event 0x7b6c4b278d165a6b33958f8ea5dfb00c8c9d4d0acf1985bef5d10786898bc3e7.
//
// Solidity: event BidRevealed(bytes32 indexed hash, address indexed owner, uint256 value, uint8 status)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) FilterBidRevealed(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*AuctionRegistrarContractBidRevealedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.FilterLogs(opts, "BidRevealed", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractBidRevealedIterator{contract: _AuctionRegistrarContract.contract, event: "BidRevealed", logs: logs, sub: sub}, nil
}

// WatchBidRevealed is a free log subscription operation binding the contract event 0x7b6c4b278d165a6b33958f8ea5dfb00c8c9d4d0acf1985bef5d10786898bc3e7.
//
// Solidity: event BidRevealed(bytes32 indexed hash, address indexed owner, uint256 value, uint8 status)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) WatchBidRevealed(opts *bind.WatchOpts, sink chan<- *AuctionRegistrarContractBidRevealed, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.WatchLogs(opts, "BidRevealed", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRegistrarContractBidRevealed)
				if err := _AuctionRegistrarContract.contract.UnpackLog(event, "BidRevealed", log); err != nil {
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

// AuctionRegistrarContractHashInvalidatedIterator is returned from FilterHashInvalidated and is used to iterate over the raw logs and unpacked data for HashInvalidated events raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractHashInvalidatedIterator struct {
	Event *AuctionRegistrarContractHashInvalidated // Event containing the contract specifics and raw log

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
func (it *AuctionRegistrarContractHashInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRegistrarContractHashInvalidated)
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
		it.Event = new(AuctionRegistrarContractHashInvalidated)
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
func (it *AuctionRegistrarContractHashInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRegistrarContractHashInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRegistrarContractHashInvalidated represents a HashInvalidated event raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractHashInvalidated struct {
	Hash             [32]byte
	Name             common.Hash
	Value            *big.Int
	RegistrationDate *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHashInvalidated is a free log retrieval operation binding the contract event 0x1f9c649fe47e58bb60f4e52f0d90e4c47a526c9f90c5113df842c025970b66ad.
//
// Solidity: event HashInvalidated(bytes32 indexed hash, string indexed name, uint256 value, uint256 registrationDate)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) FilterHashInvalidated(opts *bind.FilterOpts, hash [][32]byte, name []string) (*AuctionRegistrarContractHashInvalidatedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.FilterLogs(opts, "HashInvalidated", hashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractHashInvalidatedIterator{contract: _AuctionRegistrarContract.contract, event: "HashInvalidated", logs: logs, sub: sub}, nil
}

// WatchHashInvalidated is a free log subscription operation binding the contract event 0x1f9c649fe47e58bb60f4e52f0d90e4c47a526c9f90c5113df842c025970b66ad.
//
// Solidity: event HashInvalidated(bytes32 indexed hash, string indexed name, uint256 value, uint256 registrationDate)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) WatchHashInvalidated(opts *bind.WatchOpts, sink chan<- *AuctionRegistrarContractHashInvalidated, hash [][32]byte, name []string) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.WatchLogs(opts, "HashInvalidated", hashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRegistrarContractHashInvalidated)
				if err := _AuctionRegistrarContract.contract.UnpackLog(event, "HashInvalidated", log); err != nil {
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

// AuctionRegistrarContractHashRegisteredIterator is returned from FilterHashRegistered and is used to iterate over the raw logs and unpacked data for HashRegistered events raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractHashRegisteredIterator struct {
	Event *AuctionRegistrarContractHashRegistered // Event containing the contract specifics and raw log

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
func (it *AuctionRegistrarContractHashRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRegistrarContractHashRegistered)
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
		it.Event = new(AuctionRegistrarContractHashRegistered)
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
func (it *AuctionRegistrarContractHashRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRegistrarContractHashRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRegistrarContractHashRegistered represents a HashRegistered event raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractHashRegistered struct {
	Hash             [32]byte
	Owner            common.Address
	Value            *big.Int
	RegistrationDate *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHashRegistered is a free log retrieval operation binding the contract event 0x0f0c27adfd84b60b6f456b0e87cdccb1e5fb9603991588d87fa99f5b6b61e670.
//
// Solidity: event HashRegistered(bytes32 indexed hash, address indexed owner, uint256 value, uint256 registrationDate)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) FilterHashRegistered(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*AuctionRegistrarContractHashRegisteredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.FilterLogs(opts, "HashRegistered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractHashRegisteredIterator{contract: _AuctionRegistrarContract.contract, event: "HashRegistered", logs: logs, sub: sub}, nil
}

// WatchHashRegistered is a free log subscription operation binding the contract event 0x0f0c27adfd84b60b6f456b0e87cdccb1e5fb9603991588d87fa99f5b6b61e670.
//
// Solidity: event HashRegistered(bytes32 indexed hash, address indexed owner, uint256 value, uint256 registrationDate)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) WatchHashRegistered(opts *bind.WatchOpts, sink chan<- *AuctionRegistrarContractHashRegistered, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.WatchLogs(opts, "HashRegistered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRegistrarContractHashRegistered)
				if err := _AuctionRegistrarContract.contract.UnpackLog(event, "HashRegistered", log); err != nil {
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

// AuctionRegistrarContractHashReleasedIterator is returned from FilterHashReleased and is used to iterate over the raw logs and unpacked data for HashReleased events raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractHashReleasedIterator struct {
	Event *AuctionRegistrarContractHashReleased // Event containing the contract specifics and raw log

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
func (it *AuctionRegistrarContractHashReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRegistrarContractHashReleased)
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
		it.Event = new(AuctionRegistrarContractHashReleased)
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
func (it *AuctionRegistrarContractHashReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRegistrarContractHashReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRegistrarContractHashReleased represents a HashReleased event raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractHashReleased struct {
	Hash  [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHashReleased is a free log retrieval operation binding the contract event 0x292b79b9246fa2c8e77d3fe195b251f9cb839d7d038e667c069ee7708c631e16.
//
// Solidity: event HashReleased(bytes32 indexed hash, uint256 value)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) FilterHashReleased(opts *bind.FilterOpts, hash [][32]byte) (*AuctionRegistrarContractHashReleasedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.FilterLogs(opts, "HashReleased", hashRule)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractHashReleasedIterator{contract: _AuctionRegistrarContract.contract, event: "HashReleased", logs: logs, sub: sub}, nil
}

// WatchHashReleased is a free log subscription operation binding the contract event 0x292b79b9246fa2c8e77d3fe195b251f9cb839d7d038e667c069ee7708c631e16.
//
// Solidity: event HashReleased(bytes32 indexed hash, uint256 value)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) WatchHashReleased(opts *bind.WatchOpts, sink chan<- *AuctionRegistrarContractHashReleased, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.WatchLogs(opts, "HashReleased", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRegistrarContractHashReleased)
				if err := _AuctionRegistrarContract.contract.UnpackLog(event, "HashReleased", log); err != nil {
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

// AuctionRegistrarContractNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractNewBidIterator struct {
	Event *AuctionRegistrarContractNewBid // Event containing the contract specifics and raw log

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
func (it *AuctionRegistrarContractNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRegistrarContractNewBid)
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
		it.Event = new(AuctionRegistrarContractNewBid)
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
func (it *AuctionRegistrarContractNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRegistrarContractNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRegistrarContractNewBid represents a NewBid event raised by the AuctionRegistrarContract contract.
type AuctionRegistrarContractNewBid struct {
	Hash    [32]byte
	Bidder  common.Address
	Deposit *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xb556ff269c1b6714f432c36431e2041d28436a73b6c3f19c021827bbdc6bfc29.
//
// Solidity: event NewBid(bytes32 indexed hash, address indexed bidder, uint256 deposit)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) FilterNewBid(opts *bind.FilterOpts, hash [][32]byte, bidder []common.Address) (*AuctionRegistrarContractNewBidIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.FilterLogs(opts, "NewBid", hashRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionRegistrarContractNewBidIterator{contract: _AuctionRegistrarContract.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xb556ff269c1b6714f432c36431e2041d28436a73b6c3f19c021827bbdc6bfc29.
//
// Solidity: event NewBid(bytes32 indexed hash, address indexed bidder, uint256 deposit)
func (_AuctionRegistrarContract *AuctionRegistrarContractFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *AuctionRegistrarContractNewBid, hash [][32]byte, bidder []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _AuctionRegistrarContract.contract.WatchLogs(opts, "NewBid", hashRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRegistrarContractNewBid)
				if err := _AuctionRegistrarContract.contract.UnpackLog(event, "NewBid", log); err != nil {
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
