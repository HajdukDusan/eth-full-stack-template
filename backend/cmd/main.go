package main

import (
	"backend/contracts/PaymentContract"
	"fmt"
	"os"

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

	rpcUrl := os.Getenv("ALCHEMY_RPC_WS_URL")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	paymentContract, err := PaymentContract.NewPaymentContract(
		common.HexToAddress(PaymentContract.Address),
		client,
	)

	fmt.Println(paymentContract)
}
