import { expect, use } from "chai"
import { deployments, ethers } from "hardhat"
import { Signer, BigNumber } from "ethers"
import {
    deployMockContract,
    MockContract,
} from "@ethereum-waffle/mock-contract"
import { solidity, MockProvider } from "ethereum-waffle"

import { TotemVesting } from "../src/types/TotemVesting"
import { TotemVesting__factory } from "../src/types/factories/TotemVesting__factory"

const {
    advanceBlockWithTime,
    takeSnapshot,
    revertProvider,
} = require("./utils")

use(solidity)

describe("TotemVesting", async () => {
    let signer: Signer

    let totemVesting: TotemVesting
    let mockTotem: MockContract // Mock TotemToken

    const MONTH_INTERVAL = 60 * 60 * 24 * 30

    beforeEach(async () => {
        ;[signer] = await ethers.getSigners()

        const TotemToken = await deployments.getArtifact("TotemToken")
        mockTotem = await deployMockContract(signer, TotemToken.abi)

        totemVesting = await deployTotemVesting()
    })

    const deployTotemVesting = async (
        _signer?: Signer
    ): Promise<TotemVesting> => {
        const totemVestingFactory = new TotemVesting__factory(_signer || signer)
        const totemVesting = await totemVestingFactory.deploy(
            mockTotem.address,
            1500,
            MONTH_INTERVAL,
            10,
            3
        )

        return totemVesting
    }

    it("initialize contract correctly", async () => {
        const totemVestingFactory = new TotemVesting__factory(signer)

        await expect(
            totemVestingFactory.deploy(
                mockTotem.address,
                0,
                MONTH_INTERVAL,
                10,
                3
            ),
            "totalAmount validation"
        ).to.be.revertedWith("")

        await expect(
            totemVestingFactory.deploy(mockTotem.address, 1000, 0, 10, 3),
            "withdrawInterval validation"
        ).to.be.revertedWith("")

        await expect(
            totemVestingFactory.deploy(
                mockTotem.address,
                1000,
                MONTH_INTERVAL,
                0,
                3
            ),
            "releasePeriods validation"
        ).to.be.revertedWith("")
    })

    it("add recipient validation", async () => {
        const snapshotId = await takeSnapshot(signer.provider)

        const [walletA, walletB, walletC] = new MockProvider().getWallets()

        await expect(
            totemVesting.addRecipient(walletA.address, 0),
            "addRecipient validation: no zero amount"
        ).to.be.revertedWith("")
        await expect(
            totemVesting.addRecipient(walletA.address, 1501),
            "addRecipient validation: can't higher than total amount"
        ).to.be.revertedWith("")

        await totemVesting.addRecipient(walletA.address, 1000)
        await totemVesting.addRecipient(walletB.address, 500)
        await expect(
            totemVesting.addRecipient(walletC.address, 1),
            "addRecipient validation: no more tokens left"
        ).to.be.revertedWith("")

        await totemVesting.setStartTime(
            BigNumber.from(new Date().getTime()).div(1000).add(200)
        )

        await advanceBlockWithTime(signer.provider, 200)

        await expect(
            totemVesting.addRecipient(walletA.address, 1000),
            "addRecipient validation: can't set after vesting starts"
        ).to.be.revertedWith("")

        await revertProvider(signer.provider, snapshotId)
    })

    it("set startTime validation", async () => {
        const snapshotId = await takeSnapshot(signer.provider)

        await expect(
            totemVesting.setStartTime(
                BigNumber.from(new Date().getTime()).div(1000).sub(100)
            ),
            "set startTime validation: can't set previous time"
        ).to.be.revertedWith("")

        await totemVesting.setStartTime(
            BigNumber.from(new Date().getTime()).div(1000).add(200)
        )
        await advanceBlockWithTime(signer.provider, 200)

        await expect(
            totemVesting.setStartTime(
                BigNumber.from(new Date().getTime()).div(1000).add(100)
            ),
            "set startTime validation: can't set after vesting starts"
        ).to.be.revertedWith("")

        await revertProvider(signer.provider, snapshotId)
    })

    it("calculate vested/withdrawable correctly", async () => {
        const signerAddress = await signer.getAddress()
        const snapshotId = await takeSnapshot(signer.provider)

        await totemVesting.addRecipient(signerAddress, 1000)

        expect(await totemVesting.vested(signerAddress)).to.equal(0)

        await totemVesting.setStartTime(
            BigNumber.from(new Date().getTime()).div(1000).add(200)
        )

        expect(await totemVesting.vested(signerAddress)).to.equal(0)

        await advanceBlockWithTime(signer.provider, 200)
        expect(await totemVesting.vested(signerAddress)).to.equal(0)

        await advanceBlockWithTime(signer.provider, MONTH_INTERVAL * 3)
        expect(await totemVesting.vested(signerAddress)).to.equal(100)

        await advanceBlockWithTime(signer.provider, MONTH_INTERVAL)
        expect(await totemVesting.vested(signerAddress)).to.equal(200)

        await mockTotem.mock.transfer.returns(true)

        await totemVesting.withdraw()
        expect(await totemVesting.vested(signerAddress)).to.equal(200)
        expect(await totemVesting.withdrawable(signerAddress)).to.equal(0)

        await revertProvider(signer.provider, snapshotId)
    })
})
