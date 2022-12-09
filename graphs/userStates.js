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