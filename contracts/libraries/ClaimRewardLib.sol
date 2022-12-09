// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./CalculateRewardLib.sol";
import "../libraries/BasisPoints.sol";
import "../Staking/StakingPoolStorageStructure.sol";

library ClaimRewardLib {

    using CalculateRewardLib for *;
    using BasisPoints for uint256; 
    using SafeMath for uint256;

    uint256 public constant foo = 0;

    function withdrawStakingReturn(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker,
        StakingPoolStorageStructure.LibParams storage _lps
    )
        public 
    {
        
        if (_staker.length == 0) return;

        for (uint256 i = 0; i < _staker.length; i++) {
            uint256 rewardPerStake = CalculateRewardLib._getStakingRewardPerStake(
                _staker, 
                i, 
                _lps);

            _staker[i].lastWithdrawalTime = block.timestamp;
            _staker[i].amountWithdrawn = _staker[i].amountWithdrawn.add(
                rewardPerStake
            );
        }
    }

    function withdrawPrize(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker
    ) 
        public
    {
        if (_staker.length == 0) return;

        for (uint256 i = 0; i < _staker.length; i++) {
            _staker[i].prizeRewardWithdrawn = true;
        }
    }

    function withdrawStakedBalance(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker
    )
        public 
    {
        
        if (_staker.length == 0) return;

        for (uint256 i = 0; i < _staker.length; i++) {
            _staker[i].didUnstake = true;
        }
    }

    function getStakingReturn(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker,
        StakingPoolStorageStructure.LibParams storage _lps  
    ) 
        public 
        view 
        returns (uint256) 
    {
        if (_staker.length == 0) return 0;

        uint256 reward = 0;
        for (uint256 i = 0; i < _staker.length; i++) {
            uint256 rewardPerStake = CalculateRewardLib._getStakingRewardPerStake(
                _staker,
                i, 
                _lps
            );

            reward = reward.add(rewardPerStake);
        }

        return reward;
    }

    function getPrize(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        StakingPoolStorageStructure.LibParams storage _lps,
        StakingPoolStorageStructure.PrizeRewardRate[] storage _prizeRewardRates
    )
        public
        view
        returns (uint256, uint256)
    {
        if (!_lps.isMatured) return (0, 0);

        if (_staker.length == 0) return (0, 0);

        uint256 maturingWrappedTokenPrizeAmount =
            (_lps.usdPrizeAmount.mul(10**_lps.oracleDecimals)).div(_lps.maturingPrice);

        uint256 reward = 0;
        uint256 wrappedTokenReward = 0;

        for (uint256 i = 0; i < _staker.length; i++) {
            if (!_staker[i].prizeRewardWithdrawn) {

                uint256 _percent = CalculateRewardLib._getPercentageReward(
                    _staker[i].rank,
                    _prizeRewardRates
                );

                reward = reward.add(
                            _lps.prizeAmount.mulBP(_percent)
                        );

                wrappedTokenReward = wrappedTokenReward.add(
                            maturingWrappedTokenPrizeAmount
                                .mulBP(_percent)
                        );        
            }
        }

        if (_lps.collaborativeReward > 0) {
            reward = reward.addBP(_lps.collaborativeReward);
            wrappedTokenReward = wrappedTokenReward.addBP(_lps.collaborativeReward);
        }

        return (reward, wrappedTokenReward);
    }

}

