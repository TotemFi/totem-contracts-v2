```mermaid
%%{init: {'securityLevel': 'loose', 'theme':'neutral'}}%%
graph TD
    A((Owner Starts)) --> B[\9 Vesting Addresses\]
    B--> |TXN|C[[TotemToken.sol: setDistributionTeamAddresses]]
    C-->|Success|D(Vesting Addresses Are Set)
    C --> |Failure|A
    D --> |TXN|E[[TotemToken: distributeTokens]]
    E --> |Successful|F(Totem Vesting Is Done)
    E --> |Failure|D
    F --> G[\StakingPoolFactory Address\]
    G --> |TXN|H[[RewardManager.sol: addOperator]]
    G --> |TXN|I[[RewardManager.sol: addRewarder]]
    H --> J(StakingPoolFactory Has The Required Roles)
    I --> J
    J --> K((Ready To Create Staking Pools))
```