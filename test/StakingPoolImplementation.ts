import { expect, use } from "chai";
import { deployments, ethers } from "hardhat";
import { Address } from "hardhat-deploy/dist/types";
import { Signer, BigNumberish, BigNumber } from "ethers";
import { deployMockContract,MockContract } from "@ethereum-waffle/mock-contract";
import { solidity } from "ethereum-waffle";
 
import  UniswapRouter from "@uniswap/v2-periphery/build/UniswapV2Router02.json";

import { WrappedERC20Token } from "../src/types/WrappedERC20Token";
import { WrappedERC20Token__factory } from "../src/types/factories/WrappedERC20Token__factory";

import { Contract } from "@ethersproject/contracts";
import { StakingPoolProxy } from "../src/types/StakingPoolProxy";
import { StakingPoolProxy__factory } from "../src/types/factories/StakingPoolProxy__factory";

import { StakingPoolImplementation } from "../src/types/StakingPoolImplementation";
import { StakingPoolImplementation__factory } from "../src/types/factories/StakingPoolImplementation__factory";
import { StakingPoolImplementationLibraryAddresses } from "../src/types/factories/StakingPoolImplementation__factory";
import { StakingPoolImplementationForTest } from "../src/types/StakingPoolImplementationForTest";
import { StakingPoolImplementationForTest__factory } from "../src/types/factories/StakingPoolImplementationForTest__factory";
import { StakingPoolImplementationForTestLibraryAddresses } from "../src/types/factories/StakingPoolImplementationForTest__factory";

import { CalculateRewardLib } from "../src/types/CalculateRewardLib";
import { CalculateRewardLib__factory } from "../src/types/factories/CalculateRewardLib__factory";

import { ClaimRewardLib } from "../src/types/ClaimRewardLib";
import { ClaimRewardLib__factory } from "../src/types/factories/ClaimRewardLib__factory";

import { IndexedClaimRewardLib } from "../src/types/IndexedClaimRewardLib";
import { IndexedClaimRewardLib__factory } from "../src/types/factories/IndexedClaimRewardLib__factory";

import { StakingPoolFactory } from "../src/types/StakingPoolFactory";
import { StakingPoolFactory__factory } from "../src/types/factories/StakingPoolFactory__factory";

import { TotemToken } from "../src/types/TotemToken";
import { TotemToken__factory } from "../src/types/factories/TotemToken__factory";

import { RewardManager } from "../src/types/RewardManager";
import { RewardManager__factory } from "../src/types/factories/RewardManager__factory";

import { WrappedTokenDistributorUpgradeable } from "../src/types/WrappedTokenDistributorUpgradeable";
import { WrappedTokenDistributorUpgradeable__factory } from "../src/types/factories/WrappedTokenDistributorUpgradeable__factory";

const {
    advanceBlockWithTime,
    takeSnapshot,
    revertProvider,
} = require("./utils");

use(solidity);

describe("StakingPoolImplementation", async () => {
    let snapshotId: any;

    let signer1: Signer;
    let signer2: Signer;
    let signer3: Signer;
    let signer4: Signer;
    let superAdmin: Signer;

    let signer1Address: Address;
    let signer2Address: Address;
    let signer3Address: Address;
    let signer4Address: Address;
    let superAdminAddress: Address;

    let totemToken: TotemToken;
    let rewardManager: RewardManager;

    let wrappedTokenDistributor: WrappedTokenDistributorUpgradeable;
    let mockUniswapRouter: MockContract;

    let mockAggregatorContract: MockContract;
    let mockAggregatorContractForPrize: MockContract;

    let usdContract: WrappedERC20Token;
    let bep20Token: WrappedERC20Token;

    let stakingPoolImplementation: StakingPoolImplementation;
    let stakingPoolImplementationForTest: StakingPoolImplementationForTest;
    let stakingPoolProxy: StakingPoolProxy;
    let stakingPoolFactoryInstance: StakingPoolFactory;
    let stakingPool: Contract;

    const zeroUSDVariables: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
    ] = [
        0,
        1296000,
        432000,
        BigNumber.from(10).pow(18).mul(125000),
        6000,
        BigNumber.from(10).pow(18).mul(4500),
        0,
        2000,
        175,
        300,
        BigNumber.from(10).pow(18).mul(250),
        8
    ];

    const variables: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
    ] = [
        0,
        1296000,
        432000,
        BigNumber.from(10).pow(18).mul(125000),
        6000,
        BigNumber.from(10).pow(18).mul(4500),
        BigNumber.from(10).pow(18).mul(1000),
        2000,
        175,
        300,
        BigNumber.from(10).pow(18).mul(250),
        8
    ];

    let ranks: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
    ] = [1, 2, 3, 10, 25, 0, 0, 0];

    let percentages: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
    ] = [3750, 2000, 1000, 250, 100, 0, 0, 0];

    const nilAddress = "0x0000000000000000000000000000000000000000";

    const zeroAddresses: [
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string
    ] = [
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
        nilAddress,
    ];

    const zeroIndexes: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
    ] = [
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0,
    ];

    const fiftyThousandBTCPrice = BigNumber.from(10).pow(8).mul(50000);
    const fiftyFiveThousandBTCPrice = BigNumber.from(10).pow(8).mul(55000);
    const sixtyThousandBTCPrice = BigNumber.from(10).pow(8).mul(60000);

    const oneHundred = BigNumber.from(10).pow(18).mul(100);
    const oneHundredTaxed = BigNumber.from(10).pow(18).mul(97);

    const oneThousand = BigNumber.from(10).pow(18).mul(1000);
    const oneThousandTaxed = BigNumber.from(10).pow(18).mul(970);

    const threeThousand = BigNumber.from(10).pow(18).mul(3000);
    const threeThousandTaxed = BigNumber.from(10).pow(18).mul(2900);

    const fiveThousand = BigNumber.from(10).pow(18).mul(5000);
    const fiveThousandTaxed = BigNumber.from(10).pow(18).mul(4850);

    const tenThousand = BigNumber.from(10).pow(18).mul(10000);
    const tenThousandTaxed = BigNumber.from(10).pow(18).mul(9700);

    const hundredThousand = BigNumber.from(10).pow(18).mul(100000);
    const hundredThousandTaxed = BigNumber.from(10).pow(18).mul(97000);

    const deployRewardManager = async (
        _signer?: Signer
    ): Promise<RewardManager> => {
        const rewardManagerFactory = new RewardManager__factory(
            _signer || signer1
        );
        const totemTokenInstance = await rewardManagerFactory.deploy(
            totemToken.address
        );
        return totemTokenInstance;
    };

    const deployTotemToken = async (_signer?: Signer): Promise<TotemToken> => {
        const totemTokenFactory = new TotemToken__factory(_signer || signer1);
        const totemToken = await totemTokenFactory.deploy();
        return totemToken;
    };

    const deploybep20Token = async (
        _signer?: Signer
    ): Promise<WrappedERC20Token> => {
        const bep20TokenFactory = new WrappedERC20Token__factory(
            _signer || signer1
        );
        const bep20Token = await bep20TokenFactory.deploy(
            "Wrapped Token",
            "WTKN"
        );
        return bep20Token;
    };

    const deployFiatTokenV2 = async (
        _signer?: Signer
    ): Promise<WrappedERC20Token> => {
        const fiatTokenV2Factory = new WrappedERC20Token__factory(
            _signer || signer1
        );
        const fiatTokenV2 = await fiatTokenV2Factory.deploy(
            "Binance USD Token",
            "BUSD"
        );
        return fiatTokenV2;
    };

    const deployWrappedTokenDistributor = async (
        _signer?: Signer
    ): Promise<WrappedTokenDistributorUpgradeable> => {
        const WrappedTokenDistributorFactory =
            new WrappedTokenDistributorUpgradeable__factory(_signer || signer1);
        const WrappedTokenDistributor =
            await WrappedTokenDistributorFactory.deploy();

        WrappedTokenDistributor.__WrappedTokenDistributor_initialize(
            mockUniswapRouter.address,
            usdContract.address,
            bep20Token.address
        );

        return WrappedTokenDistributor;
    };

    const deployCalculateRewardLib = async (
        _signer?: Signer
    ): Promise<CalculateRewardLib> => {
        const CalculateRewardLibFactory = new CalculateRewardLib__factory(
            _signer || signer1
        );
        const CalculateRewardLib = await CalculateRewardLibFactory.deploy();

        return CalculateRewardLib;
    };

    const deployClaimRewardLib = async (
        _signer?: Signer
    ): Promise<ClaimRewardLib> => {
        const ClaimRewardLibFactory = new ClaimRewardLib__factory(
            _signer || signer1
        );
        const ClaimRewardLib = await ClaimRewardLibFactory.deploy();

        return ClaimRewardLib;
    };

    const deployIndexedClaimRewardLib = async (
        _signer?: Signer
    ): Promise<IndexedClaimRewardLib> => {
        const IndexedClaimRewardLibFactory = new IndexedClaimRewardLib__factory(
            _signer || signer1
        );
        const IndexedClaimRewardLib =
            await IndexedClaimRewardLibFactory.deploy();

        return IndexedClaimRewardLib;
    };

    const getStakingPoolImplementationFactory = async (
        _signer?: Signer
    ): Promise<StakingPoolImplementation__factory> => {
        const CalculateRewardLib = await deployCalculateRewardLib();
        const ClaimRewardLib = await deployClaimRewardLib();
        const IndexedClaimRewardLib = await deployIndexedClaimRewardLib();

        let linkLibraryAddresses: StakingPoolImplementationLibraryAddresses;
        
        linkLibraryAddresses = {
            __$60b6e19b11029f04c25a43d0435d86b2be$__: CalculateRewardLib.address,
            __$171d2da81d2a08eb52330533595f0c4998$__: ClaimRewardLib.address,
            __$7fd944d8f8efff7c575121b9e933203c2c$__: IndexedClaimRewardLib.address,
        };
        const StakingPoolImplementationFactory =
            new StakingPoolImplementation__factory(
                linkLibraryAddresses,
                _signer || signer1
            );

        return StakingPoolImplementationFactory;
    };

    const getStakingPoolImplementationV2Factory = async (
        _signer?: Signer
    ): Promise<StakingPoolImplementationForTest__factory> => {
        const CalculateRewardLib = await deployCalculateRewardLib();
        const ClaimRewardLib = await deployClaimRewardLib();
        const IndexedClaimRewardLib = await deployIndexedClaimRewardLib();

        let linkLibraryAddresses: StakingPoolImplementationForTestLibraryAddresses;
        
        linkLibraryAddresses = {
            __$60b6e19b11029f04c25a43d0435d86b2be$__: CalculateRewardLib.address,
            __$171d2da81d2a08eb52330533595f0c4998$__: ClaimRewardLib.address,
            __$7fd944d8f8efff7c575121b9e933203c2c$__: IndexedClaimRewardLib.address,
        };
        const StakingPoolImplementationV2Factory =
            new StakingPoolImplementationForTest__factory(
                linkLibraryAddresses,
                _signer || signer1
            );

        return StakingPoolImplementationV2Factory;
    };

    const deployStakingPoolImplementation = async (
        _signer?: Signer
    ): Promise<StakingPoolImplementation> => {
        const StakingPoolImplementationFactory =
            await getStakingPoolImplementationFactory();

        const stakingPoolImplementationInstance =
            await StakingPoolImplementationFactory.deploy();

        return stakingPoolImplementationInstance;
    };

    const deployStakingPoolFactory = async (
        _signer?: Signer
    ): Promise<StakingPoolFactory> => {
        const stakingPoolFactoryFactory = new StakingPoolFactory__factory(
            _signer || signer1
        );
        const stakingPoolFactoryInstance =
            await stakingPoolFactoryFactory.deploy(
                totemToken.address,
                rewardManager.address,
                mockUniswapRouter.address,
                usdContract.address,
                stakingPoolImplementation.address,
                superAdminAddress
            );
        return stakingPoolFactoryInstance;
    };

    const deployStakingPoolProxy = async (
        _signer?: Signer
    ): Promise<StakingPoolProxy> => {
        const stakingPoolProxyFactory = new StakingPoolProxy__factory(
            _signer || signer1
        );

        let launchDate = Math.round(new Date().getTime() / 1000) + 5000;
        variables[0] = launchDate;

        const stakingPoolProxy = await stakingPoolProxyFactory.deploy();

        await stakingPoolProxy.upgradeTo(stakingPoolImplementation.address);

        await mockAggregatorContract.mock.decimals.returns(8);

        await stakingPoolProxy.initialize(
            "ETH",
            "Wolf",
            totemToken.address,
            rewardManager.address,
            signer1Address,
            [
                mockUniswapRouter.address,
                usdContract.address,
                mockAggregatorContract.address,
                bep20Token.address,
            ],
            variables,
            ranks,
            percentages,
            true
        );

        return stakingPoolProxy;
    };

    const deployStakingPoolProxyWithFactory = async(): Promise<any[]> => {
        let stakingPool: Contract;

        let launchDate = Math.round(new Date().getTime() / 1000) + 5000;
            variables[0] = launchDate;

        
        await mockAggregatorContract.mock.decimals.returns(8);
        
        let stakingPoolProxy = await stakingPoolFactoryInstance.createPoolProxy(
            mockAggregatorContract.address,
            bep20Token.address,
            "ETH",
            "Wolf",
            variables,
            ranks,
            percentages,
            false
        );

            const stakingPoolProxyAdr = (await (stakingPoolProxy).wait()).logs[0].address;

            const stakingPoolProxyFactory = new StakingPoolProxy__factory(
                signer1
            );

            stakingPool = stakingPoolProxyFactory.attach(
                stakingPoolProxyAdr
            );

            let returnValue = [stakingPool, stakingPoolProxyAdr];

            return returnValue;
    };

    const deployStakingPoolProxyWithNewVars = async(): Promise<any> => {
        const stakingPoolProxyInstance = new StakingPoolProxy__factory(
            signer1
        );

        let launchDate = Math.round(new Date().getTime() / 1000) + 5000;
        variables[0] = launchDate;

        const stakingPoolProxy = await stakingPoolProxyInstance.deploy();

        stakingPoolProxy.initialize(
            "ETH",
            "Wolf",
            totemToken.address,
            rewardManager.address,
            signer1Address,
            [
                mockUniswapRouter.address,
                usdContract.address,
                nilAddress,
                bep20Token.address,
            ],
            variables,
            ranks,
            percentages,
            true
        );

        await rewardManager.addPool(stakingPoolProxy.address);
        await totemToken.setTaxExemptStatus(stakingPoolProxy.address, true);

        stakingPoolProxy.upgradeTo(stakingPoolImplementation.address);

        const StakingPoolImplementationFactory =
            await getStakingPoolImplementationFactory();

        let stakingPoolInstance: Contract
        stakingPoolInstance = StakingPoolImplementationFactory.attach(
            stakingPoolProxy.address
        );

        return stakingPoolInstance;
    }

    before("deploy contracts", async () => {
        [signer1, signer2, signer3, signer4, superAdmin] = await ethers.getSigners();

        signer1Address = await signer1.getAddress();
        signer2Address = await signer2.getAddress();
        signer3Address = await signer3.getAddress();
        signer4Address = await signer4.getAddress();
        superAdminAddress = await superAdmin.getAddress();

        mockUniswapRouter = await deployMockContract(
            signer1,
            UniswapRouter.abi
        );

        const aggregatorContract = await deployments.getArtifact(
            "AggregatorV3Interface"
        );
        mockAggregatorContract = await deployMockContract(
            signer1,
            aggregatorContract.abi
        );

        mockAggregatorContractForPrize = await deployMockContract(
            signer1,
            aggregatorContract.abi
        );

        bep20Token = await deploybep20Token();
        usdContract = await deployFiatTokenV2();

        totemToken = await deployTotemToken();
        rewardManager = await deployRewardManager();
        wrappedTokenDistributor = await deployWrappedTokenDistributor();

        stakingPoolImplementation = await deployStakingPoolImplementation();
        stakingPoolFactoryInstance = await deployStakingPoolFactory();

        await totemToken.setDistributionTeamsAddresses(
            signer1Address,
            rewardManager.address,
            signer1Address,
            signer1Address,
            signer1Address,
            signer1Address,
            signer1Address,
            signer1Address,
            signer1Address
        );
        await totemToken.distributeTokens();

        let launchDate = Math.round(new Date().getTime() / 1000) + 5000;
        variables[0] = launchDate;

        await rewardManager.addOperator(stakingPoolFactoryInstance.address);
        await rewardManager.addRewarder(stakingPoolFactoryInstance.address);
    });

    beforeEach("deploy a new staking pool proxy", async () => {
        snapshotId = await takeSnapshot(signer1.provider);

        stakingPoolProxy = await deployStakingPoolProxy();

        stakingPool = stakingPoolImplementation.attach(
            stakingPoolProxy.address
        );

        await rewardManager.addPool(stakingPoolProxy.address);
        await totemToken.setTaxExemptStatus(stakingPoolProxy.address, true);
    });

    afterEach(async () => {
        await revertProvider(signer1.provider, snapshotId);
    });

    describe("StakingPoolFactory: createPoolProxy", async () => {
        beforeEach("getting time for launch date", () => {
            let launchDate = Math.round(new Date().getTime() / 1000) + 5000;
            variables[0] = launchDate;
        });

        it("only poolCreator accounts can create new pools", async () => {
            let signer2FactoryInstance = stakingPoolFactoryInstance.connect(signer2)
      
            await expect(
                signer2FactoryInstance.createPoolProxy(
                    mockAggregatorContract.address,
                    bep20Token.address,
                    "ETH",
                    "Wolf",
                    variables,
                    ranks,
                    percentages,
                    false
                )
            ).to.be.revertedWith("PoolCreatorRole: caller does not have the PoolCreator role")
        });

        it("successfully launching a new staking pool proxy", async () => {
            await expect(
                stakingPoolFactoryInstance.createPoolProxy(
                    mockAggregatorContract.address,
                    bep20Token.address,
                    "ETH",
                    "Wolf",
                    variables,
                    ranks,
                    percentages,
                    false
                )
            ).to.emit(stakingPoolFactoryInstance, "PoolCreated");
        });
        it("launching a staking pool proxy with no oracle", async () => {
            await expect(
                stakingPoolFactoryInstance.createPoolProxy(
                    nilAddress,
                    bep20Token.address,
                    "ETH",
                    "Wolf",
                    variables,
                    ranks,
                    percentages,
                    false
                )
            ).to.emit(stakingPoolFactoryInstance, "PoolCreated");
        });

        describe("zero usdPrizeAmount", async() => {
            before("changing usdPrizeAmount to zero", () => {
                variables[6] = BigNumber.from(0);
            });
            it("launching a TOTM-only prize pool proxy, no wrapped token, no oracle", async () => {
                await expect(
                    stakingPoolFactoryInstance.createPoolProxy(
                        nilAddress,
                        bep20Token.address,
                        "ETH",
                        "Wolf",
                        variables,
                        ranks,
                        percentages,
                        false
                    )
                ).to.emit(stakingPoolFactoryInstance, "PoolCreated");
            });
            after("changing usdPrizeAmount back to default", () => {
                variables[6] = BigNumber.from(10).pow(18).mul(1000);
            });
        })

        describe("custom oracle decimal", async() => {
            before("changing oracle decimal",() => {
                variables[11] = 6;
            })
            it("adjust a custom oracle decimal when oracle is zero", async () => {
                
                // let stakingPoolInstance: [Contract];

                let stakingPoolInstance = await deployStakingPoolProxyWithNewVars();
                expect(await stakingPoolInstance.oracleDecimals()).to.equal(6);
            });
            after("changing oracle decimal back to 8",() => {
                variables[11] = 8;
            })
        });

        describe("ranks and percentages check", async () => {
            let newRanks: [
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish
            ];
            let newPercentages: [
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish,
                BigNumberish
            ];
            beforeEach("using new ranks and percentages", () => {
                newRanks = [1, 2, 3, 10, 25, 0, 0, 0];
                newPercentages = [3750, 2000, 1000, 250, 100, 0, 0, 0];
            })

            it("reverts because ranks are not correctly ordered", async() => {
                newRanks[1] = 3;
                newRanks[2] = 2;

                await expect(
                    stakingPoolFactoryInstance.createPoolProxy(
                        mockAggregatorContract.address,
                        bep20Token.address,
                        "ETH",
                        "Wolf",
                        variables,
                        newRanks,
                        newPercentages,
                        false
                    )
                ).to.be.revertedWith("wrong order of ranks")

            })
            
            it("reverts because the last rank is not 25", async () => {
                newRanks[4] = 26;

                await expect(
                    stakingPoolFactoryInstance.createPoolProxy(
                        mockAggregatorContract.address,
                        bep20Token.address,
                        "ETH",
                        "Wolf",
                        variables,
                        newRanks,
                        newPercentages,
                        false
                    )
                ).to.be.revertedWith("last rank must be less than or equal to 25")
            })
            
            it("reverts because sum of reward percentage is not equal to 100%", async () => {
                newPercentages[4] = 90;

                await expect(
                    stakingPoolFactoryInstance.createPoolProxy(
                        mockAggregatorContract.address,
                        bep20Token.address,
                        "ETH",
                        "Wolf",
                        variables,
                        newRanks,
                        newPercentages,
                        false
                    )
                ).to.be.revertedWith("reward percentages must be 100%")
            })
            
            it("reverts because rank and percentage have length mismatch", async () => {
                newPercentages[4] = 0
                await expect(
                    stakingPoolFactoryInstance.createPoolProxy(
                        mockAggregatorContract.address,
                        bep20Token.address,
                        "ETH",
                        "Wolf",
                        variables,
                        newRanks,
                        newPercentages,
                        false
                    )
                ).to.be.revertedWith("ranks and percentages length mismatch");
            })
        })
    });

    describe("StakingPoolProxy", async() => {
        let stakingPoolInstance: Contract;
        let stakingPoolProxyAdr: string;
        beforeEach("deploy a new stakingPoolProxy within the factory", async() => {
            [stakingPoolInstance, stakingPoolProxyAdr] = await deployStakingPoolProxyWithFactory();
        });
        it("checks if the stakingPoolProxy owner is the superOwner", async() => {

            expect(await stakingPoolInstance.owner()).to.equal(superAdminAddress);

            expect(await stakingPoolInstance.poolCreator()).to.equal(signer1Address);
        });

        it("can not upgrade implementation because upgrade is not enabled", async() => {

            await expect(
                stakingPoolInstance.upgradeTo(stakingPoolImplementation.address)
            ).to.be.revertedWith("Upgrade is not enabled yet");

            expect(await stakingPoolInstance.upgradeEnabled()).to.be.false;
        });

        it("no one can call enableUpgrade other than the superAdmin", async() => {

            await expect(
                stakingPoolInstance.enableUpgrade()
            ).to.be.revertedWith("Ownable: caller is not the owner");

            let superAdminAsStakingPool = stakingPoolInstance.connect(superAdmin);

            await superAdminAsStakingPool.enableUpgrade();

            expect(await stakingPoolInstance.upgradeEnabled()).to.be.true;
        });

        it("can not upgrade because the address provided is the old implementation address", async() => {
            let superAdminAsStakingPool = stakingPoolInstance.connect(superAdmin);

            await superAdminAsStakingPool.enableUpgrade();

            await expect(
                stakingPoolInstance.upgradeTo(stakingPoolImplementation.address)
            ).to.be.revertedWith("Is already the implementation")
        });

        it("successfully upgrades stakingPoolImplementation", async() => {
            let superAdminAsStakingPool = stakingPoolInstance.connect(superAdmin);

            await superAdminAsStakingPool.enableUpgrade();

            const StakingPoolImplementationV2Factory =
                await getStakingPoolImplementationV2Factory();

            const stakingPoolImplementationV2Instance =
                await StakingPoolImplementationV2Factory.deploy();

            await expect(
                stakingPoolInstance.upgradeTo(stakingPoolImplementationV2Instance.address)
            ).to.emit(stakingPoolInstance, "ImplementationUpgraded")
        });

        it("unlocks a locked pool by upgrading the implementation", async() => {

            let stakingPool = stakingPoolImplementation.attach(
                stakingPoolProxyAdr
            );

            await totemToken.setTaxExemptStatus(stakingPoolProxyAdr, true);

            await advanceBlockWithTime(signer1.provider, 5050);

            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner1 = stakingPool.connect(signer1);

            expect(await stakingPoolAsSigner1.lockPool()).to.emit(stakingPool, "PoolLocked");

            let isLocked = await stakingPool.isLocked();

            expect(isLocked).to.be.true;

            // pool is locked and staking reverts
            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            let totemTokenAsSigner3 = totemToken.connect(signer3);
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            await expect(
                stakingPoolAsSigner3.stake(
                    BigNumber.from(10).pow(18).mul(1000),
                    BigNumber.from(10).pow(8).mul(60000)
                )
            ).to.be.revertedWith("0310 Pool is locked");

            // unlocking the pool
            let superAdminAsStakingPool = stakingPoolInstance.connect(superAdmin);

            await superAdminAsStakingPool.enableUpgrade();

            const StakingPoolImplementationV2Factory =
                await getStakingPoolImplementationV2Factory();

            const stakingPoolImplementationV2Instance =
                await StakingPoolImplementationV2Factory.deploy();

            await expect(
                stakingPoolInstance.upgradeTo(stakingPoolImplementationV2Instance.address)
            ).to.emit(stakingPoolInstance, "ImplementationUpgraded")


            stakingPool = stakingPoolImplementationV2Instance.attach(
                stakingPoolProxyAdr
            );

            stakingPoolAsSigner1 = stakingPool.connect(signer1);

            expect(await stakingPoolAsSigner1.lockPool()).to.emit(stakingPool, "PoolUnLocked");

            isLocked = await stakingPool.isLocked();

            expect(isLocked).to.be.false;

        });

        it('transfers ownership from signer1 to superAdmin', async() => {
            await stakingPoolProxy.transferOwnership(superAdminAddress);
        });
    })

    describe("setActivationStatus", async () => {

        it("pool can't be activated by a non-operator account ", async () => {
          let signer2stakingPool = stakingPool.connect(signer2)
    
          await expect(
            signer2stakingPool.setActivationStatus(true)
          ).to.be.revertedWith("0300 caller is not a pool creator")
        })
    
        it("pool can be activated by the operators whenever they wants", async () => {
          await stakingPool.setActivationStatus(true)
    
          expect(
            await stakingPool.isActive()
          ).to.equal(true)
        })
    })

    describe("stake", async () => {
        it("users can't stake when the launch date is not reached", async () => {
            let signer2TotemToken = totemToken.connect(signer2);
            let signer2stakingPool = stakingPool.connect(signer2);

            await totemToken.transfer(signer2Address, oneHundred);

            await signer2TotemToken.approve(stakingPool.address, oneHundred);
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(oneHundred);

            await expect(
                signer2stakingPool.stake(oneHundred, fiftyThousandBTCPrice)
            ).to.be.revertedWith("0313 pool is not active");
        });

        it("users can't stake when we reach to launch date but the pool is not activated", async () => {
            let signer2TotemToken = totemToken.connect(signer2);
            let signer2stakingPool = stakingPool.connect(signer2);

            await totemToken.transfer(signer2Address, oneHundred);

            await advanceBlockWithTime(signer1.provider, 5001);

            await signer2TotemToken.approve(stakingPool.address, oneHundred);
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(oneHundred);

            await expect(
                signer2stakingPool.stake(oneHundred, fiftyThousandBTCPrice)
            ).to.be.revertedWith("0313 pool is not active");
        });

        it("users can stake when the pool is activated and is not locked", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);
      
            let signer2TotemToken = totemToken.connect(signer2)
            let signer2stakingPool = stakingPool.connect(signer2)
      
            await totemToken.transfer(signer2Address, oneThousand)
      
            await signer2TotemToken.approve(stakingPool.address, oneThousand)
            expect(
              await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(oneThousand)
      
            expect (
              await signer2stakingPool.stake(oneThousand, fiftyThousandBTCPrice)
            ).to.emit(signer2stakingPool, "Stake")
        })

        it("successfully staking in the launched pool", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(25000)
            );

            let totemTokenAsSigner2 = totemToken.connect(signer2);
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(25000)
            );

            let userBalanceBeforeStake = await totemToken.balanceOf(
                signer2Address
            );
            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            await expect(
                stakingPoolAsSigner2.stake(
                    BigNumber.from(10).pow(18).mul(25000),
                    BigNumber.from(10).pow(8).mul(60005)
                )
            ).to.emit(stakingPoolAsSigner2, "Stake");

            let userBalanceAfterStake = await totemToken.balanceOf(
                signer2Address
            );
            expect(userBalanceBeforeStake.sub(userBalanceAfterStake)).to.equal(
                BigNumber.from(10).pow(18).mul(25000)
            );

            let prediction = await stakingPoolAsSigner2.predictions(
                signer2Address,
                0
            );

            // Staking Tax is 3%
            expect(prediction.stakedBalance).to.equal(
                BigNumber.from(10).pow(18).mul(24250)
            );
            expect(prediction.pricePrediction).to.equal(
                BigNumber.from(10).pow(8).mul(60005)
            );
        });

        it("reverts because can't stake less than the minimum staking amount", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(200)
            );

            let totemTokenAsSigner3 = totemToken.connect(signer3);
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(200)
            );

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);

            await expect(
                stakingPoolAsSigner3.stake(
                    BigNumber.from(10).pow(18).mul(200),
                    BigNumber.from(10).pow(8).mul(60000)
                )
            ).to.be.revertedWith("0311 Amount can't be less than the minimum");

            let stakingPooltotalStakedBalance = await stakingPool.totalStaked();
            expect(stakingPooltotalStakedBalance).to.equal(BigNumber.from(0));
        });

        it("reverts because can't stake above size allocation + limitRange", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            // size allocation is 125000 + 250
            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(130000)
            );

            let totemTokenAsSigner3 = totemToken.connect(signer3);
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(130000)
            );

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);

            await expect(
                stakingPoolAsSigner3.stake(
                    BigNumber.from(10).pow(18).mul(130000),
                    BigNumber.from(10).pow(8).mul(60000)
                )
            ).to.be.revertedWith("0312 Can't stake above size allocation");

            let stakingPoolTotalStakedBalance = await stakingPool.totalStaked();
            expect(stakingPoolTotalStakedBalance).to.equal(BigNumber.from(0));

            // it is within the limit range, so staking is successful (129000*0.97=125130)
            await expect(
                stakingPoolAsSigner3.stake(
                    BigNumber.from(10).pow(18).mul(129000),
                    BigNumber.from(10).pow(8).mul(60000)
                )
            ).to.emit(stakingPoolAsSigner3, "Stake");

            stakingPoolTotalStakedBalance = await stakingPool.totalStaked();
            expect(stakingPoolTotalStakedBalance).to.equal(
                BigNumber.from(10).pow(18).mul(125130)
            );
        });

        it("pool gets locked and can't stake anymore", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(131000)
            );

            let totemTokenAsSigner2 = totemToken.connect(signer2);
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(131000)
            );

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            await expect(
                stakingPoolAsSigner2.stake(
                    BigNumber.from(10).pow(18).mul(128864),
                    BigNumber.from(10).pow(8).mul(60000)
                )
            ).to.emit(stakingPoolAsSigner2, "Stake");

            // can stake more and pool is not locked
            await expect(
                stakingPoolAsSigner2.stake(
                    BigNumber.from(10).pow(18).mul(251),
                    BigNumber.from(10).pow(8).mul(60000)
                )
            ).to.emit(stakingPoolAsSigner2, "Stake");

            // staked amount is above size allocation so the pool is locked
            let isLocked = await stakingPool.isLocked();
            expect(isLocked).to.be.true;

            // pool is locked and staking reverts

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            await expect(
                stakingPoolAsSigner3.stake(
                    BigNumber.from(10).pow(18).mul(1000),
                    BigNumber.from(10).pow(8).mul(60000)
                )
            ).to.be.revertedWith("0310 Pool is locked");
        });
    });

    describe("batchStake", async () => {
        const stakingAmounts: [
            BigNumberish,
            BigNumberish,
            BigNumberish
        ] = [
            oneThousand,
            oneThousand,
            oneThousand
        ];

        const predictions: [
            BigNumberish,
            BigNumberish,
            BigNumberish
        ] = [
            fiftyThousandBTCPrice,
            sixtyThousandBTCPrice,
            fiftyFiveThousandBTCPrice
        ]
        it("users can't batch stake when the launch date is not reached", async () => {
            let signer2TotemToken = totemToken.connect(signer2);
            let signer2stakingPool = stakingPool.connect(signer2);

            await totemToken.transfer(signer2Address, threeThousand);

            await signer2TotemToken.approve(stakingPool.address, threeThousand);
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(threeThousand);

            await expect(
                signer2stakingPool.batchStake(stakingAmounts, predictions)
            ).to.be.revertedWith("0313 pool is not active");
        });

        it("successfully batch staking in the launched pool", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            await totemToken.transfer(
                signer2Address,
                threeThousand
            );

            let totemTokenAsSigner2 = totemToken.connect(signer2);
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                threeThousand
            );

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            await expect(
                stakingPoolAsSigner2.batchStake(
                    stakingAmounts,
                    predictions
                )
            ).to.emit(stakingPoolAsSigner2, "BatchStake");

            let prediction1 = await stakingPoolAsSigner2.predictions(
                signer2Address,
                0
            );

            // Staking Tax is 3%
            expect(prediction1.stakedBalance).to.equal(
                oneThousandTaxed
            );
            expect(prediction1.pricePrediction).to.equal(
                predictions[0]
            );

            let prediction2 = await stakingPoolAsSigner2.predictions(
                signer2Address,
                1
            );

            // Staking Tax is 3%
            expect(prediction2.stakedBalance).to.equal(
                oneThousandTaxed
            );
            expect(prediction2.pricePrediction).to.equal(
                predictions[1]
            );

            let stakingPooltotalStakedBalance = await stakingPool.totalStaked();
            expect(stakingPooltotalStakedBalance).to.equal(BigNumber.from(10).pow(18).mul(2910));
        });

        it("reverts because stakingAmount and predictions arrays don't have the same length", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let signer2TotemToken = totemToken.connect(signer2);
            let signer2stakingPool = stakingPool.connect(signer2);

            const stakingAmounts: [
                BigNumberish,
                BigNumberish
            ] = [
                oneThousand,
                oneThousand
            ];

            await totemToken.transfer(signer2Address, threeThousand);

            await signer2TotemToken.approve(stakingPool.address, threeThousand);
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(threeThousand);

            await expect(
                signer2stakingPool.batchStake(stakingAmounts, predictions)
            ).to.be.revertedWith("0315 stakingAmount and predictions length mismatch");
        });

        it("reverts because can't batch stake less than the minimum staking amount", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let signer2TotemToken = totemToken.connect(signer2);
            let signer2stakingPool = stakingPool.connect(signer2);

            const stakingAmounts: [
                BigNumberish,
                BigNumberish,
                BigNumberish
            ] = [
                BigNumber.from(10).pow(18).mul(200),
                oneThousand,
                oneThousand
            ];  

            await totemToken.transfer(signer2Address, threeThousand);

            await signer2TotemToken.approve(stakingPool.address, threeThousand);
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(threeThousand);

            await expect(
                signer2stakingPool.batchStake(stakingAmounts, predictions)
            ).to.be.revertedWith("0311 Amount can't be less than the minimum");
        });

        it("reverts because can't batch stake above size allocation", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            // size allocation is 125000, limit range is 135309
            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(136310)
            );

            let totemTokenAsSigner3 = totemToken.connect(signer3);
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(136310)
            );

            const stakingAmounts: [
                BigNumberish,
                BigNumberish,
                BigNumberish
            ] = [
                BigNumber.from(10).pow(18).mul(100000),
                BigNumber.from(10).pow(18).mul(35310),
                oneThousand
            ];  

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);

            await expect(
                stakingPoolAsSigner3.batchStake(stakingAmounts, predictions)
            ).to.be.revertedWith("0312 Can't stake above size allocation");
        });
    });

    describe("lockPool", async () => {
        it("pool can't be locked by a non-operator account ", async () => {
            let signer2stakingPool = stakingPool.connect(signer2)
            await expect(
                signer2stakingPool.lockPool()
            ).to.be.revertedWith("0300 caller is not a pool creator")
        })
    
        it("pool can be locked by operators whenever they want", async () => {
            await expect(
            stakingPool.lockPool()
          ).to.emit(stakingPool, "PoolLocked")
        })
    })

    describe("claimReward", async() => {
        it("non-staker accouts shouldn't get anything when calling claimReward", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);
        
            let signer2TotemToken = await totemToken.connect(signer2);
        
            let signer2stakingPool = await stakingPool.connect(signer2);
            let signer3stakingPool = await stakingPool.connect(signer3);
        
            await totemToken.transfer(signer2Address, oneThousand);
        
            await signer2TotemToken.approve(stakingPool.address, oneThousand);
            await signer2stakingPool.stake(oneThousand, fiftyThousandBTCPrice);
        
            expect(await totemToken.balanceOf(signer3Address)).to.equal(0);
        
            await signer3stakingPool.claimReward();
        
            expect(await totemToken.balanceOf(signer3Address)).to.equal(0);
        });
        
        it("successfully claims reward", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            let totemTokenAsSigner2 = totemToken.connect(signer2);

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            let totemTokenAsSigner3 = totemToken.connect(signer3);

            let stakingPoolAsSigner4 = stakingPool.connect(signer4);
            let totemTokenAsSigner4 = totemToken.connect(signer4);

            await mockAggregatorContract.mock.decimals.returns(8);
            await mockAggregatorContract.mock.latestRoundData.returns(
                "73786976294838208680",
                BigNumber.from(10).pow(8).mul(60000),
                "1620053717",
                "1620053717",
                "73786976294838208680"
            );

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner2.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60020)
            );

            await advanceBlockWithTime(signer1.provider, 1000);

            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer3Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner3.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(59995)
            );

            await advanceBlockWithTime(signer1.provider, 1000);

            await totemToken.transfer(
                signer4Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner4.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer4Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner4.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60001)
            );

            let stakingPooltotalStakedBalance = await stakingPool.totalStaked();
            expect(stakingPooltotalStakedBalance).to.equal(
                BigNumber.from(10).pow(18).mul(2910)
            );

            await bep20Token.transfer(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            await advanceBlockWithTime(signer1.provider, 430000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            let newAddresses = zeroAddresses;
            newAddresses[0] = signer2Address;
            newAddresses[1] = signer3Address;
            newAddresses[2] = signer4Address;

            // let sortedStakersIndexed = zeroIndexes

            await stakingPool.updateMaturingPrice(
                BigNumber.from(10).pow(18).mul(50000),
                true,
                0,
                nilAddress,
                0
            );
            await stakingPool.setSortedStakers(newAddresses, zeroIndexes);
            await stakingPool.endPool();

            let isCollabReward =
                await stakingPoolAsSigner3.collaborativeReward();
            if (isCollabReward) {
                //First winner reward checks
                let winner = await stakingPool.sortedStakers(0);
                expect(winner.stakerAddress).to.equal(signer2Address);

                expect(await stakingPool.maturityTime()).to.equal(1296000);

                expect(await stakingPool.lockTime()).to.equal(432000);

                expect(await stakingPool.sizeAllocation()).to.equal(
                    BigNumber.from(10).pow(18).mul(125000)
                );

                expect(await stakingPool.stakeApr()).to.equal(6000);

                expect(await stakingPool.prizeAmount()).to.equal(
                    BigNumber.from(10).pow(18).mul(4500)
                );

                await stakingPoolAsSigner2.claimReward();
                expect(
                    (
                        await totemToken.balanceOf(signer2Address)
                    ).div(10 ** 15).toNumber()
                ).to.be.lessThan(2685090);

                expect(await bep20Token.balanceOf(signer2Address)).to.equal(
                    "6249999999999999"
                );

                //Second winner rewards check
                let secondWinner = await stakingPool.sortedStakers(1);
                expect(secondWinner.stakerAddress).to.equal(signer3Address);

                await stakingPoolAsSigner3.claimReward();
                expect(
                    (
                        await totemToken.balanceOf(secondWinner.stakerAddress)
                    ).div(10 ** 15).toNumber()
                ).to.be.lessThan(1897575);
                expect(
                    await bep20Token.balanceOf(secondWinner.stakerAddress)
                ).to.equal("3333333333333333");
                //Third winner rewards check
                let thirdWinner = await stakingPool.sortedStakers(2);
                expect(thirdWinner.stakerAddress).to.equal(signer4Address);

                await stakingPoolAsSigner4.claimReward();
                expect(
                    (
                        await totemToken.balanceOf(thirdWinner.stakerAddress)
                    ).div(10 ** 15).toNumber()
                ).to.be.lessThan(1447560);
                expect(
                    await bep20Token.balanceOf(thirdWinner.stakerAddress)
                ).to.equal("1666666666666666");
            } else {
                let winner = await stakingPool.stakers(0);
                expect(winner.stakerAddress).to.equal(signer2Address);

                await stakingPoolAsSigner2.claimReward();
                expect(await totemToken.balanceOf(signer2Address)).is.lessThan(
                    BigNumber.from(10).pow(18).mul(2685)
                );
                expect(await bep20Token.balanceOf(signer2Address)).to.equal(
                    "624999"
                );

                let secondWinner = await stakingPool.stakers(1);
                expect(secondWinner.stakerAddress).to.equal(signer3Address);

                await stakingPoolAsSigner3.claimReward();
                expect(
                    await totemToken.balanceOf(secondWinner.stakerAddress)
                ).is.lessThan(BigNumber.from(10).pow(18).mul(1898));
                expect(
                    await bep20Token.balanceOf(secondWinner.stakerAddress)
                ).to.equal("3333");

                let thirdWinner = await stakingPool.stakers(2);
                expect(thirdWinner.stakerAddress).to.equal(signer4Address);

                await stakingPoolAsSigner4.claimReward();
                expect(
                    await totemToken.balanceOf(thirdWinner.stakerAddress)
                ).is.lessThan(BigNumber.from(10).pow(18).mul(1447));
                expect(
                    await bep20Token.balanceOf(thirdWinner.stakerAddress)
                ).to.equal("1666");
            }
        });
    });

    describe("indexedClaimReward", async() => {
        it("successfully claims reward at index 0", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            let totemTokenAsSigner2 = totemToken.connect(signer2);

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            let totemTokenAsSigner3 = totemToken.connect(signer3);

            let stakingPoolAsSigner4 = stakingPool.connect(signer4);
            let totemTokenAsSigner4 = totemToken.connect(signer4);

            await mockAggregatorContract.mock.decimals.returns(8);
            await mockAggregatorContract.mock.latestRoundData.returns(
                "73786976294838208680",
                BigNumber.from(10).pow(8).mul(60000),
                "1620053717",
                "1620053717",
                "73786976294838208680"
            );

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner2.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60020)
            );

            await advanceBlockWithTime(signer1.provider, 1000);

            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer3Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner3.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(59995)
            );

            await advanceBlockWithTime(signer1.provider, 1000);

            await totemToken.transfer(
                signer4Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner4.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer4Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner4.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60001)
            );

            let stakingPooltotalStakedBalance = await stakingPool.totalStaked();
            expect(stakingPooltotalStakedBalance).to.equal(
                BigNumber.from(10).pow(18).mul(2910)
            );

            await bep20Token.transfer(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            await advanceBlockWithTime(signer1.provider, 430000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            let newAddresses = zeroAddresses;
            newAddresses[0] = signer2Address;
            newAddresses[1] = signer3Address;
            newAddresses[2] = signer4Address;

            await stakingPool.updateMaturingPrice(
                BigNumber.from(10).pow(18).mul(50000),
                true,
                0,
                nilAddress,
                0
            );
            await stakingPool.setSortedStakers(newAddresses, zeroIndexes);
            await stakingPool.endPool();

            let isCollabReward =
                await stakingPoolAsSigner3.collaborativeReward();
            if (isCollabReward) {
                //First winner reward checks
                let winner = await stakingPool.sortedStakers(0);
                expect(winner.stakerAddress).to.equal(signer2Address);

                expect(await stakingPool.maturityTime()).to.equal(1296000);

                expect(await stakingPool.lockTime()).to.equal(432000);

                expect(await stakingPool.sizeAllocation()).to.equal(
                    BigNumber.from(10).pow(18).mul(125000)
                );

                expect(await stakingPool.stakeApr()).to.equal(6000);

                expect(await stakingPool.prizeAmount()).to.equal(
                    BigNumber.from(10).pow(18).mul(4500)
                );

                await stakingPoolAsSigner2.indexedClaimReward(0);
                expect(
                    (
                        await totemToken.balanceOf(signer2Address)
                    ).div(10 ** 15).toNumber()
                ).to.be.lessThan(2685090);

                expect(await bep20Token.balanceOf(signer2Address)).to.equal(
                    "6249999999999999"
                );

                //Second winner rewards check
                let secondWinner = await stakingPool.sortedStakers(1);
                expect(secondWinner.stakerAddress).to.equal(signer3Address);

                await stakingPoolAsSigner3.claimReward();
                expect(
                    (
                        await totemToken.balanceOf(secondWinner.stakerAddress)
                    ).div(10 ** 15).toNumber()
                ).to.be.lessThan(1897575);
                expect(
                    await bep20Token.balanceOf(secondWinner.stakerAddress)
                ).to.equal("3333333333333333");
                //Third winner rewards check
                let thirdWinner = await stakingPool.sortedStakers(2);
                expect(thirdWinner.stakerAddress).to.equal(signer4Address);

                await stakingPoolAsSigner4.claimReward();
                expect(
                    (
                        await totemToken.balanceOf(thirdWinner.stakerAddress)
                    ).div(10 ** 15).toNumber()
                ).to.be.lessThan(1447560);
                expect(
                    await bep20Token.balanceOf(thirdWinner.stakerAddress)
                ).to.equal("1666666666666666");
            } else {
                let winner = await stakingPool.stakers(0);
                expect(winner.stakerAddress).to.equal(signer2Address);

                await stakingPoolAsSigner2.claimReward();
                expect(await totemToken.balanceOf(signer2Address)).is.lessThan(
                    BigNumber.from(10).pow(18).mul(2685)
                );
                expect(await bep20Token.balanceOf(signer2Address)).to.equal(
                    "624999"
                );

                let secondWinner = await stakingPool.stakers(1);
                expect(secondWinner.stakerAddress).to.equal(signer3Address);

                await stakingPoolAsSigner3.claimReward();
                expect(
                    await totemToken.balanceOf(secondWinner.stakerAddress)
                ).is.lessThan(BigNumber.from(10).pow(18).mul(1898));
                expect(
                    await bep20Token.balanceOf(secondWinner.stakerAddress)
                ).to.equal("3333");

                let thirdWinner = await stakingPool.stakers(2);
                expect(thirdWinner.stakerAddress).to.equal(signer4Address);

                await stakingPoolAsSigner4.claimReward();
                expect(
                    await totemToken.balanceOf(thirdWinner.stakerAddress)
                ).is.lessThan(BigNumber.from(10).pow(18).mul(1447));
                expect(
                    await bep20Token.balanceOf(thirdWinner.stakerAddress)
                ).to.equal("1666");
            }
        });
    });

    describe("purchaseWrappedToken", async () => {
        let deadLine: number;
        beforeEach("set a deadline", () => {
            deadLine = Math.round(new Date().getTime() / 1000) + 500;
        });

        it("successfully swap Wrapped Tokens for USD Tokens", async () => {
            await mockUniswapRouter.mock.getAmountsOut
                .withArgs(BigNumber.from(10).pow(18).mul(1000), [
                    usdContract.address,
                    bep20Token.address,
                ])
                .returns([
                    BigNumber.from(10).pow(18).mul(1000),
                    BigNumber.from(10).pow(12).mul(40000),
                ]);

            expect(
                await wrappedTokenDistributor.getEstimatedWrappedTokenForUSD(
                    BigNumber.from(10).pow(18).mul(1000)
                )
            ).to.equal(BigNumber.from(10).pow(12).mul(40000));

            await mockUniswapRouter.mock.swapExactTokensForTokens
                .withArgs(
                    BigNumber.from(10).pow(18).mul(1000),
                    BigNumber.from(10).pow(12).mul(38800),
                    [usdContract.address, bep20Token.address],
                    stakingPool.address,
                    deadLine
                )
                .returns([BigNumber.from(10).pow(12).mul(38800)]);
            let swapRouterAddress =
                await wrappedTokenDistributor.getSwapRouter();
            await stakingPool.purchaseWrappedToken(
                BigNumber.from(10).pow(18).mul(1000),
                deadLine
            );
            expect(
                await usdContract.allowance(
                    stakingPool.address,
                    swapRouterAddress
                )
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await bep20Token.transfer(
                stakingPool.address,
                BigNumber.from(10).pow(12).mul(38800)
            );
            expect(await stakingPool.getWrappedTokenBalance()).to.equal(
                BigNumber.from(10).pow(12).mul(38800)
            );
        });

        it("reverts because pool is not usd rewarder", async () => {
            const stakingPoolProxyInstance = new StakingPoolProxy__factory(
                signer1
            );

            let launchDate = Math.round(new Date().getTime() / 1000) + 5000;
            zeroUSDVariables[0] = launchDate;

            const stakingPoolProxy = await stakingPoolProxyInstance.deploy();

            stakingPoolProxy.initialize(
                "ETH",
                "Wolf",
                totemToken.address,
                rewardManager.address,
                signer1Address,
                [
                    mockUniswapRouter.address,
                    usdContract.address,
                    mockAggregatorContract.address,
                    bep20Token.address,
                ],
                zeroUSDVariables,
                ranks,
                percentages,
                true
            );

            await rewardManager.addPool(stakingPoolProxy.address);
            await totemToken.setTaxExemptStatus(stakingPoolProxy.address, true);

            stakingPoolProxy.upgradeTo(stakingPoolImplementation.address);

            const StakingPoolImplementationFactory =
                await getStakingPoolImplementationFactory();
            let stakingPoolInstance = StakingPoolImplementationFactory.attach(
                stakingPoolProxy.address
            );

            await expect(
                stakingPoolInstance.purchaseWrappedToken(
                    BigNumber.from(10).pow(18).mul(1000),
                    deadLine
                )
            ).to.be.revertedWith("0340 The pool is only TOTM rewarder");
        });

        it("reverts because the usd amount to purchase wrapped token is zero", async () => {
            await expect(
                stakingPool.purchaseWrappedToken(BigNumber.from(0), deadLine)
            ).to.be.revertedWith("0341 Amount can't be zero");
        });

        it("reverts because there is not enough time to purchase wrapped token", async () => {
            await advanceBlockWithTime(
                signer1.provider,
                Math.round(new Date().getTime() / 1000) + 100
            );
            await expect(
                stakingPool.purchaseWrappedToken(
                    BigNumber.from(10).pow(18).mul(1000),
                    deadLine
                )
            ).to.be.revertedWith("0342 Deadline is low");
        });
    });

    describe("updateMaturingPrice", async() => {
        it("oracle updates the maturing price", async () => {
            await mockAggregatorContract.mock.decimals.returns(8);
            await mockAggregatorContract.mock.latestRoundData.returns(
                "73786976294838208680",
                BigNumber.from(10).pow(8).mul(60000),
                "1620053717",
                "1620053717",
                "73786976294838208680"
            );

            await advanceBlockWithTime(signer1.provider, 5050);
            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            // Does not set the price to 50000, because oracle is specified
            await stakingPool.updateMaturingPrice(
                BigNumber.from(10).pow(8).mul(50000),
                true,
                0,
                nilAddress,
                0
            );

            let stakingPoolMaturingPrice = await stakingPool.maturingPrice();
            expect(stakingPoolMaturingPrice).to.equal(
                BigNumber.from(10).pow(8).mul(60000)
            );
        });

        it("pool creator updates the maturing price because no oracle is specified", async () => {
            const stakingPoolProxyInstance = new StakingPoolProxy__factory(
                signer1
            );

            let launchDate = Math.round(new Date().getTime() / 1000) + 5000;
            variables[0] = launchDate;

            const stakingPoolProxy = await stakingPoolProxyInstance.deploy();

            stakingPoolProxy.initialize(
                "ETH",
                "Wolf",
                totemToken.address,
                rewardManager.address,
                signer1Address,
                [
                    mockUniswapRouter.address,
                    usdContract.address,
                    nilAddress,
                    bep20Token.address,
                ],
                variables,
                ranks,
                percentages,
                true
            );

            await rewardManager.addPool(stakingPoolProxy.address);
            await totemToken.setTaxExemptStatus(stakingPoolProxy.address, true);

            stakingPoolProxy.upgradeTo(stakingPoolImplementation.address);

            const StakingPoolImplementationFactory =
                await getStakingPoolImplementationFactory();
            let stakingPoolInstance = StakingPoolImplementationFactory.attach(
                stakingPoolProxy.address
            );

            await advanceBlockWithTime(signer1.provider, 5050);
            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            // Sets price to 59000, because oracle is not specified and updating maturing price is manual
            await stakingPoolInstance.updateMaturingPrice(
                BigNumber.from(10).pow(8).mul(59000),
                true,
                0,
                nilAddress,
                0
            );

            let stakingPoolMaturingPrice =
                await stakingPoolInstance.maturingPrice();
            expect(stakingPoolMaturingPrice).to.equal(
                BigNumber.from(10).pow(8).mul(59000)
            );
        });

        it("reverts because pool is not mature yet", async () => {
            await advanceBlockWithTime(signer1.provider, 500000);

            await expect(
                stakingPool.updateMaturingPrice(
                    BigNumber.from(10).pow(8).mul(50000),
                    true,
                0,
                nilAddress,
                0
                )
            ).to.be.revertedWith(
                "0350 Can't set maturing price before the maturity time"
            );
        });

        it("prediction oracle and prize oracle are different and specified", async () => {
            await mockAggregatorContract.mock.decimals.returns(8);
            await mockAggregatorContract.mock.latestRoundData.returns(
                "73786976294838208680",
                BigNumber.from(10).pow(8).mul(60000),
                "1620053717",
                "1620053717",
                "73786976294838208680"
            );

            await mockAggregatorContractForPrize.mock.decimals.returns(8);
            await mockAggregatorContractForPrize.mock.latestRoundData.returns(
                "73786976294838208680",
                BigNumber.from(10).pow(8).mul(3000),
                "1620053717",
                "1620053717",
                "73786976294838208680"
            );


            await advanceBlockWithTime(signer1.provider, 5050);
            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            await stakingPool.updateMaturingPrice(
                BigNumber.from(10).pow(8).mul(50000),
                false,
                0,
                mockAggregatorContractForPrize.address,
                0
            );

            let stakingPoolMaturingPrice = await stakingPool.maturingPrice();
            expect(stakingPoolMaturingPrice).to.equal(
                BigNumber.from(10).pow(8).mul(60000)
            );

            let stakingPoolPrizeMaturingPrice;
            let stakingPoolPrizeOracleDecimals;
            [stakingPoolPrizeMaturingPrice, stakingPoolPrizeOracleDecimals] = await stakingPool.getLPS();

            expect(stakingPoolPrizeMaturingPrice).to.equal(
                BigNumber.from(10).pow(8).mul(3000)
            );
            expect(stakingPoolPrizeOracleDecimals).to.equal(
                BigNumber.from(8)
            );
        });

        it("prediction oracle and prize oracle are different but prize oracle does not work", async () => {
            await mockAggregatorContract.mock.decimals.returns(8);
            await mockAggregatorContract.mock.latestRoundData.returns(
                "73786976294838208680",
                BigNumber.from(10).pow(8).mul(60000),
                "1620053717",
                "1620053717",
                "73786976294838208680"
            );


            await advanceBlockWithTime(signer1.provider, 5050);
            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            await stakingPool.updateMaturingPrice(
                0,
                false,
                BigNumber.from(10).pow(8).mul(3100),
                nilAddress,
                6
            );

            let stakingPoolMaturingPrice = await stakingPool.maturingPrice();
            expect(stakingPoolMaturingPrice).to.equal(
                BigNumber.from(10).pow(8).mul(60000)
            );

            let stakingPoolPrizeMaturingPrice;
            let stakingPoolPrizeOracleDecimals;
            [stakingPoolPrizeMaturingPrice, stakingPoolPrizeOracleDecimals] = await stakingPool.getLPS();

            expect(stakingPoolPrizeMaturingPrice).to.equal(
                BigNumber.from(10).pow(8).mul(3100)
            );
            expect(stakingPoolPrizeOracleDecimals).to.equal(
                BigNumber.from(6)
            );
        });
    });

    describe("stakingPool: getStakers", async () => {
        it("check the stakers", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let signer2TotemToken = await totemToken.connect(signer2);
            let signer3TotemToken = await totemToken.connect(signer3);
            let signer4TotemToken = await totemToken.connect(signer4);

            let signer2stakingPool = await stakingPool.connect(signer2);
            let signer3stakingPool = await stakingPool.connect(signer3);
            let signer4stakingPool = await stakingPool.connect(signer4);

            await totemToken.transfer(signer2Address, tenThousand);
            await totemToken.transfer(signer3Address, oneThousand);
            await totemToken.transfer(signer4Address, oneThousand);

            await signer2TotemToken.approve(stakingPool.address, oneThousand);
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(oneThousand);
            await signer2stakingPool.stake(oneThousand, fiftyThousandBTCPrice);

            await signer2TotemToken.approve(stakingPool.address, oneThousand);
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(oneThousand);
            await signer2stakingPool.stake(oneThousand, sixtyThousandBTCPrice);

            await signer3TotemToken.approve(stakingPool.address, oneThousand);
            expect(
                await totemToken.allowance(signer3Address, stakingPool.address)
            ).to.equal(oneThousand);
            await signer3stakingPool.stake(oneThousand, fiftyThousandBTCPrice);

            await signer4TotemToken.approve(stakingPool.address, oneThousand);
            expect(
                await totemToken.allowance(signer4Address, stakingPool.address)
            ).to.equal(oneThousand);
            await signer4stakingPool.stake(oneThousand, fiftyThousandBTCPrice);

            const addrsAndIndexes = await stakingPool.getStakers();

            expect(addrsAndIndexes[0][2]).to.equal(signer3Address);
        });
    });

    describe("stakingPool: setSortedStakers", async () => {
        it("setSortedStakers can't be called by non-operator accounts", async () => {
            let signer2stakingPool = stakingPool.connect(signer2);

            await expect(
                signer2stakingPool.setSortedStakers(zeroAddresses, zeroIndexes)
            ).to.be.revertedWith("0300 caller is not a pool creator");
        });

        it("reverts because setSortedStakers can not be called before maturity date", async() => {
            let newAddresses = zeroAddresses;

            await expect(
                stakingPool.setSortedStakers(newAddresses, zeroIndexes)
            ).to.be.revertedWith("0390 Can't set sorted stakers before the maturity time");
        });

        it("calls setSortedStakers successfully", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            let totemTokenAsSigner2 = totemToken.connect(signer2);

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            let totemTokenAsSigner3 = totemToken.connect(signer3);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await stakingPoolAsSigner2.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60001)
            );

            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await stakingPoolAsSigner3.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(59980)
            );

            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            let newAddresses = zeroAddresses;
            newAddresses[0] = signer2Address;
            newAddresses[1] = signer3Address;
            newAddresses[2] = nilAddress;

            expect(
                await stakingPool.setSortedStakers(newAddresses, zeroIndexes)
            ).to.emit(stakingPool, "PoolSorted");

            expect((await stakingPool.sortedStakers(0)).stakerAddress).to.equal(
                signer2Address
            );
        });

        it("previous sortedStakers is removed when setSortedStakers is called again", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            let totemTokenAsSigner2 = totemToken.connect(signer2);

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            let totemTokenAsSigner3 = totemToken.connect(signer3);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await stakingPoolAsSigner2.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60001)
            );

            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await stakingPoolAsSigner3.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(59980)
            );

            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            let newAddresses = zeroAddresses;
            newAddresses[0] = signer2Address;
            newAddresses[1] = signer3Address;

            expect(
                await stakingPool.setSortedStakers(newAddresses, zeroIndexes)
            ).to.emit(stakingPool, "PoolSorted");

            expect((await stakingPool.sortedStakers(0)).stakerAddress).to.equal(
                signer2Address
            );

            newAddresses[0] = signer3Address;
            newAddresses[1] = signer2Address;

            expect(
                await stakingPool.setSortedStakers(newAddresses, zeroIndexes)
            ).to.emit(stakingPool, "PoolSorted");

            expect((await stakingPool.sortedStakers(0)).stakerAddress).to.equal(
                signer3Address
            );
            expect((await stakingPool.sortedStakers(1)).stakerAddress).to.equal(
                signer2Address
            );
        });

        it("reverts because sorted stakers can not be more than the last rank", async() => {
            ranks = [1, 2, 0, 0, 0, 0, 0, 0];
            percentages = [8000, 2000, 0, 0, 0, 0, 0, 0];

            stakingPoolProxy = await deployStakingPoolProxy();

            stakingPool = stakingPoolImplementation.attach(
                stakingPoolProxy.address
            );

            await rewardManager.addPool(stakingPoolProxy.address);
            await totemToken.setTaxExemptStatus(stakingPoolProxy.address, true);

            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            let totemTokenAsSigner2 = totemToken.connect(signer2);

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            let totemTokenAsSigner3 = totemToken.connect(signer3);

            let stakingPoolAsSigner4 = stakingPool.connect(signer4);
            let totemTokenAsSigner4 = totemToken.connect(signer4);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await stakingPoolAsSigner2.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60001)
            );

            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await stakingPoolAsSigner3.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(59980)
            );

            await totemToken.transfer(
                signer4Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner4.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await stakingPoolAsSigner4.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60008)
            );

            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            let newAddresses = zeroAddresses;
            newAddresses[0] = signer2Address;
            newAddresses[1] = signer3Address;
            newAddresses[2] = signer4Address;

            await expect(
                stakingPool.setSortedStakers(newAddresses, zeroIndexes)
            ).to.be.revertedWith("number of sorted stakers must be less than or equal to the last rank")

            ranks = [1, 2, 3, 10, 25, 0, 0, 0];
            percentages = [3750, 2000, 1000, 250, 100, 0, 0, 0];
        })
    });

    describe("endPool", async() => {
        it("successfully ends the pool", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            let totemTokenAsSigner2 = totemToken.connect(signer2);

            let stakingPoolAsSigner3 = stakingPool.connect(signer3);
            let totemTokenAsSigner3 = totemToken.connect(signer3);

            let stakingPoolAsSigner4 = stakingPool.connect(signer4);
            let totemTokenAsSigner4 = totemToken.connect(signer4);

            await mockAggregatorContract.mock.decimals.returns(8);
            await mockAggregatorContract.mock.latestRoundData.returns(
                "73786976294838208680",
                BigNumber.from(10).pow(8).mul(60000),
                "1620053717",
                "1620053717",
                "73786976294838208680"
            );

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner2.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60001)
            );

            await totemToken.transfer(
                signer3Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner3.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer3Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner3.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(59980)
            );

            await totemToken.transfer(
                signer4Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner4.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer4Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner4.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60021)
            );

            expect(
                await stakingPool.averagePricePrediction()
            ).to.be.equal(6000066666666)

            await bep20Token.transfer(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            let addresses = zeroAddresses;
            addresses[0] = signer2Address;
            addresses[1] = signer3Address;
            addresses[2] = signer4Address;

            let sortedStakersIndexed = zeroIndexes;

            await expect(
                stakingPool.setSortedStakers(addresses, sortedStakersIndexed)
            ).to.emit(stakingPool, "PoolSorted");

            await expect(stakingPool.endPool()).to.emit(
                stakingPool,
                "PoolMatured"
            );

            let winner = await stakingPool.stakers(0);
            expect(winner.stakerAddress).to.equal(signer2Address);
        });

        it("reverts because can't end pool before maturity time", async () => {
            await advanceBlockWithTime(signer1.provider, 5001);
            await stakingPool.setActivationStatus(true);

            await advanceBlockWithTime(signer1.provider, 500000);
            await expect(stakingPool.endPool()).to.be.revertedWith(
                "0360 Can't end pool before the maturity time"
            );
        });

        it("reverts because wrapped token is not available", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);
            await expect(stakingPool.endPool()).to.be.revertedWith(
                "0361 WrappedToken Rewards not available"
            );
        });

        it("reverts because stakers are not sorted yet", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let stakingPoolAsSigner2 = stakingPool.connect(signer2);
            let totemTokenAsSigner2 = totemToken.connect(signer2);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await totemTokenAsSigner2.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));
            await stakingPoolAsSigner2.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(60001)
            );

            await bep20Token.transfer(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            await advanceBlockWithTime(signer1.provider, 432000);
            await advanceBlockWithTime(signer1.provider, 1296000);

            await expect(stakingPool.endPool()).to.be.revertedWith(
                "0362 first should sort"
            );
        });
    });

    describe("declareEmergency", async () => {
        it("reverts because the caller is not pool creator", async () => {
            let signer2stakingPool = stakingPool.connect(signer2);

            await expect(
                signer2stakingPool.declareEmergency()
            ).to.be.revertedWith("0300 caller is not a pool creator");
        });

        it("successfully declaring an emergency", async () => {
            await stakingPool.declareEmergency();

            expect(await stakingPool.isAnEmergency()).to.be.true;
        });
    });

    describe("emergentWithdraw", async () => {
        it("reverts because stakers can't call emergentWithdraw when it's not an emergency", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let signer2TotemToken = totemToken.connect(signer2);
            let signer2stakingPool = stakingPool.connect(signer2);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            await signer2TotemToken.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            expect(
                await totemToken.allowance(signer2Address, stakingPool.address)
            ).to.equal(BigNumber.from(10).pow(18).mul(1000));

            await signer2stakingPool.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(50000)
            );

            await expect(
                signer2stakingPool.emergentWithdraw()
            ).to.be.revertedWith("it's not an emergency");
        });

        it("stakers can call emergentWithdraw in case of emergency", async () => {
            await advanceBlockWithTime(signer1.provider, 5050);
            await stakingPool.setActivationStatus(true);

            let signer2TotemToken = totemToken.connect(signer2);
            let signer2stakingPool = stakingPool.connect(signer2);

            await totemToken.transfer(
                signer2Address,
                BigNumber.from(10).pow(18).mul(1000)
            );

            await signer2TotemToken.approve(
                stakingPool.address,
                BigNumber.from(10).pow(18).mul(1000)
            );
            await signer2stakingPool.stake(
                BigNumber.from(10).pow(18).mul(1000),
                BigNumber.from(10).pow(8).mul(50000)
            );

            expect(await totemToken.balanceOf(signer2Address)).to.equal(0);

            await stakingPool.declareEmergency();

            await signer2stakingPool.emergentWithdraw();

            expect(await totemToken.balanceOf(signer2Address)).to.equal(
                BigNumber.from(10).pow(18).mul(970)
            );
        });
    });

    describe("withdrawStuckTokens", async () => {
        it("withdrawStuckTokens can't be called by non-operator accounts", async () => {
            let signer2stakingPool = stakingPool.connect(signer2);
    
            await expect(
                signer2stakingPool.withdrawStuckTokens(
                    bep20Token.address,
                    oneHundred,
                    signer2Address
                )
            ).to.be.revertedWith("0300 caller is not a pool creator");
        });
    
        it("Totem can't be withdrawn from stakingPool", async () => {
            await expect(
                stakingPool.withdrawStuckTokens(
                    totemToken.address,
                    oneHundred,
                    signer2Address
                )
            ).to.be.revertedWith("0370 totems can not be transfered");
        });
    
        it("stuck tokens can be withdrawn from the stakingPool", async () => {
            await bep20Token.transfer(stakingPool.address, oneHundred);
    
            expect(await bep20Token.balanceOf(signer4Address)).to.equal(0);
    
            await stakingPool.withdrawStuckTokens(
                bep20Token.address,
                oneHundred,
                signer4Address
            );
    
            expect(await bep20Token.balanceOf(signer4Address)).to.equal(oneHundred);
        });
    });
    
});