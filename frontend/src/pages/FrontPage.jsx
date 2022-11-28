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

import StupidContract from "../contracts/StupidContract";

const FrontPage = () => {
  const viewIndexInputRef = useRef();
  const txMsgInputRef = useRef();

  const [provider, setProvider] = useState(null);
  const [networkChainId, setChainId] = useState(null);

  const [signer, setSigner] = useState(null);
  const [walletAddress, setWalletAddress] = useState(null);
  const [userBalance, setUserBalance] = useState(null);

  const stupidContract = new ethers.Contract(
    StupidContract.Address,
    StupidContract.ABI
  );

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
          stupidContract.connect(signer);
          setSigner(signer);
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

    setSigner(null);
    setWalletAddress(null);
    userBalance(null);
  }

  async function contractView() {
    const contract = new ethers.Contract(
      StupidContract.Address,
      StupidContract.ABI,
      provider
    );

    const result = await contract.stupidRegistry(
      viewIndexInputRef.current.value
    );

    alert(result);
  }

  async function contractTx() {
    // send tx
    let tx = await stupidContract
      .connect(signer)
      .AddToRegistry(txMsgInputRef.current.value, { value: 1000 });

    alert("tx sent with hash: " + tx.hash);

    // wait for tx to be mined
    const receipt = await tx.wait();
    if (receipt.status !== 1) {
      alert("Tx failed!");
      return;
    }

    // get logs from tx receipt
    const logs = parseLogs(receipt.logs, StupidContract.ABI);
    logs.forEach((log) => {
      console.log(log);
    });
  }

  function parseLogs(logs, abi) {
    let abiInterface = new ethers.utils.Interface(abi);
    return logs.map((log) => abiInterface.parseLog(log));
  }

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
        </Box>
      </Box>
    </>
  );
};

export default FrontPage;
