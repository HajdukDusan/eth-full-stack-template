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
	"github.com/schollz/progressbar/v3"
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

	events, err := GetEvents(
		mainnetClient,
		nil,
		nil,
		[]EventWrapper{
			{
				"Transfer",
				[]string{"address", "address", "uint256"},
				func(log types.Log) (interface{}, error) {
					return erc20API.ParseTransfer(log)
				},
			},
			{
				"Approval",
				[]string{"address", "address", "uint256"},
				func(log types.Log) (interface{}, error) {
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

type EventWrapper struct {
	Name        string
	Args        []string
	ParseMethod func(types.Log) (interface{}, error)
}

func GetEvents(client *ethclient.Client, startBlock *big.Int, endBlock *big.Int, events []EventWrapper) ([]interface{}, error) {

	eventSigs := generateSignatures(events)

	var fromBlock *big.Int = setBigInt(startBlock)
	var toBlock *big.Int = setBigInt(endBlock)

	result := make([]interface{}, 0)

	stats := make([]int64, len(eventSigs))

	bar, err := setupProgressBar(client, fromBlock, toBlock)
	if err != nil {
		return nil, err
	}

	for {
		logs, newFromBlock, newToBlock, err := fetchEvents(client, fromBlock, toBlock, eventSigs)
		if err != nil {
			return nil, err
		}
		fromBlock = newFromBlock
		toBlock = newToBlock

		logs_arr := make([]interface{}, len(logs))

		for log_indx, vLog := range logs {

			// find the matching event
			for event_indx, eventSigHash := range eventSigs {

				if vLog.Topics[0].Hex() == eventSigHash.Hex() {

					parsedEvent, err := events[event_indx].ParseMethod(vLog)
					if err != nil {
						return nil, err
					}
					logs_arr[log_indx] = parsedEvent
					stats[event_indx]++
					break
				}
			}
		}

		if toBlock == nil {
			bar.Finish()
		} else {
			bar.Add64(toBlock.Int64() - fromBlock.Int64())
		}

		// fmt.Println(fromBlock.String(), " -> ", toBlock.String(), " | Logs: ", len(logs_arr))

		result = append(result, logs_arr...)

		// if this is the end, break
		if toBlock.String() == endBlock.String() {
			break
		}

		fromBlock = fromBlock.Add(toBlock, big.NewInt(1))
		toBlock = setBigInt(endBlock)
	}

	for indx := range stats {
		fmt.Println(events[indx].Name+" events:", stats[indx])
	}

	return result, nil
}

func setupProgressBar(client *ethclient.Client, startBlock *big.Int, endBlock *big.Int) (*progressbar.ProgressBar, error) {

	var start int64 = 0
	var end int64 = 0

	if startBlock != nil {
		start = startBlock.Int64()
	}

	if endBlock != nil {
		end = endBlock.Int64()
	} else {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		end = header.Number.Int64()
	}

	return progressbar.NewOptions64(
		end-start,
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionSetItsString("blocks"),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionSetWidth(25),
		progressbar.OptionSetDescription("Fetching logs"),
	), nil
}

func generateSignatures(events []EventWrapper) []common.Hash {

	result := make([]common.Hash, len(events))

	for event_indx, event := range events {

		sig := event.Name + "("

		for arg_indx, arg := range event.Args {
			sig += arg
			if arg_indx != len(event.Args)-1 {
				sig += ","
			} else {
				sig += ")"
			}
		}

		result[event_indx] = crypto.Keccak256Hash([]byte(sig))
	}

	return result
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
				if (fromBlock != nil && fromBlock.Int64() == first) && (toBlock != nil && toBlock.Int64() == second) {
					break
				}

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
