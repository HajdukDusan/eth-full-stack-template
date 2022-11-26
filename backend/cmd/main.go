package main

import (
	"backend/internal/examples"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {

	// Load Env File
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	privKey := os.Getenv("GOERLI_PRIVATE_KEY")
	// privKey := os.Getenv("LOCALHOST_PRIVATE_KEY")

	// rpcUrl := os.Getenv("ALCHEMY_MAINNET_RPC_WS_URL")
	rpcUrl := os.Getenv("ALCHEMY_GOERLI_RPC_WS_URL")
	// rpcUrl := os.Getenv("LOCALHOST_RPC_WS_URL")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	examples.SendContractTx(client, privKey)
	// examples.callViewFunc(client)
	// examples.getLogs(client)
}

// normal transfer tx example needed
