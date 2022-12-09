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

// StakingPoolFactoryMetaData contains all meta data concerning the StakingPoolFactory contract.
var StakingPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"},{\"internalType\":\"contractIRewardManager\",\"name\":\"_rewardManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingPoolImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_superAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewStakingPoolImplemnetationWasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewSuperAdminWasSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"wrappedTokenSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"poolType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[12]\",\"name\":\"variables\",\"type\":\"uint256[12]\"},{\"indexed\":false,\"internalType\":\"uint256[8]\",\"name\":\"ranks\",\"type\":\"uint256[8]\"},{\"indexed\":false,\"internalType\":\"uint256[8]\",\"name\":\"percentages\",\"type\":\"uint256[8]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnhancedEnabled\",\"type\":\"bool\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PoolCreatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PoolCreatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPoolCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wrappedTokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_wrappedTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_poolType\",\"type\":\"string\"},{\"internalType\":\"uint256[12]\",\"name\":\"_variables\",\"type\":\"uint256[12]\"},{\"internalType\":\"uint256[8]\",\"name\":\"_ranks\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[8]\",\"name\":\"_percentages\",\"type\":\"uint256[8]\"},{\"internalType\":\"bool\",\"name\":\"isEnhancedEnabled\",\"type\":\"bool\"}],\"name\":\"createPoolProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPoolCreator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePoolCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStakingPoolTaxRate\",\"type\":\"uint256\"}],\"name\":\"setDefaultTaxRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ImpAdr\",\"type\":\"address\"}],\"name\":\"setNewStakingPoolImplementationAdr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_superAdmin\",\"type\":\"address\"}],\"name\":\"setNewSuperAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPoolImplementationAdr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPoolTaxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractITotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingPoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolFactoryMetaData.ABI instead.
var StakingPoolFactoryABI = StakingPoolFactoryMetaData.ABI

// StakingPoolFactory is an auto generated Go binding around an Ethereum contract.
type StakingPoolFactory struct {
	StakingPoolFactoryCaller     // Read-only binding to the contract
	StakingPoolFactoryTransactor // Write-only binding to the contract
	StakingPoolFactoryFilterer   // Log filterer for contract events
}

// StakingPoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolFactorySession struct {
	Contract     *StakingPoolFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakingPoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolFactoryCallerSession struct {
	Contract *StakingPoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StakingPoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolFactoryTransactorSession struct {
	Contract     *StakingPoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StakingPoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolFactoryRaw struct {
	Contract *StakingPoolFactory // Generic contract binding to access the raw methods on
}

// StakingPoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolFactoryCallerRaw struct {
	Contract *StakingPoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolFactoryTransactorRaw struct {
	Contract *StakingPoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPoolFactory creates a new instance of StakingPoolFactory, bound to a specific deployed contract.
func NewStakingPoolFactory(address common.Address, backend bind.ContractBackend) (*StakingPoolFactory, error) {
	contract, err := bindStakingPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactory{StakingPoolFactoryCaller: StakingPoolFactoryCaller{contract: contract}, StakingPoolFactoryTransactor: StakingPoolFactoryTransactor{contract: contract}, StakingPoolFactoryFilterer: StakingPoolFactoryFilterer{contract: contract}}, nil
}

// NewStakingPoolFactoryCaller creates a new read-only instance of StakingPoolFactory, bound to a specific deployed contract.
func NewStakingPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolFactoryCaller, error) {
	contract, err := bindStakingPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryCaller{contract: contract}, nil
}

// NewStakingPoolFactoryTransactor creates a new write-only instance of StakingPoolFactory, bound to a specific deployed contract.
func NewStakingPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolFactoryTransactor, error) {
	contract, err := bindStakingPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryTransactor{contract: contract}, nil
}

// NewStakingPoolFactoryFilterer creates a new log filterer instance of StakingPoolFactory, bound to a specific deployed contract.
func NewStakingPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolFactoryFilterer, error) {
	contract, err := bindStakingPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryFilterer{contract: contract}, nil
}

// bindStakingPoolFactory binds a generic wrapper to an already deployed contract.
func bindStakingPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolFactory *StakingPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolFactory.Contract.StakingPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolFactory *StakingPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.StakingPoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolFactory *StakingPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.StakingPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolFactory *StakingPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolFactory *StakingPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolFactory *StakingPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.contract.Transact(opts, method, params...)
}

// IsPoolCreator is a free data retrieval call binding the contract method 0x9a94775d.
//
// Solidity: function isPoolCreator(address account) view returns(bool)
func (_StakingPoolFactory *StakingPoolFactoryCaller) IsPoolCreator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "isPoolCreator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolCreator is a free data retrieval call binding the contract method 0x9a94775d.
//
// Solidity: function isPoolCreator(address account) view returns(bool)
func (_StakingPoolFactory *StakingPoolFactorySession) IsPoolCreator(account common.Address) (bool, error) {
	return _StakingPoolFactory.Contract.IsPoolCreator(&_StakingPoolFactory.CallOpts, account)
}

// IsPoolCreator is a free data retrieval call binding the contract method 0x9a94775d.
//
// Solidity: function isPoolCreator(address account) view returns(bool)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) IsPoolCreator(account common.Address) (bool, error) {
	return _StakingPoolFactory.Contract.IsPoolCreator(&_StakingPoolFactory.CallOpts, account)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolFactory *StakingPoolFactoryCaller) MinimumStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "minimumStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolFactory *StakingPoolFactorySession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolFactory.Contract.MinimumStakeAmount(&_StakingPoolFactory.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolFactory.Contract.MinimumStakeAmount(&_StakingPoolFactory.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolFactory *StakingPoolFactorySession) RewardManager() (common.Address, error) {
	return _StakingPoolFactory.Contract.RewardManager(&_StakingPoolFactory.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) RewardManager() (common.Address, error) {
	return _StakingPoolFactory.Contract.RewardManager(&_StakingPoolFactory.CallOpts)
}

// StakingPoolImplementationAdr is a free data retrieval call binding the contract method 0xbc8a0669.
//
// Solidity: function stakingPoolImplementationAdr() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCaller) StakingPoolImplementationAdr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "stakingPoolImplementationAdr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPoolImplementationAdr is a free data retrieval call binding the contract method 0xbc8a0669.
//
// Solidity: function stakingPoolImplementationAdr() view returns(address)
func (_StakingPoolFactory *StakingPoolFactorySession) StakingPoolImplementationAdr() (common.Address, error) {
	return _StakingPoolFactory.Contract.StakingPoolImplementationAdr(&_StakingPoolFactory.CallOpts)
}

// StakingPoolImplementationAdr is a free data retrieval call binding the contract method 0xbc8a0669.
//
// Solidity: function stakingPoolImplementationAdr() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) StakingPoolImplementationAdr() (common.Address, error) {
	return _StakingPoolFactory.Contract.StakingPoolImplementationAdr(&_StakingPoolFactory.CallOpts)
}

// StakingPoolTaxRate is a free data retrieval call binding the contract method 0x441f5c65.
//
// Solidity: function stakingPoolTaxRate() view returns(uint256)
func (_StakingPoolFactory *StakingPoolFactoryCaller) StakingPoolTaxRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "stakingPoolTaxRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingPoolTaxRate is a free data retrieval call binding the contract method 0x441f5c65.
//
// Solidity: function stakingPoolTaxRate() view returns(uint256)
func (_StakingPoolFactory *StakingPoolFactorySession) StakingPoolTaxRate() (*big.Int, error) {
	return _StakingPoolFactory.Contract.StakingPoolTaxRate(&_StakingPoolFactory.CallOpts)
}

// StakingPoolTaxRate is a free data retrieval call binding the contract method 0x441f5c65.
//
// Solidity: function stakingPoolTaxRate() view returns(uint256)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) StakingPoolTaxRate() (*big.Int, error) {
	return _StakingPoolFactory.Contract.StakingPoolTaxRate(&_StakingPoolFactory.CallOpts)
}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCaller) SuperAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "superAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_StakingPoolFactory *StakingPoolFactorySession) SuperAdmin() (common.Address, error) {
	return _StakingPoolFactory.Contract.SuperAdmin(&_StakingPoolFactory.CallOpts)
}

// SuperAdmin is a free data retrieval call binding the contract method 0x29575f6a.
//
// Solidity: function superAdmin() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) SuperAdmin() (common.Address, error) {
	return _StakingPoolFactory.Contract.SuperAdmin(&_StakingPoolFactory.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_StakingPoolFactory *StakingPoolFactorySession) SwapRouter() (common.Address, error) {
	return _StakingPoolFactory.Contract.SwapRouter(&_StakingPoolFactory.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) SwapRouter() (common.Address, error) {
	return _StakingPoolFactory.Contract.SwapRouter(&_StakingPoolFactory.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolFactory *StakingPoolFactorySession) TotemToken() (common.Address, error) {
	return _StakingPoolFactory.Contract.TotemToken(&_StakingPoolFactory.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) TotemToken() (common.Address, error) {
	return _StakingPoolFactory.Contract.TotemToken(&_StakingPoolFactory.CallOpts)
}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCaller) UsdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolFactory.contract.Call(opts, &out, "usdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_StakingPoolFactory *StakingPoolFactorySession) UsdToken() (common.Address, error) {
	return _StakingPoolFactory.Contract.UsdToken(&_StakingPoolFactory.CallOpts)
}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_StakingPoolFactory *StakingPoolFactoryCallerSession) UsdToken() (common.Address, error) {
	return _StakingPoolFactory.Contract.UsdToken(&_StakingPoolFactory.CallOpts)
}

// AddPoolCreator is a paid mutator transaction binding the contract method 0x8c2d741c.
//
// Solidity: function addPoolCreator(address account) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactor) AddPoolCreator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.contract.Transact(opts, "addPoolCreator", account)
}

// AddPoolCreator is a paid mutator transaction binding the contract method 0x8c2d741c.
//
// Solidity: function addPoolCreator(address account) returns()
func (_StakingPoolFactory *StakingPoolFactorySession) AddPoolCreator(account common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.AddPoolCreator(&_StakingPoolFactory.TransactOpts, account)
}

// AddPoolCreator is a paid mutator transaction binding the contract method 0x8c2d741c.
//
// Solidity: function addPoolCreator(address account) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactorSession) AddPoolCreator(account common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.AddPoolCreator(&_StakingPoolFactory.TransactOpts, account)
}

// CreatePoolProxy is a paid mutator transaction binding the contract method 0xceba6f8d.
//
// Solidity: function createPoolProxy(address _oracleContract, address _wrappedTokenContract, string _wrappedTokenSymbol, string _poolType, uint256[12] _variables, uint256[8] _ranks, uint256[8] _percentages, bool isEnhancedEnabled) returns(address)
func (_StakingPoolFactory *StakingPoolFactoryTransactor) CreatePoolProxy(opts *bind.TransactOpts, _oracleContract common.Address, _wrappedTokenContract common.Address, _wrappedTokenSymbol string, _poolType string, _variables [12]*big.Int, _ranks [8]*big.Int, _percentages [8]*big.Int, isEnhancedEnabled bool) (*types.Transaction, error) {
	return _StakingPoolFactory.contract.Transact(opts, "createPoolProxy", _oracleContract, _wrappedTokenContract, _wrappedTokenSymbol, _poolType, _variables, _ranks, _percentages, isEnhancedEnabled)
}

// CreatePoolProxy is a paid mutator transaction binding the contract method 0xceba6f8d.
//
// Solidity: function createPoolProxy(address _oracleContract, address _wrappedTokenContract, string _wrappedTokenSymbol, string _poolType, uint256[12] _variables, uint256[8] _ranks, uint256[8] _percentages, bool isEnhancedEnabled) returns(address)
func (_StakingPoolFactory *StakingPoolFactorySession) CreatePoolProxy(_oracleContract common.Address, _wrappedTokenContract common.Address, _wrappedTokenSymbol string, _poolType string, _variables [12]*big.Int, _ranks [8]*big.Int, _percentages [8]*big.Int, isEnhancedEnabled bool) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.CreatePoolProxy(&_StakingPoolFactory.TransactOpts, _oracleContract, _wrappedTokenContract, _wrappedTokenSymbol, _poolType, _variables, _ranks, _percentages, isEnhancedEnabled)
}

// CreatePoolProxy is a paid mutator transaction binding the contract method 0xceba6f8d.
//
// Solidity: function createPoolProxy(address _oracleContract, address _wrappedTokenContract, string _wrappedTokenSymbol, string _poolType, uint256[12] _variables, uint256[8] _ranks, uint256[8] _percentages, bool isEnhancedEnabled) returns(address)
func (_StakingPoolFactory *StakingPoolFactoryTransactorSession) CreatePoolProxy(_oracleContract common.Address, _wrappedTokenContract common.Address, _wrappedTokenSymbol string, _poolType string, _variables [12]*big.Int, _ranks [8]*big.Int, _percentages [8]*big.Int, isEnhancedEnabled bool) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.CreatePoolProxy(&_StakingPoolFactory.TransactOpts, _oracleContract, _wrappedTokenContract, _wrappedTokenSymbol, _poolType, _variables, _ranks, _percentages, isEnhancedEnabled)
}

// RenouncePoolCreator is a paid mutator transaction binding the contract method 0xe281cc7e.
//
// Solidity: function renouncePoolCreator() returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactor) RenouncePoolCreator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolFactory.contract.Transact(opts, "renouncePoolCreator")
}

// RenouncePoolCreator is a paid mutator transaction binding the contract method 0xe281cc7e.
//
// Solidity: function renouncePoolCreator() returns()
func (_StakingPoolFactory *StakingPoolFactorySession) RenouncePoolCreator() (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.RenouncePoolCreator(&_StakingPoolFactory.TransactOpts)
}

// RenouncePoolCreator is a paid mutator transaction binding the contract method 0xe281cc7e.
//
// Solidity: function renouncePoolCreator() returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactorSession) RenouncePoolCreator() (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.RenouncePoolCreator(&_StakingPoolFactory.TransactOpts)
}

// SetDefaultTaxRate is a paid mutator transaction binding the contract method 0xd53b245b.
//
// Solidity: function setDefaultTaxRate(uint256 newStakingPoolTaxRate) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactor) SetDefaultTaxRate(opts *bind.TransactOpts, newStakingPoolTaxRate *big.Int) (*types.Transaction, error) {
	return _StakingPoolFactory.contract.Transact(opts, "setDefaultTaxRate", newStakingPoolTaxRate)
}

// SetDefaultTaxRate is a paid mutator transaction binding the contract method 0xd53b245b.
//
// Solidity: function setDefaultTaxRate(uint256 newStakingPoolTaxRate) returns()
func (_StakingPoolFactory *StakingPoolFactorySession) SetDefaultTaxRate(newStakingPoolTaxRate *big.Int) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetDefaultTaxRate(&_StakingPoolFactory.TransactOpts, newStakingPoolTaxRate)
}

// SetDefaultTaxRate is a paid mutator transaction binding the contract method 0xd53b245b.
//
// Solidity: function setDefaultTaxRate(uint256 newStakingPoolTaxRate) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactorSession) SetDefaultTaxRate(newStakingPoolTaxRate *big.Int) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetDefaultTaxRate(&_StakingPoolFactory.TransactOpts, newStakingPoolTaxRate)
}

// SetNewStakingPoolImplementationAdr is a paid mutator transaction binding the contract method 0x49a84429.
//
// Solidity: function setNewStakingPoolImplementationAdr(address _ImpAdr) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactor) SetNewStakingPoolImplementationAdr(opts *bind.TransactOpts, _ImpAdr common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.contract.Transact(opts, "setNewStakingPoolImplementationAdr", _ImpAdr)
}

// SetNewStakingPoolImplementationAdr is a paid mutator transaction binding the contract method 0x49a84429.
//
// Solidity: function setNewStakingPoolImplementationAdr(address _ImpAdr) returns()
func (_StakingPoolFactory *StakingPoolFactorySession) SetNewStakingPoolImplementationAdr(_ImpAdr common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetNewStakingPoolImplementationAdr(&_StakingPoolFactory.TransactOpts, _ImpAdr)
}

// SetNewStakingPoolImplementationAdr is a paid mutator transaction binding the contract method 0x49a84429.
//
// Solidity: function setNewStakingPoolImplementationAdr(address _ImpAdr) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactorSession) SetNewStakingPoolImplementationAdr(_ImpAdr common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetNewStakingPoolImplementationAdr(&_StakingPoolFactory.TransactOpts, _ImpAdr)
}

// SetNewSuperAdmin is a paid mutator transaction binding the contract method 0xb6c3ef27.
//
// Solidity: function setNewSuperAdmin(address _superAdmin) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactor) SetNewSuperAdmin(opts *bind.TransactOpts, _superAdmin common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.contract.Transact(opts, "setNewSuperAdmin", _superAdmin)
}

// SetNewSuperAdmin is a paid mutator transaction binding the contract method 0xb6c3ef27.
//
// Solidity: function setNewSuperAdmin(address _superAdmin) returns()
func (_StakingPoolFactory *StakingPoolFactorySession) SetNewSuperAdmin(_superAdmin common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetNewSuperAdmin(&_StakingPoolFactory.TransactOpts, _superAdmin)
}

// SetNewSuperAdmin is a paid mutator transaction binding the contract method 0xb6c3ef27.
//
// Solidity: function setNewSuperAdmin(address _superAdmin) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactorSession) SetNewSuperAdmin(_superAdmin common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetNewSuperAdmin(&_StakingPoolFactory.TransactOpts, _superAdmin)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactor) SetSwapRouter(opts *bind.TransactOpts, _swapRouter common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.contract.Transact(opts, "setSwapRouter", _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_StakingPoolFactory *StakingPoolFactorySession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetSwapRouter(&_StakingPoolFactory.TransactOpts, _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_StakingPoolFactory *StakingPoolFactoryTransactorSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _StakingPoolFactory.Contract.SetSwapRouter(&_StakingPoolFactory.TransactOpts, _swapRouter)
}

// StakingPoolFactoryNewStakingPoolImplemnetationWasSetIterator is returned from FilterNewStakingPoolImplemnetationWasSet and is used to iterate over the raw logs and unpacked data for NewStakingPoolImplemnetationWasSet events raised by the StakingPoolFactory contract.
type StakingPoolFactoryNewStakingPoolImplemnetationWasSetIterator struct {
	Event *StakingPoolFactoryNewStakingPoolImplemnetationWasSet // Event containing the contract specifics and raw log

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
func (it *StakingPoolFactoryNewStakingPoolImplemnetationWasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolFactoryNewStakingPoolImplemnetationWasSet)
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
		it.Event = new(StakingPoolFactoryNewStakingPoolImplemnetationWasSet)
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
func (it *StakingPoolFactoryNewStakingPoolImplemnetationWasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolFactoryNewStakingPoolImplemnetationWasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolFactoryNewStakingPoolImplemnetationWasSet represents a NewStakingPoolImplemnetationWasSet event raised by the StakingPoolFactory contract.
type StakingPoolFactoryNewStakingPoolImplemnetationWasSet struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewStakingPoolImplemnetationWasSet is a free log retrieval operation binding the contract event 0x3199ab49607559567913d4714288fee1e692c31829ed4722a6fd064adbdc5b54.
//
// Solidity: event NewStakingPoolImplemnetationWasSet()
func (_StakingPoolFactory *StakingPoolFactoryFilterer) FilterNewStakingPoolImplemnetationWasSet(opts *bind.FilterOpts) (*StakingPoolFactoryNewStakingPoolImplemnetationWasSetIterator, error) {

	logs, sub, err := _StakingPoolFactory.contract.FilterLogs(opts, "NewStakingPoolImplemnetationWasSet")
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryNewStakingPoolImplemnetationWasSetIterator{contract: _StakingPoolFactory.contract, event: "NewStakingPoolImplemnetationWasSet", logs: logs, sub: sub}, nil
}

// WatchNewStakingPoolImplemnetationWasSet is a free log subscription operation binding the contract event 0x3199ab49607559567913d4714288fee1e692c31829ed4722a6fd064adbdc5b54.
//
// Solidity: event NewStakingPoolImplemnetationWasSet()
func (_StakingPoolFactory *StakingPoolFactoryFilterer) WatchNewStakingPoolImplemnetationWasSet(opts *bind.WatchOpts, sink chan<- *StakingPoolFactoryNewStakingPoolImplemnetationWasSet) (event.Subscription, error) {

	logs, sub, err := _StakingPoolFactory.contract.WatchLogs(opts, "NewStakingPoolImplemnetationWasSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolFactoryNewStakingPoolImplemnetationWasSet)
				if err := _StakingPoolFactory.contract.UnpackLog(event, "NewStakingPoolImplemnetationWasSet", log); err != nil {
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

// ParseNewStakingPoolImplemnetationWasSet is a log parse operation binding the contract event 0x3199ab49607559567913d4714288fee1e692c31829ed4722a6fd064adbdc5b54.
//
// Solidity: event NewStakingPoolImplemnetationWasSet()
func (_StakingPoolFactory *StakingPoolFactoryFilterer) ParseNewStakingPoolImplemnetationWasSet(log types.Log) (*StakingPoolFactoryNewStakingPoolImplemnetationWasSet, error) {
	event := new(StakingPoolFactoryNewStakingPoolImplemnetationWasSet)
	if err := _StakingPoolFactory.contract.UnpackLog(event, "NewStakingPoolImplemnetationWasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolFactoryNewSuperAdminWasSetIterator is returned from FilterNewSuperAdminWasSet and is used to iterate over the raw logs and unpacked data for NewSuperAdminWasSet events raised by the StakingPoolFactory contract.
type StakingPoolFactoryNewSuperAdminWasSetIterator struct {
	Event *StakingPoolFactoryNewSuperAdminWasSet // Event containing the contract specifics and raw log

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
func (it *StakingPoolFactoryNewSuperAdminWasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolFactoryNewSuperAdminWasSet)
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
		it.Event = new(StakingPoolFactoryNewSuperAdminWasSet)
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
func (it *StakingPoolFactoryNewSuperAdminWasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolFactoryNewSuperAdminWasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolFactoryNewSuperAdminWasSet represents a NewSuperAdminWasSet event raised by the StakingPoolFactory contract.
type StakingPoolFactoryNewSuperAdminWasSet struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewSuperAdminWasSet is a free log retrieval operation binding the contract event 0xb25686d5de2544aa0f0235ea99dd20e962286bc9cec2fb371b226a637be461f7.
//
// Solidity: event NewSuperAdminWasSet()
func (_StakingPoolFactory *StakingPoolFactoryFilterer) FilterNewSuperAdminWasSet(opts *bind.FilterOpts) (*StakingPoolFactoryNewSuperAdminWasSetIterator, error) {

	logs, sub, err := _StakingPoolFactory.contract.FilterLogs(opts, "NewSuperAdminWasSet")
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryNewSuperAdminWasSetIterator{contract: _StakingPoolFactory.contract, event: "NewSuperAdminWasSet", logs: logs, sub: sub}, nil
}

// WatchNewSuperAdminWasSet is a free log subscription operation binding the contract event 0xb25686d5de2544aa0f0235ea99dd20e962286bc9cec2fb371b226a637be461f7.
//
// Solidity: event NewSuperAdminWasSet()
func (_StakingPoolFactory *StakingPoolFactoryFilterer) WatchNewSuperAdminWasSet(opts *bind.WatchOpts, sink chan<- *StakingPoolFactoryNewSuperAdminWasSet) (event.Subscription, error) {

	logs, sub, err := _StakingPoolFactory.contract.WatchLogs(opts, "NewSuperAdminWasSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolFactoryNewSuperAdminWasSet)
				if err := _StakingPoolFactory.contract.UnpackLog(event, "NewSuperAdminWasSet", log); err != nil {
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

// ParseNewSuperAdminWasSet is a log parse operation binding the contract event 0xb25686d5de2544aa0f0235ea99dd20e962286bc9cec2fb371b226a637be461f7.
//
// Solidity: event NewSuperAdminWasSet()
func (_StakingPoolFactory *StakingPoolFactoryFilterer) ParseNewSuperAdminWasSet(log types.Log) (*StakingPoolFactoryNewSuperAdminWasSet, error) {
	event := new(StakingPoolFactoryNewSuperAdminWasSet)
	if err := _StakingPoolFactory.contract.UnpackLog(event, "NewSuperAdminWasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the StakingPoolFactory contract.
type StakingPoolFactoryPoolCreatedIterator struct {
	Event *StakingPoolFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *StakingPoolFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolFactoryPoolCreated)
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
		it.Event = new(StakingPoolFactoryPoolCreated)
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
func (it *StakingPoolFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolFactoryPoolCreated represents a PoolCreated event raised by the StakingPoolFactory contract.
type StakingPoolFactoryPoolCreated struct {
	Pool               common.Address
	WrappedTokenSymbol string
	PoolType           string
	Variables          [12]*big.Int
	Ranks              [8]*big.Int
	Percentages        [8]*big.Int
	IsEnhancedEnabled  bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xd69cf3c5f6280df773a744806dc6920bc09109ef4a352c0846f5ef35965dca12.
//
// Solidity: event PoolCreated(address indexed pool, string wrappedTokenSymbol, string poolType, uint256[12] variables, uint256[8] ranks, uint256[8] percentages, bool isEnhancedEnabled)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, pool []common.Address) (*StakingPoolFactoryPoolCreatedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _StakingPoolFactory.contract.FilterLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryPoolCreatedIterator{contract: _StakingPoolFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xd69cf3c5f6280df773a744806dc6920bc09109ef4a352c0846f5ef35965dca12.
//
// Solidity: event PoolCreated(address indexed pool, string wrappedTokenSymbol, string poolType, uint256[12] variables, uint256[8] ranks, uint256[8] percentages, bool isEnhancedEnabled)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *StakingPoolFactoryPoolCreated, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _StakingPoolFactory.contract.WatchLogs(opts, "PoolCreated", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolFactoryPoolCreated)
				if err := _StakingPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0xd69cf3c5f6280df773a744806dc6920bc09109ef4a352c0846f5ef35965dca12.
//
// Solidity: event PoolCreated(address indexed pool, string wrappedTokenSymbol, string poolType, uint256[12] variables, uint256[8] ranks, uint256[8] percentages, bool isEnhancedEnabled)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) ParsePoolCreated(log types.Log) (*StakingPoolFactoryPoolCreated, error) {
	event := new(StakingPoolFactoryPoolCreated)
	if err := _StakingPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolFactoryPoolCreatorAddedIterator is returned from FilterPoolCreatorAdded and is used to iterate over the raw logs and unpacked data for PoolCreatorAdded events raised by the StakingPoolFactory contract.
type StakingPoolFactoryPoolCreatorAddedIterator struct {
	Event *StakingPoolFactoryPoolCreatorAdded // Event containing the contract specifics and raw log

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
func (it *StakingPoolFactoryPoolCreatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolFactoryPoolCreatorAdded)
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
		it.Event = new(StakingPoolFactoryPoolCreatorAdded)
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
func (it *StakingPoolFactoryPoolCreatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolFactoryPoolCreatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolFactoryPoolCreatorAdded represents a PoolCreatorAdded event raised by the StakingPoolFactory contract.
type StakingPoolFactoryPoolCreatorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolCreatorAdded is a free log retrieval operation binding the contract event 0xac89bb7b3d0c5a763a97f31bc75f8faee00426e7b235f02ca76897d55caf7b61.
//
// Solidity: event PoolCreatorAdded(address indexed account)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) FilterPoolCreatorAdded(opts *bind.FilterOpts, account []common.Address) (*StakingPoolFactoryPoolCreatorAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakingPoolFactory.contract.FilterLogs(opts, "PoolCreatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryPoolCreatorAddedIterator{contract: _StakingPoolFactory.contract, event: "PoolCreatorAdded", logs: logs, sub: sub}, nil
}

// WatchPoolCreatorAdded is a free log subscription operation binding the contract event 0xac89bb7b3d0c5a763a97f31bc75f8faee00426e7b235f02ca76897d55caf7b61.
//
// Solidity: event PoolCreatorAdded(address indexed account)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) WatchPoolCreatorAdded(opts *bind.WatchOpts, sink chan<- *StakingPoolFactoryPoolCreatorAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakingPoolFactory.contract.WatchLogs(opts, "PoolCreatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolFactoryPoolCreatorAdded)
				if err := _StakingPoolFactory.contract.UnpackLog(event, "PoolCreatorAdded", log); err != nil {
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

// ParsePoolCreatorAdded is a log parse operation binding the contract event 0xac89bb7b3d0c5a763a97f31bc75f8faee00426e7b235f02ca76897d55caf7b61.
//
// Solidity: event PoolCreatorAdded(address indexed account)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) ParsePoolCreatorAdded(log types.Log) (*StakingPoolFactoryPoolCreatorAdded, error) {
	event := new(StakingPoolFactoryPoolCreatorAdded)
	if err := _StakingPoolFactory.contract.UnpackLog(event, "PoolCreatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolFactoryPoolCreatorRemovedIterator is returned from FilterPoolCreatorRemoved and is used to iterate over the raw logs and unpacked data for PoolCreatorRemoved events raised by the StakingPoolFactory contract.
type StakingPoolFactoryPoolCreatorRemovedIterator struct {
	Event *StakingPoolFactoryPoolCreatorRemoved // Event containing the contract specifics and raw log

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
func (it *StakingPoolFactoryPoolCreatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolFactoryPoolCreatorRemoved)
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
		it.Event = new(StakingPoolFactoryPoolCreatorRemoved)
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
func (it *StakingPoolFactoryPoolCreatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolFactoryPoolCreatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolFactoryPoolCreatorRemoved represents a PoolCreatorRemoved event raised by the StakingPoolFactory contract.
type StakingPoolFactoryPoolCreatorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolCreatorRemoved is a free log retrieval operation binding the contract event 0x2bc71cc60b543df5fdd80ef5a45e1cdb344843769b649e9d901de3be24aeb24e.
//
// Solidity: event PoolCreatorRemoved(address indexed account)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) FilterPoolCreatorRemoved(opts *bind.FilterOpts, account []common.Address) (*StakingPoolFactoryPoolCreatorRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakingPoolFactory.contract.FilterLogs(opts, "PoolCreatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolFactoryPoolCreatorRemovedIterator{contract: _StakingPoolFactory.contract, event: "PoolCreatorRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolCreatorRemoved is a free log subscription operation binding the contract event 0x2bc71cc60b543df5fdd80ef5a45e1cdb344843769b649e9d901de3be24aeb24e.
//
// Solidity: event PoolCreatorRemoved(address indexed account)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) WatchPoolCreatorRemoved(opts *bind.WatchOpts, sink chan<- *StakingPoolFactoryPoolCreatorRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakingPoolFactory.contract.WatchLogs(opts, "PoolCreatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolFactoryPoolCreatorRemoved)
				if err := _StakingPoolFactory.contract.UnpackLog(event, "PoolCreatorRemoved", log); err != nil {
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

// ParsePoolCreatorRemoved is a log parse operation binding the contract event 0x2bc71cc60b543df5fdd80ef5a45e1cdb344843769b649e9d901de3be24aeb24e.
//
// Solidity: event PoolCreatorRemoved(address indexed account)
func (_StakingPoolFactory *StakingPoolFactoryFilterer) ParsePoolCreatorRemoved(log types.Log) (*StakingPoolFactoryPoolCreatorRemoved, error) {
	event := new(StakingPoolFactoryPoolCreatorRemoved)
	if err := _StakingPoolFactory.contract.UnpackLog(event, "PoolCreatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
