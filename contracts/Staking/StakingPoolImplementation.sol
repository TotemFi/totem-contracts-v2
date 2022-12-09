// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./StakingPoolStorageStructure.sol";

contract StakingPoolImplementation is StakingPoolStorageStructure {
    using BasisPoints for uint256;
    using SafeMath for uint256;
    using CalculateRewardLib for *;
    using IndexedClaimRewardLib for *;
    using ClaimRewardLib for *;
    using PriceConsumer for *;

    modifier onlyPoolCreator() {
        require(
            _msgSender() == poolCreator,
            "0300 caller is not a pool creator"
        );
        _;
    }

    event Stake(
        address indexed user,
        uint256 amount,
        uint256 pricePrediction,
        uint256 indexOfStake
    );

    // TODO: add the index of stakes for BatchStake event too
    event BatchStake(
        address indexed user,
        uint256[] stakeAmounts,
        uint256[] predictions,
        uint256 indexOfFirstStake,
        uint256 indexOfLastStake
    );

    event WithdrawStakingReturn(address indexed user, uint256 stakingReturn);
    event WithdrawTotemPrize(address indexed user, uint256 totemPrize);
    event WithdrawWrappedTokenPrize(
        address indexed user,
        uint256 wrappedTokenPrize
    );
    event Unstake(address indexed user, uint256 amount);

    event PoolActivated();
    event PoolDeactivated();
    event PoolLocked();
    event PoolSorted();
    event PoolMatured();
    event PoolDeleted();

    function setActivationStatus(bool _activationStatus)
        external
        onlyPoolCreator
    {
        require(
            isActive != _activationStatus,
            "Not changing the activation status"
        );
        isActive = _activationStatus;

        if (isActive) emit PoolActivated();
        else emit PoolDeactivated();
    }

    function stake(uint256 _amount, uint256 _pricePrediction) external {
        require(
            isActive && block.timestamp > launchDate,
            "0313 pool is not active"
        );
        require(
            block.timestamp < (launchDate + lockTime),
            "0316 Can not stake after lock date"
        );
        require(!isLocked, "0310 Pool is locked");
        require(
            _amount >= minimumStakeAmount,
            "0311 Amount can't be less than the minimum"
        );

        // uint256 limitRange = sizeAllocation.mul(sizeLimitRangeRate).div(100);
        uint256 limitRange = minimumStakeAmount;
        uint256 taxRate = totemToken.taxRate();
        uint256 tax = totemToken.taxExempt(_msgSender())
            ? 0
            : _amount.mulBP(taxRate);

        require(
            totalStaked.add(_amount).sub(tax) <= sizeAllocation.add(limitRange),
            "0312 Can't stake above size allocation"
        );

        uint256 stakeTaxAmount;

        (stakeTaxAmount, _amount) = getStakingTax(_amount, taxRate);

        totemToken.transferFrom(
            _msgSender(),
            address(this),
            (_amount + stakeTaxAmount)
        );

        /// @dev This is to remove token tax (not staking tax) from the amount
        _amount = _amount.sub(tax);

        if (stakeTaxAmount > 0)
            totemToken.transfer(totemToken.taxationWallet(), stakeTaxAmount);

        uint256 indexOfStake = predictions[_msgSender()].length;

        _stake(_msgSender(), _amount, _pricePrediction);

        totalStaked = totalStaked.add(_amount);

        if (totalStaked >= sizeAllocation) {
            _lockPool();
        }

        emit Stake(_msgSender(), _amount, _pricePrediction, indexOfStake);
    }

    function batchStake(
        uint256[] calldata _stakingAmounts,
        uint256[] calldata _predictions
    ) external {
        require(
            isActive && block.timestamp > launchDate,
            "0313 pool is not active"
        );
        require(
            block.timestamp < (launchDate + lockTime),
            "0316 Can not stake after lock date"
        );
        require(!isLocked, "0310 Pool is locked");
        require(
            _stakingAmounts.length == _predictions.length,
            "0315 stakingAmount and predictions length mismatch"
        );

        uint256 totalStakingAmount = 0;

        for (uint256 i; i < _stakingAmounts.length; i++) {
            require(
                _stakingAmounts[i] >= minimumStakeAmount,
                "0311 Amount can't be less than the minimum"
            );
            totalStakingAmount = totalStakingAmount.add(_stakingAmounts[i]);
        }

        // uint256 limitRange = sizeAllocation.mul(sizeLimitRangeRate).div(100);
        uint256 limitRange = minimumStakeAmount;
        uint256 taxRate = totemToken.taxRate();
        uint256 tax = totemToken.taxExempt(_msgSender())
            ? 0
            : totalStakingAmount.mulBP(taxRate);

        require(
            totalStaked.add(totalStakingAmount).sub(tax) <=
                sizeAllocation.add(limitRange),
            "0312 Can't stake above size allocation"
        );

        uint256 stakeTaxAmount;

        (stakeTaxAmount, totalStakingAmount) = getStakingTax(
            totalStakingAmount,
            taxRate
        );

        totemToken.transferFrom(
            _msgSender(),
            address(this),
            (totalStakingAmount + stakeTaxAmount)
        );

        /// @dev This is to remove token tax (not staking tax) from the amount
        totalStakingAmount = totalStakingAmount.sub(tax);

        if (stakeTaxAmount > 0)
            totemToken.transfer(totemToken.taxationWallet(), stakeTaxAmount);

        uint256 indexOfFirstStake = predictions[_msgSender()].length;

        for (uint256 i; i < _stakingAmounts.length; i++) {
            uint256 stakingAmount = _stakingAmounts[i];

            tax = totemToken.taxExempt(_msgSender())
                ? 0
                : stakingAmount.mulBP(taxRate);

            (stakeTaxAmount, stakingAmount) = getStakingTax(
                stakingAmount,
                taxRate
            );

            /// @dev This is to remove token tax (not staking tax) from the amount
            stakingAmount = stakingAmount.sub(tax);

            _stake(_msgSender(), stakingAmount, _predictions[i]);

            totalStaked = totalStaked.add(stakingAmount);
        }

        uint256 indexOfLastStake = predictions[_msgSender()].length - 1;

        if (totalStaked >= sizeAllocation) {
            _lockPool();
        }

        emit BatchStake(
            _msgSender(),
            _stakingAmounts,
            _predictions,
            indexOfFirstStake,
            indexOfLastStake
        );
    }

    function _stake(
        address _staker,
        uint256 _amount,
        uint256 _pricePrediction
    ) internal {
        stakers.push(
            Staker({stakerAddress: _staker, index: predictions[_staker].length})
        );

        predictions[_staker].push(
            StakeWithPrediction({
                stakedBalance: _amount,
                stakedTime: block.timestamp,
                amountWithdrawn: 0,
                lastWithdrawalTime: block.timestamp,
                pricePrediction: _pricePrediction,
                difference: type(uint256).max,
                rank: type(uint256).max,
                prizeRewardWithdrawn: false,
                didUnstake: false
            })
        );

        _getAveragePricePrediction(_pricePrediction, _amount);
    }

    function getStakingTax(uint256 amount, uint256 tokenTaxRate)
        public
        view
        returns (uint256, uint256)
    {
        uint256 newStakeTaxRate = stakeTaxRate > tokenTaxRate
            ? stakeTaxRate.sub(tokenTaxRate)
            : 0;
        if (newStakeTaxRate == 0) {
            return (0, amount);
        }
        return (
            amount.mulBP(newStakeTaxRate),
            amount.sub(amount.mulBP(newStakeTaxRate))
        );
    }

    function claimReward() external {
        uint256 stakingReturn = ClaimRewardLib.getStakingReturn(
            predictions[_msgSender()],
            lps
        );

        (uint256 totemPrize, uint256 wrappedTokenPrize) = ClaimRewardLib
            .getPrize(predictions[_msgSender()], lps, prizeRewardRates);

        uint256 withdrawableTotemReward = totemPrize + stakingReturn;

        if (isMatured) {
            if (usdPrizeAmount > 0) {
                if (wrappedTokenPrize > 0) {
                    /// @dev Not the actual withdraw, only updating the array in the mapping
                    ClaimRewardLib.withdrawPrize(predictions[_msgSender()]);

                    require(
                        wrappedToken.transfer(_msgSender(), wrappedTokenPrize),
                        "0320"
                    );

                    emit WithdrawWrappedTokenPrize(
                        _msgSender(),
                        wrappedTokenPrize
                    );
                }
            }

            if (totemPrize > 0) {
                ClaimRewardLib.withdrawPrize(predictions[_msgSender()]);
            }

            uint256 stakedBalance = CalculateRewardLib.getTotalStakedBalance(
                predictions[_msgSender()]
            );

            if (stakedBalance > 0) {
                ClaimRewardLib.withdrawStakedBalance(predictions[_msgSender()]);

                totemToken.transfer(_msgSender(), stakedBalance);

                emit Unstake(_msgSender(), stakedBalance);
            }
        }

        /// @dev before maturity, totemPrize is always zero
        if (withdrawableTotemReward > 0) {
            /// @dev Send the token reward only when rewardManager has the enough funds
            require(
                totemToken.balanceOf(address(rewardManager)) >=
                    withdrawableTotemReward,
                "Not enough balance in reward manager"
            );

            ClaimRewardLib.withdrawStakingReturn(
                predictions[_msgSender()],
                lps
            );

            rewardManager.rewardUser(_msgSender(), withdrawableTotemReward);

            emit WithdrawStakingReturn(_msgSender(), stakingReturn);
            emit WithdrawTotemPrize(_msgSender(), totemPrize);
        }
    }

    function indexedClaimReward(uint256 stakeIndex) external {
        require(
            predictions[_msgSender()].length >= stakeIndex,
            "Index does not exist"
        );
        require(
            predictions[_msgSender()].length != 0,
            "User does not have any stakes"
        );

        uint256 stakingReturn = IndexedClaimRewardLib.getIndexedStakingReturn(
            predictions[_msgSender()],
            stakeIndex,
            lps
        );

        (uint256 totemPrize, uint256 wrappedTokenPrize) = IndexedClaimRewardLib
            .getIndexedPrize(
                predictions[_msgSender()],
                stakeIndex,
                lps,
                prizeRewardRates
            );

        uint256 withdrawableTotemReward = totemPrize + stakingReturn;

        if (isMatured) {
            if (usdPrizeAmount > 0) {
                if (wrappedTokenPrize > 0) {
                    IndexedClaimRewardLib.withdrawIndexedPrize(
                        predictions[_msgSender()],
                        stakeIndex
                    );

                    require(
                        wrappedToken.transfer(_msgSender(), wrappedTokenPrize),
                        "0330"
                    );

                    emit WithdrawWrappedTokenPrize(
                        _msgSender(),
                        wrappedTokenPrize
                    );
                }
            }

            if (totemPrize > 0) {
                IndexedClaimRewardLib.withdrawIndexedPrize(
                    predictions[_msgSender()],
                    stakeIndex
                );
            }

            uint256 stakedBalance = IndexedClaimRewardLib
                .getIndexedStakedBalance(predictions[_msgSender()], stakeIndex);

            if (stakedBalance > 0) {
                IndexedClaimRewardLib.withdrawIndexedStakedBalance(
                    predictions[_msgSender()],
                    stakeIndex
                );

                totemToken.transfer(_msgSender(), stakedBalance);

                emit Unstake(_msgSender(), stakedBalance);
            }
        }

        /// @dev before maturity, totemPrize is always zero
        if (withdrawableTotemReward > 0) {
            /// @dev Send the token reward only when rewardManager has the enough funds
            require(
                totemToken.balanceOf(address(rewardManager)) >=
                    withdrawableTotemReward,
                "Not enough balance in reward manager"
            );

            IndexedClaimRewardLib.withdrawIndexedStakingReturn(
                predictions[_msgSender()],
                stakeIndex,
                lps
            );

            rewardManager.rewardUser(_msgSender(), withdrawableTotemReward);

            emit WithdrawStakingReturn(_msgSender(), stakingReturn);
            emit WithdrawTotemPrize(_msgSender(), totemPrize);
        }
    }

    function purchaseWrappedToken(uint256 usdAmount, uint256 deadline)
        external
        onlyPoolCreator
    {
        //TODO: require usdAmount to be more than usdPrizeAmount, to have enough rewards
        require(usdPrizeAmount > 0, "0340 The pool is only TOTM rewarder");

        require(usdAmount > 0, "0341 Amount can't be zero");

        require(deadline >= block.timestamp, "0342 Deadline is low");

        address swapRouterAddress = getSwapRouter();
        approveTokens(swapRouterAddress, usdAmount);

        uint256 wrappedTokenAmount = getEstimatedWrappedTokenForUSD(usdAmount);

        uint256 wrappedTokenAmountWithSlippage = wrappedTokenAmount.sub(
            wrappedTokenAmount.mulBP(300)
        );

        transferTokensThroughSwap(
            address(this),
            usdAmount,
            wrappedTokenAmountWithSlippage,
            deadline
        );
    }

    function getWrappedTokenBalance() public view returns (uint256) {
        return wrappedToken.balanceOf(address(this));
    }

    function lockPool() public virtual onlyPoolCreator {
        _lockPool();
    }

    function _lockPool() internal {
        isLocked = true;

        emit PoolLocked();
    }

    /**
     * @param _predictionPrice is ignored if oracle is not zero address.When there is no oracle,
             _predictionPrice is the maturingPrice and is set manually by the pool creator
     * @param _prizePrice is ignored if oracle is not zero address.When there is no oracle,
             _predictionPrice is the maturingPrice and is set manually by the pool creator
    */
    function updateMaturingPrice(
        uint256 _predictionPrice,
        bool _samePredictionPrizeToken,
        uint256 _prizePrice,
        address _oracleContract,
        uint256 _oracleDecimals
    ) external onlyPoolCreator {
        require(
            block.timestamp >= launchDate + lockTime + maturityTime,
            "0350 Can't set maturing price before the maturity time"
        );

        if (_samePredictionPrizeToken) {
            if (oracleContract != address(0)) {
                maturingPrice = PriceConsumer.getLatestPrice(oracleContract);
                lps.maturingPrice = maturingPrice;
                lps.oracleDecimals = oracleDecimals;
            } else {
                maturingPrice = _predictionPrice;
                lps.maturingPrice = _predictionPrice;
                lps.oracleDecimals = oracleDecimals;
            }
        } else {
            if (oracleContract != address(0)) {
                if (_oracleContract != address(0)) {
                    maturingPrice = PriceConsumer.getLatestPrice(
                        oracleContract
                    );
                    lps.maturingPrice = PriceConsumer.getLatestPrice(
                        _oracleContract
                    );
                    lps.oracleDecimals = PriceConsumer.getDecimals(
                        _oracleContract
                    );
                } else {
                    maturingPrice = PriceConsumer.getLatestPrice(
                        oracleContract
                    );
                    lps.maturingPrice = _prizePrice;
                    lps.oracleDecimals = _oracleDecimals;
                }
            } else {
                if (_oracleContract != address(0)) {
                    maturingPrice = _predictionPrice;
                    lps.maturingPrice = PriceConsumer.getLatestPrice(
                        _oracleContract
                    );
                    lps.oracleDecimals = PriceConsumer.getDecimals(
                        _oracleContract
                    );
                } else {
                    maturingPrice = _predictionPrice;
                    lps.maturingPrice = _prizePrice;
                    lps.oracleDecimals = _oracleDecimals;
                }
            }
        }
    }

    function getLPS() public view returns (uint256, uint256) {
        return (lps.maturingPrice, lps.oracleDecimals);
    }

    function getPrizeTokenDecimals(address _oracleContract)
        public
        view
        returns (uint256)
    {
        return PriceConsumer.getDecimals(_oracleContract);
    }

    /**
     * @notice Sets oracle to zero in case it was given incorrectly by the owner,
     *         or it is not available
     */
    function setOracleToZero() external onlyPoolCreator {
        oracleContract = address(0);
    }

    /**
     * @notice Sets oracle to zero in case it was given incorrectly by the owner,
     *         or it is not available
     */

    function setSortedStakers(
        address[25] calldata addrArray,
        uint256[25] calldata indexArray
    ) external onlyPoolCreator {
        require(
            block.timestamp >= launchDate + lockTime + maturityTime,
            "0390 Can't set sorted stakers before the maturity time"
        );

        if (sortedStakers.length != 0) {
            delete sortedStakers;
        }

        uint256 i;
        for (i = 0; i < addrArray.length; i++) {
            /// @dev The first 0 address means the other addresses are also 0 so they won't be checked
            if (addrArray[i] == address(0)) break;

            sortedStakers.push(
                Staker({stakerAddress: addrArray[i], index: indexArray[i]})
            );

            predictions[addrArray[i]][indexArray[i]].rank = i + 1;
        }

        require(
            prizeRewardRates[prizeRewardRates.length - 1].rank >= i,
            "number of sorted stakers must be less than or equal to the last rank"
        );

        emit PoolSorted();
    }

    function endPool() external onlyPoolCreator {
        require(
            block.timestamp >= launchDate + lockTime + maturityTime,
            "0360 Can't end pool before the maturity time"
        );

        //TODO: check to see if there is enough USD to buy the wrapped token with, the mimimum USD
        // must be usdPrizeAmount, if there is not, do not allow endPool
        if (usdPrizeAmount > 0) {
            require(
                getWrappedTokenBalance() != 0,
                "0361 WrappedToken Rewards not available"
            );
        }

        if (stakers.length > 0) {
            require(sortedStakers.length != 0, "0362 first should sort");
        }

        /**
         *  @dev potentialCollabReward allows the admin to set the collaborateive reward
         *  @notice the collaborative reward is only given to the pools which the average price
         *          predicted has the accuracy of 25$
         */
        if (potentialCollabReward > 0) {
            if (
                getDifference(averagePricePrediction, collaborativeRange) == 0
            ) {
                collaborativeReward = potentialCollabReward;
                lps.collaborativeReward = collaborativeReward;
            }
        }

        isLocked = true;
        isMatured = true;
        lps.isMatured = isMatured;

        emit PoolMatured();
    }

    function getDifference(uint256 prediction, uint256 _range)
        public
        view
        returns (uint256)
    {
        if (_range > prediction) return 0;

        if (prediction > maturingPrice) {
            if (prediction.sub(_range) <= maturingPrice) return 0;
            else return prediction.sub(_range).sub(maturingPrice);
        } else {
            if (prediction.add(_range) >= maturingPrice) return 0;
            else return maturingPrice.sub(prediction.add(_range));
        }
    }

    /**
     * @notice Gets the avgerage price prediction for calculating collaborative reward
     */
    function _getAveragePricePrediction(uint256 _prediction, uint256 _amount)
        internal
    {
        uint256 predictionsCount = stakers.length - 1;

        if (predictionsCount == 0) averagePricePrediction = _prediction;

        averagePricePrediction = averagePricePrediction
            .mul(totalStaked)
            .add(_prediction.mul(_amount))
            .div(totalStaked.add(_amount));
    }

    function deletePool() external onlyPoolCreator {
        isDeleted = true;
        emit PoolDeleted();
    }

    function getStakers()
        public
        view
        returns (address[] memory, uint256[] memory)
    {
        address[] memory addrs = new address[](stakers.length);
        uint256[] memory indexes = new uint256[](stakers.length);

        for (uint256 i = 0; i < stakers.length; i++) {
            addrs[i] = stakers[i].stakerAddress;
            indexes[i] = stakers[i].index;
        }

        return (addrs, indexes);
    }

    function getStakingReward(address _staker) public view returns (uint256) {
        uint256 reward = ClaimRewardLib.getStakingReturn(
            predictions[_staker],
            lps
        );

        return reward;
    }

    function getIndexedStakingReward(address _staker, uint256 _stakeIndex)
        public
        view
        returns (uint256)
    {
        uint256 reward = IndexedClaimRewardLib.getIndexedStakingReturn(
            predictions[_staker],
            _stakeIndex,
            lps
        );

        return reward;
    }

    function getPrize(address _staker) public view returns (uint256, uint256) {
        (uint256 reward, uint256 wrappedTokenReward) = ClaimRewardLib.getPrize(
            predictions[_staker],
            lps,
            prizeRewardRates
        );

        return (reward, wrappedTokenReward);
    }

    function getIndexedPrize(address _staker, uint256 _stakeIndex)
        public
        view
        returns (uint256, uint256)
    {
        (uint256 reward, uint256 wrappedTokenReward) = IndexedClaimRewardLib
            .getIndexedPrize(
                predictions[_staker],
                _stakeIndex,
                lps,
                prizeRewardRates
            );

        return (reward, wrappedTokenReward);
    }

    /**  
     * @notice hasUnStaked return true if the user staked in the pool and then 
            has unStaked it (claimed)
    */
    function hasUnStaked(address staker, uint256 stakeIndex)
        external
        view
        returns (bool)
    {
        StakeWithPrediction[] memory userStakes = predictions[staker];

        require(
            userStakes.length > 0,
            "0380 this address didn't stake in this pool"
        );

        require(stakeIndex < userStakes.length, "0381 this index exceeds");

        if (userStakes[stakeIndex].didUnstake) {
            return true;
        }
        return false;
    }

    function withdrawStuckTokens(
        address _stuckToken,
        uint256 amount,
        address receiver
    ) external onlyPoolCreator {
        require(
            _stuckToken != address(totemToken),
            "0370 totems can not be transfered"
        );
        IERC20 stuckToken = IERC20(_stuckToken);
        stuckToken.transfer(receiver, amount);
    }

    function declareEmergency() external onlyPoolCreator {
        isActive = false;
        isAnEmergency = true;

        _lockPool();
    }

    function emergentWithdraw() external {
        require(isAnEmergency, "it's not an emergency");

        uint256 stakedBalance = CalculateRewardLib.getTotalStakedBalance(
            predictions[_msgSender()]
        );
        if (stakedBalance > 0) {
            ClaimRewardLib.withdrawStakedBalance(predictions[_msgSender()]);

            totemToken.transfer(_msgSender(), stakedBalance);

            emit Unstake(_msgSender(), stakedBalance);
        }
    }
}
