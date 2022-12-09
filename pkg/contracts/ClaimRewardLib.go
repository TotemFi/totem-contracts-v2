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

// ClaimRewardLibMetaData contains all meta data concerning the ClaimRewardLib contract.
var ClaimRewardLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"foo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ClaimRewardLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimRewardLibMetaData.ABI instead.
var ClaimRewardLibABI = ClaimRewardLibMetaData.ABI

// ClaimRewardLib is an auto generated Go binding around an Ethereum contract.
type ClaimRewardLib struct {
	ClaimRewardLibCaller     // Read-only binding to the contract
	ClaimRewardLibTransactor // Write-only binding to the contract
	ClaimRewardLibFilterer   // Log filterer for contract events
}

// ClaimRewardLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimRewardLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimRewardLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimRewardLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimRewardLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimRewardLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimRewardLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimRewardLibSession struct {
	Contract     *ClaimRewardLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimRewardLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimRewardLibCallerSession struct {
	Contract *ClaimRewardLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ClaimRewardLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimRewardLibTransactorSession struct {
	Contract     *ClaimRewardLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ClaimRewardLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimRewardLibRaw struct {
	Contract *ClaimRewardLib // Generic contract binding to access the raw methods on
}

// ClaimRewardLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimRewardLibCallerRaw struct {
	Contract *ClaimRewardLibCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimRewardLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimRewardLibTransactorRaw struct {
	Contract *ClaimRewardLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimRewardLib creates a new instance of ClaimRewardLib, bound to a specific deployed contract.
func NewClaimRewardLib(address common.Address, backend bind.ContractBackend) (*ClaimRewardLib, error) {
	contract, err := bindClaimRewardLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimRewardLib{ClaimRewardLibCaller: ClaimRewardLibCaller{contract: contract}, ClaimRewardLibTransactor: ClaimRewardLibTransactor{contract: contract}, ClaimRewardLibFilterer: ClaimRewardLibFilterer{contract: contract}}, nil
}

// NewClaimRewardLibCaller creates a new read-only instance of ClaimRewardLib, bound to a specific deployed contract.
func NewClaimRewardLibCaller(address common.Address, caller bind.ContractCaller) (*ClaimRewardLibCaller, error) {
	contract, err := bindClaimRewardLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimRewardLibCaller{contract: contract}, nil
}

// NewClaimRewardLibTransactor creates a new write-only instance of ClaimRewardLib, bound to a specific deployed contract.
func NewClaimRewardLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimRewardLibTransactor, error) {
	contract, err := bindClaimRewardLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimRewardLibTransactor{contract: contract}, nil
}

// NewClaimRewardLibFilterer creates a new log filterer instance of ClaimRewardLib, bound to a specific deployed contract.
func NewClaimRewardLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimRewardLibFilterer, error) {
	contract, err := bindClaimRewardLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimRewardLibFilterer{contract: contract}, nil
}

// bindClaimRewardLib binds a generic wrapper to an already deployed contract.
func bindClaimRewardLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimRewardLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimRewardLib *ClaimRewardLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimRewardLib.Contract.ClaimRewardLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimRewardLib *ClaimRewardLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimRewardLib.Contract.ClaimRewardLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimRewardLib *ClaimRewardLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimRewardLib.Contract.ClaimRewardLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimRewardLib *ClaimRewardLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimRewardLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimRewardLib *ClaimRewardLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimRewardLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimRewardLib *ClaimRewardLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimRewardLib.Contract.contract.Transact(opts, method, params...)
}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() view returns(uint256)
func (_ClaimRewardLib *ClaimRewardLibCaller) Foo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimRewardLib.contract.Call(opts, &out, "foo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() view returns(uint256)
func (_ClaimRewardLib *ClaimRewardLibSession) Foo() (*big.Int, error) {
	return _ClaimRewardLib.Contract.Foo(&_ClaimRewardLib.CallOpts)
}

// Foo is a free data retrieval call binding the contract method 0xc2985578.
//
// Solidity: function foo() view returns(uint256)
func (_ClaimRewardLib *ClaimRewardLibCallerSession) Foo() (*big.Int, error) {
	return _ClaimRewardLib.Contract.Foo(&_ClaimRewardLib.CallOpts)
}
