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

// StrategicVestingMetaData contains all meta data concerning the StrategicVesting contract.
var StrategicVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LOCK_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELEASE_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"addRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStartTimeSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releasePeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartTime\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unallocatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"vested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountVested\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StrategicVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategicVestingMetaData.ABI instead.
var StrategicVestingABI = StrategicVestingMetaData.ABI

// StrategicVesting is an auto generated Go binding around an Ethereum contract.
type StrategicVesting struct {
	StrategicVestingCaller     // Read-only binding to the contract
	StrategicVestingTransactor // Write-only binding to the contract
	StrategicVestingFilterer   // Log filterer for contract events
}

// StrategicVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategicVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategicVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategicVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategicVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategicVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategicVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategicVestingSession struct {
	Contract     *StrategicVesting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategicVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategicVestingCallerSession struct {
	Contract *StrategicVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StrategicVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategicVestingTransactorSession struct {
	Contract     *StrategicVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StrategicVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategicVestingRaw struct {
	Contract *StrategicVesting // Generic contract binding to access the raw methods on
}

// StrategicVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategicVestingCallerRaw struct {
	Contract *StrategicVestingCaller // Generic read-only contract binding to access the raw methods on
}

// StrategicVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategicVestingTransactorRaw struct {
	Contract *StrategicVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategicVesting creates a new instance of StrategicVesting, bound to a specific deployed contract.
func NewStrategicVesting(address common.Address, backend bind.ContractBackend) (*StrategicVesting, error) {
	contract, err := bindStrategicVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategicVesting{StrategicVestingCaller: StrategicVestingCaller{contract: contract}, StrategicVestingTransactor: StrategicVestingTransactor{contract: contract}, StrategicVestingFilterer: StrategicVestingFilterer{contract: contract}}, nil
}

// NewStrategicVestingCaller creates a new read-only instance of StrategicVesting, bound to a specific deployed contract.
func NewStrategicVestingCaller(address common.Address, caller bind.ContractCaller) (*StrategicVestingCaller, error) {
	contract, err := bindStrategicVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategicVestingCaller{contract: contract}, nil
}

// NewStrategicVestingTransactor creates a new write-only instance of StrategicVesting, bound to a specific deployed contract.
func NewStrategicVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategicVestingTransactor, error) {
	contract, err := bindStrategicVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategicVestingTransactor{contract: contract}, nil
}

// NewStrategicVestingFilterer creates a new log filterer instance of StrategicVesting, bound to a specific deployed contract.
func NewStrategicVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategicVestingFilterer, error) {
	contract, err := bindStrategicVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategicVestingFilterer{contract: contract}, nil
}

// bindStrategicVesting binds a generic wrapper to an already deployed contract.
func bindStrategicVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StrategicVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategicVesting *StrategicVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategicVesting.Contract.StrategicVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategicVesting *StrategicVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategicVesting.Contract.StrategicVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategicVesting *StrategicVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategicVesting.Contract.StrategicVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategicVesting *StrategicVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategicVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategicVesting *StrategicVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategicVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategicVesting *StrategicVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategicVesting.Contract.contract.Transact(opts, method, params...)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) LOCKPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "LOCK_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) LOCKPERIODS() (*big.Int, error) {
	return _StrategicVesting.Contract.LOCKPERIODS(&_StrategicVesting.CallOpts)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) LOCKPERIODS() (*big.Int, error) {
	return _StrategicVesting.Contract.LOCKPERIODS(&_StrategicVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) RELEASEPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "RELEASE_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) RELEASEPERIODS() (*big.Int, error) {
	return _StrategicVesting.Contract.RELEASEPERIODS(&_StrategicVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) RELEASEPERIODS() (*big.Int, error) {
	return _StrategicVesting.Contract.RELEASEPERIODS(&_StrategicVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) TOTALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "TOTAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) TOTALAMOUNT() (*big.Int, error) {
	return _StrategicVesting.Contract.TOTALAMOUNT(&_StrategicVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) TOTALAMOUNT() (*big.Int, error) {
	return _StrategicVesting.Contract.TOTALAMOUNT(&_StrategicVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) WITHDRAWINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "WITHDRAW_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _StrategicVesting.Contract.WITHDRAWINTERVAL(&_StrategicVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _StrategicVesting.Contract.WITHDRAWINTERVAL(&_StrategicVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_StrategicVesting *StrategicVestingCaller) IsStartTimeSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "isStartTimeSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_StrategicVesting *StrategicVestingSession) IsStartTimeSet() (bool, error) {
	return _StrategicVesting.Contract.IsStartTimeSet(&_StrategicVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_StrategicVesting *StrategicVestingCallerSession) IsStartTimeSet() (bool, error) {
	return _StrategicVesting.Contract.IsStartTimeSet(&_StrategicVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) LockPeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "lockPeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) LockPeriods() (*big.Int, error) {
	return _StrategicVesting.Contract.LockPeriods(&_StrategicVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) LockPeriods() (*big.Int, error) {
	return _StrategicVesting.Contract.LockPeriods(&_StrategicVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategicVesting *StrategicVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategicVesting *StrategicVestingSession) Owner() (common.Address, error) {
	return _StrategicVesting.Contract.Owner(&_StrategicVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrategicVesting *StrategicVestingCallerSession) Owner() (common.Address, error) {
	return _StrategicVesting.Contract.Owner(&_StrategicVesting.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_StrategicVesting *StrategicVestingCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "recipients", arg0)

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
func (_StrategicVesting *StrategicVestingSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _StrategicVesting.Contract.Recipients(&_StrategicVesting.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_StrategicVesting *StrategicVestingCallerSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _StrategicVesting.Contract.Recipients(&_StrategicVesting.CallOpts, arg0)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) ReleasePeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "releasePeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) ReleasePeriods() (*big.Int, error) {
	return _StrategicVesting.Contract.ReleasePeriods(&_StrategicVesting.CallOpts)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) ReleasePeriods() (*big.Int, error) {
	return _StrategicVesting.Contract.ReleasePeriods(&_StrategicVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) StartTime() (*big.Int, error) {
	return _StrategicVesting.Contract.StartTime(&_StrategicVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) StartTime() (*big.Int, error) {
	return _StrategicVesting.Contract.StartTime(&_StrategicVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) TotalAmount() (*big.Int, error) {
	return _StrategicVesting.Contract.TotalAmount(&_StrategicVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) TotalAmount() (*big.Int, error) {
	return _StrategicVesting.Contract.TotalAmount(&_StrategicVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StrategicVesting *StrategicVestingCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StrategicVesting *StrategicVestingSession) TotemToken() (common.Address, error) {
	return _StrategicVesting.Contract.TotemToken(&_StrategicVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StrategicVesting *StrategicVestingCallerSession) TotemToken() (common.Address, error) {
	return _StrategicVesting.Contract.TotemToken(&_StrategicVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) UnallocatedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "unallocatedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) UnallocatedAmount() (*big.Int, error) {
	return _StrategicVesting.Contract.UnallocatedAmount(&_StrategicVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) UnallocatedAmount() (*big.Int, error) {
	return _StrategicVesting.Contract.UnallocatedAmount(&_StrategicVesting.CallOpts)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_StrategicVesting *StrategicVestingCaller) Vested(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "vested", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_StrategicVesting *StrategicVestingSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _StrategicVesting.Contract.Vested(&_StrategicVesting.CallOpts, beneficiary)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_StrategicVesting *StrategicVestingCallerSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _StrategicVesting.Contract.Vested(&_StrategicVesting.CallOpts, beneficiary)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_StrategicVesting *StrategicVestingCaller) WithdrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "withdrawInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_StrategicVesting *StrategicVestingSession) WithdrawInterval() (*big.Int, error) {
	return _StrategicVesting.Contract.WithdrawInterval(&_StrategicVesting.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_StrategicVesting *StrategicVestingCallerSession) WithdrawInterval() (*big.Int, error) {
	return _StrategicVesting.Contract.WithdrawInterval(&_StrategicVesting.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_StrategicVesting *StrategicVestingCaller) Withdrawable(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategicVesting.contract.Call(opts, &out, "withdrawable", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_StrategicVesting *StrategicVestingSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _StrategicVesting.Contract.Withdrawable(&_StrategicVesting.CallOpts, beneficiary)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_StrategicVesting *StrategicVestingCallerSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _StrategicVesting.Contract.Withdrawable(&_StrategicVesting.CallOpts, beneficiary)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_StrategicVesting *StrategicVestingTransactor) AddRecipient(opts *bind.TransactOpts, _newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _StrategicVesting.contract.Transact(opts, "addRecipient", _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_StrategicVesting *StrategicVestingSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _StrategicVesting.Contract.AddRecipient(&_StrategicVesting.TransactOpts, _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_StrategicVesting *StrategicVestingTransactorSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _StrategicVesting.Contract.AddRecipient(&_StrategicVesting.TransactOpts, _newRecipient, _totalAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategicVesting *StrategicVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategicVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategicVesting *StrategicVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategicVesting.Contract.RenounceOwnership(&_StrategicVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrategicVesting *StrategicVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrategicVesting.Contract.RenounceOwnership(&_StrategicVesting.TransactOpts)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_StrategicVesting *StrategicVestingTransactor) SetStartTime(opts *bind.TransactOpts, _newStartTime *big.Int) (*types.Transaction, error) {
	return _StrategicVesting.contract.Transact(opts, "setStartTime", _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_StrategicVesting *StrategicVestingSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _StrategicVesting.Contract.SetStartTime(&_StrategicVesting.TransactOpts, _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_StrategicVesting *StrategicVestingTransactorSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _StrategicVesting.Contract.SetStartTime(&_StrategicVesting.TransactOpts, _newStartTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategicVesting *StrategicVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrategicVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategicVesting *StrategicVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategicVesting.Contract.TransferOwnership(&_StrategicVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrategicVesting *StrategicVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrategicVesting.Contract.TransferOwnership(&_StrategicVesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StrategicVesting *StrategicVestingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategicVesting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StrategicVesting *StrategicVestingSession) Withdraw() (*types.Transaction, error) {
	return _StrategicVesting.Contract.Withdraw(&_StrategicVesting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StrategicVesting *StrategicVestingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _StrategicVesting.Contract.Withdraw(&_StrategicVesting.TransactOpts)
}

// StrategicVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrategicVesting contract.
type StrategicVestingOwnershipTransferredIterator struct {
	Event *StrategicVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StrategicVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategicVestingOwnershipTransferred)
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
		it.Event = new(StrategicVestingOwnershipTransferred)
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
func (it *StrategicVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategicVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategicVestingOwnershipTransferred represents a OwnershipTransferred event raised by the StrategicVesting contract.
type StrategicVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategicVesting *StrategicVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrategicVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategicVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrategicVestingOwnershipTransferredIterator{contract: _StrategicVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrategicVesting *StrategicVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrategicVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrategicVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategicVestingOwnershipTransferred)
				if err := _StrategicVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StrategicVesting *StrategicVestingFilterer) ParseOwnershipTransferred(log types.Log) (*StrategicVestingOwnershipTransferred, error) {
	event := new(StrategicVestingOwnershipTransferred)
	if err := _StrategicVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategicVestingStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the StrategicVesting contract.
type StrategicVestingStartTimeSetIterator struct {
	Event *StrategicVestingStartTimeSet // Event containing the contract specifics and raw log

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
func (it *StrategicVestingStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategicVestingStartTimeSet)
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
		it.Event = new(StrategicVestingStartTimeSet)
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
func (it *StrategicVestingStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategicVestingStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategicVestingStartTimeSet represents a StartTimeSet event raised by the StrategicVesting contract.
type StrategicVestingStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_StrategicVesting *StrategicVestingFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*StrategicVestingStartTimeSetIterator, error) {

	logs, sub, err := _StrategicVesting.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &StrategicVestingStartTimeSetIterator{contract: _StrategicVesting.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_StrategicVesting *StrategicVestingFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *StrategicVestingStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _StrategicVesting.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategicVestingStartTimeSet)
				if err := _StrategicVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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
func (_StrategicVesting *StrategicVestingFilterer) ParseStartTimeSet(log types.Log) (*StrategicVestingStartTimeSet, error) {
	event := new(StrategicVestingStartTimeSet)
	if err := _StrategicVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategicVestingVestingScheduleRegisteredIterator is returned from FilterVestingScheduleRegistered and is used to iterate over the raw logs and unpacked data for VestingScheduleRegistered events raised by the StrategicVesting contract.
type StrategicVestingVestingScheduleRegisteredIterator struct {
	Event *StrategicVestingVestingScheduleRegistered // Event containing the contract specifics and raw log

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
func (it *StrategicVestingVestingScheduleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategicVestingVestingScheduleRegistered)
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
		it.Event = new(StrategicVestingVestingScheduleRegistered)
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
func (it *StrategicVestingVestingScheduleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategicVestingVestingScheduleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategicVestingVestingScheduleRegistered represents a VestingScheduleRegistered event raised by the StrategicVesting contract.
type StrategicVestingVestingScheduleRegistered struct {
	RegisteredAddress common.Address
	TotalAmount       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleRegistered is a free log retrieval operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_StrategicVesting *StrategicVestingFilterer) FilterVestingScheduleRegistered(opts *bind.FilterOpts) (*StrategicVestingVestingScheduleRegisteredIterator, error) {

	logs, sub, err := _StrategicVesting.contract.FilterLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return &StrategicVestingVestingScheduleRegisteredIterator{contract: _StrategicVesting.contract, event: "VestingScheduleRegistered", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleRegistered is a free log subscription operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_StrategicVesting *StrategicVestingFilterer) WatchVestingScheduleRegistered(opts *bind.WatchOpts, sink chan<- *StrategicVestingVestingScheduleRegistered) (event.Subscription, error) {

	logs, sub, err := _StrategicVesting.contract.WatchLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategicVestingVestingScheduleRegistered)
				if err := _StrategicVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
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
func (_StrategicVesting *StrategicVestingFilterer) ParseVestingScheduleRegistered(log types.Log) (*StrategicVestingVestingScheduleRegistered, error) {
	event := new(StrategicVestingVestingScheduleRegistered)
	if err := _StrategicVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategicVestingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the StrategicVesting contract.
type StrategicVestingWithdrawIterator struct {
	Event *StrategicVestingWithdraw // Event containing the contract specifics and raw log

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
func (it *StrategicVestingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategicVestingWithdraw)
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
		it.Event = new(StrategicVestingWithdraw)
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
func (it *StrategicVestingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategicVestingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategicVestingWithdraw represents a Withdraw event raised by the StrategicVesting contract.
type StrategicVestingWithdraw struct {
	RegisteredAddress common.Address
	AmountWithdrawn   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_StrategicVesting *StrategicVestingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*StrategicVestingWithdrawIterator, error) {

	logs, sub, err := _StrategicVesting.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &StrategicVestingWithdrawIterator{contract: _StrategicVesting.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_StrategicVesting *StrategicVestingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StrategicVestingWithdraw) (event.Subscription, error) {

	logs, sub, err := _StrategicVesting.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategicVestingWithdraw)
				if err := _StrategicVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_StrategicVesting *StrategicVestingFilterer) ParseWithdraw(log types.Log) (*StrategicVestingWithdraw, error) {
	event := new(StrategicVestingWithdraw)
	if err := _StrategicVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
