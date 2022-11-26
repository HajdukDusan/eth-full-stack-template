package gef

import (
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
	"github.com/schollz/progressbar/v3"
)

func getLogsInternal(client *ethclient.Client, contractAddresses []string, startBlock *big.Int, endBlock *big.Int, events []EventWrapper) ([]interface{}, error) {

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
		logs, newFromBlock, newToBlock, err := fetchEvents(client, contractAddresses, fromBlock, toBlock, eventSigs)
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
		updateProgressBar(bar, fromBlock, toBlock)

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

func updateProgressBar(bar *progressbar.ProgressBar, fromBlock *big.Int, toBlock *big.Int) {
	if toBlock == nil {
		bar.Finish()
	} else {
		bar.Add64(toBlock.Int64() - fromBlock.Int64())
	}
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

func convertAddresses(addresses []string) []common.Address {
	result := make([]common.Address, len(addresses))

	for indx := range addresses {
		result[indx] = common.HexToAddress(addresses[indx])
	}

	return result
}

func getQuery(fromBlock *big.Int, toBlock *big.Int, addresses []common.Address, eventSigs []common.Hash) ethereum.FilterQuery {

	return ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: addresses,
		Topics:    [][]common.Hash{eventSigs},
	}
}

func subscribeToEventsInternal(client *ethclient.Client, contractAddresses []string, events []EventWrapper) {

	addresses := convertAddresses(contractAddresses)
	eventSigs := generateSignatures(events)

	query := getQuery(nil, nil, addresses, eventSigs)

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

func fetchEvents(client *ethclient.Client, contractAddresses []string, fromBlock *big.Int, toBlock *big.Int, eventSigs []common.Hash) ([]types.Log, *big.Int, *big.Int, error) {

	addresses := convertAddresses(contractAddresses)

	for {
		query := getQuery(fromBlock, toBlock, addresses, eventSigs)

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
