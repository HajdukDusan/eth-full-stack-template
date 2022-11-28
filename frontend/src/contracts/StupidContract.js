import { ethers } from "ethers"
const ABI= [
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_entryFee",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "_description",
        "type": "string"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "inputs": [],
    "name": "StupidContract__FeeNotEnough",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "StupidContract__NotOwner",
    "type": "error"
  },
  {
    "inputs": [],
    "name": "StupidContract__TransferFailed",
    "type": "error"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "index",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "timestamp",
        "type": "uint256"
      }
    ],
    "name": "StupidEvent",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "string",
        "name": "_message",
        "type": "string"
      }
    ],
    "name": "AddToRegistry",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "MessageEntryFee",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "PullStupidFees",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "StupidContractDescription",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "stupidRegistry",
    "outputs": [
      {
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "internalType": "string",
        "name": "message",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]
const Address = "0x707531c9999AaeF9232C8FEfBA31FBa4cB78d84a"
const Contract = new ethers.Contract(Address, ABI)
export {Contract, Address, ABI}