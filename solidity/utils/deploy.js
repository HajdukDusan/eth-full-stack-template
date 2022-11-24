const { run, ethers } = require("hardhat")
const { exportContract } = require("./export")

// deploy contract, and export for frontend and backend
// returns deployed contract address
// args: (ContractName, ContractArgs..)
async function deployAndExport(contractName, ...args) {

    console.log("----------------------------------------------------")
    console.log("Deploying contract " + contractName + "...")

    const contractAddress = await deployContract(contractName, ...args)

    console.log("\nExporting...");

    exportContract(contractName, contractAddress)

    console.log("Done.");

    return contractAddress
}

// args: (ContractName, ContractArgs..)
async function deployContract(contractName, ...args) {

    const contractFactory = await ethers.getContractFactory(contractName);

    const contract = await contractFactory.deploy(...args);

    console.log("\ntx:")
    console.log("\t hash: " + contract.deployTransaction.hash)
    console.log("\t nonce: " + contract.deployTransaction.nonce)

    console.log("\t from: " + contract.deployTransaction.from)
    console.log("\t to: " + contract.deployTransaction.to)

    console.log("\t gasPrice: " + ethers.utils.formatEther(contract.deployTransaction.gasPrice))
    console.log("\t gasLimit: " + ethers.utils.formatEther(contract.deployTransaction.gasLimit))
    console.log("\t value: " + ethers.utils.formatEther(contract.deployTransaction.value) + "\n")

    const receipt = await contract.deployTransaction.wait()

    console.log("\nreceipt:")
    console.log(`\tdeployed at ${contract.address}`)
    console.log('\ttx cost:', ethers.utils.formatEther(receipt.gasUsed.mul(receipt.effectiveGasPrice)))

    return contract.address
}

module.exports = { deployContract, deployAndExport }