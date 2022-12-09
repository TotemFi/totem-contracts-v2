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

// PrivateSaleVestingMetaData contains all meta data concerning the PrivateSaleVesting contract.
var PrivateSaleVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_UNLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELEASE_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"addRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStartTimeSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releasePeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartTime\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unallocatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"vested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountVested\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PrivateSaleVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use PrivateSaleVestingMetaData.ABI instead.
var PrivateSaleVestingABI = PrivateSaleVestingMetaData.ABI

// PrivateSaleVesting is an auto generated Go binding around an Ethereum contract.
type PrivateSaleVesting struct {
	PrivateSaleVestingCaller     // Read-only binding to the contract
	PrivateSaleVestingTransactor // Write-only binding to the contract
	PrivateSaleVestingFilterer   // Log filterer for contract events
}

// PrivateSaleVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivateSaleVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateSaleVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivateSaleVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateSaleVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivateSaleVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateSaleVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivateSaleVestingSession struct {
	Contract     *PrivateSaleVesting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PrivateSaleVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivateSaleVestingCallerSession struct {
	Contract *PrivateSaleVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PrivateSaleVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivateSaleVestingTransactorSession struct {
	Contract     *PrivateSaleVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PrivateSaleVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivateSaleVestingRaw struct {
	Contract *PrivateSaleVesting // Generic contract binding to access the raw methods on
}

// PrivateSaleVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivateSaleVestingCallerRaw struct {
	Contract *PrivateSaleVestingCaller // Generic read-only contract binding to access the raw methods on
}

// PrivateSaleVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivateSaleVestingTransactorRaw struct {
	Contract *PrivateSaleVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivateSaleVesting creates a new instance of PrivateSaleVesting, bound to a specific deployed contract.
func NewPrivateSaleVesting(address common.Address, backend bind.ContractBackend) (*PrivateSaleVesting, error) {
	contract, err := bindPrivateSaleVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVesting{PrivateSaleVestingCaller: PrivateSaleVestingCaller{contract: contract}, PrivateSaleVestingTransactor: PrivateSaleVestingTransactor{contract: contract}, PrivateSaleVestingFilterer: PrivateSaleVestingFilterer{contract: contract}}, nil
}

// NewPrivateSaleVestingCaller creates a new read-only instance of PrivateSaleVesting, bound to a specific deployed contract.
func NewPrivateSaleVestingCaller(address common.Address, caller bind.ContractCaller) (*PrivateSaleVestingCaller, error) {
	contract, err := bindPrivateSaleVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVestingCaller{contract: contract}, nil
}

// NewPrivateSaleVestingTransactor creates a new write-only instance of PrivateSaleVesting, bound to a specific deployed contract.
func NewPrivateSaleVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivateSaleVestingTransactor, error) {
	contract, err := bindPrivateSaleVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVestingTransactor{contract: contract}, nil
}

// NewPrivateSaleVestingFilterer creates a new log filterer instance of PrivateSaleVesting, bound to a specific deployed contract.
func NewPrivateSaleVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivateSaleVestingFilterer, error) {
	contract, err := bindPrivateSaleVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVestingFilterer{contract: contract}, nil
}

// bindPrivateSaleVesting binds a generic wrapper to an already deployed contract.
func bindPrivateSaleVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateSaleVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateSaleVesting *PrivateSaleVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrivateSaleVesting.Contract.PrivateSaleVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateSaleVesting *PrivateSaleVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.PrivateSaleVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateSaleVesting *PrivateSaleVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.PrivateSaleVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateSaleVesting *PrivateSaleVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrivateSaleVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateSaleVesting *PrivateSaleVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateSaleVesting *PrivateSaleVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.contract.Transact(opts, method, params...)
}

// INITIALUNLOCK is a free data retrieval call binding the contract method 0x4b3a5fa8.
//
// Solidity: function INITIAL_UNLOCK() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) INITIALUNLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "INITIAL_UNLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALUNLOCK is a free data retrieval call binding the contract method 0x4b3a5fa8.
//
// Solidity: function INITIAL_UNLOCK() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) INITIALUNLOCK() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.INITIALUNLOCK(&_PrivateSaleVesting.CallOpts)
}

// INITIALUNLOCK is a free data retrieval call binding the contract method 0x4b3a5fa8.
//
// Solidity: function INITIAL_UNLOCK() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) INITIALUNLOCK() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.INITIALUNLOCK(&_PrivateSaleVesting.CallOpts)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) LOCKPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "LOCK_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) LOCKPERIODS() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.LOCKPERIODS(&_PrivateSaleVesting.CallOpts)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) LOCKPERIODS() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.LOCKPERIODS(&_PrivateSaleVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) RELEASEPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "RELEASE_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) RELEASEPERIODS() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.RELEASEPERIODS(&_PrivateSaleVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) RELEASEPERIODS() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.RELEASEPERIODS(&_PrivateSaleVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) TOTALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "TOTAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) TOTALAMOUNT() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.TOTALAMOUNT(&_PrivateSaleVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) TOTALAMOUNT() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.TOTALAMOUNT(&_PrivateSaleVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) WITHDRAWINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "WITHDRAW_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.WITHDRAWINTERVAL(&_PrivateSaleVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.WITHDRAWINTERVAL(&_PrivateSaleVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) IsStartTimeSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "isStartTimeSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_PrivateSaleVesting *PrivateSaleVestingSession) IsStartTimeSet() (bool, error) {
	return _PrivateSaleVesting.Contract.IsStartTimeSet(&_PrivateSaleVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) IsStartTimeSet() (bool, error) {
	return _PrivateSaleVesting.Contract.IsStartTimeSet(&_PrivateSaleVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) LockPeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "lockPeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) LockPeriods() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.LockPeriods(&_PrivateSaleVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) LockPeriods() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.LockPeriods(&_PrivateSaleVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrivateSaleVesting *PrivateSaleVestingSession) Owner() (common.Address, error) {
	return _PrivateSaleVesting.Contract.Owner(&_PrivateSaleVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) Owner() (common.Address, error) {
	return _PrivateSaleVesting.Contract.Owner(&_PrivateSaleVesting.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "recipients", arg0)

	outstruct := new(struct {
		TotalAmount     *big.Int
		AmountWithdrawn *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountWithdrawn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_PrivateSaleVesting *PrivateSaleVestingSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _PrivateSaleVesting.Contract.Recipients(&_PrivateSaleVesting.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _PrivateSaleVesting.Contract.Recipients(&_PrivateSaleVesting.CallOpts, arg0)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) ReleasePeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "releasePeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) ReleasePeriods() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.ReleasePeriods(&_PrivateSaleVesting.CallOpts)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) ReleasePeriods() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.ReleasePeriods(&_PrivateSaleVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) StartTime() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.StartTime(&_PrivateSaleVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) StartTime() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.StartTime(&_PrivateSaleVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) TotalAmount() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.TotalAmount(&_PrivateSaleVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) TotalAmount() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.TotalAmount(&_PrivateSaleVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_PrivateSaleVesting *PrivateSaleVestingSession) TotemToken() (common.Address, error) {
	return _PrivateSaleVesting.Contract.TotemToken(&_PrivateSaleVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) TotemToken() (common.Address, error) {
	return _PrivateSaleVesting.Contract.TotemToken(&_PrivateSaleVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) UnallocatedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "unallocatedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) UnallocatedAmount() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.UnallocatedAmount(&_PrivateSaleVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) UnallocatedAmount() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.UnallocatedAmount(&_PrivateSaleVesting.CallOpts)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) Vested(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "vested", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_PrivateSaleVesting *PrivateSaleVestingSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _PrivateSaleVesting.Contract.Vested(&_PrivateSaleVesting.CallOpts, beneficiary)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _PrivateSaleVesting.Contract.Vested(&_PrivateSaleVesting.CallOpts, beneficiary)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) WithdrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "withdrawInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingSession) WithdrawInterval() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.WithdrawInterval(&_PrivateSaleVesting.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) WithdrawInterval() (*big.Int, error) {
	return _PrivateSaleVesting.Contract.WithdrawInterval(&_PrivateSaleVesting.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_PrivateSaleVesting *PrivateSaleVestingCaller) Withdrawable(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrivateSaleVesting.contract.Call(opts, &out, "withdrawable", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_PrivateSaleVesting *PrivateSaleVestingSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _PrivateSaleVesting.Contract.Withdrawable(&_PrivateSaleVesting.CallOpts, beneficiary)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_PrivateSaleVesting *PrivateSaleVestingCallerSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _PrivateSaleVesting.Contract.Withdrawable(&_PrivateSaleVesting.CallOpts, beneficiary)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactor) AddRecipient(opts *bind.TransactOpts, _newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _PrivateSaleVesting.contract.Transact(opts, "addRecipient", _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_PrivateSaleVesting *PrivateSaleVestingSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.AddRecipient(&_PrivateSaleVesting.TransactOpts, _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactorSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.AddRecipient(&_PrivateSaleVesting.TransactOpts, _newRecipient, _totalAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateSaleVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivateSaleVesting *PrivateSaleVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.RenounceOwnership(&_PrivateSaleVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.RenounceOwnership(&_PrivateSaleVesting.TransactOpts)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactor) SetStartTime(opts *bind.TransactOpts, _newStartTime *big.Int) (*types.Transaction, error) {
	return _PrivateSaleVesting.contract.Transact(opts, "setStartTime", _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_PrivateSaleVesting *PrivateSaleVestingSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.SetStartTime(&_PrivateSaleVesting.TransactOpts, _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactorSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.SetStartTime(&_PrivateSaleVesting.TransactOpts, _newStartTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PrivateSaleVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrivateSaleVesting *PrivateSaleVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.TransferOwnership(&_PrivateSaleVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.TransferOwnership(&_PrivateSaleVesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateSaleVesting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PrivateSaleVesting *PrivateSaleVestingSession) Withdraw() (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.Withdraw(&_PrivateSaleVesting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PrivateSaleVesting *PrivateSaleVestingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PrivateSaleVesting.Contract.Withdraw(&_PrivateSaleVesting.TransactOpts)
}

// PrivateSaleVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PrivateSaleVesting contract.
type PrivateSaleVestingOwnershipTransferredIterator struct {
	Event *PrivateSaleVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PrivateSaleVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSaleVestingOwnershipTransferred)
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
		it.Event = new(PrivateSaleVestingOwnershipTransferred)
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
func (it *PrivateSaleVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSaleVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSaleVestingOwnershipTransferred represents a OwnershipTransferred event raised by the PrivateSaleVesting contract.
type PrivateSaleVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PrivateSaleVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrivateSaleVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVestingOwnershipTransferredIterator{contract: _PrivateSaleVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PrivateSaleVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrivateSaleVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSaleVestingOwnershipTransferred)
				if err := _PrivateSaleVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) ParseOwnershipTransferred(log types.Log) (*PrivateSaleVestingOwnershipTransferred, error) {
	event := new(PrivateSaleVestingOwnershipTransferred)
	if err := _PrivateSaleVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateSaleVestingStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the PrivateSaleVesting contract.
type PrivateSaleVestingStartTimeSetIterator struct {
	Event *PrivateSaleVestingStartTimeSet // Event containing the contract specifics and raw log

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
func (it *PrivateSaleVestingStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSaleVestingStartTimeSet)
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
		it.Event = new(PrivateSaleVestingStartTimeSet)
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
func (it *PrivateSaleVestingStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSaleVestingStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSaleVestingStartTimeSet represents a StartTimeSet event raised by the PrivateSaleVesting contract.
type PrivateSaleVestingStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*PrivateSaleVestingStartTimeSetIterator, error) {

	logs, sub, err := _PrivateSaleVesting.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVestingStartTimeSetIterator{contract: _PrivateSaleVesting.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *PrivateSaleVestingStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _PrivateSaleVesting.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSaleVestingStartTimeSet)
				if err := _PrivateSaleVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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

// ParseStartTimeSet is a log parse operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) ParseStartTimeSet(log types.Log) (*PrivateSaleVestingStartTimeSet, error) {
	event := new(PrivateSaleVestingStartTimeSet)
	if err := _PrivateSaleVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateSaleVestingVestingScheduleRegisteredIterator is returned from FilterVestingScheduleRegistered and is used to iterate over the raw logs and unpacked data for VestingScheduleRegistered events raised by the PrivateSaleVesting contract.
type PrivateSaleVestingVestingScheduleRegisteredIterator struct {
	Event *PrivateSaleVestingVestingScheduleRegistered // Event containing the contract specifics and raw log

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
func (it *PrivateSaleVestingVestingScheduleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSaleVestingVestingScheduleRegistered)
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
		it.Event = new(PrivateSaleVestingVestingScheduleRegistered)
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
func (it *PrivateSaleVestingVestingScheduleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSaleVestingVestingScheduleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSaleVestingVestingScheduleRegistered represents a VestingScheduleRegistered event raised by the PrivateSaleVesting contract.
type PrivateSaleVestingVestingScheduleRegistered struct {
	RegisteredAddress common.Address
	TotalAmount       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleRegistered is a free log retrieval operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) FilterVestingScheduleRegistered(opts *bind.FilterOpts) (*PrivateSaleVestingVestingScheduleRegisteredIterator, error) {

	logs, sub, err := _PrivateSaleVesting.contract.FilterLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVestingVestingScheduleRegisteredIterator{contract: _PrivateSaleVesting.contract, event: "VestingScheduleRegistered", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleRegistered is a free log subscription operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) WatchVestingScheduleRegistered(opts *bind.WatchOpts, sink chan<- *PrivateSaleVestingVestingScheduleRegistered) (event.Subscription, error) {

	logs, sub, err := _PrivateSaleVesting.contract.WatchLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSaleVestingVestingScheduleRegistered)
				if err := _PrivateSaleVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
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

// ParseVestingScheduleRegistered is a log parse operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) ParseVestingScheduleRegistered(log types.Log) (*PrivateSaleVestingVestingScheduleRegistered, error) {
	event := new(PrivateSaleVestingVestingScheduleRegistered)
	if err := _PrivateSaleVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateSaleVestingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PrivateSaleVesting contract.
type PrivateSaleVestingWithdrawIterator struct {
	Event *PrivateSaleVestingWithdraw // Event containing the contract specifics and raw log

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
func (it *PrivateSaleVestingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateSaleVestingWithdraw)
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
		it.Event = new(PrivateSaleVestingWithdraw)
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
func (it *PrivateSaleVestingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateSaleVestingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateSaleVestingWithdraw represents a Withdraw event raised by the PrivateSaleVesting contract.
type PrivateSaleVestingWithdraw struct {
	RegisteredAddress common.Address
	AmountWithdrawn   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*PrivateSaleVestingWithdrawIterator, error) {

	logs, sub, err := _PrivateSaleVesting.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &PrivateSaleVestingWithdrawIterator{contract: _PrivateSaleVesting.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PrivateSaleVestingWithdraw) (event.Subscription, error) {

	logs, sub, err := _PrivateSaleVesting.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateSaleVestingWithdraw)
				if err := _PrivateSaleVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_PrivateSaleVesting *PrivateSaleVestingFilterer) ParseWithdraw(log types.Log) (*PrivateSaleVestingWithdraw, error) {
	event := new(PrivateSaleVestingWithdraw)
	if err := _PrivateSaleVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
