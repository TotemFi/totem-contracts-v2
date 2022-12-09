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

// TotemVestingMetaData contains all meta data concerning the TotemVesting contract.
var TotemVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_releasePeriods\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockPeriods\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"addRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStartTimeSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releasePeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartTime\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unallocatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"vested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountVested\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TotemVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use TotemVestingMetaData.ABI instead.
var TotemVestingABI = TotemVestingMetaData.ABI

// TotemVesting is an auto generated Go binding around an Ethereum contract.
type TotemVesting struct {
	TotemVestingCaller     // Read-only binding to the contract
	TotemVestingTransactor // Write-only binding to the contract
	TotemVestingFilterer   // Log filterer for contract events
}

// TotemVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TotemVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TotemVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TotemVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TotemVestingSession struct {
	Contract     *TotemVesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TotemVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TotemVestingCallerSession struct {
	Contract *TotemVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TotemVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TotemVestingTransactorSession struct {
	Contract     *TotemVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TotemVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TotemVestingRaw struct {
	Contract *TotemVesting // Generic contract binding to access the raw methods on
}

// TotemVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TotemVestingCallerRaw struct {
	Contract *TotemVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TotemVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TotemVestingTransactorRaw struct {
	Contract *TotemVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTotemVesting creates a new instance of TotemVesting, bound to a specific deployed contract.
func NewTotemVesting(address common.Address, backend bind.ContractBackend) (*TotemVesting, error) {
	contract, err := bindTotemVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TotemVesting{TotemVestingCaller: TotemVestingCaller{contract: contract}, TotemVestingTransactor: TotemVestingTransactor{contract: contract}, TotemVestingFilterer: TotemVestingFilterer{contract: contract}}, nil
}

// NewTotemVestingCaller creates a new read-only instance of TotemVesting, bound to a specific deployed contract.
func NewTotemVestingCaller(address common.Address, caller bind.ContractCaller) (*TotemVestingCaller, error) {
	contract, err := bindTotemVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TotemVestingCaller{contract: contract}, nil
}

// NewTotemVestingTransactor creates a new write-only instance of TotemVesting, bound to a specific deployed contract.
func NewTotemVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TotemVestingTransactor, error) {
	contract, err := bindTotemVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TotemVestingTransactor{contract: contract}, nil
}

// NewTotemVestingFilterer creates a new log filterer instance of TotemVesting, bound to a specific deployed contract.
func NewTotemVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TotemVestingFilterer, error) {
	contract, err := bindTotemVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TotemVestingFilterer{contract: contract}, nil
}

// bindTotemVesting binds a generic wrapper to an already deployed contract.
func bindTotemVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TotemVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TotemVesting *TotemVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TotemVesting.Contract.TotemVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TotemVesting *TotemVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemVesting.Contract.TotemVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TotemVesting *TotemVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TotemVesting.Contract.TotemVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TotemVesting *TotemVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TotemVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TotemVesting *TotemVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TotemVesting *TotemVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TotemVesting.Contract.contract.Transact(opts, method, params...)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_TotemVesting *TotemVestingCaller) IsStartTimeSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "isStartTimeSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_TotemVesting *TotemVestingSession) IsStartTimeSet() (bool, error) {
	return _TotemVesting.Contract.IsStartTimeSet(&_TotemVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_TotemVesting *TotemVestingCallerSession) IsStartTimeSet() (bool, error) {
	return _TotemVesting.Contract.IsStartTimeSet(&_TotemVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_TotemVesting *TotemVestingCaller) LockPeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "lockPeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_TotemVesting *TotemVestingSession) LockPeriods() (*big.Int, error) {
	return _TotemVesting.Contract.LockPeriods(&_TotemVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_TotemVesting *TotemVestingCallerSession) LockPeriods() (*big.Int, error) {
	return _TotemVesting.Contract.LockPeriods(&_TotemVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemVesting *TotemVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemVesting *TotemVestingSession) Owner() (common.Address, error) {
	return _TotemVesting.Contract.Owner(&_TotemVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemVesting *TotemVestingCallerSession) Owner() (common.Address, error) {
	return _TotemVesting.Contract.Owner(&_TotemVesting.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_TotemVesting *TotemVestingCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "recipients", arg0)

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
func (_TotemVesting *TotemVestingSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _TotemVesting.Contract.Recipients(&_TotemVesting.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_TotemVesting *TotemVestingCallerSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _TotemVesting.Contract.Recipients(&_TotemVesting.CallOpts, arg0)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_TotemVesting *TotemVestingCaller) ReleasePeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "releasePeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_TotemVesting *TotemVestingSession) ReleasePeriods() (*big.Int, error) {
	return _TotemVesting.Contract.ReleasePeriods(&_TotemVesting.CallOpts)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_TotemVesting *TotemVestingCallerSession) ReleasePeriods() (*big.Int, error) {
	return _TotemVesting.Contract.ReleasePeriods(&_TotemVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_TotemVesting *TotemVestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_TotemVesting *TotemVestingSession) StartTime() (*big.Int, error) {
	return _TotemVesting.Contract.StartTime(&_TotemVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_TotemVesting *TotemVestingCallerSession) StartTime() (*big.Int, error) {
	return _TotemVesting.Contract.StartTime(&_TotemVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TotemVesting *TotemVestingCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TotemVesting *TotemVestingSession) TotalAmount() (*big.Int, error) {
	return _TotemVesting.Contract.TotalAmount(&_TotemVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TotemVesting *TotemVestingCallerSession) TotalAmount() (*big.Int, error) {
	return _TotemVesting.Contract.TotalAmount(&_TotemVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_TotemVesting *TotemVestingCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_TotemVesting *TotemVestingSession) TotemToken() (common.Address, error) {
	return _TotemVesting.Contract.TotemToken(&_TotemVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_TotemVesting *TotemVestingCallerSession) TotemToken() (common.Address, error) {
	return _TotemVesting.Contract.TotemToken(&_TotemVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_TotemVesting *TotemVestingCaller) UnallocatedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "unallocatedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_TotemVesting *TotemVestingSession) UnallocatedAmount() (*big.Int, error) {
	return _TotemVesting.Contract.UnallocatedAmount(&_TotemVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_TotemVesting *TotemVestingCallerSession) UnallocatedAmount() (*big.Int, error) {
	return _TotemVesting.Contract.UnallocatedAmount(&_TotemVesting.CallOpts)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_TotemVesting *TotemVestingCaller) Vested(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "vested", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_TotemVesting *TotemVestingSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _TotemVesting.Contract.Vested(&_TotemVesting.CallOpts, beneficiary)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_TotemVesting *TotemVestingCallerSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _TotemVesting.Contract.Vested(&_TotemVesting.CallOpts, beneficiary)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_TotemVesting *TotemVestingCaller) WithdrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "withdrawInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_TotemVesting *TotemVestingSession) WithdrawInterval() (*big.Int, error) {
	return _TotemVesting.Contract.WithdrawInterval(&_TotemVesting.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_TotemVesting *TotemVestingCallerSession) WithdrawInterval() (*big.Int, error) {
	return _TotemVesting.Contract.WithdrawInterval(&_TotemVesting.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_TotemVesting *TotemVestingCaller) Withdrawable(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TotemVesting.contract.Call(opts, &out, "withdrawable", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_TotemVesting *TotemVestingSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _TotemVesting.Contract.Withdrawable(&_TotemVesting.CallOpts, beneficiary)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_TotemVesting *TotemVestingCallerSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _TotemVesting.Contract.Withdrawable(&_TotemVesting.CallOpts, beneficiary)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_TotemVesting *TotemVestingTransactor) AddRecipient(opts *bind.TransactOpts, _newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _TotemVesting.contract.Transact(opts, "addRecipient", _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_TotemVesting *TotemVestingSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _TotemVesting.Contract.AddRecipient(&_TotemVesting.TransactOpts, _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_TotemVesting *TotemVestingTransactorSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _TotemVesting.Contract.AddRecipient(&_TotemVesting.TransactOpts, _newRecipient, _totalAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemVesting *TotemVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemVesting *TotemVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TotemVesting.Contract.RenounceOwnership(&_TotemVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemVesting *TotemVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TotemVesting.Contract.RenounceOwnership(&_TotemVesting.TransactOpts)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_TotemVesting *TotemVestingTransactor) SetStartTime(opts *bind.TransactOpts, _newStartTime *big.Int) (*types.Transaction, error) {
	return _TotemVesting.contract.Transact(opts, "setStartTime", _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_TotemVesting *TotemVestingSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _TotemVesting.Contract.SetStartTime(&_TotemVesting.TransactOpts, _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_TotemVesting *TotemVestingTransactorSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _TotemVesting.Contract.SetStartTime(&_TotemVesting.TransactOpts, _newStartTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemVesting *TotemVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TotemVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemVesting *TotemVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemVesting.Contract.TransferOwnership(&_TotemVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemVesting *TotemVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemVesting.Contract.TransferOwnership(&_TotemVesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TotemVesting *TotemVestingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemVesting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TotemVesting *TotemVestingSession) Withdraw() (*types.Transaction, error) {
	return _TotemVesting.Contract.Withdraw(&_TotemVesting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TotemVesting *TotemVestingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _TotemVesting.Contract.Withdraw(&_TotemVesting.TransactOpts)
}

// TotemVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TotemVesting contract.
type TotemVestingOwnershipTransferredIterator struct {
	Event *TotemVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TotemVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemVestingOwnershipTransferred)
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
		it.Event = new(TotemVestingOwnershipTransferred)
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
func (it *TotemVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TotemVesting contract.
type TotemVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TotemVesting *TotemVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TotemVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TotemVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TotemVestingOwnershipTransferredIterator{contract: _TotemVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TotemVesting *TotemVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TotemVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TotemVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemVestingOwnershipTransferred)
				if err := _TotemVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TotemVesting *TotemVestingFilterer) ParseOwnershipTransferred(log types.Log) (*TotemVestingOwnershipTransferred, error) {
	event := new(TotemVestingOwnershipTransferred)
	if err := _TotemVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemVestingStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the TotemVesting contract.
type TotemVestingStartTimeSetIterator struct {
	Event *TotemVestingStartTimeSet // Event containing the contract specifics and raw log

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
func (it *TotemVestingStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemVestingStartTimeSet)
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
		it.Event = new(TotemVestingStartTimeSet)
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
func (it *TotemVestingStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemVestingStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemVestingStartTimeSet represents a StartTimeSet event raised by the TotemVesting contract.
type TotemVestingStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_TotemVesting *TotemVestingFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*TotemVestingStartTimeSetIterator, error) {

	logs, sub, err := _TotemVesting.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &TotemVestingStartTimeSetIterator{contract: _TotemVesting.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_TotemVesting *TotemVestingFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *TotemVestingStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _TotemVesting.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemVestingStartTimeSet)
				if err := _TotemVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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
func (_TotemVesting *TotemVestingFilterer) ParseStartTimeSet(log types.Log) (*TotemVestingStartTimeSet, error) {
	event := new(TotemVestingStartTimeSet)
	if err := _TotemVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemVestingVestingScheduleRegisteredIterator is returned from FilterVestingScheduleRegistered and is used to iterate over the raw logs and unpacked data for VestingScheduleRegistered events raised by the TotemVesting contract.
type TotemVestingVestingScheduleRegisteredIterator struct {
	Event *TotemVestingVestingScheduleRegistered // Event containing the contract specifics and raw log

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
func (it *TotemVestingVestingScheduleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemVestingVestingScheduleRegistered)
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
		it.Event = new(TotemVestingVestingScheduleRegistered)
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
func (it *TotemVestingVestingScheduleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemVestingVestingScheduleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemVestingVestingScheduleRegistered represents a VestingScheduleRegistered event raised by the TotemVesting contract.
type TotemVestingVestingScheduleRegistered struct {
	RegisteredAddress common.Address
	TotalAmount       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleRegistered is a free log retrieval operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_TotemVesting *TotemVestingFilterer) FilterVestingScheduleRegistered(opts *bind.FilterOpts) (*TotemVestingVestingScheduleRegisteredIterator, error) {

	logs, sub, err := _TotemVesting.contract.FilterLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return &TotemVestingVestingScheduleRegisteredIterator{contract: _TotemVesting.contract, event: "VestingScheduleRegistered", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleRegistered is a free log subscription operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_TotemVesting *TotemVestingFilterer) WatchVestingScheduleRegistered(opts *bind.WatchOpts, sink chan<- *TotemVestingVestingScheduleRegistered) (event.Subscription, error) {

	logs, sub, err := _TotemVesting.contract.WatchLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemVestingVestingScheduleRegistered)
				if err := _TotemVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
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
func (_TotemVesting *TotemVestingFilterer) ParseVestingScheduleRegistered(log types.Log) (*TotemVestingVestingScheduleRegistered, error) {
	event := new(TotemVestingVestingScheduleRegistered)
	if err := _TotemVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemVestingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the TotemVesting contract.
type TotemVestingWithdrawIterator struct {
	Event *TotemVestingWithdraw // Event containing the contract specifics and raw log

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
func (it *TotemVestingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemVestingWithdraw)
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
		it.Event = new(TotemVestingWithdraw)
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
func (it *TotemVestingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemVestingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemVestingWithdraw represents a Withdraw event raised by the TotemVesting contract.
type TotemVestingWithdraw struct {
	RegisteredAddress common.Address
	AmountWithdrawn   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_TotemVesting *TotemVestingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TotemVestingWithdrawIterator, error) {

	logs, sub, err := _TotemVesting.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TotemVestingWithdrawIterator{contract: _TotemVesting.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_TotemVesting *TotemVestingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TotemVestingWithdraw) (event.Subscription, error) {

	logs, sub, err := _TotemVesting.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemVestingWithdraw)
				if err := _TotemVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_TotemVesting *TotemVestingFilterer) ParseWithdraw(log types.Log) (*TotemVestingWithdraw, error) {
	event := new(TotemVestingWithdraw)
	if err := _TotemVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
