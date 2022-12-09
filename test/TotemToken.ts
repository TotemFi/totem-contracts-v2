import { expect, use } from "chai"
import { ethers } from "hardhat"
import { Signer } from "ethers"
import { solidity, MockProvider } from "ethereum-waffle"

import { TotemToken } from "../src/types/TotemToken"
import { TotemToken__factory } from "../src/types/factories/TotemToken__factory"

use(solidity)

describe("TotemToken", async () => {
    let signer: Signer

    let totemToken: TotemToken

    beforeEach(async () => {
        ;[signer] = await ethers.getSigners()
        totemToken = await deployTotemToken()
    })

    const deployTotemToken = async (_signer?: Signer): Promise<TotemToken> => {
        const totemTokenFactory = new TotemToken__factory(_signer || signer)
        const totemToken = await totemTokenFactory.deploy()
        return totemToken
    }

    it("initialize contract correctly", async () => {
        expect(await totemToken.name()).to.equal("Totem Token")
        expect(await totemToken.symbol()).to.equal("TOTM")
        expect(await totemToken.decimals()).to.equal(18)

        expect(await totemToken.balanceOf(totemToken.address)).to.equal(
            await totemToken.INITIAL_SUPPLY()
        )
    })

    it("distribute tokens correctly", async () => {
        const [
            communityWallet,
            stakingWallet,
            advisorsWallet,
            seedWallet,
            teamWallet,
            publicWallet,
            strategicWallet,
            liquidityWallet,
            privateWallet,
        ] = new MockProvider().getWallets()

        await totemToken.setDistributionTeamsAddresses(
            communityWallet.address,
            stakingWallet.address,
            liquidityWallet.address,
            publicWallet.address,
            advisorsWallet.address,
            seedWallet.address,
            privateWallet.address,
            teamWallet.address,
            strategicWallet.address
        )
        await totemToken.distributeTokens()

        expect(await totemToken.balanceOf(communityWallet.address)).to.equal(
            await totemToken.COMMUNITY_DEVELOPMENT()
        )
        expect(await totemToken.balanceOf(stakingWallet.address)).to.equal(
            await totemToken.STAKING_REWARDS()
        )
        expect(await totemToken.balanceOf(liquidityWallet.address)).to.equal(
            await totemToken.LIQUIDITY_POOL()
        )
        expect(await totemToken.balanceOf(publicWallet.address)).to.equal(
            await totemToken.PUBLIC_SALE()
        )
        expect(await totemToken.balanceOf(advisorsWallet.address)).to.equal(
            await totemToken.ADVISORS()
        )
        expect(await totemToken.balanceOf(seedWallet.address)).to.equal(
            await totemToken.SEED_INVESTMENT()
        )
        expect(await totemToken.balanceOf(privateWallet.address)).to.equal(
            await totemToken.PRIVATE_SALE()
        )
        expect(await totemToken.balanceOf(teamWallet.address)).to.equal(
            await totemToken.TEAM_ALLOCATION()
        )
        expect(await totemToken.balanceOf(strategicWallet.address)).to.equal(
            await totemToken.STRATEGIC_ROUND()
        )

        await expect(
            totemToken.distributeTokens(),
            "should distribute only once"
        ).to.be.revertedWith("")
    })

    it("taxation applies correctly", async () => {
        await expect(totemToken.setTaxRate(10000)).to.be.revertedWith("")

        const [
            walletA,
            taxationWallet,
            otherWallet,
        ] = new MockProvider().getWallets()

        const signerAddress = await signer.getAddress()

        // Get some tokens to signer
        await totemToken.setDistributionTeamsAddresses(
            signerAddress,
            otherWallet.address,
            otherWallet.address,
            otherWallet.address,
            otherWallet.address,
            otherWallet.address,
            otherWallet.address,
            otherWallet.address,
            otherWallet.address
        )
        await totemToken.distributeTokens()

        const initialBalance = await totemToken.balanceOf(signerAddress)

        // Signer(owner) should be set as tax exempt by default
        await totemToken.transfer(walletA.address, 1000)
        expect(await totemToken.balanceOf(walletA.address)).to.equal(1000)
        expect(await totemToken.balanceOf(signerAddress)).to.equal(
            initialBalance.sub(1000)
        )

        // Remove signer from tax exempt, tax will be sent to the signer(owner) by default
        await totemToken.setTaxExemptStatus(signerAddress, false)
        await totemToken.transfer(walletA.address, 1000)
        expect(await totemToken.balanceOf(walletA.address)).to.equal(1970)
        expect(await totemToken.balanceOf(signerAddress)).to.equal(
            initialBalance.sub(2000).add(30)
        )

        // Set custmo taxation wallet
        await totemToken.setTaxationWallet(taxationWallet.address)
        await totemToken.transfer(walletA.address, 1000)
        expect(await totemToken.balanceOf(walletA.address)).to.equal(2940)
        expect(await totemToken.balanceOf(taxationWallet.address)).to.equal(30)

        // Set custom tax rate
        await totemToken.setTaxRate(500)
        await totemToken.transfer(walletA.address, 1000)
        expect(await totemToken.balanceOf(walletA.address)).to.equal(3890)
        expect(await totemToken.balanceOf(taxationWallet.address)).to.equal(80)

        // Again turn on tax exempt for signer
        await totemToken.setTaxExemptStatus(signerAddress, true)
        await totemToken.transfer(walletA.address, 1000)
        expect(await totemToken.balanceOf(walletA.address)).to.equal(4890)
        expect(await totemToken.balanceOf(taxationWallet.address)).to.equal(80)

        // Taxation for transferFrom
        await totemToken.approve(signerAddress, 10000)

        await totemToken.transferFrom(signerAddress, walletA.address, 1000)
        expect(await totemToken.balanceOf(walletA.address)).to.equal(5890)
        expect(await totemToken.balanceOf(taxationWallet.address)).to.equal(80)

        await totemToken.setTaxExemptStatus(signerAddress, false)
        await totemToken.transferFrom(signerAddress, walletA.address, 1000)
        expect(await totemToken.balanceOf(walletA.address)).to.equal(6840)
        expect(await totemToken.balanceOf(taxationWallet.address)).to.equal(130)
    })
})
