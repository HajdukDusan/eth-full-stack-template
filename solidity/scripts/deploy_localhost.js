const { verify } = require("../utils/verify")
const { ethers } = require("hardhat");
const { deployAndExport } = require("../utils/deploy");

//npx hardhat node
//npx hardhat run --network localhost .\scripts\deploy_localhost.js
async function main() {

    // deploy and export contract
    await deployAndExport("PaymentContract", "0x14dC79964da2C08b23698B3D3cc7Ca32193d9955")
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });