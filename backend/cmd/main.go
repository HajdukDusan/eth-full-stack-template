package main

import (
	"backend/contracts/ERC20"
	"backend/pkg/gef"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {

	// Load Env File
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	rpcUrl := os.Getenv("ALCHEMY_MAINNET_RPC_WS_URL")

	mainnetClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	erc20API, err := ERC20.NewERC20(
		common.HexToAddress(ERC20.Address),
		mainnetClient,
	)

	events, err := gef.GetLogs(
		mainnetClient,
		nil,
		nil,
		[]gef.EventWrapper{
			{
				Name: "Transfer",
				Args: []string{"address", "address", "uint256"},
				ParseMethod: func(log types.Log) (interface{}, error) {
					return erc20API.ParseTransfer(log)
				},
			},
			{
				Name: "Approval",
				Args: []string{"address", "address", "uint256"},
				ParseMethod: func(log types.Log) (interface{}, error) {
					return erc20API.ParseApproval(log)
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(events))

	// for _, event := range events {
	// 	switch obj := event.(type) {
	// 	case *ERC20.ERC20Transfer:
	// 		fmt.Println("Transfer")
	// 		fmt.Println("From:", obj.From)
	// 		fmt.Println("To:", obj.To)
	// 		fmt.Println("Tokens", obj.Tokens)
	// 	case *ERC20.ERC20Approval:
	// 		fmt.Println("Approve")
	// 		fmt.Println("Spender:", obj.Spender)
	// 		fmt.Println("TokenOwner:", obj.TokenOwner)
	// 		fmt.Println("Tokens", obj.Tokens)
	// 	default:
	// 		fmt.Printf("Strange Object")
	// 	}
	// }
}
