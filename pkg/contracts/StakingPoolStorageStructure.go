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

// StakingPoolStorageStructureMetaData contains all meta data concerning the StakingPoolStorageStructure contract.
var StakingPoolStorageStructureMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApproveTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedBTC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferTokens\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapRouterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BUSDContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WrappedTokenContractAddress\",\"type\":\"address\"}],\"name\":\"__WrappedTokenDistributor_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averagePricePrediction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getEstimatedWrappedTokenForUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPathForUSDToWrappedToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAnEmergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnhancedEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMatured\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"launchDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPrizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeApr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collaborativeReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEnhancedEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMatured\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialCollabReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"predictions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastWithdrawalTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePrediction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difference\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"prizeRewardWithdrawn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didUnstake\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prizeRewardRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeLimitRangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sortedStakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeTaxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPoolImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractITotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingPoolStorageStructureABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolStorageStructureMetaData.ABI instead.
var StakingPoolStorageStructureABI = StakingPoolStorageStructureMetaData.ABI

// StakingPoolStorageStructure is an auto generated Go binding around an Ethereum contract.
type StakingPoolStorageStructure struct {
	StakingPoolStorageStructureCaller     // Read-only binding to the contract
	StakingPoolStorageStructureTransactor // Write-only binding to the contract
	StakingPoolStorageStructureFilterer   // Log filterer for contract events
}

// StakingPoolStorageStructureCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolStorageStructureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolStorageStructureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolStorageStructureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolStorageStructureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolStorageStructureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolStorageStructureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolStorageStructureSession struct {
	Contract     *StakingPoolStorageStructure // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StakingPoolStorageStructureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolStorageStructureCallerSession struct {
	Contract *StakingPoolStorageStructureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// StakingPoolStorageStructureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolStorageStructureTransactorSession struct {
	Contract     *StakingPoolStorageStructureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// StakingPoolStorageStructureRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolStorageStructureRaw struct {
	Contract *StakingPoolStorageStructure // Generic contract binding to access the raw methods on
}

// StakingPoolStorageStructureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolStorageStructureCallerRaw struct {
	Contract *StakingPoolStorageStructureCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolStorageStructureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolStorageStructureTransactorRaw struct {
	Contract *StakingPoolStorageStructureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPoolStorageStructure creates a new instance of StakingPoolStorageStructure, bound to a specific deployed contract.
func NewStakingPoolStorageStructure(address common.Address, backend bind.ContractBackend) (*StakingPoolStorageStructure, error) {
	contract, err := bindStakingPoolStorageStructure(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructure{StakingPoolStorageStructureCaller: StakingPoolStorageStructureCaller{contract: contract}, StakingPoolStorageStructureTransactor: StakingPoolStorageStructureTransactor{contract: contract}, StakingPoolStorageStructureFilterer: StakingPoolStorageStructureFilterer{contract: contract}}, nil
}

// NewStakingPoolStorageStructureCaller creates a new read-only instance of StakingPoolStorageStructure, bound to a specific deployed contract.
func NewStakingPoolStorageStructureCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolStorageStructureCaller, error) {
	contract, err := bindStakingPoolStorageStructure(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureCaller{contract: contract}, nil
}

// NewStakingPoolStorageStructureTransactor creates a new write-only instance of StakingPoolStorageStructure, bound to a specific deployed contract.
func NewStakingPoolStorageStructureTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolStorageStructureTransactor, error) {
	contract, err := bindStakingPoolStorageStructure(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureTransactor{contract: contract}, nil
}

// NewStakingPoolStorageStructureFilterer creates a new log filterer instance of StakingPoolStorageStructure, bound to a specific deployed contract.
func NewStakingPoolStorageStructureFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolStorageStructureFilterer, error) {
	contract, err := bindStakingPoolStorageStructure(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureFilterer{contract: contract}, nil
}

// bindStakingPoolStorageStructure binds a generic wrapper to an already deployed contract.
func bindStakingPoolStorageStructure(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolStorageStructureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolStorageStructure *StakingPoolStorageStructureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolStorageStructure.Contract.StakingPoolStorageStructureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolStorageStructure *StakingPoolStorageStructureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.StakingPoolStorageStructureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolStorageStructure *StakingPoolStorageStructureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.StakingPoolStorageStructureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolStorageStructure.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.contract.Transact(opts, method, params...)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) AveragePricePrediction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "averagePricePrediction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.AveragePricePrediction(&_StakingPoolStorageStructure.CallOpts)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.AveragePricePrediction(&_StakingPoolStorageStructure.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) CollaborativeRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "collaborativeRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.CollaborativeRange(&_StakingPoolStorageStructure.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.CollaborativeRange(&_StakingPoolStorageStructure.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) CollaborativeReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "collaborativeReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.CollaborativeReward(&_StakingPoolStorageStructure.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.CollaborativeReward(&_StakingPoolStorageStructure.CallOpts)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) GetEstimatedWrappedTokenForUSD(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "getEstimatedWrappedTokenForUSD", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolStorageStructure.CallOpts, _amount)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolStorageStructure.CallOpts, _amount)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) GetPathForUSDToWrappedToken(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "getPathForUSDToWrappedToken")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolStorageStructure.Contract.GetPathForUSDToWrappedToken(&_StakingPoolStorageStructure.CallOpts)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolStorageStructure.Contract.GetPathForUSDToWrappedToken(&_StakingPoolStorageStructure.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) GetSwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "getSwapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.GetSwapRouter(&_StakingPoolStorageStructure.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.GetSwapRouter(&_StakingPoolStorageStructure.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) GetUSDBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "getUSDBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.GetUSDBalance(&_StakingPoolStorageStructure.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.GetUSDBalance(&_StakingPoolStorageStructure.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) GetUSDToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "getUSDToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) GetUSDToken() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.GetUSDToken(&_StakingPoolStorageStructure.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) GetUSDToken() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.GetUSDToken(&_StakingPoolStorageStructure.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) IsActive() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsActive(&_StakingPoolStorageStructure.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) IsActive() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsActive(&_StakingPoolStorageStructure.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) IsAnEmergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "isAnEmergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) IsAnEmergency() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsAnEmergency(&_StakingPoolStorageStructure.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) IsAnEmergency() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsAnEmergency(&_StakingPoolStorageStructure.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) IsDeleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "isDeleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) IsDeleted() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsDeleted(&_StakingPoolStorageStructure.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) IsDeleted() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsDeleted(&_StakingPoolStorageStructure.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) IsEnhancedEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "isEnhancedEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsEnhancedEnabled(&_StakingPoolStorageStructure.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsEnhancedEnabled(&_StakingPoolStorageStructure.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) IsLocked() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsLocked(&_StakingPoolStorageStructure.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) IsLocked() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsLocked(&_StakingPoolStorageStructure.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) IsMatured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "isMatured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) IsMatured() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsMatured(&_StakingPoolStorageStructure.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) IsMatured() (bool, error) {
	return _StakingPoolStorageStructure.Contract.IsMatured(&_StakingPoolStorageStructure.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) LaunchDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "launchDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) LaunchDate() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.LaunchDate(&_StakingPoolStorageStructure.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) LaunchDate() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.LaunchDate(&_StakingPoolStorageStructure.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) LockTime() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.LockTime(&_StakingPoolStorageStructure.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) LockTime() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.LockTime(&_StakingPoolStorageStructure.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) Lps(opts *bind.CallOpts) (struct {
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
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "lps")

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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) Lps() (struct {
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
	return _StakingPoolStorageStructure.Contract.Lps(&_StakingPoolStorageStructure.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) Lps() (struct {
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
	return _StakingPoolStorageStructure.Contract.Lps(&_StakingPoolStorageStructure.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) MaturingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "maturingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.MaturingPrice(&_StakingPoolStorageStructure.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.MaturingPrice(&_StakingPoolStorageStructure.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) MaturityTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "maturityTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) MaturityTime() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.MaturityTime(&_StakingPoolStorageStructure.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) MaturityTime() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.MaturityTime(&_StakingPoolStorageStructure.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) MinimumStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "minimumStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.MinimumStakeAmount(&_StakingPoolStorageStructure.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.MinimumStakeAmount(&_StakingPoolStorageStructure.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) OracleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "oracleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) OracleContract() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.OracleContract(&_StakingPoolStorageStructure.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) OracleContract() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.OracleContract(&_StakingPoolStorageStructure.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) OracleDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "oracleDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.OracleDecimals(&_StakingPoolStorageStructure.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.OracleDecimals(&_StakingPoolStorageStructure.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) Owner() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.Owner(&_StakingPoolStorageStructure.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) Owner() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.Owner(&_StakingPoolStorageStructure.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) PoolCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "poolCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) PoolCreator() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.PoolCreator(&_StakingPoolStorageStructure.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) PoolCreator() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.PoolCreator(&_StakingPoolStorageStructure.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) PoolType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) PoolType() (string, error) {
	return _StakingPoolStorageStructure.Contract.PoolType(&_StakingPoolStorageStructure.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) PoolType() (string, error) {
	return _StakingPoolStorageStructure.Contract.PoolType(&_StakingPoolStorageStructure.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) PotentialCollabReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "potentialCollabReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.PotentialCollabReward(&_StakingPoolStorageStructure.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.PotentialCollabReward(&_StakingPoolStorageStructure.CallOpts)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) Predictions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
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
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "predictions", arg0, arg1)

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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
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
	return _StakingPoolStorageStructure.Contract.Predictions(&_StakingPoolStorageStructure.CallOpts, arg0, arg1)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
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
	return _StakingPoolStorageStructure.Contract.Predictions(&_StakingPoolStorageStructure.CallOpts, arg0, arg1)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) PrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "prizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.PrizeAmount(&_StakingPoolStorageStructure.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.PrizeAmount(&_StakingPoolStorageStructure.CallOpts)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) PrizeRewardRates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "prizeRewardRates", arg0)

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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolStorageStructure.Contract.PrizeRewardRates(&_StakingPoolStorageStructure.CallOpts, arg0)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolStorageStructure.Contract.PrizeRewardRates(&_StakingPoolStorageStructure.CallOpts, arg0)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) RewardManager() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.RewardManager(&_StakingPoolStorageStructure.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) RewardManager() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.RewardManager(&_StakingPoolStorageStructure.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) SizeAllocation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "sizeAllocation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.SizeAllocation(&_StakingPoolStorageStructure.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.SizeAllocation(&_StakingPoolStorageStructure.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) SizeLimitRangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "sizeLimitRangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.SizeLimitRangeRate(&_StakingPoolStorageStructure.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.SizeLimitRangeRate(&_StakingPoolStorageStructure.CallOpts)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) SortedStakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "sortedStakers", arg0)

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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolStorageStructure.Contract.SortedStakers(&_StakingPoolStorageStructure.CallOpts, arg0)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolStorageStructure.Contract.SortedStakers(&_StakingPoolStorageStructure.CallOpts, arg0)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) StakeApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "stakeApr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) StakeApr() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.StakeApr(&_StakingPoolStorageStructure.CallOpts)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) StakeApr() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.StakeApr(&_StakingPoolStorageStructure.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) StakeTaxRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "stakeTaxRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.StakeTaxRate(&_StakingPoolStorageStructure.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.StakeTaxRate(&_StakingPoolStorageStructure.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "stakers", arg0)

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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolStorageStructure.Contract.Stakers(&_StakingPoolStorageStructure.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolStorageStructure.Contract.Stakers(&_StakingPoolStorageStructure.CallOpts, arg0)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) StakingPoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "stakingPoolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.StakingPoolImplementation(&_StakingPoolStorageStructure.CallOpts)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.StakingPoolImplementation(&_StakingPoolStorageStructure.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) TotalStaked() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.TotalStaked(&_StakingPoolStorageStructure.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) TotalStaked() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.TotalStaked(&_StakingPoolStorageStructure.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) TotemToken() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.TotemToken(&_StakingPoolStorageStructure.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) TotemToken() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.TotemToken(&_StakingPoolStorageStructure.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) UpgradeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "upgradeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) UpgradeEnabled() (bool, error) {
	return _StakingPoolStorageStructure.Contract.UpgradeEnabled(&_StakingPoolStorageStructure.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) UpgradeEnabled() (bool, error) {
	return _StakingPoolStorageStructure.Contract.UpgradeEnabled(&_StakingPoolStorageStructure.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) UsdPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "usdPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.UsdPrizeAmount(&_StakingPoolStorageStructure.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolStorageStructure.Contract.UsdPrizeAmount(&_StakingPoolStorageStructure.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) WrappedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "wrappedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) WrappedToken() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.WrappedToken(&_StakingPoolStorageStructure.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) WrappedToken() (common.Address, error) {
	return _StakingPoolStorageStructure.Contract.WrappedToken(&_StakingPoolStorageStructure.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCaller) WrappedTokenSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolStorageStructure.contract.Call(opts, &out, "wrappedTokenSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolStorageStructure.Contract.WrappedTokenSymbol(&_StakingPoolStorageStructure.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureCallerSession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolStorageStructure.Contract.WrappedTokenSymbol(&_StakingPoolStorageStructure.CallOpts)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactor) WrappedTokenDistributorInitialize(opts *bind.TransactOpts, swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.contract.Transact(opts, "__WrappedTokenDistributor_initialize", swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.WrappedTokenDistributorInitialize(&_StakingPoolStorageStructure.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactorSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.WrappedTokenDistributorInitialize(&_StakingPoolStorageStructure.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.RenounceOwnership(&_StakingPoolStorageStructure.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.RenounceOwnership(&_StakingPoolStorageStructure.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.TransferOwnership(&_StakingPoolStorageStructure.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolStorageStructure *StakingPoolStorageStructureTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolStorageStructure.Contract.TransferOwnership(&_StakingPoolStorageStructure.TransactOpts, newOwner)
}

// StakingPoolStorageStructureApproveTokensIterator is returned from FilterApproveTokens and is used to iterate over the raw logs and unpacked data for ApproveTokens events raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureApproveTokensIterator struct {
	Event *StakingPoolStorageStructureApproveTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolStorageStructureApproveTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolStorageStructureApproveTokens)
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
		it.Event = new(StakingPoolStorageStructureApproveTokens)
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
func (it *StakingPoolStorageStructureApproveTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolStorageStructureApproveTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolStorageStructureApproveTokens represents a ApproveTokens event raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureApproveTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApproveTokens is a free log retrieval operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) FilterApproveTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolStorageStructureApproveTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.FilterLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureApproveTokensIterator{contract: _StakingPoolStorageStructure.contract, event: "ApproveTokens", logs: logs, sub: sub}, nil
}

// WatchApproveTokens is a free log subscription operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) WatchApproveTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolStorageStructureApproveTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.WatchLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolStorageStructureApproveTokens)
				if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) ParseApproveTokens(log types.Log) (*StakingPoolStorageStructureApproveTokens, error) {
	event := new(StakingPoolStorageStructureApproveTokens)
	if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolStorageStructureDistributedBTCIterator is returned from FilterDistributedBTC and is used to iterate over the raw logs and unpacked data for DistributedBTC events raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureDistributedBTCIterator struct {
	Event *StakingPoolStorageStructureDistributedBTC // Event containing the contract specifics and raw log

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
func (it *StakingPoolStorageStructureDistributedBTCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolStorageStructureDistributedBTC)
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
		it.Event = new(StakingPoolStorageStructureDistributedBTC)
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
func (it *StakingPoolStorageStructureDistributedBTCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolStorageStructureDistributedBTCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolStorageStructureDistributedBTC represents a DistributedBTC event raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureDistributedBTC struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedBTC is a free log retrieval operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) FilterDistributedBTC(opts *bind.FilterOpts, to []common.Address) (*StakingPoolStorageStructureDistributedBTCIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.FilterLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureDistributedBTCIterator{contract: _StakingPoolStorageStructure.contract, event: "DistributedBTC", logs: logs, sub: sub}, nil
}

// WatchDistributedBTC is a free log subscription operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) WatchDistributedBTC(opts *bind.WatchOpts, sink chan<- *StakingPoolStorageStructureDistributedBTC, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.WatchLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolStorageStructureDistributedBTC)
				if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) ParseDistributedBTC(log types.Log) (*StakingPoolStorageStructureDistributedBTC, error) {
	event := new(StakingPoolStorageStructureDistributedBTC)
	if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolStorageStructureOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureOwnershipTransferredIterator struct {
	Event *StakingPoolStorageStructureOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingPoolStorageStructureOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolStorageStructureOwnershipTransferred)
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
		it.Event = new(StakingPoolStorageStructureOwnershipTransferred)
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
func (it *StakingPoolStorageStructureOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolStorageStructureOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolStorageStructureOwnershipTransferred represents a OwnershipTransferred event raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingPoolStorageStructureOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureOwnershipTransferredIterator{contract: _StakingPoolStorageStructure.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingPoolStorageStructureOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolStorageStructureOwnershipTransferred)
				if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) ParseOwnershipTransferred(log types.Log) (*StakingPoolStorageStructureOwnershipTransferred, error) {
	event := new(StakingPoolStorageStructureOwnershipTransferred)
	if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolStorageStructureReceivedTokensIterator is returned from FilterReceivedTokens and is used to iterate over the raw logs and unpacked data for ReceivedTokens events raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureReceivedTokensIterator struct {
	Event *StakingPoolStorageStructureReceivedTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolStorageStructureReceivedTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolStorageStructureReceivedTokens)
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
		it.Event = new(StakingPoolStorageStructureReceivedTokens)
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
func (it *StakingPoolStorageStructureReceivedTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolStorageStructureReceivedTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolStorageStructureReceivedTokens represents a ReceivedTokens event raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureReceivedTokens struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedTokens is a free log retrieval operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) FilterReceivedTokens(opts *bind.FilterOpts, from []common.Address) (*StakingPoolStorageStructureReceivedTokensIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.FilterLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureReceivedTokensIterator{contract: _StakingPoolStorageStructure.contract, event: "ReceivedTokens", logs: logs, sub: sub}, nil
}

// WatchReceivedTokens is a free log subscription operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) WatchReceivedTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolStorageStructureReceivedTokens, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.WatchLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolStorageStructureReceivedTokens)
				if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) ParseReceivedTokens(log types.Log) (*StakingPoolStorageStructureReceivedTokens, error) {
	event := new(StakingPoolStorageStructureReceivedTokens)
	if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolStorageStructureTransferTokensIterator is returned from FilterTransferTokens and is used to iterate over the raw logs and unpacked data for TransferTokens events raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureTransferTokensIterator struct {
	Event *StakingPoolStorageStructureTransferTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolStorageStructureTransferTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolStorageStructureTransferTokens)
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
		it.Event = new(StakingPoolStorageStructureTransferTokens)
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
func (it *StakingPoolStorageStructureTransferTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolStorageStructureTransferTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolStorageStructureTransferTokens represents a TransferTokens event raised by the StakingPoolStorageStructure contract.
type StakingPoolStorageStructureTransferTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferTokens is a free log retrieval operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) FilterTransferTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolStorageStructureTransferTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.FilterLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolStorageStructureTransferTokensIterator{contract: _StakingPoolStorageStructure.contract, event: "TransferTokens", logs: logs, sub: sub}, nil
}

// WatchTransferTokens is a free log subscription operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) WatchTransferTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolStorageStructureTransferTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolStorageStructure.contract.WatchLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolStorageStructureTransferTokens)
				if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "TransferTokens", log); err != nil {
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
func (_StakingPoolStorageStructure *StakingPoolStorageStructureFilterer) ParseTransferTokens(log types.Log) (*StakingPoolStorageStructureTransferTokens, error) {
	event := new(StakingPoolStorageStructureTransferTokens)
	if err := _StakingPoolStorageStructure.contract.UnpackLog(event, "TransferTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
