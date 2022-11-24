const { ethers } = require("hardhat");
const { exportContract } = require("../utils/export");

//npx hardhat run .\scripts\export_existing.js
async function main() {

    // export existing contract
    await exportContract("ERC20", "0xc3761EB917CD790B30dAD99f6Cc5b4Ff93C4F9eA");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });