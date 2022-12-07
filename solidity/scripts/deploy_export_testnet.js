const { verify } = require("../utils/verify")
const { ethers } = require("hardhat");
const { deployAndExport } = require("../utils/deploy");

//npx hardhat run --network goerli .\scripts\deploy_export_testnet.js
async function main() {

    // deploy and export contract
    const address = await deployAndExport("StupidContract", 1000, "This is the stupid contract description");

    // verify contract on etherscan
    //await verify(address, [1000, "This is the stupid contract description"])
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });