const { verify } = require("../utils/verify")
const { ethers } = require("hardhat");
const { deployAndExport } = require("../utils/deploy");

//npx hardhat node
//npx hardhat run --network localhost .\scripts\deploy_export_localhost.js
async function main() {

    // deploy and export contract
    await deployAndExport("StupidContract", 1000, "This is the stupid contract description");

}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });