package examples

import (
	"backend/contracts/ERC20"
	"backend/contracts/StupidContract"
	"backend/pkg/gef"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendContractTx(client *ethclient.Client, privateKey string) {
	stupidContractAPI, err := StupidContract.NewStupidContract(
		common.HexToAddress(StupidContract.Address),
		client,
	)
	if err != nil {
		log.Fatal(err)
	}

	// send tx
	tx, err := gef.SendTx(
		client,
		privateKey,
		big.NewInt(100),
		func(txOpts *bind.TransactOpts) (*types.Transaction, error) {
			return stupidContractAPI.AddToRegistry(txOpts, "moj parametar")
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// wait for tx to be mined
	receipt, err := gef.WaitTxReceipt(client, tx)
	if err != nil {
		log.Fatal(err)
	}

	// receipt will return emited events
	for _, log := range receipt.Logs {

		stupidEvent, err := stupidContractAPI.ParseStupidEvent(*log)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(stupidEvent)
	}
}

func CallViewFunc(client *ethclient.Client) {

	stupidContractAPI, err := StupidContract.NewStupidContract(
		common.HexToAddress(StupidContract.Address),
		client,
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := stupidContractAPI.StupidContractDescription(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func GetLogs(client *ethclient.Client) {

	erc20API, err := ERC20.NewERC20(
		common.HexToAddress(ERC20.Address),
		client,
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := gef.GetLogs(
		client,
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
