<div align="center">

  <img src="https://github.com/get-icon/geticon/raw/master/icons/ethereum.svg" alt="logo" width="120px" height="120px" height="auto" />
  <br/>
  <h1>Ethereum Full Stack Project Template</h1>

  <h3>
  Made for fast and simple development of full stack Ethereum based apps.
  </h3>

<br />

<a href="https://ethereum.org/" title="Ethereum"><img src="https://github.com/get-icon/geticon/raw/master/icons/ethereum.svg" alt="Ethereum" width="32px" height="32px"></a>
<a href="https://docs.soliditylang.org/en/v0.8.17/" title="Solidity"><img src="https://cdn.worldvectorlogo.com/logos/solidity.svg" alt="Solidity" width="32px" height="32px"></a>
<a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript" title="JS"><img src="https://github.com/get-icon/geticon/raw/master/icons/javascript.svg" alt="JS" width="32px" height="32px"></a>
<a href="https://go.dev/" title="Golang"><img src="https://github.com/get-icon/geticon/raw/master/icons/go.svg" alt="Golang" width="32px" height="32px"></a>
<a href="https://reactjs.org/" title="React"><img src="https://github.com/get-icon/geticon/raw/master/icons/react.svg" alt="React" width="32px" height="32px"></a>
<a href="https://vuejs.org/" title="Vue.js"><img src="https://github.com/get-icon/geticon/raw/master/icons/vue.svg" alt="Vue.js" width="32px" height="32px"></a>
<a href="https://code.visualstudio.com/" title="Visual Studio Code"><img src="https://github.com/get-icon/geticon/raw/master/icons/visual-studio-code.svg" alt="Visual Studio Code" width="32px" height="32px"></a>

  </div>

<br />

# About

This is a template project for developing and deploying full stack apps that use and interact with smart contracts. It uses automation to update the stack with the most current contract deployments which enables fast and efficient development.

# Getting Started

To run this collection of projects, you will need to add and fill `.env` files for all parts of the stack. The examples of how `.env` files should look like are provided in `.env.dev` files in the same directory.

To export and run the backend project you will need to install golang along with `go-ethereum` and `solcjs`.

### Stack Explained

- `Solidity`: Contains a hardhat project with utility scripts for fast and automated deployment of smart contracts. The scripts folder contains examples of how to deploy and export contracts.

- `Backend`: Contains a golang project with most common examples of smart contract interactions using the custom `gef` package. (gef is a utility package for simple smart contract interactions using the `go-ethereum` package)

- `Frontend`: The frontend folder contains a create-react-app project with most common smart contract interactions.

# Deploying Contracts

## Deploying Contracts to a Local Blockchain

Using a local blockchain instance is the best way to start developing web3 apps. The local blockchain should fork a live Ethereum blockchain so you can interact with existing contracts on that blockchain. The changes you make to those contracts will reset when you stop running the local blockchain, so they are not permanent and are only used when developing your projects.

First we need to run the local blockchain by spinning up a local node. To spin up a local node run the following command in the console:

```bash
  ...\eth-full-stack-template\solidity> npx hardhat node
```

To deploy a contract on the local node we will create a js script in the scripts folder.

The js script will call the `deployContract(contractName, contractArgs..)` function.

```javascript
async function main() {
  // deploy the contract
  await deployContract("StupidContract", 1000, "My example");
}
```

First argument is the exact name of the solidity contract. It is followed by the arguments of the smart contract constructor.

The `deployContract` function returns the address of the deployed contract.

Run the js script:

```bash
  ...\eth-full-stack-template\solidity> npx hardhat run --network localhost .\scripts\deploy_localhost.js
```

With this command we specify the localhost network and the script to run.

In order for the backend and frontend to work we need to export the contract data. This can be done separately from deployment with the `exportContract(contractName, contractAddress)` function. You can also export contract right after you deploy them with the `deployAndExport(contractName, contractArgs..)` function.

```javascript
async function main() {
  // deploy and export contract
  await deployAndExport("StupidContract", 1000, "My example");
}
```

## Deploy Contracts to a Live Blockchain

You can deploy contracts to a live blockchain like `Mainnet` or the `Goerli` testnet with the following command

```bash
  ...\eth-full-stack-template\solidity> npx hardhat run --network goerli .\scripts\deploy_export_testnet.js
```

Just specify one of the networks from `hardhat.config.js`. The address associated with the private key in the `.env` file should have available ether in order to deploy the contract.

You can verify the contracts code on `Etherscan` by calling the `verify(contractAddress, contractArgs[])` function.

```javascript
async function main() {
  // deploy and export contract
  const address = await deployAndExport("StupidContract", 1000, "My example");

  // verify contract on etherscan
  await verify(address, [1000, "My example"]);
}
```

# Interacting with Contracts

In order for the backend and fronted to work you need to export the contract you wish to interact with.

## Backend

The exported contract creates a `contracts` folder in the backend directory with the generated golang package that contains an `API` and an `address` of the deployed contract.

In order to run the backend project you need to specify the network as an argument.
```bash
  ...\eth-full-stack-template\backend> go run .\cmd\main.go LOCALHOST
```

The network can be either LOCALHOST, TESTNET or MAINNET.

The backend project contains these examples of the `gef` package usage:
- Create a normal transaction
- Create a contract transaction (calling a contract function) 
- Call a contract view function
- Fetch contract logs (supports unlimited number of logs)
- Subscribe to contract events (only if your `.env` file contains web socket urls)

Most RPC nodes will return a maximum of 10,000 logs at a time, so the fetch function of the `gef` package will call the RPC node multiple times in order to fetch all logs.

## Frontend

The exported contract creates a js script in the `src/contracts` folder that contains an instance of a contract class from the `Ethers.js` library along with the contracts `ABI` and `address`. 

All blockchain interactions go through the provider of the browser extension wallet.

The frontend project contains these examples:
- Connect and disconnect a browser extension wallet
- Create a contract transaction (calling a contract function) 
- Call a contract view function
- Fetch contracts logs (up to 10k logs)
- Listen for incoming events
