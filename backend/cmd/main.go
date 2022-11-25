package main

import (
	"backend/contracts/ERC20"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
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

	// paymentContract, err := PaymentContract.NewPaymentContract(
	// 	common.HexToAddress(PaymentContract.Address),
	// 	client,
	// )

	erc20, err := ERC20.NewERC20(
		common.HexToAddress(ERC20.Address),
		client,
	)

	err = GetEvents(
		client,
		erc20,
		big.NewInt(16041981),
		nil,
		[]string{"Transfer(address,address,uint256)", "Approval(address,address,uint256)"},
	)

	if err != nil {
		log.Fatal(err)
	}

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

func GetEvents(client *ethclient.Client, erc20Contract *ERC20.ERC20, startBlock *big.Int, endBlock *big.Int, eventSignatures []string) error {

	hashedEventSigs := make([]common.Hash, len(eventSignatures))

	for indx := range eventSignatures {
		hashedEventSigs[indx] = crypto.Keccak256Hash([]byte(eventSignatures[indx]))
	}

	// contractAbi, err := abi.JSON(strings.NewReader(string(ERC20.ERC20ABI)))
	// if err != nil {
	// 	return err
	// }

	eventSum := 0

	var fromBlock *big.Int
	var toBlock *big.Int

	if startBlock != nil {
		fromBlock = new(big.Int).Set(startBlock)
	}
	if endBlock != nil {
		toBlock = new(big.Int).Set(endBlock)
	}

	for {

		query := ethereum.FilterQuery{
			FromBlock: fromBlock,
			ToBlock:   toBlock,
			Addresses: []common.Address{common.HexToAddress(ERC20.Address)},
			Topics:    [][]common.Hash{hashedEventSigs},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {

			if err.Error()[0:27] == "Log response size exceeded." {
				split := strings.Split(err.Error(), "[")
				numbers := strings.Split(split[1], " ")
				first, err := strconv.ParseInt(strings.Replace(numbers[0], ",", "", -1)[2:], 16, 64)
				if err != nil {
					return err
				}
				second, err := strconv.ParseInt(strings.Replace(numbers[1], "]", "", -1)[2:], 16, 64)
				if err != nil {
					return err
				}

				fromBlock = big.NewInt(first)
				toBlock = big.NewInt(second)
				continue
			} else {
				return err
			}
		}

		for _, vLog := range logs {
			// fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			// fmt.Printf("Log Index: %d\n", vLog.Index)

			for _, eventSigHash := range hashedEventSigs {

				if vLog.Topics[0].Hex() == eventSigHash.Hex() {

					transferEvent, err := erc20Contract.ParseTransfer(vLog)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Printf("From: %s\n", transferEvent.From.String())
					fmt.Printf("To: %s\n", transferEvent.To.String())
					fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

					break
				}
			}

			// switch vLog.Topics[0].Hex() {
			// case logTransferSigHash.Hex():

			// 	transferEvent, err := erc20Contract.ParseTransfer(vLog)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}

			// 	// transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			// 	// transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			// 	fmt.Printf("From: %s\n", transferEvent.From.String())
			// 	fmt.Printf("To: %s\n", transferEvent.To.String())
			// 	fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

			// case logApprovalSigHash.Hex():

			// 	var approvalEvent LogApproval

			// 	approvalEvent, err := contractAbi.Unpack("Approval", vLog.Data)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}

			// 	approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			// 	approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			// 	// fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
			// }
		}

		eventSum += len(logs)

		fmt.Println(fromBlock.String(), " -> ", toBlock.String(), " | Logs: ", len(logs))

		// if this is the end, break
		if toBlock.String() == endBlock.String() {
			break
		}

		fromBlock = fromBlock.Add(toBlock, big.NewInt(1))
		toBlock = nil
		if endBlock != nil {
			toBlock = new(big.Int).Set(endBlock)
		}
	}

	fmt.Println(eventSum)

	return nil
}
