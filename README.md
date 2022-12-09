# TotemFi Multi-assets Smart-contracts

This repository consists of smart contracts required for TotemFi Multi-asset DApp with Token contract, Vesting contracts and Staking contracts.

The flow charts provided below are generated with mermaid.js package and can be used to understand the flows of the contracts:

### Admin States
![Admin States](https://github.com/TotemFi/totem-contracts-v2/blob/feature/TOT-754/add-mermaidjs-charts-to-readme/graphs/adminStates.png)

### mermaidjs code to generate this chart:
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


### User States
![User States](https://github.com/TotemFi/totem-contracts-v2/blob/feature/TOT-754/add-mermaidjs-charts-to-readme/graphs/usersStates.png)

### mermaidjs code to generate this chart:
```mermaid
%%{init: {'securityLevel': 'loose', 'theme':'neutral'}}%%
graph TD
    A((User Starts)) --> B(Not Staked)
    B --> C[\stakingAmount - prediction\]
    C --> |TXN|D[[approve]]
    D --> |Success|E[[stake]]
    D --> |Failure|B
    E --> |Success|F(Staked & Pool Is Open) 
    R((Other Users))-.-> F
    E --> |Failure|D
    F --> |Lock Pool|G[Processing Winners & Rewards]
    G --> |End Pool|H(Pool Is Ended)
    H --> |TXN|I[[claimReward]]
    I --> M{Is Winner?}
    M --> |Yes| J{Is Pool USD Prize Enabled?}
    J --> |No| K[Claim Stake and Totem Reward]
    J --> |Yes| L[Claim Stake and Totem Reward and Wrapped Token Reward]
    M --> |No| O[Claim Stake]
    O --> P(UnStaked)
    K --> P
    L --> P
    P --> Q((End))
```

### Staking Pool States
![StakingPool States](https://github.com/TotemFi/totem-contracts-v2/blob/feature/TOT-754/add-mermaidjs-charts-to-readme/graphs/stakingPoolStates.png)

### mermaidjs code to generate this chart:
```mermaid
%%{init: {'securityLevel': 'loose', 'theme':'neutral'}}%%
graph TD
    A((Owner Starts Pool Creation)) --> B[\20 Parameters for Staking Pool\]
    B --> |TXN|C[[StakingPoolFactory: create]]
    C --> |Success|D(Staking Pool Is Created & Is Open)
    C-->|Failure|A
    D--> F[Users Stake And Predict]
    F --> E{Is Pool Locked}
    E--> |No|H{Is LockTime Reached?}
    H --> |Yes/TXN|I[[lockPool]]
    H --> |No| F
    I --> |Success|J(Staking Pool Is Locked)
    J --> Y{Is MaturityTime Reached?}
    Y --> |No|J
    I --> |Failure|E
    E --> |Yes| J
    Y --> |Yes/TXN|K[[updateMaturingPrice]]
    K --> |Success|L(Maturing Price Is Set)
    K --> |Failure|K
    L --> |TXN|M[[setSortedStakers]]
    M --> |Successful|N(Winners Are Sorted)
    M --> |Failure|L 
    N --> Q{Is StakingPool USD Prize Rewarder?}
    Q --> |Yes|S{Is Swap Router Available?}
    S --> |No|X{{Pool Creator Sends The Required Amount Of Wrapped Tokens To Staking Pool's Address}}
    S --> |Yes|O{{Pool Creator Sends USDC Prize Amount To Staking Pool's Address}}
    O --> |Successful| P(USDC Balance Is Available) 
    O --> |Failure|O
    X --> |Failure|X
    X --> |Successful| T
    P --> |TXN|R[[Purchase Wrapped Token]]
    R --> |Successful|T(Wrapped Token Is Available)
    R --> |Failure| P
    T --> |TXN|U[[endPool]]
    U --> |Successful| V[Pool Is Ended & Users Can Claim Their Stakes And Rewards]
    V --> W((End))
    Q --> |No|U
    U --> |Failure|N
```


#### How to run

-   Deploy to testnet (goerli in the below example)

```
yarn
PROJECT_ID="infura id" PRIVATE_KEY"priavete key to deploy the contracts from" yarn deploy:goerli
```

-   Deploy to local (requires ganache to run on local)


You can find more information in the `scripts-info` section in the `package.json` file.


#### How to deploy

-   Deploy to bsc_testnet

```
yarn clean

yarn prepare

NODE_ENV=default PRIVATE_KEY=<your-private-key> yarn deploy:bsc_testnet
```

-   Deploy to bsc mainnet


```
NODE_ENV=production PRIVATE_KEY=<your-private-key> yarn deploy:bsc
```

-   Deploy to Polygon (Matic) mainnet


```
NODE_ENV=polygon_production PRIVATE_KEY=<your-private-key> yarn deploy:matic
```


#### How to verify a contract

- Verify on bsc_testnet

You will first need an API key from your account on https://bscscan.com/login

then run 

```
ETHERSCAN_KEY=<your-api-key> npx hardhat verify --network bsc_testnet <contract-address> "constructor's-first-arg" "constructor's-second-arg" "constructor's-third-arg" "constructor's-forth-arg"
```

#### How to work with smart contracts 

- Information on how to use contracts can be found on the confluence page (the document below should be updated for this new project)

https://totemfi.atlassian.net/wiki/spaces/TP/pages/84574258/Instruction+and+information+for+using+TotemFi+smart+contracts


