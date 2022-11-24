const { verify } = require("../utils/verify")
const { ethers } = require("hardhat");
const { deployAndExport } = require("../utils/deploy");

//npx hardhat run --network goerli .\scripts\deploy_testnet.js
async function main() {

    // deploy and export contract
    const address = await deployAndExport("PaymentContract", "0x14dC79964da2C08b23698B3D3cc7Ca32193d9955")

    // verify contract on etherscan
    await verify(address, ["0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"])
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });