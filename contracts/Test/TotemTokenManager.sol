// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./TotemToken.sol";
import "../Role/Roles.sol";

contract TotemTokenManager is Context, Ownable {

    TotemToken totemToken;

    using Roles for Roles.Role;

    Roles.Role private _managers;

    event ManagerAdded(address indexed account);
    event ManagerRemoved(address indexed account);

    constructor(TotemToken _totemToken) {
        if (!isManager(_msgSender())) {
            _addManager(_msgSender());
        }

        totemToken = _totemToken;
    }

    modifier onlyManager() {
        require(
            isManager(_msgSender()),
            "0100 caller does not have the Manager role"
        );
        _;
    }

    function isManager(address account) public view returns (bool) {
        return _managers.has(account);
    }

    function addManager(address account) public onlyOwner {
        _addManager(account);
    }

    function removeManager(address account) public onlyOwner {
        _removeManager(account);
    }

    function renounceManager() public {
        _removeManager(_msgSender());
    }

    function _addManager(address account) internal {
        _managers.add(account);
        emit ManagerAdded(account);
    }

    function _removeManager(address account) internal {
        _managers.remove(account);
        emit ManagerRemoved(account);
    }

    function setLocker(address _locker) external onlyOwner {
        totemToken.setLocker(_locker);
    }

    function setDistributionTeamsAddresses(
        address _CommunityDevelopmentAddr,
        address _StakingRewardsAddr,
        address _LiquidityPoolAddr,
        address _PublicSaleAddr,
        address _AdvisorsAddr,
        address _SeedInvestmentAddr,
        address _PrivateSaleAddr,
        address _TeamAllocationAddr,
        address _StrategicRoundAddr
    ) public onlyOwner {
        totemToken.setDistributionTeamsAddresses(
            _CommunityDevelopmentAddr,
            _StakingRewardsAddr,
            _LiquidityPoolAddr,
            _PublicSaleAddr,
            _AdvisorsAddr,
            _SeedInvestmentAddr,
            _PrivateSaleAddr,
            _TeamAllocationAddr,
            _StrategicRoundAddr
        );
    }

    function distributeTokens() public onlyOwner {
        totemToken.distributeTokens();
    }

    function setTaxRate(uint256 newTaxRate) public onlyOwner {
        totemToken.setTaxRate(newTaxRate);
    }

    function setTaxExemptStatus(address account, bool status) public onlyManager {
        totemToken.setTaxExemptStatus(account, status);
    }

    function setTaxationWallet(address newTaxationWallet) public onlyOwner {
        totemToken.setTaxationWallet(newTaxationWallet);
    }


    function renounceOwnership() public override onlyOwner {
        // renounceOwnership is over written and do nothing
    }

    function totemTokenTransferOwnership(address newOwner) public onlyOwner {
        totemToken.transferOwnership(newOwner);
    }
}