const { verify } = require("../utils/verify")
const { ethers } = require("hardhat");
const { deployContract } = require("../utils/deploy");

//npx hardhat node
//npx hardhat run --network localhost .\scripts\deploy_localhost.js
async function main() {

    // deploy and export contract
    await deployContract("StupidContract", 1000, "This is the stupid contract description");

}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });