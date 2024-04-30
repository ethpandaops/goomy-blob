// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package largetx

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

// LooperMetaData contains all meta data concerning the Looper contract.
var LooperMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"loop_it\",\"inputs\":[{\"name\":\"no_of_times_to_loop\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"}]",
}

// LooperABI is the input ABI used to generate the binding from.
// Deprecated: Use LooperMetaData.ABI instead.
var LooperABI = LooperMetaData.ABI

// Looper is an auto generated Go binding around an Ethereum contract.
type Looper struct {
	LooperCaller     // Read-only binding to the contract
	LooperTransactor // Write-only binding to the contract
	LooperFilterer   // Log filterer for contract events
}

// LooperCaller is an auto generated read-only Go binding around an Ethereum contract.
type LooperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LooperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LooperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LooperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LooperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LooperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LooperSession struct {
	Contract     *Looper           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LooperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LooperCallerSession struct {
	Contract *LooperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LooperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LooperTransactorSession struct {
	Contract     *LooperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LooperRaw is an auto generated low-level Go binding around an Ethereum contract.
type LooperRaw struct {
	Contract *Looper // Generic contract binding to access the raw methods on
}

// LooperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LooperCallerRaw struct {
	Contract *LooperCaller // Generic read-only contract binding to access the raw methods on
}

// LooperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LooperTransactorRaw struct {
	Contract *LooperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLooper creates a new instance of Looper, bound to a specific deployed contract.
func NewLooper(address common.Address, backend bind.ContractBackend) (*Looper, error) {
	contract, err := bindLooper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Looper{LooperCaller: LooperCaller{contract: contract}, LooperTransactor: LooperTransactor{contract: contract}, LooperFilterer: LooperFilterer{contract: contract}}, nil
}

// NewLooperCaller creates a new read-only instance of Looper, bound to a specific deployed contract.
func NewLooperCaller(address common.Address, caller bind.ContractCaller) (*LooperCaller, error) {
	contract, err := bindLooper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LooperCaller{contract: contract}, nil
}

// NewLooperTransactor creates a new write-only instance of Looper, bound to a specific deployed contract.
func NewLooperTransactor(address common.Address, transactor bind.ContractTransactor) (*LooperTransactor, error) {
	contract, err := bindLooper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LooperTransactor{contract: contract}, nil
}

// NewLooperFilterer creates a new log filterer instance of Looper, bound to a specific deployed contract.
func NewLooperFilterer(address common.Address, filterer bind.ContractFilterer) (*LooperFilterer, error) {
	contract, err := bindLooper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LooperFilterer{contract: contract}, nil
}

// bindLooper binds a generic wrapper to an already deployed contract.
func bindLooper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LooperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Looper *LooperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Looper.Contract.LooperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Looper *LooperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Looper.Contract.LooperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Looper *LooperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Looper.Contract.LooperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Looper *LooperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Looper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Looper *LooperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Looper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Looper *LooperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Looper.Contract.contract.Transact(opts, method, params...)
}

// LoopIt is a paid mutator transaction binding the contract method 0x8f491c65.
//
// Solidity: function loop_it(uint64 no_of_times_to_loop) returns(uint256[])
func (_Looper *LooperTransactor) LoopIt(opts *bind.TransactOpts, no_of_times_to_loop uint64) (*types.Transaction, error) {
	return _Looper.contract.Transact(opts, "loop_it", no_of_times_to_loop)
}

// LoopIt is a paid mutator transaction binding the contract method 0x8f491c65.
//
// Solidity: function loop_it(uint64 no_of_times_to_loop) returns(uint256[])
func (_Looper *LooperSession) LoopIt(no_of_times_to_loop uint64) (*types.Transaction, error) {
	return _Looper.Contract.LoopIt(&_Looper.TransactOpts, no_of_times_to_loop)
}

// LoopIt is a paid mutator transaction binding the contract method 0x8f491c65.
//
// Solidity: function loop_it(uint64 no_of_times_to_loop) returns(uint256[])
func (_Looper *LooperTransactorSession) LoopIt(no_of_times_to_loop uint64) (*types.Transaction, error) {
	return _Looper.Contract.LoopIt(&_Looper.TransactOpts, no_of_times_to_loop)
}
