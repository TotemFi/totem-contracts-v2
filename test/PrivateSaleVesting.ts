import { expect, use } from "chai"
import { deployments, ethers } from "hardhat"
import { Signer, BigNumber } from "ethers"
import {
    deployMockContract,
    MockContract,
} from "@ethereum-waffle/mock-contract"
import { solidity } from "ethereum-waffle"

import { PrivateSaleVesting } from "../src/types/PrivateSaleVesting"
import { PrivateSaleVesting__factory } from "../src/types/factories/PrivateSaleVesting__factory"

const {
    advanceBlockWithTime,
    takeSnapshot,
    revertProvider,
} = require("./utils")

use(solidity)

describe("PrivateSaleVesting", async () => {
    let signer: Signer

    let privateSaleVesting: PrivateSaleVesting
    let mockTotem: MockContract // Mock TotemToken

    const MONTH_INTERVAL = 60 * 60 * 24 * 30

    beforeEach(async () => {
        ;[signer] = await ethers.getSigners()

        const TotemToken = await deployments.getArtifact("TotemToken")
        mockTotem = await deployMockContract(signer, TotemToken.abi)

        privateSaleVesting = await deployPrivateSaleVesting()
    })

    const deployPrivateSaleVesting = async (
        _signer?: Signer
    ): Promise<PrivateSaleVesting> => {
        const PrivateSaleVestingFactory = new PrivateSaleVesting__factory(
            _signer || signer
        )
        const PrivateSaleVesting = await PrivateSaleVestingFactory.deploy(
            mockTotem.address
        )

        return PrivateSaleVesting
    }

    it("calculate vested/withdrawable correctly", async () => {
        const signerAddress = await signer.getAddress()
        const snapshotId = await takeSnapshot(signer.provider)

        await privateSaleVesting.addRecipient(signerAddress, 1000)

        expect(await privateSaleVesting.vested(signerAddress)).to.equal(0)

        privateSaleVesting.setStartTime(
            BigNumber.from(new Date().getTime()).div(1000).add(100)
        )

        expect(await privateSaleVesting.vested(signerAddress)).to.equal(0)

        await advanceBlockWithTime(signer.provider, 150)
        expect(await privateSaleVesting.vested(signerAddress)).to.equal(225)

        await advanceBlockWithTime(signer.provider, MONTH_INTERVAL)
        expect(await privateSaleVesting.vested(signerAddress)).to.equal(418)

        await advanceBlockWithTime(signer.provider, MONTH_INTERVAL)
        expect(await privateSaleVesting.vested(signerAddress)).to.equal(611)

        await mockTotem.mock.transfer.returns(true)

        await privateSaleVesting.withdraw()
        expect(await privateSaleVesting.vested(signerAddress)).to.equal(611)
        expect(await privateSaleVesting.withdrawable(signerAddress)).to.equal(0)

        await revertProvider(signer.provider, snapshotId)
    })
})
