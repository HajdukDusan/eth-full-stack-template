package main

import (
	"backend/contracts/ERC20"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	mainnetClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	erc20, err := ERC20.NewERC20(
		common.HexToAddress(ERC20.Address),
		mainnetClient,
	)

	events, err := GetEvents(
		mainnetClient,
		erc20,
		big.NewInt(6383820),
		big.NewInt(6383840),
		[]string{"Transfer(address,address,uint256)", "Approval(address,address,uint256)"},
		[]func(types.Log) (interface{}, error){
			func(log types.Log) (interface{}, error) {
				return erc20.ParseTransfer(log)
			},
			func(log types.Log) (interface{}, error) {
				return erc20.ParseApproval(log)
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {

		// fmt.Println(reflect.TypeOf(event))

		// transferEvent, ok := event.(*ERC20.ERC20Transfer)
		// if ok {
		// 	fmt.Println("Transfer")
		// 	fmt.Println("From:", transferEvent.From)
		// 	fmt.Println("To:", transferEvent.To)
		// 	fmt.Println("Tokens", transferEvent.Tokens)
		// }

		switch obj := event.(type) {
		case *ERC20.ERC20Transfer:
			fmt.Println("Transfer")
			fmt.Println("From:", obj.From)
			fmt.Println("To:", obj.To)
			fmt.Println("Tokens", obj.Tokens)
		case *ERC20.ERC20Approval:
			fmt.Println("Approve")
			fmt.Println("Spender:", obj.Spender)
			fmt.Println("TokenOwner:", obj.TokenOwner)
			fmt.Println("Tokens", obj.Tokens)
		default:
			fmt.Printf("Strange Object")
		}
	}

	// fmt.Println(len(events))
}

func GetEvents(client *ethclient.Client, erc20Contract *ERC20.ERC20, startBlock *big.Int, endBlock *big.Int, eventSignatures []string, parseEvent []func(types.Log) (interface{}, error)) ([]interface{}, error) {

	hashedEventSigs := make([]common.Hash, len(eventSignatures))

	for indx := range eventSignatures {
		hashedEventSigs[indx] = crypto.Keccak256Hash([]byte(eventSignatures[indx]))
	}

	var fromBlock *big.Int = setBigInt(startBlock)
	var toBlock *big.Int = setBigInt(endBlock)

	result := make([]interface{}, 0)

	for {
		logs, newFromBlock, newToBlock, err := fetchEvents(client, fromBlock, toBlock, hashedEventSigs)
		if err != nil {
			return nil, err
		}
		fromBlock = newFromBlock
		toBlock = newToBlock

		logs_arr := make([]interface{}, len(logs))

		for log_indx, vLog := range logs {
			// fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			// fmt.Printf("Log Index: %d\n", vLog.Index)

			// find the matching event
			for event_indx, eventSigHash := range hashedEventSigs {

				if vLog.Topics[0].Hex() == eventSigHash.Hex() {

					parsedEvent, err := parseEvent[event_indx](vLog)
					if err != nil {
						return nil, err
					}
					logs_arr[log_indx] = parsedEvent
					break
				}
			}
		}

		fmt.Println(fromBlock.String(), " -> ", toBlock.String(), " | Logs: ", len(logs_arr))

		result = append(result, logs_arr...)

		// if this is the end, break
		if toBlock.String() == endBlock.String() {
			break
		}

		fromBlock = fromBlock.Add(toBlock, big.NewInt(1))

		toBlock = setBigInt(endBlock)
	}

	return result, nil
}

func setBigInt(value *big.Int) *big.Int {

	if value != nil {
		return new(big.Int).Set(value)
	}

	return nil
}

func fetchEvents(client *ethclient.Client, fromBlock *big.Int, toBlock *big.Int, hashedEventSigs []common.Hash) ([]types.Log, *big.Int, *big.Int, error) {

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
					return nil, nil, nil, err
				}
				second, err := strconv.ParseInt(strings.Replace(numbers[1], "]", "", -1)[2:], 16, 64)
				if err != nil {
					return nil, nil, nil, err
				}

				// break loop if no change in params
				// if (fromBlock != nil && fromBlock.Int64() == first) && (toBlock != nil && toBlock.Int64() == second) {
				// 	break
				// }

				fromBlock = big.NewInt(first)
				toBlock = big.NewInt(second)

				continue

			} else {
				return nil, nil, nil, err
			}
		}

		return logs, fromBlock, toBlock, nil
	}

	return nil, nil, nil, errors.New("Fetching resulted in infinite loop!")
}
