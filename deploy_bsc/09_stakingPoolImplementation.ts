import { HardhatRuntimeEnvironment } from "hardhat/types"
import { DeployFunction } from "hardhat-deploy/types"

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
    const { deployments, getNamedAccounts } = hre
    const { deploy } = deployments
    const { deployer } = await getNamedAccounts()

    const CalculateRewardLib = await deploy("CalculateRewardLib", {
        from: deployer,
        log: true,
        skipIfAlreadyDeployed: true,
    })

    const ClaimRewardLib = await deploy("ClaimRewardLib", {
        from: deployer,
        log: true,
        skipIfAlreadyDeployed: true,
    })

    const IndexedClaimRewardLib = await deploy("IndexedClaimRewardLib", {
        from: deployer,
        log: true,
        skipIfAlreadyDeployed: true,
    })

    await deploy("StakingPoolImplementation", {
        from: deployer,
        log: true,
        skipIfAlreadyDeployed: true,
        libraries: {
            "CalculateRewardLib": CalculateRewardLib.address,
            "ClaimRewardLib": ClaimRewardLib.address,
            "IndexedClaimRewardLib": IndexedClaimRewardLib.address
        }
    })
}

export default func
export const tags = ["StakingPoolImplementation"]
