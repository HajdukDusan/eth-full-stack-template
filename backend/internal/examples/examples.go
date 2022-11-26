package examples

import (
	"backend/contracts/StupidContract"
	"backend/pkg/gef"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// subscripcija na event

func SendNormalTx(client *ethclient.Client, privateKey string, receiverAddress string) {

	// send tx
	_, err := gef.SendNormalTx(
		client,
		privateKey,
		big.NewInt(1000),
		receiverAddress,
	)
	if err != nil {
		log.Fatal(err)
	}

	// cant wait for receipt on localhost becase the tx finishes before wait starts
	// receipt, err := gef.WaitTxReceipt(client, tx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func SendContractTx(client *ethclient.Client, privateKey string, stupidContractAPI *StupidContract.StupidContract) {

	// send tx
	tx, err := gef.SendContractTx(
		client,
		privateKey,
		big.NewInt(1000),
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

func CallContractViewFunc(client *ethclient.Client, stupidContractAPI *StupidContract.StupidContract) {

	result, err := stupidContractAPI.StupidContractDescription(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func GetContractLogs(client *ethclient.Client, stupidContractAPI *StupidContract.StupidContract) {

	events, err := gef.GetLogs(
		client,
		[]string{StupidContract.Address},
		nil,
		nil,
		[]gef.EventWrapper{
			{
				Name: "StupidEvent",
				Args: []string{"uint256", "address", "uint256"},
				ParseMethod: func(log types.Log) (interface{}, error) {
					return stupidContractAPI.ParseStupidEvent(log)
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		switch obj := event.(type) {
		case *StupidContract.StupidContractStupidEvent:
			fmt.Println("StupidEvent")
			fmt.Println("\tIndex:", obj.Index)
			fmt.Println("\tSender:", obj.Sender)
			fmt.Println("\tTimestamp:", obj.Timestamp)
		default:
			fmt.Printf("Unexpected log object")
		}
	}
}
