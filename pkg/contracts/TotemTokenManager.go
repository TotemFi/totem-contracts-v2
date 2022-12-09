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

// TotemTokenManagerMetaData contains all meta data concerning the TotemTokenManager contract.
var TotemTokenManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_CommunityDevelopmentAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_StakingRewardsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_LiquidityPoolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_PublicSaleAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_AdvisorsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_SeedInvestmentAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_PrivateSaleAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_TeamAllocationAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_StrategicRoundAddr\",\"type\":\"address\"}],\"name\":\"setDistributionTeamsAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_locker\",\"type\":\"address\"}],\"name\":\"setLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setTaxExemptStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTaxRate\",\"type\":\"uint256\"}],\"name\":\"setTaxRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTaxationWallet\",\"type\":\"address\"}],\"name\":\"setTaxationWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"totemTokenTransferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TotemTokenManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TotemTokenManagerMetaData.ABI instead.
var TotemTokenManagerABI = TotemTokenManagerMetaData.ABI

// TotemTokenManager is an auto generated Go binding around an Ethereum contract.
type TotemTokenManager struct {
	TotemTokenManagerCaller     // Read-only binding to the contract
	TotemTokenManagerTransactor // Write-only binding to the contract
	TotemTokenManagerFilterer   // Log filterer for contract events
}

// TotemTokenManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TotemTokenManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemTokenManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TotemTokenManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemTokenManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TotemTokenManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemTokenManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TotemTokenManagerSession struct {
	Contract     *TotemTokenManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TotemTokenManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TotemTokenManagerCallerSession struct {
	Contract *TotemTokenManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TotemTokenManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TotemTokenManagerTransactorSession struct {
	Contract     *TotemTokenManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TotemTokenManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TotemTokenManagerRaw struct {
	Contract *TotemTokenManager // Generic contract binding to access the raw methods on
}

// TotemTokenManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TotemTokenManagerCallerRaw struct {
	Contract *TotemTokenManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TotemTokenManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TotemTokenManagerTransactorRaw struct {
	Contract *TotemTokenManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTotemTokenManager creates a new instance of TotemTokenManager, bound to a specific deployed contract.
func NewTotemTokenManager(address common.Address, backend bind.ContractBackend) (*TotemTokenManager, error) {
	contract, err := bindTotemTokenManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TotemTokenManager{TotemTokenManagerCaller: TotemTokenManagerCaller{contract: contract}, TotemTokenManagerTransactor: TotemTokenManagerTransactor{contract: contract}, TotemTokenManagerFilterer: TotemTokenManagerFilterer{contract: contract}}, nil
}

// NewTotemTokenManagerCaller creates a new read-only instance of TotemTokenManager, bound to a specific deployed contract.
func NewTotemTokenManagerCaller(address common.Address, caller bind.ContractCaller) (*TotemTokenManagerCaller, error) {
	contract, err := bindTotemTokenManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TotemTokenManagerCaller{contract: contract}, nil
}

// NewTotemTokenManagerTransactor creates a new write-only instance of TotemTokenManager, bound to a specific deployed contract.
func NewTotemTokenManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TotemTokenManagerTransactor, error) {
	contract, err := bindTotemTokenManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TotemTokenManagerTransactor{contract: contract}, nil
}

// NewTotemTokenManagerFilterer creates a new log filterer instance of TotemTokenManager, bound to a specific deployed contract.
func NewTotemTokenManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TotemTokenManagerFilterer, error) {
	contract, err := bindTotemTokenManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TotemTokenManagerFilterer{contract: contract}, nil
}

// bindTotemTokenManager binds a generic wrapper to an already deployed contract.
func bindTotemTokenManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TotemTokenManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TotemTokenManager *TotemTokenManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TotemTokenManager.Contract.TotemTokenManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TotemTokenManager *TotemTokenManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.TotemTokenManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TotemTokenManager *TotemTokenManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.TotemTokenManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TotemTokenManager *TotemTokenManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TotemTokenManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TotemTokenManager *TotemTokenManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TotemTokenManager *TotemTokenManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.contract.Transact(opts, method, params...)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_TotemTokenManager *TotemTokenManagerCaller) IsManager(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _TotemTokenManager.contract.Call(opts, &out, "isManager", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_TotemTokenManager *TotemTokenManagerSession) IsManager(account common.Address) (bool, error) {
	return _TotemTokenManager.Contract.IsManager(&_TotemTokenManager.CallOpts, account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_TotemTokenManager *TotemTokenManagerCallerSession) IsManager(account common.Address) (bool, error) {
	return _TotemTokenManager.Contract.IsManager(&_TotemTokenManager.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemTokenManager *TotemTokenManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemTokenManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemTokenManager *TotemTokenManagerSession) Owner() (common.Address, error) {
	return _TotemTokenManager.Contract.Owner(&_TotemTokenManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemTokenManager *TotemTokenManagerCallerSession) Owner() (common.Address, error) {
	return _TotemTokenManager.Contract.Owner(&_TotemTokenManager.CallOpts)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address account) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) AddManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "addManager", account)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address account) returns()
func (_TotemTokenManager *TotemTokenManagerSession) AddManager(account common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.AddManager(&_TotemTokenManager.TransactOpts, account)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address account) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) AddManager(account common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.AddManager(&_TotemTokenManager.TransactOpts, account)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) DistributeTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "distributeTokens")
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_TotemTokenManager *TotemTokenManagerSession) DistributeTokens() (*types.Transaction, error) {
	return _TotemTokenManager.Contract.DistributeTokens(&_TotemTokenManager.TransactOpts)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) DistributeTokens() (*types.Transaction, error) {
	return _TotemTokenManager.Contract.DistributeTokens(&_TotemTokenManager.TransactOpts)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address account) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) RemoveManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "removeManager", account)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address account) returns()
func (_TotemTokenManager *TotemTokenManagerSession) RemoveManager(account common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.RemoveManager(&_TotemTokenManager.TransactOpts, account)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address account) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) RemoveManager(account common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.RemoveManager(&_TotemTokenManager.TransactOpts, account)
}

// RenounceManager is a paid mutator transaction binding the contract method 0xf8b91abe.
//
// Solidity: function renounceManager() returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) RenounceManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "renounceManager")
}

// RenounceManager is a paid mutator transaction binding the contract method 0xf8b91abe.
//
// Solidity: function renounceManager() returns()
func (_TotemTokenManager *TotemTokenManagerSession) RenounceManager() (*types.Transaction, error) {
	return _TotemTokenManager.Contract.RenounceManager(&_TotemTokenManager.TransactOpts)
}

// RenounceManager is a paid mutator transaction binding the contract method 0xf8b91abe.
//
// Solidity: function renounceManager() returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) RenounceManager() (*types.Transaction, error) {
	return _TotemTokenManager.Contract.RenounceManager(&_TotemTokenManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemTokenManager *TotemTokenManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TotemTokenManager.Contract.RenounceOwnership(&_TotemTokenManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TotemTokenManager.Contract.RenounceOwnership(&_TotemTokenManager.TransactOpts)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) SetDistributionTeamsAddresses(opts *bind.TransactOpts, _CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "setDistributionTeamsAddresses", _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_TotemTokenManager *TotemTokenManagerSession) SetDistributionTeamsAddresses(_CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetDistributionTeamsAddresses(&_TotemTokenManager.TransactOpts, _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) SetDistributionTeamsAddresses(_CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetDistributionTeamsAddresses(&_TotemTokenManager.TransactOpts, _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) SetLocker(opts *bind.TransactOpts, _locker common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "setLocker", _locker)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_TotemTokenManager *TotemTokenManagerSession) SetLocker(_locker common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetLocker(&_TotemTokenManager.TransactOpts, _locker)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) SetLocker(_locker common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetLocker(&_TotemTokenManager.TransactOpts, _locker)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) SetTaxExemptStatus(opts *bind.TransactOpts, account common.Address, status bool) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "setTaxExemptStatus", account, status)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_TotemTokenManager *TotemTokenManagerSession) SetTaxExemptStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetTaxExemptStatus(&_TotemTokenManager.TransactOpts, account, status)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) SetTaxExemptStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetTaxExemptStatus(&_TotemTokenManager.TransactOpts, account, status)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) SetTaxRate(opts *bind.TransactOpts, newTaxRate *big.Int) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "setTaxRate", newTaxRate)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_TotemTokenManager *TotemTokenManagerSession) SetTaxRate(newTaxRate *big.Int) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetTaxRate(&_TotemTokenManager.TransactOpts, newTaxRate)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) SetTaxRate(newTaxRate *big.Int) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetTaxRate(&_TotemTokenManager.TransactOpts, newTaxRate)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) SetTaxationWallet(opts *bind.TransactOpts, newTaxationWallet common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "setTaxationWallet", newTaxationWallet)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_TotemTokenManager *TotemTokenManagerSession) SetTaxationWallet(newTaxationWallet common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetTaxationWallet(&_TotemTokenManager.TransactOpts, newTaxationWallet)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) SetTaxationWallet(newTaxationWallet common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.SetTaxationWallet(&_TotemTokenManager.TransactOpts, newTaxationWallet)
}

// TotemTokenTransferOwnership is a paid mutator transaction binding the contract method 0xabd7b03f.
//
// Solidity: function totemTokenTransferOwnership(address newOwner) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) TotemTokenTransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "totemTokenTransferOwnership", newOwner)
}

// TotemTokenTransferOwnership is a paid mutator transaction binding the contract method 0xabd7b03f.
//
// Solidity: function totemTokenTransferOwnership(address newOwner) returns()
func (_TotemTokenManager *TotemTokenManagerSession) TotemTokenTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.TotemTokenTransferOwnership(&_TotemTokenManager.TransactOpts, newOwner)
}

// TotemTokenTransferOwnership is a paid mutator transaction binding the contract method 0xabd7b03f.
//
// Solidity: function totemTokenTransferOwnership(address newOwner) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) TotemTokenTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.TotemTokenTransferOwnership(&_TotemTokenManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemTokenManager *TotemTokenManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemTokenManager *TotemTokenManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.TransferOwnership(&_TotemTokenManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemTokenManager *TotemTokenManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemTokenManager.Contract.TransferOwnership(&_TotemTokenManager.TransactOpts, newOwner)
}

// TotemTokenManagerManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the TotemTokenManager contract.
type TotemTokenManagerManagerAddedIterator struct {
	Event *TotemTokenManagerManagerAdded // Event containing the contract specifics and raw log

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
func (it *TotemTokenManagerManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemTokenManagerManagerAdded)
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
		it.Event = new(TotemTokenManagerManagerAdded)
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
func (it *TotemTokenManagerManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemTokenManagerManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemTokenManagerManagerAdded represents a ManagerAdded event raised by the TotemTokenManager contract.
type TotemTokenManagerManagerAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed account)
func (_TotemTokenManager *TotemTokenManagerFilterer) FilterManagerAdded(opts *bind.FilterOpts, account []common.Address) (*TotemTokenManagerManagerAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TotemTokenManager.contract.FilterLogs(opts, "ManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &TotemTokenManagerManagerAddedIterator{contract: _TotemTokenManager.contract, event: "ManagerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed account)
func (_TotemTokenManager *TotemTokenManagerFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *TotemTokenManagerManagerAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TotemTokenManager.contract.WatchLogs(opts, "ManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemTokenManagerManagerAdded)
				if err := _TotemTokenManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
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

// ParseManagerAdded is a log parse operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed account)
func (_TotemTokenManager *TotemTokenManagerFilterer) ParseManagerAdded(log types.Log) (*TotemTokenManagerManagerAdded, error) {
	event := new(TotemTokenManagerManagerAdded)
	if err := _TotemTokenManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemTokenManagerManagerRemovedIterator is returned from FilterManagerRemoved and is used to iterate over the raw logs and unpacked data for ManagerRemoved events raised by the TotemTokenManager contract.
type TotemTokenManagerManagerRemovedIterator struct {
	Event *TotemTokenManagerManagerRemoved // Event containing the contract specifics and raw log

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
func (it *TotemTokenManagerManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemTokenManagerManagerRemoved)
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
		it.Event = new(TotemTokenManagerManagerRemoved)
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
func (it *TotemTokenManagerManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemTokenManagerManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemTokenManagerManagerRemoved represents a ManagerRemoved event raised by the TotemTokenManager contract.
type TotemTokenManagerManagerRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerRemoved is a free log retrieval operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed account)
func (_TotemTokenManager *TotemTokenManagerFilterer) FilterManagerRemoved(opts *bind.FilterOpts, account []common.Address) (*TotemTokenManagerManagerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TotemTokenManager.contract.FilterLogs(opts, "ManagerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &TotemTokenManagerManagerRemovedIterator{contract: _TotemTokenManager.contract, event: "ManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchManagerRemoved is a free log subscription operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed account)
func (_TotemTokenManager *TotemTokenManagerFilterer) WatchManagerRemoved(opts *bind.WatchOpts, sink chan<- *TotemTokenManagerManagerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TotemTokenManager.contract.WatchLogs(opts, "ManagerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemTokenManagerManagerRemoved)
				if err := _TotemTokenManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
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

// ParseManagerRemoved is a log parse operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed account)
func (_TotemTokenManager *TotemTokenManagerFilterer) ParseManagerRemoved(log types.Log) (*TotemTokenManagerManagerRemoved, error) {
	event := new(TotemTokenManagerManagerRemoved)
	if err := _TotemTokenManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemTokenManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TotemTokenManager contract.
type TotemTokenManagerOwnershipTransferredIterator struct {
	Event *TotemTokenManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TotemTokenManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemTokenManagerOwnershipTransferred)
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
		it.Event = new(TotemTokenManagerOwnershipTransferred)
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
func (it *TotemTokenManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemTokenManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemTokenManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TotemTokenManager contract.
type TotemTokenManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TotemTokenManager *TotemTokenManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TotemTokenManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TotemTokenManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TotemTokenManagerOwnershipTransferredIterator{contract: _TotemTokenManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TotemTokenManager *TotemTokenManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TotemTokenManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TotemTokenManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemTokenManagerOwnershipTransferred)
				if err := _TotemTokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TotemTokenManager *TotemTokenManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TotemTokenManagerOwnershipTransferred, error) {
	event := new(TotemTokenManagerOwnershipTransferred)
	if err := _TotemTokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
