package main

import (
	"backend/contracts/StupidContract"
	"backend/internal/examples"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {

	// Load Env File
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// privKey := os.Getenv("MAINNET_PRIVATE_KEY")
	privKey := os.Getenv("GOERLI_PRIVATE_KEY")
	// privKey := os.Getenv("LOCALHOST_PRIVATE_KEY")

	// rpcUrl := os.Getenv("ALCHEMY_MAINNET_RPC_WS_URL")
	rpcUrl := os.Getenv("ALCHEMY_GOERLI_RPC_WS_URL")
	// rpcUrl := os.Getenv("LOCALHOST_RPC_WS_URL")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	stupidContractAPI, err := StupidContract.NewStupidContract(
		common.HexToAddress(StupidContract.Address),
		client,
	)
	if err != nil {
		log.Fatal(err)
	}

	// usage examples

	go examples.SubscribeToEvent(client, stupidContractAPI)

	time.Sleep(5 * time.Second)

	// examples.SendNormalTx(client, privKey, "0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	examples.SendContractTx(client, privKey, stupidContractAPI)

	time.Sleep(5 * time.Second)

	examples.SendContractTx(client, privKey, stupidContractAPI)

	time.Sleep(10 * time.Second)

	examples.SendContractTx(client, privKey, stupidContractAPI)
	// examples.GetContractLogs(client, stupidContractAPI)
	// examples.CallContractViewFunc(client, stupidContractAPI)

}
