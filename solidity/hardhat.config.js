require("@nomiclabs/hardhat-ethers");
require("@nomicfoundation/hardhat-chai-matchers");
require("@nomiclabs/hardhat-etherscan")
require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config()
require("hardhat-deploy")

const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY
const GOERLI_RPC_URL = process.env.GOERLI_RPC_URL
const MAINNET_RPC_URL = process.env.MAINNET_RPC_URL

const MAINNET_INFURA_KEY = process.env.MAINNET_INFURA_KEY

const GOERLI_PRIVATE_KEY = process.env.GOERLI_PRIVATE_KEY
const MAINNET_PRIVATE_KEY = process.env.MAINNET_PRIVATE_KEY

module.exports = {
  solidity: "0.8.17",
  networks: {
    hardhat: {
      forking: {
        url: MAINNET_INFURA_KEY
      },
      chainId: 31337,
      accounts: {
        count: 10
      }
    },
    goerli: {
      url: GOERLI_RPC_URL,
      accounts: [GOERLI_PRIVATE_KEY],
      chainId: 5,
      blockConfirmations: 1,
    },
    mainnet: {
      url: MAINNET_RPC_URL,
      accounts: [MAINNET_PRIVATE_KEY],
      chainId: 1,
      blockConfirmations: 1,
    },
  },
  etherscan: {
    apiKey: ETHERSCAN_API_KEY,
  },
  namedAccounts: {
    deployer: {
      default: 0,
    },
  },
};
