import {
  Box,
  Button,
  Input,
  Text,
  Center,
  FormLabel,
  Flex,
} from "@chakra-ui/react";
import React, { useRef, useState, useEffect } from "react";

import BigDecimal from "js-big-decimal";

import { ethers } from "ethers";

import {
  Contract as StupidContract,
  Address as StupidContractAddress,
  ABI as StupidContractABI,
} from "../contracts/StupidContract";

const FrontPage = () => {
  const viewIndexInputRef = useRef();
  const txMsgInputRef = useRef();

  const [provider, setProvider] = useState(null);
  const [networkChainId, setChainId] = useState(null);
  const [walletAddress, setWalletAddress] = useState(null);
  const [userBalance, setUserBalance] = useState(null);

  async function connect() {
    if (window.ethereum) {
      const provider = new ethers.providers.Web3Provider(window.ethereum);

      setProvider(provider);
      setChainId((await provider.getNetwork()).chainId);

      provider
        .send("eth_requestAccounts", [])
        .then(async () => {
          const signer = provider.getSigner();

          setWalletAddress(await signer.getAddress());
          setUserBalance(ethers.utils.formatEther(await signer.getBalance()));
        })
        .catch((error) => {
          alert("Request Accounts Failed");
        });
    } else {
      alert("Please Install Metamask.");
    }
  }

  async function disconnect() {
    setProvider(null);
    setChainId(null);

    setWalletAddress(null);
    setUserBalance(null);
  }

  async function contractView() {
    const result = await StupidContract.connect(
      provider.getSigner()
    ).stupidRegistry(viewIndexInputRef.current.value);

    alert(result);
  }

  async function contractTx() {
    // send tx
    let tx = await StupidContract.connect(provider.getSigner()).AddToRegistry(
      txMsgInputRef.current.value,
      { value: 1000 }
    );

    alert("tx sent with hash: " + tx.hash);

    // wait for tx to be mined
    // if on localhost it cant be trusted :)
    const receipt = await tx.wait();
    if (receipt.status !== 1) {
      alert("Tx failed!");
      return;
    }

    // get logs from tx receipt
    const logs = parseLogs(receipt.logs, StupidContractABI);
    logs.forEach((log) => {
      console.log(log);
    });
  }

  function parseLogs(logs, abi) {
    let abiInterface = new ethers.utils.Interface(abi);
    return logs.map((log) => abiInterface.parseLog(log));
  }

  // up to 10k events
  async function getLogs() {
    let filter = StupidContract.filters.StupidEvent();
    filter.fromBlock = 0;
    filter.toBlock = "latest";

    const logs = await provider.getLogs(filter);

    const parsedLogs = parseLogs(logs, StupidContractABI);

    console.log(parsedLogs);
  }

  function listenForEvent() {

    // dont run miltiple times without reseting the provider
    StupidContract.connect(provider).on("StupidEvent", (index, sender, timestamp) => {
      console.log(index);
      console.log(sender);
      console.log(timestamp);
    });
    // at first this returns the new-est event, dont know why
  }

  //TODO get historical txs

  return (
    <>
      <Box maxW="2xl" mx={"auto"} pt={1} px={{ base: 2, sm: 12, md: 17 }}>
        <Center>
          <Text fontSize="6xl">Stupid Contract Front</Text>
        </Center>
      </Box>

      <Box maxW="4xl" mx={"auto"} pt={5} px={{ base: 2, sm: 12, md: 17 }}>
        <br></br>

        <Text fontSize="2xl">Chain ID: {networkChainId}</Text>
        <Text fontSize="2xl">Connected With: {walletAddress}</Text>
        <Text fontSize="2xl">Wallet Amount: {userBalance} ETH</Text>

        <Box maxW="2xl" mx={"auto"} pt={10} px={{ base: 20, sm: 15, md: 20 }}>
          <Flex justifyContent="space-between">
            <Button onClick={connect}>Connect Wallet</Button>
            <Button onClick={disconnect}>Disconnect Wallet</Button>
          </Flex>

          <br></br>
          <br></br>

          <Button onClick={contractView}>Contract Call View Function</Button>
          <FormLabel color="white">Index parameter</FormLabel>
          <Input ref={viewIndexInputRef} type="text" />

          <br></br>
          <br></br>

          <Button onClick={contractTx}>Create Contract Tx</Button>
          <FormLabel color="white">Message parameter</FormLabel>
          <Input ref={txMsgInputRef} type="text" />

          <br></br>
          <br></br>

          <Button onClick={getLogs}>Fetch Logs</Button>

          <br></br>
          <br></br>

          <Button onClick={listenForEvent}>Listen For Event</Button>
        </Box>
      </Box>
    </>
  );
};

export default FrontPage;
