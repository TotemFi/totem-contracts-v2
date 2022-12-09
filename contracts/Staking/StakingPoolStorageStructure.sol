// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "../Distribution/USDRetriever.sol";
import "../Distribution/WrappedTokenDistributorUpgradeable.sol";

import "../libraries/BasisPoints.sol";
import "../libraries/CalculateRewardLib.sol";
import "../libraries/IndexedClaimRewardLib.sol";
import "../libraries/ClaimRewardLib.sol";
import "../libraries/PriceConsumer.sol";

import "../interfaces/ITotemToken.sol";
import "../interfaces/IRewardManager.sol";

contract StakingPoolStorageStructure is
    OwnableUpgradeable,
    USDRetriever,
    WrappedTokenDistributorUpgradeable
{
    /**
     * @notice Declared for passing the needed params to libraries.
     */
    struct LibParams {
        uint256 launchDate;
        uint256 lockTime;
        uint256 maturityTime;
        uint256 maturingPrice;
        uint256 usdPrizeAmount;
        uint256 prizeAmount;
        uint256 stakeApr;
        uint256 collaborativeReward;
        uint256 oracleDecimals;
        bool isEnhancedEnabled;
        bool isMatured;
    }

    struct StakeWithPrediction {
        uint256 stakedBalance;
        uint256 stakedTime;
        uint256 amountWithdrawn;
        uint256 lastWithdrawalTime;
        uint256 pricePrediction;
        uint256 difference;
        uint256 rank;
        bool prizeRewardWithdrawn;
        bool didUnstake;
    }

    struct Staker {
        address stakerAddress;
        uint256 index;
    }

    struct PrizeRewardRate {
        uint256 rank;
        uint256 percentage;
    }

    uint256 public constant sizeLimitRangeRate = 5;

    uint256 public launchDate;
    uint256 public lockTime;
    uint256 public maturityTime;
    uint256 public sizeAllocation;
    uint256 public stakeApr;
    uint256 public prizeAmount;
    /**
     * @notice usdPrizeAmount is the enabler of WrappedToken rewarder; If it is set to 0 
            then the pool is only TOTM rewarder.
     */
    uint256 public usdPrizeAmount;
    uint256 public stakeTaxRate;
    uint256 public minimumStakeAmount;
    uint256 public totalStaked;
    uint256 public maturingPrice;
    uint256 public potentialCollabReward;
    uint256 public collaborativeRange;
    /**
     * @notice Based on the white paper, the collaborative reward can be 20% (2000),
             25% (2500) or 35% (3500).
     */
    uint256 public collaborativeReward;
    uint256 public oracleDecimals;
    uint256 public averagePricePrediction;

    address public stakingPoolImplementation;
    address public poolCreator;
    address public oracleContract;

    bool public isAnEmergency;
    bool public isEnhancedEnabled;
    bool public isActive;
    bool public isLocked;
    bool public isMatured;
    bool public isDeleted;
    /**
     * @dev StakingPoolImplementation can't be upgraded unless superAdmin sets this flag.
     */
    bool public upgradeEnabled;

    string public wrappedTokenSymbol;
    string public poolType;

    LibParams public lps;

    PrizeRewardRate[] public prizeRewardRates;
    Staker[] public stakers;
    Staker[] public sortedStakers;

    mapping(address => StakeWithPrediction[]) public predictions;

    ITotemToken public totemToken;
    IRewardManager public rewardManager;
    IERC20 public wrappedToken;
}
