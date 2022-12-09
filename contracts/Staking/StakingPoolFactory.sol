// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../Role/PoolCreator.sol";
import "./StakingPoolProxy.sol";
import "../interfaces/IRewardManager.sol";
import "../Distribution/USDRetriever.sol";

contract StakingPoolFactory is PoolCreator {
    ITotemToken public totemToken;
    IRewardManager public rewardManager;

    address public superAdmin;
    address public stakingPoolImplementationAdr;
    address public swapRouter;
    address public usdToken;
    uint256 public stakingPoolTaxRate;
    uint256 public minimumStakeAmount;

    /**
     * @param variables The StakingPoolProxy is created with these specs:
            variables[0] = launchDate
            variables[1] = maturityTime
            variables[2] = lockTime
            variables[3] = sizeAllocation
            variables[4] = stakeApr
            variables[5] = prizeAmount
            variables[6] = usdPrizeAmount
            variables[7] = potentialCollabReward
            variables[8] = collaborativeRange
            variables[9] = stakingPoolTaxRate
            variables[10] = minimumStakeAmount
            variables[11] = oracleDeciamls
     */
    event PoolCreated(
        address indexed pool,
        string wrappedTokenSymbol,
        string poolType,
        uint256[12] variables,
        uint256[8] ranks,
        uint256[8] percentages,
        bool isEnhancedEnabled
    );

    event NewStakingPoolImplemnetationWasSet();

    event NewSuperAdminWasSet();

    constructor (
        ITotemToken _totemToken,
        IRewardManager _rewardManager,
        address _swapRouter,
        address _usdToken,
        address _stakingPoolImplementation,
        address _superAdmin
    ){
        totemToken = _totemToken;
        rewardManager = _rewardManager;
        swapRouter = _swapRouter;
        usdToken = _usdToken;
        stakingPoolImplementationAdr = _stakingPoolImplementation;
        superAdmin = _superAdmin;
         
        stakingPoolTaxRate = 300;
    }

    /**
     * @notice creates a StakingPoolProxy for the  provided stakingPoolImplementationAdr
            and initializes it so that the pool is ready to be used.
       @param _variables The StakingPoolProxy is created with these specs:
            variables[0] = launchDate
            variables[1] = maturityTime
            variables[2] = lockTime
            variables[3] = sizeAllocation
            variables[4] = stakeApr
            variables[5] = prizeAmount
            variables[6] = usdPrizeAmount
            variables[7] = potentialCollabReward
            variables[8] = collaborativeRange
            variables[9] = stakingPoolTaxRate
            variables[10] = minimumStakeAmount
            variables[11] = oracleDeciamls
    */
    function createPoolProxy(
        address _oracleContract,
        address _wrappedTokenContract,
        string memory _wrappedTokenSymbol,
        string memory _poolType,
        uint256[12] memory _variables,
        uint256[8] memory _ranks,
        uint256[8] memory _percentages,
        bool isEnhancedEnabled
    ) external onlyPoolCreator returns (address) {
        
        require(
            _ranks.length == _percentages.length,
            "length of ranks and percentages should be same"
        );

        if (_variables[9] == 0) {
            _variables[9] = stakingPoolTaxRate;
        }

        StakingPoolProxy stakingPoolProxy = new StakingPoolProxy();
        address stakingPoolProxyAdr = address(stakingPoolProxy);

        stakingPoolProxy.upgradeTo(stakingPoolImplementationAdr);

        address[4] memory addrs = [swapRouter, usdToken, _oracleContract, _wrappedTokenContract];

        _createPool( 
            addrs, 
            _wrappedTokenSymbol, 
            _poolType, 
            _variables, 
            _ranks, 
            _percentages, 
            isEnhancedEnabled,
            stakingPoolProxy
        );

        stakingPoolProxy.transferOwnership(superAdmin);

        rewardManager.addPool(stakingPoolProxyAdr);

        return stakingPoolProxyAdr;
    }

    function _createPool(
        address[4] memory _addrs,
        string memory _wrappedTokenSymbol,
        string memory _poolType,
        uint256[12] memory _variables,
        uint256[8] memory _ranks,
        uint256[8] memory _percentages,
        bool _isEnhancedEnabled,
        StakingPoolProxy _stakingPoolProxy
    ) internal {
        _stakingPoolProxy.initialize(
            _wrappedTokenSymbol,
            _poolType,
            totemToken,
            rewardManager,
            _msgSender(),
            _addrs,
            _variables,
            _ranks,
            _percentages,
            _isEnhancedEnabled
        );

        address stakingPoolProxyAdr = address(_stakingPoolProxy);

        emit PoolCreated(
            stakingPoolProxyAdr,
            _wrappedTokenSymbol,
            _poolType,
            _variables,
            _ranks,
            _percentages,
            _isEnhancedEnabled
        );
    }

    /**
     * @notice This function is called whenever we want to use a new StakingPoolImplementation
            to create our proxies for.
     * @param _ImpAdr address of the new StakingPoolImplementation contract.
    */
    function setNewStakingPoolImplementationAdr(address _ImpAdr) external onlyPoolCreator {
        require(
            stakingPoolImplementationAdr != _ImpAdr, 
            'This address is the implementation that is already being used'
        );
        stakingPoolImplementationAdr = _ImpAdr;
        emit NewStakingPoolImplemnetationWasSet();
    }

    /**
     * @notice Changes superAdmin's address so that new StakingPoolProxies have this new superAdmin
    */
    function setNewSuperAdmin(address _superAdmin) external onlyPoolCreator {
        superAdmin = _superAdmin;
        emit NewSuperAdminWasSet();
    }

    function setSwapRouter(address _swapRouter) external onlyPoolCreator {
        require(_swapRouter != address(0), "0410");
        swapRouter = _swapRouter;
    }

    function setDefaultTaxRate(uint256 newStakingPoolTaxRate)
        external
        onlyPoolCreator
    {
        require(
            newStakingPoolTaxRate < 10000,
            "0420 Tax connot be over 100% (10000 BP)"
        );
        stakingPoolTaxRate = newStakingPoolTaxRate;
    }
}