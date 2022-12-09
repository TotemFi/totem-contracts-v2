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

// StakingPoolProxyMetaData contains all meta data concerning the StakingPoolProxy contract.
var StakingPoolProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApproveTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedBTC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ImplementationUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferTokens\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapRouterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BUSDContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WrappedTokenContractAddress\",\"type\":\"address\"}],\"name\":\"__WrappedTokenDistributor_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averagePricePrediction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getEstimatedWrappedTokenForUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPathForUSDToWrappedToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_wrappedTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_poolType\",\"type\":\"string\"},{\"internalType\":\"contractITotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"},{\"internalType\":\"contractIRewardManager\",\"name\":\"_rewardManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolCreator\",\"type\":\"address\"},{\"internalType\":\"address[4]\",\"name\":\"_addrs\",\"type\":\"address[4]\"},{\"internalType\":\"uint256[12]\",\"name\":\"_variables\",\"type\":\"uint256[12]\"},{\"internalType\":\"uint256[8]\",\"name\":\"_ranks\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[8]\",\"name\":\"_percentages\",\"type\":\"uint256[8]\"},{\"internalType\":\"bool\",\"name\":\"_isEnhancedEnabled\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAnEmergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnhancedEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMatured\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"launchDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPrizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeApr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collaborativeReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEnhancedEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMatured\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialCollabReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"predictions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastWithdrawalTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePrediction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difference\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"prizeRewardWithdrawn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didUnstake\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prizeRewardRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeLimitRangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sortedStakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeTaxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPoolImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractITotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newStakingPoolImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingPoolProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolProxyMetaData.ABI instead.
var StakingPoolProxyABI = StakingPoolProxyMetaData.ABI

// StakingPoolProxy is an auto generated Go binding around an Ethereum contract.
type StakingPoolProxy struct {
	StakingPoolProxyCaller     // Read-only binding to the contract
	StakingPoolProxyTransactor // Write-only binding to the contract
	StakingPoolProxyFilterer   // Log filterer for contract events
}

// StakingPoolProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolProxySession struct {
	Contract     *StakingPoolProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingPoolProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolProxyCallerSession struct {
	Contract *StakingPoolProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StakingPoolProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolProxyTransactorSession struct {
	Contract     *StakingPoolProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StakingPoolProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolProxyRaw struct {
	Contract *StakingPoolProxy // Generic contract binding to access the raw methods on
}

// StakingPoolProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolProxyCallerRaw struct {
	Contract *StakingPoolProxyCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolProxyTransactorRaw struct {
	Contract *StakingPoolProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPoolProxy creates a new instance of StakingPoolProxy, bound to a specific deployed contract.
func NewStakingPoolProxy(address common.Address, backend bind.ContractBackend) (*StakingPoolProxy, error) {
	contract, err := bindStakingPoolProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxy{StakingPoolProxyCaller: StakingPoolProxyCaller{contract: contract}, StakingPoolProxyTransactor: StakingPoolProxyTransactor{contract: contract}, StakingPoolProxyFilterer: StakingPoolProxyFilterer{contract: contract}}, nil
}

// NewStakingPoolProxyCaller creates a new read-only instance of StakingPoolProxy, bound to a specific deployed contract.
func NewStakingPoolProxyCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolProxyCaller, error) {
	contract, err := bindStakingPoolProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyCaller{contract: contract}, nil
}

// NewStakingPoolProxyTransactor creates a new write-only instance of StakingPoolProxy, bound to a specific deployed contract.
func NewStakingPoolProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolProxyTransactor, error) {
	contract, err := bindStakingPoolProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyTransactor{contract: contract}, nil
}

// NewStakingPoolProxyFilterer creates a new log filterer instance of StakingPoolProxy, bound to a specific deployed contract.
func NewStakingPoolProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolProxyFilterer, error) {
	contract, err := bindStakingPoolProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyFilterer{contract: contract}, nil
}

// bindStakingPoolProxy binds a generic wrapper to an already deployed contract.
func bindStakingPoolProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolProxy *StakingPoolProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolProxy.Contract.StakingPoolProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolProxy *StakingPoolProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.StakingPoolProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolProxy *StakingPoolProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.StakingPoolProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolProxy *StakingPoolProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolProxy *StakingPoolProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolProxy *StakingPoolProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.contract.Transact(opts, method, params...)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) AveragePricePrediction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "averagePricePrediction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolProxy.Contract.AveragePricePrediction(&_StakingPoolProxy.CallOpts)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolProxy.Contract.AveragePricePrediction(&_StakingPoolProxy.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) CollaborativeRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "collaborativeRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolProxy.Contract.CollaborativeRange(&_StakingPoolProxy.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolProxy.Contract.CollaborativeRange(&_StakingPoolProxy.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) CollaborativeReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "collaborativeReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolProxy.Contract.CollaborativeReward(&_StakingPoolProxy.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolProxy.Contract.CollaborativeReward(&_StakingPoolProxy.CallOpts)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) GetEstimatedWrappedTokenForUSD(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "getEstimatedWrappedTokenForUSD", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolProxy.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolProxy.CallOpts, _amount)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolProxy.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolProxy.CallOpts, _amount)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolProxy *StakingPoolProxyCaller) GetPathForUSDToWrappedToken(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "getPathForUSDToWrappedToken")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolProxy *StakingPoolProxySession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolProxy.Contract.GetPathForUSDToWrappedToken(&_StakingPoolProxy.CallOpts)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolProxy *StakingPoolProxyCallerSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolProxy.Contract.GetPathForUSDToWrappedToken(&_StakingPoolProxy.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) GetSwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "getSwapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolProxy.Contract.GetSwapRouter(&_StakingPoolProxy.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolProxy.Contract.GetSwapRouter(&_StakingPoolProxy.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) GetUSDBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "getUSDBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolProxy.Contract.GetUSDBalance(&_StakingPoolProxy.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolProxy.Contract.GetUSDBalance(&_StakingPoolProxy.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) GetUSDToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "getUSDToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) GetUSDToken() (common.Address, error) {
	return _StakingPoolProxy.Contract.GetUSDToken(&_StakingPoolProxy.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) GetUSDToken() (common.Address, error) {
	return _StakingPoolProxy.Contract.GetUSDToken(&_StakingPoolProxy.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxySession) IsActive() (bool, error) {
	return _StakingPoolProxy.Contract.IsActive(&_StakingPoolProxy.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) IsActive() (bool, error) {
	return _StakingPoolProxy.Contract.IsActive(&_StakingPoolProxy.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCaller) IsAnEmergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "isAnEmergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxySession) IsAnEmergency() (bool, error) {
	return _StakingPoolProxy.Contract.IsAnEmergency(&_StakingPoolProxy.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) IsAnEmergency() (bool, error) {
	return _StakingPoolProxy.Contract.IsAnEmergency(&_StakingPoolProxy.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCaller) IsDeleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "isDeleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxySession) IsDeleted() (bool, error) {
	return _StakingPoolProxy.Contract.IsDeleted(&_StakingPoolProxy.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) IsDeleted() (bool, error) {
	return _StakingPoolProxy.Contract.IsDeleted(&_StakingPoolProxy.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCaller) IsEnhancedEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "isEnhancedEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxySession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolProxy.Contract.IsEnhancedEnabled(&_StakingPoolProxy.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolProxy.Contract.IsEnhancedEnabled(&_StakingPoolProxy.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxySession) IsLocked() (bool, error) {
	return _StakingPoolProxy.Contract.IsLocked(&_StakingPoolProxy.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) IsLocked() (bool, error) {
	return _StakingPoolProxy.Contract.IsLocked(&_StakingPoolProxy.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCaller) IsMatured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "isMatured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxySession) IsMatured() (bool, error) {
	return _StakingPoolProxy.Contract.IsMatured(&_StakingPoolProxy.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) IsMatured() (bool, error) {
	return _StakingPoolProxy.Contract.IsMatured(&_StakingPoolProxy.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) LaunchDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "launchDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) LaunchDate() (*big.Int, error) {
	return _StakingPoolProxy.Contract.LaunchDate(&_StakingPoolProxy.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) LaunchDate() (*big.Int, error) {
	return _StakingPoolProxy.Contract.LaunchDate(&_StakingPoolProxy.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) LockTime() (*big.Int, error) {
	return _StakingPoolProxy.Contract.LockTime(&_StakingPoolProxy.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) LockTime() (*big.Int, error) {
	return _StakingPoolProxy.Contract.LockTime(&_StakingPoolProxy.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolProxy *StakingPoolProxyCaller) Lps(opts *bind.CallOpts) (struct {
	LaunchDate          *big.Int
	LockTime            *big.Int
	MaturityTime        *big.Int
	MaturingPrice       *big.Int
	UsdPrizeAmount      *big.Int
	PrizeAmount         *big.Int
	StakeApr            *big.Int
	CollaborativeReward *big.Int
	OracleDecimals      *big.Int
	IsEnhancedEnabled   bool
	IsMatured           bool
}, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "lps")

	outstruct := new(struct {
		LaunchDate          *big.Int
		LockTime            *big.Int
		MaturityTime        *big.Int
		MaturingPrice       *big.Int
		UsdPrizeAmount      *big.Int
		PrizeAmount         *big.Int
		StakeApr            *big.Int
		CollaborativeReward *big.Int
		OracleDecimals      *big.Int
		IsEnhancedEnabled   bool
		IsMatured           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LaunchDate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaturityTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaturingPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UsdPrizeAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrizeAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StakeApr = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CollaborativeReward = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.OracleDecimals = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.IsEnhancedEnabled = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.IsMatured = *abi.ConvertType(out[10], new(bool)).(*bool)

	return *outstruct, err

}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolProxy *StakingPoolProxySession) Lps() (struct {
	LaunchDate          *big.Int
	LockTime            *big.Int
	MaturityTime        *big.Int
	MaturingPrice       *big.Int
	UsdPrizeAmount      *big.Int
	PrizeAmount         *big.Int
	StakeApr            *big.Int
	CollaborativeReward *big.Int
	OracleDecimals      *big.Int
	IsEnhancedEnabled   bool
	IsMatured           bool
}, error) {
	return _StakingPoolProxy.Contract.Lps(&_StakingPoolProxy.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) Lps() (struct {
	LaunchDate          *big.Int
	LockTime            *big.Int
	MaturityTime        *big.Int
	MaturingPrice       *big.Int
	UsdPrizeAmount      *big.Int
	PrizeAmount         *big.Int
	StakeApr            *big.Int
	CollaborativeReward *big.Int
	OracleDecimals      *big.Int
	IsEnhancedEnabled   bool
	IsMatured           bool
}, error) {
	return _StakingPoolProxy.Contract.Lps(&_StakingPoolProxy.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) MaturingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "maturingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolProxy.Contract.MaturingPrice(&_StakingPoolProxy.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolProxy.Contract.MaturingPrice(&_StakingPoolProxy.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) MaturityTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "maturityTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) MaturityTime() (*big.Int, error) {
	return _StakingPoolProxy.Contract.MaturityTime(&_StakingPoolProxy.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) MaturityTime() (*big.Int, error) {
	return _StakingPoolProxy.Contract.MaturityTime(&_StakingPoolProxy.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) MinimumStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "minimumStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolProxy.Contract.MinimumStakeAmount(&_StakingPoolProxy.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolProxy.Contract.MinimumStakeAmount(&_StakingPoolProxy.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) OracleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "oracleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) OracleContract() (common.Address, error) {
	return _StakingPoolProxy.Contract.OracleContract(&_StakingPoolProxy.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) OracleContract() (common.Address, error) {
	return _StakingPoolProxy.Contract.OracleContract(&_StakingPoolProxy.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) OracleDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "oracleDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolProxy.Contract.OracleDecimals(&_StakingPoolProxy.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolProxy.Contract.OracleDecimals(&_StakingPoolProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) Owner() (common.Address, error) {
	return _StakingPoolProxy.Contract.Owner(&_StakingPoolProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) Owner() (common.Address, error) {
	return _StakingPoolProxy.Contract.Owner(&_StakingPoolProxy.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) PoolCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "poolCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) PoolCreator() (common.Address, error) {
	return _StakingPoolProxy.Contract.PoolCreator(&_StakingPoolProxy.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) PoolCreator() (common.Address, error) {
	return _StakingPoolProxy.Contract.PoolCreator(&_StakingPoolProxy.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolProxy *StakingPoolProxyCaller) PoolType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolProxy *StakingPoolProxySession) PoolType() (string, error) {
	return _StakingPoolProxy.Contract.PoolType(&_StakingPoolProxy.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) PoolType() (string, error) {
	return _StakingPoolProxy.Contract.PoolType(&_StakingPoolProxy.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) PotentialCollabReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "potentialCollabReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolProxy.Contract.PotentialCollabReward(&_StakingPoolProxy.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolProxy.Contract.PotentialCollabReward(&_StakingPoolProxy.CallOpts)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolProxy *StakingPoolProxyCaller) Predictions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakedBalance        *big.Int
	StakedTime           *big.Int
	AmountWithdrawn      *big.Int
	LastWithdrawalTime   *big.Int
	PricePrediction      *big.Int
	Difference           *big.Int
	Rank                 *big.Int
	PrizeRewardWithdrawn bool
	DidUnstake           bool
}, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "predictions", arg0, arg1)

	outstruct := new(struct {
		StakedBalance        *big.Int
		StakedTime           *big.Int
		AmountWithdrawn      *big.Int
		LastWithdrawalTime   *big.Int
		PricePrediction      *big.Int
		Difference           *big.Int
		Rank                 *big.Int
		PrizeRewardWithdrawn bool
		DidUnstake           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountWithdrawn = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastWithdrawalTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PricePrediction = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Difference = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Rank = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PrizeRewardWithdrawn = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.DidUnstake = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolProxy *StakingPoolProxySession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBalance        *big.Int
	StakedTime           *big.Int
	AmountWithdrawn      *big.Int
	LastWithdrawalTime   *big.Int
	PricePrediction      *big.Int
	Difference           *big.Int
	Rank                 *big.Int
	PrizeRewardWithdrawn bool
	DidUnstake           bool
}, error) {
	return _StakingPoolProxy.Contract.Predictions(&_StakingPoolProxy.CallOpts, arg0, arg1)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBalance        *big.Int
	StakedTime           *big.Int
	AmountWithdrawn      *big.Int
	LastWithdrawalTime   *big.Int
	PricePrediction      *big.Int
	Difference           *big.Int
	Rank                 *big.Int
	PrizeRewardWithdrawn bool
	DidUnstake           bool
}, error) {
	return _StakingPoolProxy.Contract.Predictions(&_StakingPoolProxy.CallOpts, arg0, arg1)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) PrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "prizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolProxy.Contract.PrizeAmount(&_StakingPoolProxy.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolProxy.Contract.PrizeAmount(&_StakingPoolProxy.CallOpts)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolProxy *StakingPoolProxyCaller) PrizeRewardRates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "prizeRewardRates", arg0)

	outstruct := new(struct {
		Rank       *big.Int
		Percentage *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rank = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Percentage = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolProxy *StakingPoolProxySession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolProxy.Contract.PrizeRewardRates(&_StakingPoolProxy.CallOpts, arg0)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolProxy.Contract.PrizeRewardRates(&_StakingPoolProxy.CallOpts, arg0)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) RewardManager() (common.Address, error) {
	return _StakingPoolProxy.Contract.RewardManager(&_StakingPoolProxy.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) RewardManager() (common.Address, error) {
	return _StakingPoolProxy.Contract.RewardManager(&_StakingPoolProxy.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) SizeAllocation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "sizeAllocation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolProxy.Contract.SizeAllocation(&_StakingPoolProxy.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolProxy.Contract.SizeAllocation(&_StakingPoolProxy.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) SizeLimitRangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "sizeLimitRangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolProxy.Contract.SizeLimitRangeRate(&_StakingPoolProxy.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolProxy.Contract.SizeLimitRangeRate(&_StakingPoolProxy.CallOpts)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolProxy *StakingPoolProxyCaller) SortedStakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "sortedStakers", arg0)

	outstruct := new(struct {
		StakerAddress common.Address
		Index         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolProxy *StakingPoolProxySession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolProxy.Contract.SortedStakers(&_StakingPoolProxy.CallOpts, arg0)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolProxy.Contract.SortedStakers(&_StakingPoolProxy.CallOpts, arg0)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) StakeApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "stakeApr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) StakeApr() (*big.Int, error) {
	return _StakingPoolProxy.Contract.StakeApr(&_StakingPoolProxy.CallOpts)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) StakeApr() (*big.Int, error) {
	return _StakingPoolProxy.Contract.StakeApr(&_StakingPoolProxy.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) StakeTaxRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "stakeTaxRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolProxy.Contract.StakeTaxRate(&_StakingPoolProxy.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolProxy.Contract.StakeTaxRate(&_StakingPoolProxy.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolProxy *StakingPoolProxyCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		StakerAddress common.Address
		Index         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolProxy *StakingPoolProxySession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolProxy.Contract.Stakers(&_StakingPoolProxy.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolProxy.Contract.Stakers(&_StakingPoolProxy.CallOpts, arg0)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) StakingPoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "stakingPoolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolProxy.Contract.StakingPoolImplementation(&_StakingPoolProxy.CallOpts)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolProxy.Contract.StakingPoolImplementation(&_StakingPoolProxy.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) TotalStaked() (*big.Int, error) {
	return _StakingPoolProxy.Contract.TotalStaked(&_StakingPoolProxy.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) TotalStaked() (*big.Int, error) {
	return _StakingPoolProxy.Contract.TotalStaked(&_StakingPoolProxy.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) TotemToken() (common.Address, error) {
	return _StakingPoolProxy.Contract.TotemToken(&_StakingPoolProxy.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) TotemToken() (common.Address, error) {
	return _StakingPoolProxy.Contract.TotemToken(&_StakingPoolProxy.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCaller) UpgradeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "upgradeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxySession) UpgradeEnabled() (bool, error) {
	return _StakingPoolProxy.Contract.UpgradeEnabled(&_StakingPoolProxy.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) UpgradeEnabled() (bool, error) {
	return _StakingPoolProxy.Contract.UpgradeEnabled(&_StakingPoolProxy.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCaller) UsdPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "usdPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxySession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolProxy.Contract.UsdPrizeAmount(&_StakingPoolProxy.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolProxy.Contract.UsdPrizeAmount(&_StakingPoolProxy.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCaller) WrappedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "wrappedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxySession) WrappedToken() (common.Address, error) {
	return _StakingPoolProxy.Contract.WrappedToken(&_StakingPoolProxy.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) WrappedToken() (common.Address, error) {
	return _StakingPoolProxy.Contract.WrappedToken(&_StakingPoolProxy.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolProxy *StakingPoolProxyCaller) WrappedTokenSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolProxy.contract.Call(opts, &out, "wrappedTokenSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolProxy *StakingPoolProxySession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolProxy.Contract.WrappedTokenSymbol(&_StakingPoolProxy.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolProxy *StakingPoolProxyCallerSession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolProxy.Contract.WrappedTokenSymbol(&_StakingPoolProxy.CallOpts)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) WrappedTokenDistributorInitialize(opts *bind.TransactOpts, swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.Transact(opts, "__WrappedTokenDistributor_initialize", swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolProxy *StakingPoolProxySession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.WrappedTokenDistributorInitialize(&_StakingPoolProxy.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.WrappedTokenDistributorInitialize(&_StakingPoolProxy.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// DisableUpgrade is a paid mutator transaction binding the contract method 0x67fc9138.
//
// Solidity: function disableUpgrade() returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) DisableUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.Transact(opts, "disableUpgrade")
}

// DisableUpgrade is a paid mutator transaction binding the contract method 0x67fc9138.
//
// Solidity: function disableUpgrade() returns()
func (_StakingPoolProxy *StakingPoolProxySession) DisableUpgrade() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.DisableUpgrade(&_StakingPoolProxy.TransactOpts)
}

// DisableUpgrade is a paid mutator transaction binding the contract method 0x67fc9138.
//
// Solidity: function disableUpgrade() returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) DisableUpgrade() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.DisableUpgrade(&_StakingPoolProxy.TransactOpts)
}

// EnableUpgrade is a paid mutator transaction binding the contract method 0x28016f9a.
//
// Solidity: function enableUpgrade() returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) EnableUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.Transact(opts, "enableUpgrade")
}

// EnableUpgrade is a paid mutator transaction binding the contract method 0x28016f9a.
//
// Solidity: function enableUpgrade() returns()
func (_StakingPoolProxy *StakingPoolProxySession) EnableUpgrade() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.EnableUpgrade(&_StakingPoolProxy.TransactOpts)
}

// EnableUpgrade is a paid mutator transaction binding the contract method 0x28016f9a.
//
// Solidity: function enableUpgrade() returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) EnableUpgrade() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.EnableUpgrade(&_StakingPoolProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x869eb310.
//
// Solidity: function initialize(string _wrappedTokenSymbol, string _poolType, address _totemToken, address _rewardManager, address _poolCreator, address[4] _addrs, uint256[12] _variables, uint256[8] _ranks, uint256[8] _percentages, bool _isEnhancedEnabled) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) Initialize(opts *bind.TransactOpts, _wrappedTokenSymbol string, _poolType string, _totemToken common.Address, _rewardManager common.Address, _poolCreator common.Address, _addrs [4]common.Address, _variables [12]*big.Int, _ranks [8]*big.Int, _percentages [8]*big.Int, _isEnhancedEnabled bool) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.Transact(opts, "initialize", _wrappedTokenSymbol, _poolType, _totemToken, _rewardManager, _poolCreator, _addrs, _variables, _ranks, _percentages, _isEnhancedEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0x869eb310.
//
// Solidity: function initialize(string _wrappedTokenSymbol, string _poolType, address _totemToken, address _rewardManager, address _poolCreator, address[4] _addrs, uint256[12] _variables, uint256[8] _ranks, uint256[8] _percentages, bool _isEnhancedEnabled) returns()
func (_StakingPoolProxy *StakingPoolProxySession) Initialize(_wrappedTokenSymbol string, _poolType string, _totemToken common.Address, _rewardManager common.Address, _poolCreator common.Address, _addrs [4]common.Address, _variables [12]*big.Int, _ranks [8]*big.Int, _percentages [8]*big.Int, _isEnhancedEnabled bool) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.Initialize(&_StakingPoolProxy.TransactOpts, _wrappedTokenSymbol, _poolType, _totemToken, _rewardManager, _poolCreator, _addrs, _variables, _ranks, _percentages, _isEnhancedEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0x869eb310.
//
// Solidity: function initialize(string _wrappedTokenSymbol, string _poolType, address _totemToken, address _rewardManager, address _poolCreator, address[4] _addrs, uint256[12] _variables, uint256[8] _ranks, uint256[8] _percentages, bool _isEnhancedEnabled) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) Initialize(_wrappedTokenSymbol string, _poolType string, _totemToken common.Address, _rewardManager common.Address, _poolCreator common.Address, _addrs [4]common.Address, _variables [12]*big.Int, _ranks [8]*big.Int, _percentages [8]*big.Int, _isEnhancedEnabled bool) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.Initialize(&_StakingPoolProxy.TransactOpts, _wrappedTokenSymbol, _poolType, _totemToken, _rewardManager, _poolCreator, _addrs, _variables, _ranks, _percentages, _isEnhancedEnabled)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolProxy *StakingPoolProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.RenounceOwnership(&_StakingPoolProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.RenounceOwnership(&_StakingPoolProxy.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolProxy *StakingPoolProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.TransferOwnership(&_StakingPoolProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.TransferOwnership(&_StakingPoolProxy.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newStakingPoolImplementation) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) UpgradeTo(opts *bind.TransactOpts, _newStakingPoolImplementation common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.Transact(opts, "upgradeTo", _newStakingPoolImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newStakingPoolImplementation) returns()
func (_StakingPoolProxy *StakingPoolProxySession) UpgradeTo(_newStakingPoolImplementation common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.UpgradeTo(&_StakingPoolProxy.TransactOpts, _newStakingPoolImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newStakingPoolImplementation) returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) UpgradeTo(_newStakingPoolImplementation common.Address) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.UpgradeTo(&_StakingPoolProxy.TransactOpts, _newStakingPoolImplementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingPoolProxy *StakingPoolProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.Fallback(&_StakingPoolProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.Fallback(&_StakingPoolProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingPoolProxy *StakingPoolProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingPoolProxy *StakingPoolProxySession) Receive() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.Receive(&_StakingPoolProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingPoolProxy *StakingPoolProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingPoolProxy.Contract.Receive(&_StakingPoolProxy.TransactOpts)
}

// StakingPoolProxyApproveTokensIterator is returned from FilterApproveTokens and is used to iterate over the raw logs and unpacked data for ApproveTokens events raised by the StakingPoolProxy contract.
type StakingPoolProxyApproveTokensIterator struct {
	Event *StakingPoolProxyApproveTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolProxyApproveTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolProxyApproveTokens)
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
		it.Event = new(StakingPoolProxyApproveTokens)
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
func (it *StakingPoolProxyApproveTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolProxyApproveTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolProxyApproveTokens represents a ApproveTokens event raised by the StakingPoolProxy contract.
type StakingPoolProxyApproveTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApproveTokens is a free log retrieval operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) FilterApproveTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolProxyApproveTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.FilterLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyApproveTokensIterator{contract: _StakingPoolProxy.contract, event: "ApproveTokens", logs: logs, sub: sub}, nil
}

// WatchApproveTokens is a free log subscription operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) WatchApproveTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolProxyApproveTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.WatchLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolProxyApproveTokens)
				if err := _StakingPoolProxy.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
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

// ParseApproveTokens is a log parse operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) ParseApproveTokens(log types.Log) (*StakingPoolProxyApproveTokens, error) {
	event := new(StakingPoolProxyApproveTokens)
	if err := _StakingPoolProxy.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolProxyDistributedBTCIterator is returned from FilterDistributedBTC and is used to iterate over the raw logs and unpacked data for DistributedBTC events raised by the StakingPoolProxy contract.
type StakingPoolProxyDistributedBTCIterator struct {
	Event *StakingPoolProxyDistributedBTC // Event containing the contract specifics and raw log

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
func (it *StakingPoolProxyDistributedBTCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolProxyDistributedBTC)
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
		it.Event = new(StakingPoolProxyDistributedBTC)
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
func (it *StakingPoolProxyDistributedBTCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolProxyDistributedBTCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolProxyDistributedBTC represents a DistributedBTC event raised by the StakingPoolProxy contract.
type StakingPoolProxyDistributedBTC struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedBTC is a free log retrieval operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) FilterDistributedBTC(opts *bind.FilterOpts, to []common.Address) (*StakingPoolProxyDistributedBTCIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.FilterLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyDistributedBTCIterator{contract: _StakingPoolProxy.contract, event: "DistributedBTC", logs: logs, sub: sub}, nil
}

// WatchDistributedBTC is a free log subscription operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) WatchDistributedBTC(opts *bind.WatchOpts, sink chan<- *StakingPoolProxyDistributedBTC, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.WatchLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolProxyDistributedBTC)
				if err := _StakingPoolProxy.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
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

// ParseDistributedBTC is a log parse operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) ParseDistributedBTC(log types.Log) (*StakingPoolProxyDistributedBTC, error) {
	event := new(StakingPoolProxyDistributedBTC)
	if err := _StakingPoolProxy.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolProxyImplementationUpgradedIterator is returned from FilterImplementationUpgraded and is used to iterate over the raw logs and unpacked data for ImplementationUpgraded events raised by the StakingPoolProxy contract.
type StakingPoolProxyImplementationUpgradedIterator struct {
	Event *StakingPoolProxyImplementationUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingPoolProxyImplementationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolProxyImplementationUpgraded)
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
		it.Event = new(StakingPoolProxyImplementationUpgraded)
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
func (it *StakingPoolProxyImplementationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolProxyImplementationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolProxyImplementationUpgraded represents a ImplementationUpgraded event raised by the StakingPoolProxy contract.
type StakingPoolProxyImplementationUpgraded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpgraded is a free log retrieval operation binding the contract event 0x3dab95f89f74034beeac9f9fe9bec350186cb958223e1791afb0dbb64859cabd.
//
// Solidity: event ImplementationUpgraded()
func (_StakingPoolProxy *StakingPoolProxyFilterer) FilterImplementationUpgraded(opts *bind.FilterOpts) (*StakingPoolProxyImplementationUpgradedIterator, error) {

	logs, sub, err := _StakingPoolProxy.contract.FilterLogs(opts, "ImplementationUpgraded")
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyImplementationUpgradedIterator{contract: _StakingPoolProxy.contract, event: "ImplementationUpgraded", logs: logs, sub: sub}, nil
}

// WatchImplementationUpgraded is a free log subscription operation binding the contract event 0x3dab95f89f74034beeac9f9fe9bec350186cb958223e1791afb0dbb64859cabd.
//
// Solidity: event ImplementationUpgraded()
func (_StakingPoolProxy *StakingPoolProxyFilterer) WatchImplementationUpgraded(opts *bind.WatchOpts, sink chan<- *StakingPoolProxyImplementationUpgraded) (event.Subscription, error) {

	logs, sub, err := _StakingPoolProxy.contract.WatchLogs(opts, "ImplementationUpgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolProxyImplementationUpgraded)
				if err := _StakingPoolProxy.contract.UnpackLog(event, "ImplementationUpgraded", log); err != nil {
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

// ParseImplementationUpgraded is a log parse operation binding the contract event 0x3dab95f89f74034beeac9f9fe9bec350186cb958223e1791afb0dbb64859cabd.
//
// Solidity: event ImplementationUpgraded()
func (_StakingPoolProxy *StakingPoolProxyFilterer) ParseImplementationUpgraded(log types.Log) (*StakingPoolProxyImplementationUpgraded, error) {
	event := new(StakingPoolProxyImplementationUpgraded)
	if err := _StakingPoolProxy.contract.UnpackLog(event, "ImplementationUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingPoolProxy contract.
type StakingPoolProxyOwnershipTransferredIterator struct {
	Event *StakingPoolProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingPoolProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolProxyOwnershipTransferred)
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
		it.Event = new(StakingPoolProxyOwnershipTransferred)
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
func (it *StakingPoolProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolProxyOwnershipTransferred represents a OwnershipTransferred event raised by the StakingPoolProxy contract.
type StakingPoolProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolProxy *StakingPoolProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingPoolProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyOwnershipTransferredIterator{contract: _StakingPoolProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolProxy *StakingPoolProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingPoolProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolProxyOwnershipTransferred)
				if err := _StakingPoolProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingPoolProxy *StakingPoolProxyFilterer) ParseOwnershipTransferred(log types.Log) (*StakingPoolProxyOwnershipTransferred, error) {
	event := new(StakingPoolProxyOwnershipTransferred)
	if err := _StakingPoolProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolProxyReceivedTokensIterator is returned from FilterReceivedTokens and is used to iterate over the raw logs and unpacked data for ReceivedTokens events raised by the StakingPoolProxy contract.
type StakingPoolProxyReceivedTokensIterator struct {
	Event *StakingPoolProxyReceivedTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolProxyReceivedTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolProxyReceivedTokens)
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
		it.Event = new(StakingPoolProxyReceivedTokens)
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
func (it *StakingPoolProxyReceivedTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolProxyReceivedTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolProxyReceivedTokens represents a ReceivedTokens event raised by the StakingPoolProxy contract.
type StakingPoolProxyReceivedTokens struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedTokens is a free log retrieval operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) FilterReceivedTokens(opts *bind.FilterOpts, from []common.Address) (*StakingPoolProxyReceivedTokensIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.FilterLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyReceivedTokensIterator{contract: _StakingPoolProxy.contract, event: "ReceivedTokens", logs: logs, sub: sub}, nil
}

// WatchReceivedTokens is a free log subscription operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) WatchReceivedTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolProxyReceivedTokens, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.WatchLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolProxyReceivedTokens)
				if err := _StakingPoolProxy.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
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

// ParseReceivedTokens is a log parse operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) ParseReceivedTokens(log types.Log) (*StakingPoolProxyReceivedTokens, error) {
	event := new(StakingPoolProxyReceivedTokens)
	if err := _StakingPoolProxy.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolProxyTransferTokensIterator is returned from FilterTransferTokens and is used to iterate over the raw logs and unpacked data for TransferTokens events raised by the StakingPoolProxy contract.
type StakingPoolProxyTransferTokensIterator struct {
	Event *StakingPoolProxyTransferTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolProxyTransferTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolProxyTransferTokens)
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
		it.Event = new(StakingPoolProxyTransferTokens)
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
func (it *StakingPoolProxyTransferTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolProxyTransferTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolProxyTransferTokens represents a TransferTokens event raised by the StakingPoolProxy contract.
type StakingPoolProxyTransferTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferTokens is a free log retrieval operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) FilterTransferTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolProxyTransferTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.FilterLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolProxyTransferTokensIterator{contract: _StakingPoolProxy.contract, event: "TransferTokens", logs: logs, sub: sub}, nil
}

// WatchTransferTokens is a free log subscription operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) WatchTransferTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolProxyTransferTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolProxy.contract.WatchLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolProxyTransferTokens)
				if err := _StakingPoolProxy.contract.UnpackLog(event, "TransferTokens", log); err != nil {
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

// ParseTransferTokens is a log parse operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolProxy *StakingPoolProxyFilterer) ParseTransferTokens(log types.Log) (*StakingPoolProxyTransferTokens, error) {
	event := new(StakingPoolProxyTransferTokens)
	if err := _StakingPoolProxy.contract.UnpackLog(event, "TransferTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
