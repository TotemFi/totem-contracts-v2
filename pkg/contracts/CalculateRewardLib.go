// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// CalculateRewardLibMetaData contains all meta data concerning the CalculateRewardLib contract.
var CalculateRewardLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"foo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CalculateRewardLibABI is the input ABI used to generate the binding from.
// Deprecated: Use CalculateRewardLibMetaData.ABI instead.
var CalculateRewardLibABI = CalculateRewardLibMetaData.ABI

// CalculateRewardLib is an auto generated Go binding around an Ethereum contract.
type CalculateRewardLib struct {
	CalculateRewardLibCaller     // Read-only binding to the contract
	CalculateRewardLibTransactor // Write-only binding to the contract
	CalculateRewardLibFilterer   // Log filterer for contract events
}

// CalculateRewardLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalculateRewardLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculateRewardLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalculateRewardLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculateRewardLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalculateRewardLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculateRewardLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalculateRewardLibSession struct {
	Contract     *CalculateRewardLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CalculateRewardLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalculateRewardLibCallerSession struct {
	Contract *CalculateRewardLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CalculateRewardLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalculateRewardLibTransactorSession struct {
	Contract     *CalculateRewardLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CalculateRewardLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalculateRewardLibRaw struct {
	Contract *CalculateRewardLib // Generic contract binding to access the raw methods on
}

// CalculateRewardLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalculateRewardLibCallerRaw struct {
	Contract *CalculateRewardLibCaller // Generic read-only contract binding to access the raw methods on
}

// CalculateRewardLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalculateRewardLibTransactorRaw struct {
	Contract *CalculateRewardLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalculateRewardLib creates a new instance of CalculateRewardLib, bound to a specific deployed contract.
func NewCalculateRewardLib(address common.Address, backend bind.ContractBackend) (*CalculateRewardLib, error) {
	contract, err := bindCalculateRewardLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CalculateRewardLib{CalculateRewardLibCaller: CalculateRewardLibCaller{contract: contract}, CalculateRewardLibTransactor: CalculateRewardLibTransactor{contract: contract}, CalculateRewardLibFilterer: CalculateRewardLibFilterer{contract: contract}}, nil
}

// NewCalculateRewardLibCaller creates a new read-only instance of CalculateRewardLib, bound to a specific deployed contract.
func NewCalculateRewardLibCaller(address common.Address, caller bind.ContractCaller) (*CalculateRewardLibCaller, error) {
	contract, err := bindCalculateRewardLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalculateRewardLibCaller{contract: contract}, nil
}

// NewCalculateRewardLibTransactor creates a new write-only instance of CalculateRewardLib, bound to a specific deployed contract.
func NewCalculateRewardLibTransactor(address common.Address, transactor bind.ContractTransactor) (*CalculateRewardLibTransactor, error) {
	contract, err := bindCalculateRewardLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalculateRewardLibTransactor{contract: contract}, nil
}

// NewCalculateRewardLibFilterer creates a new log filterer instance of CalculateRewardLib, bound to a specific deployed contract.
func NewCalculateRewardLibFilterer(address common.Address, filterer bind.ContractFilterer) (*CalculateRewardLibFilterer, error) {
	contract, err := bindCalculateRewardLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalculateRewardLibFilterer{contract: contract}, nil
}

// bindCalculateRewardLib binds a generic wrapper to an already deployed contract.
func bindCalculateRewardLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalculateRewardLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CalculateRewardLib *CalculateRewardLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CalculateRewardLib.Contract.CalculateRewardLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CalculateRewardLib *CalculateRewardLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CalculateRewardLib.Contract.CalculateRewardLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CalculateRewardLib *CalculateRewardLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CalculateRewardLib.Contract.CalculateRewardLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CalculateRewardLib *CalculateRewardLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CalculateRewardLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CalculateRewardLib *CalculateRewardLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CalculateRewardLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CalculateRewardLib *CalculateRewardLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CalculateRewardLib.Contract.contract.Transact(opts, method, params...)
}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() view returns(uint256)
func (_CalculateRewardLib *CalculateRewardLibCaller) Foo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CalculateRewardLib.contract.Call(opts, &out, "foo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() view returns(uint256)
func (_CalculateRewardLib *CalculateRewardLibSession) Foo() (*big.Int, error) {
	return _CalculateRewardLib.Contract.Foo(&_CalculateRewardLib.CallOpts)
}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() view returns(uint256)
func (_CalculateRewardLib *CalculateRewardLibCallerSession) Foo() (*big.Int, error) {
	return _CalculateRewardLib.Contract.Foo(&_CalculateRewardLib.CallOpts)
}
