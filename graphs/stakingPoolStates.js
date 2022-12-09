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