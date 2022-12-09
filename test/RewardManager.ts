import { expect, use } from "chai"
import { deployments, ethers } from "hardhat"
import { Signer } from "ethers"
import {
    deployMockContract,
    MockContract,
} from "@ethereum-waffle/mock-contract"
import { solidity, MockProvider } from "ethereum-waffle"

import { RewardManager } from "../src/types/RewardManager"
import { RewardManager__factory } from "../src/types/factories/RewardManager__factory"

use(solidity)

describe("RewardManager", async () => {
    let signer: Signer

    let rewardManager: RewardManager
    let mockTotem: MockContract

    const deployRewardManager = async (
        _signer?: Signer
    ): Promise<RewardManager> => {
        const rewardFactory = new RewardManager__factory(_signer || signer)
        const rewardManager = await rewardFactory.deploy(mockTotem.address)
        return rewardManager
    }

    beforeEach(async () => {
        ;[signer] = await ethers.getSigners()

        const TotemToken = await deployments.getArtifact("TotemToken")
        mockTotem = await deployMockContract(signer, TotemToken.abi)
        rewardManager = await deployRewardManager()
    })

    it("addPool function can only be called by operator", async () => {
        const [wallet] = new MockProvider().getWallets()

        await mockTotem.mock.balanceOf.returns(50000)
        await mockTotem.mock.transfer.returns(true)

        // Owner is the operator by default
        await rewardManager.addPool(wallet.address)

        await rewardManager.renounceOperator()
        await expect(
            rewardManager.addPool(wallet.address),
            "function can only be called by operator"
        ).to.be.revertedWith(
            "OperatorRole: caller does not have the Operator role"
        )
    })

    it("rewardUser function can only be called by rewarder", async () => {
        const [wallet] = new MockProvider().getWallets()

        await mockTotem.mock.balanceOf.returns(50000)
        await mockTotem.mock.transfer.returns(true)

        // Owner is the rewarder by default
        await rewardManager.rewardUser(wallet.address, 100)

        await rewardManager.renounceRewarder()
        await expect(
            rewardManager.rewardUser(wallet.address, 100),
            "function can only be called by rewarder"
        ).to.be.revertedWith(
            "RewarderRole: caller does not have the Rewarder role"
        )
    })
})
