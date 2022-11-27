const path = require("path");
const fs = require("fs");
const cmd = require('node-cmd');

const exportContract = (contractName, address) => {

    exportForFront(contractName, address)

    exportForBackend(contractName, address)
};

const exportForFront = (contractName, address) => {

    runShellCmd(
        "cd .. &" +
        "cd ./frontend &" +
        "mkdir contracts"
    )

    try {
        const dir = path.resolve(
            __dirname,
            "../artifacts/contracts/" +
            contractName +
            ".sol/" +
            contractName +
            ".json"
        );
        const file = fs.readFileSync(dir, "utf8");
        const json = JSON.parse(file);
        const abi = json.abi;

        const dirWrite = path.resolve(
            __dirname,
            "../../frontend/src/contracts/" + contractName + ".js"
        );

        fs.writeFileSync(dirWrite, "const ABI= " + JSON.stringify(abi, null, 2) + "\nconst Address = \"" + address + "\"\nexport default {Address, ABI}", (err) => {
            if (err) {
                console.log(err);
            }
        });

        return abi;
    } catch (e) {
        console.log(e);
    }
}

const exportForBackend = (contractName, contractAddress) => {

    runShellCmd(
        "cd .. &" +
        "cd ./backend &" +
        "mkdir contracts &" +
        "cd ./contracts &" +
        "mkdir " + contractName + "&" +
        "cd ./" + contractName + "&" +
        "mkdir abi &" +
        "mkdir bin &" +
        "cd ../../.. &" +
        "solcjs --optimize --abi ./solidity/contracts/" + contractName + ".sol -o ./backend/contracts/" + contractName + "/abi &" +
        "solcjs --optimize --bin ./solidity/contracts/" + contractName + ".sol -o ./backend/contracts/" + contractName + "/bin &" +
        "abigen " +
        "--abi=./backend/contracts/" + contractName + "/abi/solidity_contracts_" + contractName + "_sol_" + contractName + ".abi " +
        "--bin=./backend/contracts/" + contractName + "/bin/solidity_contracts_" + contractName + "_sol_" + contractName + ".bin " +
        "--pkg=" + contractName + " --out=./backend/contracts/" + contractName + "/api.go"
    )

    try {
        const dirWrite = path.resolve(
            __dirname,
            "../../backend/contracts/" + contractName + "/address.go"
        );

        fs.writeFileSync(dirWrite, "package " + contractName + "\n\nvar Address = \"" + contractAddress + "\"", (err) => {
            if (err) {
                console.log(err);
            }
        });
    } catch (e) {
        console.log(e);
    }
}

const runShellCmd = (command) => {
    const executedCmd = cmd.runSync(command);

    if (executedCmd.err) {
        console.log(executedCmd.err)
    }

    if (executedCmd.stderr) {
        console.log(executedCmd.stderr)
    }

    if (executedCmd.data) {
        console.log(executedCmd.data)
    }
}



module.exports = { exportContract };