import { HardhatRuntimeEnvironment } from "hardhat/types"
import { DeployFunction } from "hardhat-deploy/types"
import config from 'config'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
    const { deployments, getNamedAccounts } = hre
    const { deploy } = deployments
    const { deployer } = await getNamedAccounts()
    const TotemToken = await deployments.get("TotemToken")
    const RewardManager = await deployments.get("RewardManager")

    // ! These addresses are just for testing purpose on bsc testnet.
    // ! Need to change when deploying to mainnet
    const swapRouter = config.get("swapRouterAddress")
    const busdAddress = config.get("contracts.usdAddresses.busd")
    const superAdmin = config.get("superAdmin")

    const StakingPoolImplementation = await deployments.get("StakingPoolImplementation");


    await deploy("StakingPoolFactory", {
        from: deployer,
        log: true,
        skipIfAlreadyDeployed: true,
        args: [
            TotemToken.address,
            RewardManager.address,
            swapRouter,
            busdAddress,
            StakingPoolImplementation.address,
            superAdmin
        ],
    })
}

export default func
export const tags = ["StakingPoolFactory"]
