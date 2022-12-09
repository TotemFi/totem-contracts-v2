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

// StakingPoolImplementationMetaData contains all meta data concerning the StakingPoolImplementation contract.
var StakingPoolImplementationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApproveTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"stakeAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"predictions\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexOfFirstStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexOfLastStake\",\"type\":\"uint256\"}],\"name\":\"BatchStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedBTC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolMatured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolSorted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePrediction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexOfStake\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingReturn\",\"type\":\"uint256\"}],\"name\":\"WithdrawStakingReturn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totemPrize\",\"type\":\"uint256\"}],\"name\":\"WithdrawTotemPrize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrappedTokenPrize\",\"type\":\"uint256\"}],\"name\":\"WithdrawWrappedTokenPrize\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapRouterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BUSDContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WrappedTokenContractAddress\",\"type\":\"address\"}],\"name\":\"__WrappedTokenDistributor_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averagePricePrediction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_predictions\",\"type\":\"uint256[]\"}],\"name\":\"batchStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"declareEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deletePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergentWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prediction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_range\",\"type\":\"uint256\"}],\"name\":\"getDifference\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getEstimatedWrappedTokenForUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeIndex\",\"type\":\"uint256\"}],\"name\":\"getIndexedPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeIndex\",\"type\":\"uint256\"}],\"name\":\"getIndexedStakingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPathForUSDToWrappedToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleContract\",\"type\":\"address\"}],\"name\":\"getPrizeTokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenTaxRate\",\"type\":\"uint256\"}],\"name\":\"getStakingTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"}],\"name\":\"hasUnStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"}],\"name\":\"indexedClaimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAnEmergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnhancedEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMatured\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"launchDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPrizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeApr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collaborativeReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEnhancedEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMatured\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialCollabReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"predictions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastWithdrawalTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePrediction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difference\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"prizeRewardWithdrawn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didUnstake\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prizeRewardRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"purchaseWrappedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_activationStatus\",\"type\":\"bool\"}],\"name\":\"setActivationStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setOracleToZero\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[25]\",\"name\":\"addrArray\",\"type\":\"address[25]\"},{\"internalType\":\"uint256[25]\",\"name\":\"indexArray\",\"type\":\"uint256[25]\"}],\"name\":\"setSortedStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeLimitRangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sortedStakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePrediction\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeTaxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPoolImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractITotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_predictionPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_samePredictionPrizeToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_prizePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_oracleContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_oracleDecimals\",\"type\":\"uint256\"}],\"name\":\"updateMaturingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stuckToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdrawStuckTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingPoolImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolImplementationMetaData.ABI instead.
var StakingPoolImplementationABI = StakingPoolImplementationMetaData.ABI

// StakingPoolImplementation is an auto generated Go binding around an Ethereum contract.
type StakingPoolImplementation struct {
	StakingPoolImplementationCaller     // Read-only binding to the contract
	StakingPoolImplementationTransactor // Write-only binding to the contract
	StakingPoolImplementationFilterer   // Log filterer for contract events
}

// StakingPoolImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolImplementationSession struct {
	Contract     *StakingPoolImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StakingPoolImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolImplementationCallerSession struct {
	Contract *StakingPoolImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// StakingPoolImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolImplementationTransactorSession struct {
	Contract     *StakingPoolImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// StakingPoolImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolImplementationRaw struct {
	Contract *StakingPoolImplementation // Generic contract binding to access the raw methods on
}

// StakingPoolImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolImplementationCallerRaw struct {
	Contract *StakingPoolImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolImplementationTransactorRaw struct {
	Contract *StakingPoolImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPoolImplementation creates a new instance of StakingPoolImplementation, bound to a specific deployed contract.
func NewStakingPoolImplementation(address common.Address, backend bind.ContractBackend) (*StakingPoolImplementation, error) {
	contract, err := bindStakingPoolImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementation{StakingPoolImplementationCaller: StakingPoolImplementationCaller{contract: contract}, StakingPoolImplementationTransactor: StakingPoolImplementationTransactor{contract: contract}, StakingPoolImplementationFilterer: StakingPoolImplementationFilterer{contract: contract}}, nil
}

// NewStakingPoolImplementationCaller creates a new read-only instance of StakingPoolImplementation, bound to a specific deployed contract.
func NewStakingPoolImplementationCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolImplementationCaller, error) {
	contract, err := bindStakingPoolImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationCaller{contract: contract}, nil
}

// NewStakingPoolImplementationTransactor creates a new write-only instance of StakingPoolImplementation, bound to a specific deployed contract.
func NewStakingPoolImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolImplementationTransactor, error) {
	contract, err := bindStakingPoolImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationTransactor{contract: contract}, nil
}

// NewStakingPoolImplementationFilterer creates a new log filterer instance of StakingPoolImplementation, bound to a specific deployed contract.
func NewStakingPoolImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolImplementationFilterer, error) {
	contract, err := bindStakingPoolImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationFilterer{contract: contract}, nil
}

// bindStakingPoolImplementation binds a generic wrapper to an already deployed contract.
func bindStakingPoolImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolImplementation *StakingPoolImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolImplementation.Contract.StakingPoolImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolImplementation *StakingPoolImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.StakingPoolImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolImplementation *StakingPoolImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.StakingPoolImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolImplementation *StakingPoolImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolImplementation *StakingPoolImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolImplementation *StakingPoolImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.contract.Transact(opts, method, params...)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) AveragePricePrediction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "averagePricePrediction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.AveragePricePrediction(&_StakingPoolImplementation.CallOpts)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.AveragePricePrediction(&_StakingPoolImplementation.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) CollaborativeRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "collaborativeRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.CollaborativeRange(&_StakingPoolImplementation.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.CollaborativeRange(&_StakingPoolImplementation.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) CollaborativeReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "collaborativeReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.CollaborativeReward(&_StakingPoolImplementation.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.CollaborativeReward(&_StakingPoolImplementation.CallOpts)
}

// GetDifference is a free data retrieval call binding the contract method 0x4c7c77fc.
//
// Solidity: function getDifference(uint256 prediction, uint256 _range) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetDifference(opts *bind.CallOpts, prediction *big.Int, _range *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getDifference", prediction, _range)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDifference is a free data retrieval call binding the contract method 0x4c7c77fc.
//
// Solidity: function getDifference(uint256 prediction, uint256 _range) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetDifference(prediction *big.Int, _range *big.Int) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetDifference(&_StakingPoolImplementation.CallOpts, prediction, _range)
}

// GetDifference is a free data retrieval call binding the contract method 0x4c7c77fc.
//
// Solidity: function getDifference(uint256 prediction, uint256 _range) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetDifference(prediction *big.Int, _range *big.Int) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetDifference(&_StakingPoolImplementation.CallOpts, prediction, _range)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetEstimatedWrappedTokenForUSD(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getEstimatedWrappedTokenForUSD", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolImplementation.CallOpts, _amount)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolImplementation.CallOpts, _amount)
}

// GetIndexedPrize is a free data retrieval call binding the contract method 0xb7b30170.
//
// Solidity: function getIndexedPrize(address _staker, uint256 _stakeIndex) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetIndexedPrize(opts *bind.CallOpts, _staker common.Address, _stakeIndex *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getIndexedPrize", _staker, _stakeIndex)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetIndexedPrize is a free data retrieval call binding the contract method 0xb7b30170.
//
// Solidity: function getIndexedPrize(address _staker, uint256 _stakeIndex) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetIndexedPrize(_staker common.Address, _stakeIndex *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetIndexedPrize(&_StakingPoolImplementation.CallOpts, _staker, _stakeIndex)
}

// GetIndexedPrize is a free data retrieval call binding the contract method 0xb7b30170.
//
// Solidity: function getIndexedPrize(address _staker, uint256 _stakeIndex) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetIndexedPrize(_staker common.Address, _stakeIndex *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetIndexedPrize(&_StakingPoolImplementation.CallOpts, _staker, _stakeIndex)
}

// GetIndexedStakingRewardA is a free data retrieval call binding the contract method 0x2342e480.
//
// Solidity: function getIndexedStakingReward(address _staker, uint256 _stakeIndex) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetIndexedStakingRewardA(opts *bind.CallOpts, _staker common.Address, _stakeIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getIndexedStakingReward", _staker, _stakeIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexedStakingRewardA is a free data retrieval call binding the contract method 0x2342e480.
//
// Solidity: function getIndexedStakingReward(address _staker, uint256 _stakeIndex) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetIndexedStakingRewardA(_staker common.Address, _stakeIndex *big.Int) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetIndexedStakingRewardA(&_StakingPoolImplementation.CallOpts, _staker, _stakeIndex)
}

// GetIndexedStakingRewardA is a free data retrieval call binding the contract method 0x2342e480.
//
// Solidity: function getIndexedStakingReward(address _staker, uint256 _stakeIndex) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetIndexedStakingRewardA(_staker common.Address, _stakeIndex *big.Int) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetIndexedStakingRewardA(&_StakingPoolImplementation.CallOpts, _staker, _stakeIndex)
}

// GetLPS is a free data retrieval call binding the contract method 0x9e3dd80f.
//
// Solidity: function getLPS() view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetLPS(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getLPS")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLPS is a free data retrieval call binding the contract method 0x9e3dd80f.
//
// Solidity: function getLPS() view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetLPS() (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetLPS(&_StakingPoolImplementation.CallOpts)
}

// GetLPS is a free data retrieval call binding the contract method 0x9e3dd80f.
//
// Solidity: function getLPS() view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetLPS() (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetLPS(&_StakingPoolImplementation.CallOpts)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetPathForUSDToWrappedToken(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getPathForUSDToWrappedToken")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolImplementation.Contract.GetPathForUSDToWrappedToken(&_StakingPoolImplementation.CallOpts)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolImplementation.Contract.GetPathForUSDToWrappedToken(&_StakingPoolImplementation.CallOpts)
}

// GetPrize is a free data retrieval call binding the contract method 0x2d5cd1d6.
//
// Solidity: function getPrize(address _staker) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetPrize(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getPrize", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPrize is a free data retrieval call binding the contract method 0x2d5cd1d6.
//
// Solidity: function getPrize(address _staker) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetPrize(_staker common.Address) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetPrize(&_StakingPoolImplementation.CallOpts, _staker)
}

// GetPrize is a free data retrieval call binding the contract method 0x2d5cd1d6.
//
// Solidity: function getPrize(address _staker) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetPrize(_staker common.Address) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetPrize(&_StakingPoolImplementation.CallOpts, _staker)
}

// GetPrizeTokenDecimals is a free data retrieval call binding the contract method 0xcee0c889.
//
// Solidity: function getPrizeTokenDecimals(address _oracleContract) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetPrizeTokenDecimals(opts *bind.CallOpts, _oracleContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getPrizeTokenDecimals", _oracleContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrizeTokenDecimals is a free data retrieval call binding the contract method 0xcee0c889.
//
// Solidity: function getPrizeTokenDecimals(address _oracleContract) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetPrizeTokenDecimals(_oracleContract common.Address) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetPrizeTokenDecimals(&_StakingPoolImplementation.CallOpts, _oracleContract)
}

// GetPrizeTokenDecimals is a free data retrieval call binding the contract method 0xcee0c889.
//
// Solidity: function getPrizeTokenDecimals(address _oracleContract) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetPrizeTokenDecimals(_oracleContract common.Address) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetPrizeTokenDecimals(&_StakingPoolImplementation.CallOpts, _oracleContract)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[], uint256[])
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetStakers(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getStakers")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[], uint256[])
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetStakers() ([]common.Address, []*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetStakers(&_StakingPoolImplementation.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[], uint256[])
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetStakers() ([]common.Address, []*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetStakers(&_StakingPoolImplementation.CallOpts)
}

// GetStakingRewardA is a free data retrieval call binding the contract method 0x78f4d413.
//
// Solidity: function getStakingReward(address _staker) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetStakingRewardA(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getStakingReward", _staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingRewardA is a free data retrieval call binding the contract method 0x78f4d413.
//
// Solidity: function getStakingReward(address _staker) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetStakingRewardA(_staker common.Address) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetStakingRewardA(&_StakingPoolImplementation.CallOpts, _staker)
}

// GetStakingRewardA is a free data retrieval call binding the contract method 0x78f4d413.
//
// Solidity: function getStakingReward(address _staker) view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetStakingRewardA(_staker common.Address) (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetStakingRewardA(&_StakingPoolImplementation.CallOpts, _staker)
}

// GetStakingTax is a free data retrieval call binding the contract method 0x395b3dac.
//
// Solidity: function getStakingTax(uint256 amount, uint256 tokenTaxRate) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetStakingTax(opts *bind.CallOpts, amount *big.Int, tokenTaxRate *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getStakingTax", amount, tokenTaxRate)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakingTax is a free data retrieval call binding the contract method 0x395b3dac.
//
// Solidity: function getStakingTax(uint256 amount, uint256 tokenTaxRate) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetStakingTax(amount *big.Int, tokenTaxRate *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetStakingTax(&_StakingPoolImplementation.CallOpts, amount, tokenTaxRate)
}

// GetStakingTax is a free data retrieval call binding the contract method 0x395b3dac.
//
// Solidity: function getStakingTax(uint256 amount, uint256 tokenTaxRate) view returns(uint256, uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetStakingTax(amount *big.Int, tokenTaxRate *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementation.Contract.GetStakingTax(&_StakingPoolImplementation.CallOpts, amount, tokenTaxRate)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetSwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getSwapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolImplementation.Contract.GetSwapRouter(&_StakingPoolImplementation.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolImplementation.Contract.GetSwapRouter(&_StakingPoolImplementation.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetUSDBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getUSDBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetUSDBalance(&_StakingPoolImplementation.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetUSDBalance(&_StakingPoolImplementation.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetUSDToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getUSDToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetUSDToken() (common.Address, error) {
	return _StakingPoolImplementation.Contract.GetUSDToken(&_StakingPoolImplementation.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetUSDToken() (common.Address, error) {
	return _StakingPoolImplementation.Contract.GetUSDToken(&_StakingPoolImplementation.CallOpts)
}

// GetWrappedTokenBalance is a free data retrieval call binding the contract method 0xf8e829c3.
//
// Solidity: function getWrappedTokenBalance() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) GetWrappedTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "getWrappedTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWrappedTokenBalance is a free data retrieval call binding the contract method 0xf8e829c3.
//
// Solidity: function getWrappedTokenBalance() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) GetWrappedTokenBalance() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetWrappedTokenBalance(&_StakingPoolImplementation.CallOpts)
}

// GetWrappedTokenBalance is a free data retrieval call binding the contract method 0xf8e829c3.
//
// Solidity: function getWrappedTokenBalance() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) GetWrappedTokenBalance() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.GetWrappedTokenBalance(&_StakingPoolImplementation.CallOpts)
}

// HasUnStaked is a free data retrieval call binding the contract method 0x45e91319.
//
// Solidity: function hasUnStaked(address staker, uint256 stakeIndex) view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) HasUnStaked(opts *bind.CallOpts, staker common.Address, stakeIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "hasUnStaked", staker, stakeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUnStaked is a free data retrieval call binding the contract method 0x45e91319.
//
// Solidity: function hasUnStaked(address staker, uint256 stakeIndex) view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) HasUnStaked(staker common.Address, stakeIndex *big.Int) (bool, error) {
	return _StakingPoolImplementation.Contract.HasUnStaked(&_StakingPoolImplementation.CallOpts, staker, stakeIndex)
}

// HasUnStaked is a free data retrieval call binding the contract method 0x45e91319.
//
// Solidity: function hasUnStaked(address staker, uint256 stakeIndex) view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) HasUnStaked(staker common.Address, stakeIndex *big.Int) (bool, error) {
	return _StakingPoolImplementation.Contract.HasUnStaked(&_StakingPoolImplementation.CallOpts, staker, stakeIndex)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) IsActive() (bool, error) {
	return _StakingPoolImplementation.Contract.IsActive(&_StakingPoolImplementation.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) IsActive() (bool, error) {
	return _StakingPoolImplementation.Contract.IsActive(&_StakingPoolImplementation.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) IsAnEmergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "isAnEmergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) IsAnEmergency() (bool, error) {
	return _StakingPoolImplementation.Contract.IsAnEmergency(&_StakingPoolImplementation.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) IsAnEmergency() (bool, error) {
	return _StakingPoolImplementation.Contract.IsAnEmergency(&_StakingPoolImplementation.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) IsDeleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "isDeleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) IsDeleted() (bool, error) {
	return _StakingPoolImplementation.Contract.IsDeleted(&_StakingPoolImplementation.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) IsDeleted() (bool, error) {
	return _StakingPoolImplementation.Contract.IsDeleted(&_StakingPoolImplementation.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) IsEnhancedEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "isEnhancedEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolImplementation.Contract.IsEnhancedEnabled(&_StakingPoolImplementation.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolImplementation.Contract.IsEnhancedEnabled(&_StakingPoolImplementation.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) IsLocked() (bool, error) {
	return _StakingPoolImplementation.Contract.IsLocked(&_StakingPoolImplementation.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) IsLocked() (bool, error) {
	return _StakingPoolImplementation.Contract.IsLocked(&_StakingPoolImplementation.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) IsMatured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "isMatured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) IsMatured() (bool, error) {
	return _StakingPoolImplementation.Contract.IsMatured(&_StakingPoolImplementation.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) IsMatured() (bool, error) {
	return _StakingPoolImplementation.Contract.IsMatured(&_StakingPoolImplementation.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) LaunchDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "launchDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) LaunchDate() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.LaunchDate(&_StakingPoolImplementation.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) LaunchDate() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.LaunchDate(&_StakingPoolImplementation.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) LockTime() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.LockTime(&_StakingPoolImplementation.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) LockTime() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.LockTime(&_StakingPoolImplementation.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) Lps(opts *bind.CallOpts) (struct {
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
	err := _StakingPoolImplementation.contract.Call(opts, &out, "lps")

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
func (_StakingPoolImplementation *StakingPoolImplementationSession) Lps() (struct {
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
	return _StakingPoolImplementation.Contract.Lps(&_StakingPoolImplementation.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) Lps() (struct {
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
	return _StakingPoolImplementation.Contract.Lps(&_StakingPoolImplementation.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) MaturingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "maturingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.MaturingPrice(&_StakingPoolImplementation.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.MaturingPrice(&_StakingPoolImplementation.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) MaturityTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "maturityTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) MaturityTime() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.MaturityTime(&_StakingPoolImplementation.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) MaturityTime() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.MaturityTime(&_StakingPoolImplementation.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) MinimumStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "minimumStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.MinimumStakeAmount(&_StakingPoolImplementation.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.MinimumStakeAmount(&_StakingPoolImplementation.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) OracleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "oracleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) OracleContract() (common.Address, error) {
	return _StakingPoolImplementation.Contract.OracleContract(&_StakingPoolImplementation.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) OracleContract() (common.Address, error) {
	return _StakingPoolImplementation.Contract.OracleContract(&_StakingPoolImplementation.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) OracleDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "oracleDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.OracleDecimals(&_StakingPoolImplementation.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.OracleDecimals(&_StakingPoolImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) Owner() (common.Address, error) {
	return _StakingPoolImplementation.Contract.Owner(&_StakingPoolImplementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) Owner() (common.Address, error) {
	return _StakingPoolImplementation.Contract.Owner(&_StakingPoolImplementation.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) PoolCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "poolCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) PoolCreator() (common.Address, error) {
	return _StakingPoolImplementation.Contract.PoolCreator(&_StakingPoolImplementation.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) PoolCreator() (common.Address, error) {
	return _StakingPoolImplementation.Contract.PoolCreator(&_StakingPoolImplementation.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) PoolType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolImplementation *StakingPoolImplementationSession) PoolType() (string, error) {
	return _StakingPoolImplementation.Contract.PoolType(&_StakingPoolImplementation.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) PoolType() (string, error) {
	return _StakingPoolImplementation.Contract.PoolType(&_StakingPoolImplementation.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) PotentialCollabReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "potentialCollabReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.PotentialCollabReward(&_StakingPoolImplementation.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.PotentialCollabReward(&_StakingPoolImplementation.CallOpts)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) Predictions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
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
	err := _StakingPoolImplementation.contract.Call(opts, &out, "predictions", arg0, arg1)

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
func (_StakingPoolImplementation *StakingPoolImplementationSession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
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
	return _StakingPoolImplementation.Contract.Predictions(&_StakingPoolImplementation.CallOpts, arg0, arg1)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
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
	return _StakingPoolImplementation.Contract.Predictions(&_StakingPoolImplementation.CallOpts, arg0, arg1)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) PrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "prizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.PrizeAmount(&_StakingPoolImplementation.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.PrizeAmount(&_StakingPoolImplementation.CallOpts)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) PrizeRewardRates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "prizeRewardRates", arg0)

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
func (_StakingPoolImplementation *StakingPoolImplementationSession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolImplementation.Contract.PrizeRewardRates(&_StakingPoolImplementation.CallOpts, arg0)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolImplementation.Contract.PrizeRewardRates(&_StakingPoolImplementation.CallOpts, arg0)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) RewardManager() (common.Address, error) {
	return _StakingPoolImplementation.Contract.RewardManager(&_StakingPoolImplementation.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) RewardManager() (common.Address, error) {
	return _StakingPoolImplementation.Contract.RewardManager(&_StakingPoolImplementation.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) SizeAllocation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "sizeAllocation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.SizeAllocation(&_StakingPoolImplementation.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.SizeAllocation(&_StakingPoolImplementation.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) SizeLimitRangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "sizeLimitRangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.SizeLimitRangeRate(&_StakingPoolImplementation.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.SizeLimitRangeRate(&_StakingPoolImplementation.CallOpts)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) SortedStakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "sortedStakers", arg0)

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
func (_StakingPoolImplementation *StakingPoolImplementationSession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementation.Contract.SortedStakers(&_StakingPoolImplementation.CallOpts, arg0)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementation.Contract.SortedStakers(&_StakingPoolImplementation.CallOpts, arg0)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) StakeApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "stakeApr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) StakeApr() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.StakeApr(&_StakingPoolImplementation.CallOpts)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) StakeApr() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.StakeApr(&_StakingPoolImplementation.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) StakeTaxRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "stakeTaxRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.StakeTaxRate(&_StakingPoolImplementation.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.StakeTaxRate(&_StakingPoolImplementation.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "stakers", arg0)

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
func (_StakingPoolImplementation *StakingPoolImplementationSession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementation.Contract.Stakers(&_StakingPoolImplementation.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementation.Contract.Stakers(&_StakingPoolImplementation.CallOpts, arg0)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) StakingPoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "stakingPoolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolImplementation.Contract.StakingPoolImplementation(&_StakingPoolImplementation.CallOpts)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolImplementation.Contract.StakingPoolImplementation(&_StakingPoolImplementation.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) TotalStaked() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.TotalStaked(&_StakingPoolImplementation.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) TotalStaked() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.TotalStaked(&_StakingPoolImplementation.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) TotemToken() (common.Address, error) {
	return _StakingPoolImplementation.Contract.TotemToken(&_StakingPoolImplementation.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) TotemToken() (common.Address, error) {
	return _StakingPoolImplementation.Contract.TotemToken(&_StakingPoolImplementation.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) UpgradeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "upgradeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationSession) UpgradeEnabled() (bool, error) {
	return _StakingPoolImplementation.Contract.UpgradeEnabled(&_StakingPoolImplementation.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) UpgradeEnabled() (bool, error) {
	return _StakingPoolImplementation.Contract.UpgradeEnabled(&_StakingPoolImplementation.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) UsdPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "usdPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationSession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.UsdPrizeAmount(&_StakingPoolImplementation.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementation.Contract.UsdPrizeAmount(&_StakingPoolImplementation.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) WrappedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "wrappedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationSession) WrappedToken() (common.Address, error) {
	return _StakingPoolImplementation.Contract.WrappedToken(&_StakingPoolImplementation.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) WrappedToken() (common.Address, error) {
	return _StakingPoolImplementation.Contract.WrappedToken(&_StakingPoolImplementation.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolImplementation *StakingPoolImplementationCaller) WrappedTokenSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolImplementation.contract.Call(opts, &out, "wrappedTokenSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolImplementation *StakingPoolImplementationSession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolImplementation.Contract.WrappedTokenSymbol(&_StakingPoolImplementation.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolImplementation *StakingPoolImplementationCallerSession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolImplementation.Contract.WrappedTokenSymbol(&_StakingPoolImplementation.CallOpts)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) WrappedTokenDistributorInitialize(opts *bind.TransactOpts, swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "__WrappedTokenDistributor_initialize", swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.WrappedTokenDistributorInitialize(&_StakingPoolImplementation.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.WrappedTokenDistributorInitialize(&_StakingPoolImplementation.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// BatchStake is a paid mutator transaction binding the contract method 0xd49fab8a.
//
// Solidity: function batchStake(uint256[] _stakingAmounts, uint256[] _predictions) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) BatchStake(opts *bind.TransactOpts, _stakingAmounts []*big.Int, _predictions []*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "batchStake", _stakingAmounts, _predictions)
}

// BatchStake is a paid mutator transaction binding the contract method 0xd49fab8a.
//
// Solidity: function batchStake(uint256[] _stakingAmounts, uint256[] _predictions) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) BatchStake(_stakingAmounts []*big.Int, _predictions []*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.BatchStake(&_StakingPoolImplementation.TransactOpts, _stakingAmounts, _predictions)
}

// BatchStake is a paid mutator transaction binding the contract method 0xd49fab8a.
//
// Solidity: function batchStake(uint256[] _stakingAmounts, uint256[] _predictions) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) BatchStake(_stakingAmounts []*big.Int, _predictions []*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.BatchStake(&_StakingPoolImplementation.TransactOpts, _stakingAmounts, _predictions)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) ClaimReward() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.ClaimReward(&_StakingPoolImplementation.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.ClaimReward(&_StakingPoolImplementation.TransactOpts)
}

// DeclareEmergency is a paid mutator transaction binding the contract method 0x31e244e5.
//
// Solidity: function declareEmergency() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) DeclareEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "declareEmergency")
}

// DeclareEmergency is a paid mutator transaction binding the contract method 0x31e244e5.
//
// Solidity: function declareEmergency() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) DeclareEmergency() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.DeclareEmergency(&_StakingPoolImplementation.TransactOpts)
}

// DeclareEmergency is a paid mutator transaction binding the contract method 0x31e244e5.
//
// Solidity: function declareEmergency() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) DeclareEmergency() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.DeclareEmergency(&_StakingPoolImplementation.TransactOpts)
}

// DeletePool is a paid mutator transaction binding the contract method 0x8bfbbcbe.
//
// Solidity: function deletePool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) DeletePool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "deletePool")
}

// DeletePool is a paid mutator transaction binding the contract method 0x8bfbbcbe.
//
// Solidity: function deletePool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) DeletePool() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.DeletePool(&_StakingPoolImplementation.TransactOpts)
}

// DeletePool is a paid mutator transaction binding the contract method 0x8bfbbcbe.
//
// Solidity: function deletePool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) DeletePool() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.DeletePool(&_StakingPoolImplementation.TransactOpts)
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x8a9920b7.
//
// Solidity: function emergentWithdraw() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) EmergentWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "emergentWithdraw")
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x8a9920b7.
//
// Solidity: function emergentWithdraw() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) EmergentWithdraw() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.EmergentWithdraw(&_StakingPoolImplementation.TransactOpts)
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x8a9920b7.
//
// Solidity: function emergentWithdraw() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) EmergentWithdraw() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.EmergentWithdraw(&_StakingPoolImplementation.TransactOpts)
}

// EndPool is a paid mutator transaction binding the contract method 0x2d42cf7d.
//
// Solidity: function endPool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) EndPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "endPool")
}

// EndPool is a paid mutator transaction binding the contract method 0x2d42cf7d.
//
// Solidity: function endPool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) EndPool() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.EndPool(&_StakingPoolImplementation.TransactOpts)
}

// EndPool is a paid mutator transaction binding the contract method 0x2d42cf7d.
//
// Solidity: function endPool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) EndPool() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.EndPool(&_StakingPoolImplementation.TransactOpts)
}

// IndexedClaimReward is a paid mutator transaction binding the contract method 0xb1b1e1c0.
//
// Solidity: function indexedClaimReward(uint256 stakeIndex) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) IndexedClaimReward(opts *bind.TransactOpts, stakeIndex *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "indexedClaimReward", stakeIndex)
}

// IndexedClaimReward is a paid mutator transaction binding the contract method 0xb1b1e1c0.
//
// Solidity: function indexedClaimReward(uint256 stakeIndex) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) IndexedClaimReward(stakeIndex *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.IndexedClaimReward(&_StakingPoolImplementation.TransactOpts, stakeIndex)
}

// IndexedClaimReward is a paid mutator transaction binding the contract method 0xb1b1e1c0.
//
// Solidity: function indexedClaimReward(uint256 stakeIndex) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) IndexedClaimReward(stakeIndex *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.IndexedClaimReward(&_StakingPoolImplementation.TransactOpts, stakeIndex)
}

// LockPool is a paid mutator transaction binding the contract method 0x4026478e.
//
// Solidity: function lockPool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) LockPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "lockPool")
}

// LockPool is a paid mutator transaction binding the contract method 0x4026478e.
//
// Solidity: function lockPool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) LockPool() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.LockPool(&_StakingPoolImplementation.TransactOpts)
}

// LockPool is a paid mutator transaction binding the contract method 0x4026478e.
//
// Solidity: function lockPool() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) LockPool() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.LockPool(&_StakingPoolImplementation.TransactOpts)
}

// PurchaseWrappedToken is a paid mutator transaction binding the contract method 0x4ddc1bdc.
//
// Solidity: function purchaseWrappedToken(uint256 usdAmount, uint256 deadline) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) PurchaseWrappedToken(opts *bind.TransactOpts, usdAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "purchaseWrappedToken", usdAmount, deadline)
}

// PurchaseWrappedToken is a paid mutator transaction binding the contract method 0x4ddc1bdc.
//
// Solidity: function purchaseWrappedToken(uint256 usdAmount, uint256 deadline) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) PurchaseWrappedToken(usdAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.PurchaseWrappedToken(&_StakingPoolImplementation.TransactOpts, usdAmount, deadline)
}

// PurchaseWrappedToken is a paid mutator transaction binding the contract method 0x4ddc1bdc.
//
// Solidity: function purchaseWrappedToken(uint256 usdAmount, uint256 deadline) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) PurchaseWrappedToken(usdAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.PurchaseWrappedToken(&_StakingPoolImplementation.TransactOpts, usdAmount, deadline)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.RenounceOwnership(&_StakingPoolImplementation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.RenounceOwnership(&_StakingPoolImplementation.TransactOpts)
}

// SetActivationStatus is a paid mutator transaction binding the contract method 0xa118c9a7.
//
// Solidity: function setActivationStatus(bool _activationStatus) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) SetActivationStatus(opts *bind.TransactOpts, _activationStatus bool) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "setActivationStatus", _activationStatus)
}

// SetActivationStatus is a paid mutator transaction binding the contract method 0xa118c9a7.
//
// Solidity: function setActivationStatus(bool _activationStatus) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) SetActivationStatus(_activationStatus bool) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.SetActivationStatus(&_StakingPoolImplementation.TransactOpts, _activationStatus)
}

// SetActivationStatus is a paid mutator transaction binding the contract method 0xa118c9a7.
//
// Solidity: function setActivationStatus(bool _activationStatus) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) SetActivationStatus(_activationStatus bool) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.SetActivationStatus(&_StakingPoolImplementation.TransactOpts, _activationStatus)
}

// SetOracleToZero is a paid mutator transaction binding the contract method 0xdf7f92e0.
//
// Solidity: function setOracleToZero() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) SetOracleToZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "setOracleToZero")
}

// SetOracleToZero is a paid mutator transaction binding the contract method 0xdf7f92e0.
//
// Solidity: function setOracleToZero() returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) SetOracleToZero() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.SetOracleToZero(&_StakingPoolImplementation.TransactOpts)
}

// SetOracleToZero is a paid mutator transaction binding the contract method 0xdf7f92e0.
//
// Solidity: function setOracleToZero() returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) SetOracleToZero() (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.SetOracleToZero(&_StakingPoolImplementation.TransactOpts)
}

// SetSortedStakers is a paid mutator transaction binding the contract method 0x444a4902.
//
// Solidity: function setSortedStakers(address[25] addrArray, uint256[25] indexArray) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) SetSortedStakers(opts *bind.TransactOpts, addrArray [25]common.Address, indexArray [25]*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "setSortedStakers", addrArray, indexArray)
}

// SetSortedStakers is a paid mutator transaction binding the contract method 0x444a4902.
//
// Solidity: function setSortedStakers(address[25] addrArray, uint256[25] indexArray) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) SetSortedStakers(addrArray [25]common.Address, indexArray [25]*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.SetSortedStakers(&_StakingPoolImplementation.TransactOpts, addrArray, indexArray)
}

// SetSortedStakers is a paid mutator transaction binding the contract method 0x444a4902.
//
// Solidity: function setSortedStakers(address[25] addrArray, uint256[25] indexArray) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) SetSortedStakers(addrArray [25]common.Address, indexArray [25]*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.SetSortedStakers(&_StakingPoolImplementation.TransactOpts, addrArray, indexArray)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _amount, uint256 _pricePrediction) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int, _pricePrediction *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "stake", _amount, _pricePrediction)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _amount, uint256 _pricePrediction) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) Stake(_amount *big.Int, _pricePrediction *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.Stake(&_StakingPoolImplementation.TransactOpts, _amount, _pricePrediction)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _amount, uint256 _pricePrediction) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) Stake(_amount *big.Int, _pricePrediction *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.Stake(&_StakingPoolImplementation.TransactOpts, _amount, _pricePrediction)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.TransferOwnership(&_StakingPoolImplementation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.TransferOwnership(&_StakingPoolImplementation.TransactOpts, newOwner)
}

// UpdateMaturingPrice is a paid mutator transaction binding the contract method 0xe113af51.
//
// Solidity: function updateMaturingPrice(uint256 _predictionPrice, bool _samePredictionPrizeToken, uint256 _prizePrice, address _oracleContract, uint256 _oracleDecimals) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) UpdateMaturingPrice(opts *bind.TransactOpts, _predictionPrice *big.Int, _samePredictionPrizeToken bool, _prizePrice *big.Int, _oracleContract common.Address, _oracleDecimals *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "updateMaturingPrice", _predictionPrice, _samePredictionPrizeToken, _prizePrice, _oracleContract, _oracleDecimals)
}

// UpdateMaturingPrice is a paid mutator transaction binding the contract method 0xe113af51.
//
// Solidity: function updateMaturingPrice(uint256 _predictionPrice, bool _samePredictionPrizeToken, uint256 _prizePrice, address _oracleContract, uint256 _oracleDecimals) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) UpdateMaturingPrice(_predictionPrice *big.Int, _samePredictionPrizeToken bool, _prizePrice *big.Int, _oracleContract common.Address, _oracleDecimals *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.UpdateMaturingPrice(&_StakingPoolImplementation.TransactOpts, _predictionPrice, _samePredictionPrizeToken, _prizePrice, _oracleContract, _oracleDecimals)
}

// UpdateMaturingPrice is a paid mutator transaction binding the contract method 0xe113af51.
//
// Solidity: function updateMaturingPrice(uint256 _predictionPrice, bool _samePredictionPrizeToken, uint256 _prizePrice, address _oracleContract, uint256 _oracleDecimals) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) UpdateMaturingPrice(_predictionPrice *big.Int, _samePredictionPrizeToken bool, _prizePrice *big.Int, _oracleContract common.Address, _oracleDecimals *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.UpdateMaturingPrice(&_StakingPoolImplementation.TransactOpts, _predictionPrice, _samePredictionPrizeToken, _prizePrice, _oracleContract, _oracleDecimals)
}

// WithdrawStuckTokens is a paid mutator transaction binding the contract method 0x5d2631e2.
//
// Solidity: function withdrawStuckTokens(address _stuckToken, uint256 amount, address receiver) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactor) WithdrawStuckTokens(opts *bind.TransactOpts, _stuckToken common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.contract.Transact(opts, "withdrawStuckTokens", _stuckToken, amount, receiver)
}

// WithdrawStuckTokens is a paid mutator transaction binding the contract method 0x5d2631e2.
//
// Solidity: function withdrawStuckTokens(address _stuckToken, uint256 amount, address receiver) returns()
func (_StakingPoolImplementation *StakingPoolImplementationSession) WithdrawStuckTokens(_stuckToken common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.WithdrawStuckTokens(&_StakingPoolImplementation.TransactOpts, _stuckToken, amount, receiver)
}

// WithdrawStuckTokens is a paid mutator transaction binding the contract method 0x5d2631e2.
//
// Solidity: function withdrawStuckTokens(address _stuckToken, uint256 amount, address receiver) returns()
func (_StakingPoolImplementation *StakingPoolImplementationTransactorSession) WithdrawStuckTokens(_stuckToken common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementation.Contract.WithdrawStuckTokens(&_StakingPoolImplementation.TransactOpts, _stuckToken, amount, receiver)
}

// StakingPoolImplementationApproveTokensIterator is returned from FilterApproveTokens and is used to iterate over the raw logs and unpacked data for ApproveTokens events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationApproveTokensIterator struct {
	Event *StakingPoolImplementationApproveTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationApproveTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationApproveTokens)
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
		it.Event = new(StakingPoolImplementationApproveTokens)
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
func (it *StakingPoolImplementationApproveTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationApproveTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationApproveTokens represents a ApproveTokens event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationApproveTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApproveTokens is a free log retrieval operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterApproveTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolImplementationApproveTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationApproveTokensIterator{contract: _StakingPoolImplementation.contract, event: "ApproveTokens", logs: logs, sub: sub}, nil
}

// WatchApproveTokens is a free log subscription operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchApproveTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationApproveTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationApproveTokens)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
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
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseApproveTokens(log types.Log) (*StakingPoolImplementationApproveTokens, error) {
	event := new(StakingPoolImplementationApproveTokens)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationBatchStakeIterator is returned from FilterBatchStake and is used to iterate over the raw logs and unpacked data for BatchStake events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationBatchStakeIterator struct {
	Event *StakingPoolImplementationBatchStake // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationBatchStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationBatchStake)
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
		it.Event = new(StakingPoolImplementationBatchStake)
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
func (it *StakingPoolImplementationBatchStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationBatchStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationBatchStake represents a BatchStake event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationBatchStake struct {
	User              common.Address
	StakeAmounts      []*big.Int
	Predictions       []*big.Int
	IndexOfFirstStake *big.Int
	IndexOfLastStake  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBatchStake is a free log retrieval operation binding the contract event 0x852fe7a2ad7b2c8c0a0e9b927734feb1acc3e8bc608b8921d4c4768c6193c33b.
//
// Solidity: event BatchStake(address indexed user, uint256[] stakeAmounts, uint256[] predictions, uint256 indexOfFirstStake, uint256 indexOfLastStake)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterBatchStake(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationBatchStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "BatchStake", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationBatchStakeIterator{contract: _StakingPoolImplementation.contract, event: "BatchStake", logs: logs, sub: sub}, nil
}

// WatchBatchStake is a free log subscription operation binding the contract event 0x852fe7a2ad7b2c8c0a0e9b927734feb1acc3e8bc608b8921d4c4768c6193c33b.
//
// Solidity: event BatchStake(address indexed user, uint256[] stakeAmounts, uint256[] predictions, uint256 indexOfFirstStake, uint256 indexOfLastStake)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchBatchStake(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationBatchStake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "BatchStake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationBatchStake)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "BatchStake", log); err != nil {
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

// ParseBatchStake is a log parse operation binding the contract event 0x852fe7a2ad7b2c8c0a0e9b927734feb1acc3e8bc608b8921d4c4768c6193c33b.
//
// Solidity: event BatchStake(address indexed user, uint256[] stakeAmounts, uint256[] predictions, uint256 indexOfFirstStake, uint256 indexOfLastStake)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseBatchStake(log types.Log) (*StakingPoolImplementationBatchStake, error) {
	event := new(StakingPoolImplementationBatchStake)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "BatchStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationDistributedBTCIterator is returned from FilterDistributedBTC and is used to iterate over the raw logs and unpacked data for DistributedBTC events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationDistributedBTCIterator struct {
	Event *StakingPoolImplementationDistributedBTC // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationDistributedBTCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationDistributedBTC)
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
		it.Event = new(StakingPoolImplementationDistributedBTC)
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
func (it *StakingPoolImplementationDistributedBTCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationDistributedBTCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationDistributedBTC represents a DistributedBTC event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationDistributedBTC struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedBTC is a free log retrieval operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterDistributedBTC(opts *bind.FilterOpts, to []common.Address) (*StakingPoolImplementationDistributedBTCIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationDistributedBTCIterator{contract: _StakingPoolImplementation.contract, event: "DistributedBTC", logs: logs, sub: sub}, nil
}

// WatchDistributedBTC is a free log subscription operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchDistributedBTC(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationDistributedBTC, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationDistributedBTC)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
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
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseDistributedBTC(log types.Log) (*StakingPoolImplementationDistributedBTC, error) {
	event := new(StakingPoolImplementationDistributedBTC)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationOwnershipTransferredIterator struct {
	Event *StakingPoolImplementationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationOwnershipTransferred)
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
		it.Event = new(StakingPoolImplementationOwnershipTransferred)
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
func (it *StakingPoolImplementationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationOwnershipTransferred represents a OwnershipTransferred event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingPoolImplementationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationOwnershipTransferredIterator{contract: _StakingPoolImplementation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationOwnershipTransferred)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseOwnershipTransferred(log types.Log) (*StakingPoolImplementationOwnershipTransferred, error) {
	event := new(StakingPoolImplementationOwnershipTransferred)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationPoolActivatedIterator is returned from FilterPoolActivated and is used to iterate over the raw logs and unpacked data for PoolActivated events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolActivatedIterator struct {
	Event *StakingPoolImplementationPoolActivated // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationPoolActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationPoolActivated)
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
		it.Event = new(StakingPoolImplementationPoolActivated)
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
func (it *StakingPoolImplementationPoolActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationPoolActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationPoolActivated represents a PoolActivated event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolActivated is a free log retrieval operation binding the contract event 0x3d31b97eef590df350bb2cdbc390034c613c9f23ea6ae1906682628c85248b53.
//
// Solidity: event PoolActivated()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterPoolActivated(opts *bind.FilterOpts) (*StakingPoolImplementationPoolActivatedIterator, error) {

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "PoolActivated")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationPoolActivatedIterator{contract: _StakingPoolImplementation.contract, event: "PoolActivated", logs: logs, sub: sub}, nil
}

// WatchPoolActivated is a free log subscription operation binding the contract event 0x3d31b97eef590df350bb2cdbc390034c613c9f23ea6ae1906682628c85248b53.
//
// Solidity: event PoolActivated()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchPoolActivated(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationPoolActivated) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "PoolActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationPoolActivated)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolActivated", log); err != nil {
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

// ParsePoolActivated is a log parse operation binding the contract event 0x3d31b97eef590df350bb2cdbc390034c613c9f23ea6ae1906682628c85248b53.
//
// Solidity: event PoolActivated()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParsePoolActivated(log types.Log) (*StakingPoolImplementationPoolActivated, error) {
	event := new(StakingPoolImplementationPoolActivated)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationPoolDeactivatedIterator is returned from FilterPoolDeactivated and is used to iterate over the raw logs and unpacked data for PoolDeactivated events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolDeactivatedIterator struct {
	Event *StakingPoolImplementationPoolDeactivated // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationPoolDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationPoolDeactivated)
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
		it.Event = new(StakingPoolImplementationPoolDeactivated)
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
func (it *StakingPoolImplementationPoolDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationPoolDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationPoolDeactivated represents a PoolDeactivated event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolDeactivated is a free log retrieval operation binding the contract event 0x763d36bfedf6707954f8bb39c2011a58a47a54844f2b3cdcd326c71cad0aa753.
//
// Solidity: event PoolDeactivated()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterPoolDeactivated(opts *bind.FilterOpts) (*StakingPoolImplementationPoolDeactivatedIterator, error) {

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "PoolDeactivated")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationPoolDeactivatedIterator{contract: _StakingPoolImplementation.contract, event: "PoolDeactivated", logs: logs, sub: sub}, nil
}

// WatchPoolDeactivated is a free log subscription operation binding the contract event 0x763d36bfedf6707954f8bb39c2011a58a47a54844f2b3cdcd326c71cad0aa753.
//
// Solidity: event PoolDeactivated()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchPoolDeactivated(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationPoolDeactivated) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "PoolDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationPoolDeactivated)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolDeactivated", log); err != nil {
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

// ParsePoolDeactivated is a log parse operation binding the contract event 0x763d36bfedf6707954f8bb39c2011a58a47a54844f2b3cdcd326c71cad0aa753.
//
// Solidity: event PoolDeactivated()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParsePoolDeactivated(log types.Log) (*StakingPoolImplementationPoolDeactivated, error) {
	event := new(StakingPoolImplementationPoolDeactivated)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationPoolDeletedIterator is returned from FilterPoolDeleted and is used to iterate over the raw logs and unpacked data for PoolDeleted events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolDeletedIterator struct {
	Event *StakingPoolImplementationPoolDeleted // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationPoolDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationPoolDeleted)
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
		it.Event = new(StakingPoolImplementationPoolDeleted)
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
func (it *StakingPoolImplementationPoolDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationPoolDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationPoolDeleted represents a PoolDeleted event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolDeleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolDeleted is a free log retrieval operation binding the contract event 0x9331b2a38c658c57a39ea19b2709d928c28a12cc8fbb838d12b85148550f937d.
//
// Solidity: event PoolDeleted()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterPoolDeleted(opts *bind.FilterOpts) (*StakingPoolImplementationPoolDeletedIterator, error) {

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "PoolDeleted")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationPoolDeletedIterator{contract: _StakingPoolImplementation.contract, event: "PoolDeleted", logs: logs, sub: sub}, nil
}

// WatchPoolDeleted is a free log subscription operation binding the contract event 0x9331b2a38c658c57a39ea19b2709d928c28a12cc8fbb838d12b85148550f937d.
//
// Solidity: event PoolDeleted()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchPoolDeleted(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationPoolDeleted) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "PoolDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationPoolDeleted)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolDeleted", log); err != nil {
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

// ParsePoolDeleted is a log parse operation binding the contract event 0x9331b2a38c658c57a39ea19b2709d928c28a12cc8fbb838d12b85148550f937d.
//
// Solidity: event PoolDeleted()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParsePoolDeleted(log types.Log) (*StakingPoolImplementationPoolDeleted, error) {
	event := new(StakingPoolImplementationPoolDeleted)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationPoolLockedIterator is returned from FilterPoolLocked and is used to iterate over the raw logs and unpacked data for PoolLocked events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolLockedIterator struct {
	Event *StakingPoolImplementationPoolLocked // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationPoolLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationPoolLocked)
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
		it.Event = new(StakingPoolImplementationPoolLocked)
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
func (it *StakingPoolImplementationPoolLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationPoolLocked represents a PoolLocked event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolLocked is a free log retrieval operation binding the contract event 0x2e136745550967f28b779ce8c395ca060a2f844edfb5b06631fe677f66f9e305.
//
// Solidity: event PoolLocked()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterPoolLocked(opts *bind.FilterOpts) (*StakingPoolImplementationPoolLockedIterator, error) {

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "PoolLocked")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationPoolLockedIterator{contract: _StakingPoolImplementation.contract, event: "PoolLocked", logs: logs, sub: sub}, nil
}

// WatchPoolLocked is a free log subscription operation binding the contract event 0x2e136745550967f28b779ce8c395ca060a2f844edfb5b06631fe677f66f9e305.
//
// Solidity: event PoolLocked()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchPoolLocked(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationPoolLocked) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "PoolLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationPoolLocked)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolLocked", log); err != nil {
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

// ParsePoolLocked is a log parse operation binding the contract event 0x2e136745550967f28b779ce8c395ca060a2f844edfb5b06631fe677f66f9e305.
//
// Solidity: event PoolLocked()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParsePoolLocked(log types.Log) (*StakingPoolImplementationPoolLocked, error) {
	event := new(StakingPoolImplementationPoolLocked)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationPoolMaturedIterator is returned from FilterPoolMatured and is used to iterate over the raw logs and unpacked data for PoolMatured events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolMaturedIterator struct {
	Event *StakingPoolImplementationPoolMatured // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationPoolMaturedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationPoolMatured)
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
		it.Event = new(StakingPoolImplementationPoolMatured)
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
func (it *StakingPoolImplementationPoolMaturedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationPoolMaturedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationPoolMatured represents a PoolMatured event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolMatured struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolMatured is a free log retrieval operation binding the contract event 0xe8ef7a17c92d025d9c1ead1c6b3ca3b44d3283d0ee954f615c5c476911f629a5.
//
// Solidity: event PoolMatured()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterPoolMatured(opts *bind.FilterOpts) (*StakingPoolImplementationPoolMaturedIterator, error) {

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "PoolMatured")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationPoolMaturedIterator{contract: _StakingPoolImplementation.contract, event: "PoolMatured", logs: logs, sub: sub}, nil
}

// WatchPoolMatured is a free log subscription operation binding the contract event 0xe8ef7a17c92d025d9c1ead1c6b3ca3b44d3283d0ee954f615c5c476911f629a5.
//
// Solidity: event PoolMatured()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchPoolMatured(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationPoolMatured) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "PoolMatured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationPoolMatured)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolMatured", log); err != nil {
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

// ParsePoolMatured is a log parse operation binding the contract event 0xe8ef7a17c92d025d9c1ead1c6b3ca3b44d3283d0ee954f615c5c476911f629a5.
//
// Solidity: event PoolMatured()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParsePoolMatured(log types.Log) (*StakingPoolImplementationPoolMatured, error) {
	event := new(StakingPoolImplementationPoolMatured)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolMatured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationPoolSortedIterator is returned from FilterPoolSorted and is used to iterate over the raw logs and unpacked data for PoolSorted events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolSortedIterator struct {
	Event *StakingPoolImplementationPoolSorted // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationPoolSortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationPoolSorted)
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
		it.Event = new(StakingPoolImplementationPoolSorted)
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
func (it *StakingPoolImplementationPoolSortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationPoolSortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationPoolSorted represents a PoolSorted event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationPoolSorted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolSorted is a free log retrieval operation binding the contract event 0x5c243cc6e4f210bf2dd86b8eca0ca891d3f8ad2eb09fbf91735eb348bc199c7e.
//
// Solidity: event PoolSorted()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterPoolSorted(opts *bind.FilterOpts) (*StakingPoolImplementationPoolSortedIterator, error) {

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "PoolSorted")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationPoolSortedIterator{contract: _StakingPoolImplementation.contract, event: "PoolSorted", logs: logs, sub: sub}, nil
}

// WatchPoolSorted is a free log subscription operation binding the contract event 0x5c243cc6e4f210bf2dd86b8eca0ca891d3f8ad2eb09fbf91735eb348bc199c7e.
//
// Solidity: event PoolSorted()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchPoolSorted(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationPoolSorted) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "PoolSorted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationPoolSorted)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolSorted", log); err != nil {
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

// ParsePoolSorted is a log parse operation binding the contract event 0x5c243cc6e4f210bf2dd86b8eca0ca891d3f8ad2eb09fbf91735eb348bc199c7e.
//
// Solidity: event PoolSorted()
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParsePoolSorted(log types.Log) (*StakingPoolImplementationPoolSorted, error) {
	event := new(StakingPoolImplementationPoolSorted)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "PoolSorted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationReceivedTokensIterator is returned from FilterReceivedTokens and is used to iterate over the raw logs and unpacked data for ReceivedTokens events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationReceivedTokensIterator struct {
	Event *StakingPoolImplementationReceivedTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationReceivedTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationReceivedTokens)
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
		it.Event = new(StakingPoolImplementationReceivedTokens)
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
func (it *StakingPoolImplementationReceivedTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationReceivedTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationReceivedTokens represents a ReceivedTokens event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationReceivedTokens struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedTokens is a free log retrieval operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterReceivedTokens(opts *bind.FilterOpts, from []common.Address) (*StakingPoolImplementationReceivedTokensIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationReceivedTokensIterator{contract: _StakingPoolImplementation.contract, event: "ReceivedTokens", logs: logs, sub: sub}, nil
}

// WatchReceivedTokens is a free log subscription operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchReceivedTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationReceivedTokens, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationReceivedTokens)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
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
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseReceivedTokens(log types.Log) (*StakingPoolImplementationReceivedTokens, error) {
	event := new(StakingPoolImplementationReceivedTokens)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationStakeIterator struct {
	Event *StakingPoolImplementationStake // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationStake)
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
		it.Event = new(StakingPoolImplementationStake)
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
func (it *StakingPoolImplementationStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationStake represents a Stake event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationStake struct {
	User            common.Address
	Amount          *big.Int
	PricePrediction *big.Int
	IndexOfStake    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0xf556991011e831bcfac4f406d547e5e32cdd98267efab83935230d5f8d02c446.
//
// Solidity: event Stake(address indexed user, uint256 amount, uint256 pricePrediction, uint256 indexOfStake)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterStake(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "Stake", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationStakeIterator{contract: _StakingPoolImplementation.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0xf556991011e831bcfac4f406d547e5e32cdd98267efab83935230d5f8d02c446.
//
// Solidity: event Stake(address indexed user, uint256 amount, uint256 pricePrediction, uint256 indexOfStake)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationStake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "Stake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationStake)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0xf556991011e831bcfac4f406d547e5e32cdd98267efab83935230d5f8d02c446.
//
// Solidity: event Stake(address indexed user, uint256 amount, uint256 pricePrediction, uint256 indexOfStake)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseStake(log types.Log) (*StakingPoolImplementationStake, error) {
	event := new(StakingPoolImplementationStake)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationTransferTokensIterator is returned from FilterTransferTokens and is used to iterate over the raw logs and unpacked data for TransferTokens events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationTransferTokensIterator struct {
	Event *StakingPoolImplementationTransferTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationTransferTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationTransferTokens)
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
		it.Event = new(StakingPoolImplementationTransferTokens)
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
func (it *StakingPoolImplementationTransferTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationTransferTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationTransferTokens represents a TransferTokens event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationTransferTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferTokens is a free log retrieval operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterTransferTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolImplementationTransferTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationTransferTokensIterator{contract: _StakingPoolImplementation.contract, event: "TransferTokens", logs: logs, sub: sub}, nil
}

// WatchTransferTokens is a free log subscription operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchTransferTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationTransferTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationTransferTokens)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "TransferTokens", log); err != nil {
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
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseTransferTokens(log types.Log) (*StakingPoolImplementationTransferTokens, error) {
	event := new(StakingPoolImplementationTransferTokens)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "TransferTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationUnstakeIterator struct {
	Event *StakingPoolImplementationUnstake // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationUnstake)
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
		it.Event = new(StakingPoolImplementationUnstake)
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
func (it *StakingPoolImplementationUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationUnstake represents a Unstake event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationUnstake struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed user, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterUnstake(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationUnstakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "Unstake", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationUnstakeIterator{contract: _StakingPoolImplementation.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed user, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationUnstake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "Unstake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationUnstake)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed user, uint256 amount)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseUnstake(log types.Log) (*StakingPoolImplementationUnstake, error) {
	event := new(StakingPoolImplementationUnstake)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationWithdrawStakingReturnIterator is returned from FilterWithdrawStakingReturn and is used to iterate over the raw logs and unpacked data for WithdrawStakingReturn events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationWithdrawStakingReturnIterator struct {
	Event *StakingPoolImplementationWithdrawStakingReturn // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationWithdrawStakingReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationWithdrawStakingReturn)
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
		it.Event = new(StakingPoolImplementationWithdrawStakingReturn)
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
func (it *StakingPoolImplementationWithdrawStakingReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationWithdrawStakingReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationWithdrawStakingReturn represents a WithdrawStakingReturn event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationWithdrawStakingReturn struct {
	User          common.Address
	StakingReturn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawStakingReturn is a free log retrieval operation binding the contract event 0x115e2fe760f98d4fd0828d7741ae2070e7db700f892efbfe98448b53ac583886.
//
// Solidity: event WithdrawStakingReturn(address indexed user, uint256 stakingReturn)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterWithdrawStakingReturn(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationWithdrawStakingReturnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "WithdrawStakingReturn", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationWithdrawStakingReturnIterator{contract: _StakingPoolImplementation.contract, event: "WithdrawStakingReturn", logs: logs, sub: sub}, nil
}

// WatchWithdrawStakingReturn is a free log subscription operation binding the contract event 0x115e2fe760f98d4fd0828d7741ae2070e7db700f892efbfe98448b53ac583886.
//
// Solidity: event WithdrawStakingReturn(address indexed user, uint256 stakingReturn)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchWithdrawStakingReturn(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationWithdrawStakingReturn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "WithdrawStakingReturn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationWithdrawStakingReturn)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "WithdrawStakingReturn", log); err != nil {
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

// ParseWithdrawStakingReturn is a log parse operation binding the contract event 0x115e2fe760f98d4fd0828d7741ae2070e7db700f892efbfe98448b53ac583886.
//
// Solidity: event WithdrawStakingReturn(address indexed user, uint256 stakingReturn)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseWithdrawStakingReturn(log types.Log) (*StakingPoolImplementationWithdrawStakingReturn, error) {
	event := new(StakingPoolImplementationWithdrawStakingReturn)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "WithdrawStakingReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationWithdrawTotemPrizeIterator is returned from FilterWithdrawTotemPrize and is used to iterate over the raw logs and unpacked data for WithdrawTotemPrize events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationWithdrawTotemPrizeIterator struct {
	Event *StakingPoolImplementationWithdrawTotemPrize // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationWithdrawTotemPrizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationWithdrawTotemPrize)
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
		it.Event = new(StakingPoolImplementationWithdrawTotemPrize)
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
func (it *StakingPoolImplementationWithdrawTotemPrizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationWithdrawTotemPrizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationWithdrawTotemPrize represents a WithdrawTotemPrize event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationWithdrawTotemPrize struct {
	User       common.Address
	TotemPrize *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTotemPrize is a free log retrieval operation binding the contract event 0x723b06f6c3a94dc82019852e6aa01c0241fba5f84410b9003ddc3440ad499e59.
//
// Solidity: event WithdrawTotemPrize(address indexed user, uint256 totemPrize)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterWithdrawTotemPrize(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationWithdrawTotemPrizeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "WithdrawTotemPrize", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationWithdrawTotemPrizeIterator{contract: _StakingPoolImplementation.contract, event: "WithdrawTotemPrize", logs: logs, sub: sub}, nil
}

// WatchWithdrawTotemPrize is a free log subscription operation binding the contract event 0x723b06f6c3a94dc82019852e6aa01c0241fba5f84410b9003ddc3440ad499e59.
//
// Solidity: event WithdrawTotemPrize(address indexed user, uint256 totemPrize)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchWithdrawTotemPrize(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationWithdrawTotemPrize, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "WithdrawTotemPrize", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationWithdrawTotemPrize)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "WithdrawTotemPrize", log); err != nil {
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

// ParseWithdrawTotemPrize is a log parse operation binding the contract event 0x723b06f6c3a94dc82019852e6aa01c0241fba5f84410b9003ddc3440ad499e59.
//
// Solidity: event WithdrawTotemPrize(address indexed user, uint256 totemPrize)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseWithdrawTotemPrize(log types.Log) (*StakingPoolImplementationWithdrawTotemPrize, error) {
	event := new(StakingPoolImplementationWithdrawTotemPrize)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "WithdrawTotemPrize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationWithdrawWrappedTokenPrizeIterator is returned from FilterWithdrawWrappedTokenPrize and is used to iterate over the raw logs and unpacked data for WithdrawWrappedTokenPrize events raised by the StakingPoolImplementation contract.
type StakingPoolImplementationWithdrawWrappedTokenPrizeIterator struct {
	Event *StakingPoolImplementationWithdrawWrappedTokenPrize // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationWithdrawWrappedTokenPrizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationWithdrawWrappedTokenPrize)
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
		it.Event = new(StakingPoolImplementationWithdrawWrappedTokenPrize)
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
func (it *StakingPoolImplementationWithdrawWrappedTokenPrizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationWithdrawWrappedTokenPrizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationWithdrawWrappedTokenPrize represents a WithdrawWrappedTokenPrize event raised by the StakingPoolImplementation contract.
type StakingPoolImplementationWithdrawWrappedTokenPrize struct {
	User              common.Address
	WrappedTokenPrize *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWrappedTokenPrize is a free log retrieval operation binding the contract event 0xd93fc6d1e8779271f55cdb0d5f7b6b05be64a2226a1ad050d3036f5609c1cdcf.
//
// Solidity: event WithdrawWrappedTokenPrize(address indexed user, uint256 wrappedTokenPrize)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) FilterWithdrawWrappedTokenPrize(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationWithdrawWrappedTokenPrizeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.FilterLogs(opts, "WithdrawWrappedTokenPrize", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationWithdrawWrappedTokenPrizeIterator{contract: _StakingPoolImplementation.contract, event: "WithdrawWrappedTokenPrize", logs: logs, sub: sub}, nil
}

// WatchWithdrawWrappedTokenPrize is a free log subscription operation binding the contract event 0xd93fc6d1e8779271f55cdb0d5f7b6b05be64a2226a1ad050d3036f5609c1cdcf.
//
// Solidity: event WithdrawWrappedTokenPrize(address indexed user, uint256 wrappedTokenPrize)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) WatchWithdrawWrappedTokenPrize(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationWithdrawWrappedTokenPrize, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementation.contract.WatchLogs(opts, "WithdrawWrappedTokenPrize", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationWithdrawWrappedTokenPrize)
				if err := _StakingPoolImplementation.contract.UnpackLog(event, "WithdrawWrappedTokenPrize", log); err != nil {
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

// ParseWithdrawWrappedTokenPrize is a log parse operation binding the contract event 0xd93fc6d1e8779271f55cdb0d5f7b6b05be64a2226a1ad050d3036f5609c1cdcf.
//
// Solidity: event WithdrawWrappedTokenPrize(address indexed user, uint256 wrappedTokenPrize)
func (_StakingPoolImplementation *StakingPoolImplementationFilterer) ParseWithdrawWrappedTokenPrize(log types.Log) (*StakingPoolImplementationWithdrawWrappedTokenPrize, error) {
	event := new(StakingPoolImplementationWithdrawWrappedTokenPrize)
	if err := _StakingPoolImplementation.contract.UnpackLog(event, "WithdrawWrappedTokenPrize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
