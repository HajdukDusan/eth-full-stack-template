import { Box, Button, Input, Text, Center, FormLabel } from "@chakra-ui/react";
import React, { useRef, useState, useEffect } from "react";

import BigDecimal from "js-big-decimal";

import { ethers } from "ethers";

import StupidContract from "../contracts/StupidContract";

const FrontPage = () => {
  const [errorMessage, setErrorMessage] = useState(null);

  const [provider, setProvider] = useState(null);

  const [networkChainId, setChainId] = useState(null);
  const [defaultAccount, setDefaultAccount] = useState(null);
  const [userBalance, setUserBalance] = useState(null);

  async function connect() {
    if (window.ethereum) {
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      setProvider(provider);
      setChainId((await provider.getNetwork()).chainId);
      provider.send("eth_requestAccounts", []).then(async () => {
        await accountChangedHandler(provider.getSigner());
      });
    } else {
      setErrorMessage("Please Install Metamask!!!");
    }
  }

  async function disconnect() {
    setDefaultAccount(null);
    setUserBalance(null);
    setProvider(null);
    setChainId(null);
  }

  async function callContractViewFunc() { 

    const contract = new ethers.Contract(StupidContract.Address, StupidContract.ABI, provider)

    alert(await contract.StupidContractDescription())
  }

  async function contractTx() { 

    const contract = new ethers.Contract(StupidContract.Address, StupidContract.ABI, provider)

    alert(await contract.StupidContractDescription())
  }

  const accountChangedHandler = async (newAccount) => {
    const address = await newAccount.getAddress();
    setDefaultAccount(address);
    const balance = await newAccount.getBalance();
    setUserBalance(ethers.utils.formatEther(balance));
  };

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
        <Text fontSize="2xl">Connected With: {defaultAccount}</Text>
        <Text fontSize="2xl">Wallet Amount: {userBalance} ETH</Text>

        <Box maxW="2xl" mx={"auto"} pt={10} px={{ base: 20, sm: 15, md: 20 }}>
          <Button onClick={connect}>Connect Wallet</Button>

          <br></br>
          <br></br>

          <Button onClick={disconnect}>Disconnect Wallet</Button>

          <br></br>
          <br></br>

          <Button onClick={callContractViewFunc}>Contract Call View Function</Button>

          <FormLabel color="white">Message</FormLabel>
          <Input
            // onChange={(event) => setAddress(event.target.value)}
            type="text"
            // value={address}
          />
        </Box>
      </Box>

      {errorMessage}
    </>
  );
};

export default FrontPage;
