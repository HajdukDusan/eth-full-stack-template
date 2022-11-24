package main

import (
	"backend/contracts/ERC20"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	GetEvents(client, big.NewInt(139975), big.NewInt(6770802))

	// paymentContract, err := PaymentContract.NewPaymentContract(
	// 	common.HexToAddress(PaymentContract.Address),
	// 	client,
	// )

	// fmt.Println(paymentContract)
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func GetEvents(client *ethclient.Client, fromBlock *big.Int, toBlock *big.Int) {

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{common.HexToAddress(ERC20.Address)},
		Topics:    [][]common.Hash{{logTransferSigHash, logApprovalSigHash}},
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(ERC20.ERC20ABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			// fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			// fmt.Printf("Log Index: %d\n", vLog.Index)

			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				// fmt.Printf("Log Name: Transfer\n")

				var transferEvent LogTransfer

				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				// fmt.Printf("From: %s\n", transferEvent.From.Hex())
				// fmt.Printf("To: %s\n", transferEvent.To.Hex())
				// fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

			case logApprovalSigHash.Hex():
				// fmt.Printf("Log Name: Approval\n")

				var approvalEvent LogApproval

				err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
				approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

				// fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
				// fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
				// fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())

			default:
				fmt.Println("kita")
			}

			fmt.Println(len(logs))
		}
	}

}
