// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./StakingPoolStorageStructure.sol";

contract StakingPoolProxy is StakingPoolStorageStructure {
    using BasisPoints for uint256;
    using SafeMath for uint256;

    modifier onlyPoolCreator() {
        require(msg.sender == poolCreator, "msg.sender is not an owner");
        _;
    }

    event ImplementationUpgraded();

    /**
     * @dev poolCreator is set to the address of StakingPoolFactory here, but it will change
            to the address of the owner after initialize is called. This is to prevent any other
            entity other than the StakingPoolFactory to call initialize and upgradeTo (for the 
            first time).
            upgradeEnabled set to true so that upgradeTo can be called for the first time
            when the main impelementaiton is being set. 
    */
    constructor() {
        poolCreator = msg.sender;
        upgradeEnabled = true;
    }

    /**
     * @notice This is called in case we want to upgrade a working pool which inherits from
            the original implementation, to apply bug fixes and consider emergency cases.
    */
    function upgradeTo(address _newStakingPoolImplementation)
        external
        onlyPoolCreator
    {
        require(upgradeEnabled, "Upgrade is not enabled yet");
        require(
            stakingPoolImplementation != _newStakingPoolImplementation,
            "Is already the implementation"
        );
        _setStakingPoolImplementation(_newStakingPoolImplementation);
        upgradeEnabled = false;
    }

    /**
     * @notice StakingPoolImplementation can't be upgraded unless superAdmin sets upgradeEnabled
     */
    function enableUpgrade() external onlyOwner {
        upgradeEnabled = true;
    }

    function disableUpgrade() external onlyOwner {
        upgradeEnabled = false;
    }

    /**
     * @notice The initializer modifier is used to make sure initialize() is not called 
            more than once because we want it to act like a constructor.
       @param _addrs Addresses used by priceConsumer and WrappedTokenDistributor
                _addrs[0] = swapRouterAddress
                _addrs[1] = BUSDContractAddress
                _addrs[2] = OracleAddress
                _addrs[3] = WrappedTokenContractAddress
    */

    // TODO: gas optimization required
    function initialize(
        string memory _wrappedTokenSymbol,
        string memory _poolType,
        ITotemToken _totemToken,
        IRewardManager _rewardManager,
        address _poolCreator,
        address[4] memory _addrs,
        uint256[12] memory _variables,
        uint256[8] memory _ranks,
        uint256[8] memory _percentages,
        bool _isEnhancedEnabled
    ) public initializer onlyPoolCreator {
        /// @dev we should call inits because we don't have a constructor to do it for us
        OwnableUpgradeable.__Ownable_init();
        ContextUpgradeable.__Context_init();

        WrappedTokenDistributorUpgradeable.__WrappedTokenDistributor_initialize(
                _addrs[0],
                _addrs[1],
                _addrs[3]
            );

        require(
            _variables[0] > block.timestamp,
            "0301 launch date can't be in the past"
        );

        wrappedTokenSymbol = _wrappedTokenSymbol;
        poolType = _poolType;
        totemToken = _totemToken;
        rewardManager = _rewardManager;
        poolCreator = _poolCreator;
        setUSDToken(_addrs[1]);
        oracleContract = _addrs[2];
        wrappedToken = IERC20(_addrs[3]);

        launchDate = _variables[0];
        maturityTime = _variables[1];
        lockTime = _variables[2];
        sizeAllocation = _variables[3];
        stakeApr = _variables[4];
        prizeAmount = _variables[5];
        usdPrizeAmount = _variables[6];
        potentialCollabReward = _variables[7];
        collaborativeRange = _variables[8];
        stakeTaxRate = _variables[9];
        minimumStakeAmount = _variables[10];

        if (_addrs[2] == address(0)) oracleDecimals = _variables[11];
        else oracleDecimals = PriceConsumer.getDecimals(oracleContract);

        isEnhancedEnabled = _isEnhancedEnabled;

        // ranks validation
        uint256 rewardRates = _ranks[0].mul(_percentages[0]);
        uint256 l;
        for (l = 0; l < _ranks.length; l++) {
            if (_ranks[l] == 0) break;

            if (_ranks[l + 1] != 0) {
                require(_ranks[l] < _ranks[l + 1], "wrong order of ranks");
                rewardRates = rewardRates.add(
                    (_ranks[l + 1].sub(_ranks[l])).mul(_percentages[l + 1])
                );
            }

            prizeRewardRates.push(
                PrizeRewardRate({rank: _ranks[l], percentage: _percentages[l]})
            );
        }

        require(
            _ranks[l - 1] <= 25,
            "last rank must be less than or equal to 25"
        );

        require(
            _percentages[l - 1] != 0 && _percentages[l] == 0,
            "ranks and percentages length mismatch"
        );

        require(rewardRates == 10000, "reward percentages must be 100%");

        /**
         * @notice LibParams are set here. Some of them may change in the lifetime of a pool
                which is also considered
        */
        lps.launchDate = launchDate;
        lps.lockTime = lockTime;
        lps.maturityTime = maturityTime;
        lps.usdPrizeAmount = usdPrizeAmount;
        lps.prizeAmount = prizeAmount;
        lps.stakeApr = stakeApr;
        lps.collaborativeReward = collaborativeReward;
        lps.isEnhancedEnabled = isEnhancedEnabled;
    }

    fallback() external payable {
        address opr = stakingPoolImplementation;
        require(opr != address(0));
        assembly {
            calldatacopy(0, 0, calldatasize())
            let result := delegatecall(gas(), opr, 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            switch result
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }

    receive() external payable {}

    function _setStakingPoolImplementation(address _newStakingPool) internal {
        stakingPoolImplementation = _newStakingPool;
        emit ImplementationUpgraded();
    }
}
