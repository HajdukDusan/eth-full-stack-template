import { Box, Button, Input, Text, Center, FormLabel } from "@chakra-ui/react";
import React, { useRef, useState, useEffect } from "react";

import BigDecimal from "js-big-decimal";

import { ethers } from "ethers";

import StupidContract from "../contracts/StupidContract";

const FrontPage = () => {
  const viewIndexInputRef = useRef();
  const txMsgInputRef = useRef();

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
      setErrorMessage("Please Install Metamask.");
    }
  }

  async function disconnect() {
    setDefaultAccount(null);
    setUserBalance(null);
    setProvider(null);
    setChainId(null);
  }

  async function callContractViewFunc() {
    const contract = new ethers.Contract(
      StupidContract.Address,
      StupidContract.ABI,
      provider
    );



    const result = await contract.stupidRegistry(viewIndexInputRef.current.value);

    alert(result);
  }

  async function contractTx() {
    const signer = await provider.getSigner(defaultAccount);

    const contractWithSigner = new ethers.Contract(
      StupidContract.Address,
      StupidContract.ABI,
      signer
    );

    //const txOptions = {value: ethers.utils.parseEther("1.0")}
    const txOptions = { value: 1000 };

    let tx = await contractWithSigner.AddToRegistry(
      txMsgInputRef.current.value,
      txOptions
    );
    console.log(tx.hash);

    const receipt = await tx.wait();
    if (receipt.status !== 1) {
      alert("Tx failed!");
      return;
    }

    console.log(receipt);
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

          <Button onClick={callContractViewFunc}>
            Contract Call View Function
          </Button>
          <FormLabel color="white">Index parameter</FormLabel>
          <Input ref={viewIndexInputRef} type="text" />

          <br></br>
          <br></br>

          <Button onClick={contractTx}>Create Contract Tx</Button>
          <FormLabel color="white">Message parameter</FormLabel>
          <Input ref={txMsgInputRef} type="text" />
        </Box>
      </Box>

      {errorMessage}
    </>
  );
};

export default FrontPage;
