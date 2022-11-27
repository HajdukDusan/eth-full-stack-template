package main

import (
	"backend/contracts/StupidContract"
	"backend/internal/examples"
	"backend/pkg/gef"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func main() {

	//###########################
	client, privKey := gef.ConfigureClient(os.Args[1])
	//###########################

	//###########################
	stupidContractAPI, err := StupidContract.NewStupidContract(
		common.HexToAddress(StupidContract.Address),
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	//###########################

	//###########################
	go examples.SubscribeToEvent(client, stupidContractAPI)
	// examples.SendNormalTx(client, privKey, "0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	examples.SendContractTx(client, privKey, stupidContractAPI)
	// examples.GetContractLogs(client, stupidContractAPI)
	// examples.CallContractViewFunc(client, stupidContractAPI)

}
