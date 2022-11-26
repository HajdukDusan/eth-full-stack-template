package gef

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EventWrapper struct {
	Name        string
	Args        []string
	ParseMethod func(types.Log) (interface{}, error)
}

// GetLogs returns logs for specified events within a certain block range.
//
// Arguments:
//   - client - the ethclient client pointer
//   - startBlock - if nil start from block 0
//   - endBlock - if nil end at the most recent block
//   - events - the list of EventWrapper objects that contains the name of the event, its arg types and
//     a ParseMethod which determains how the fetched log should be parsed
func GetLogs(client *ethclient.Client, startBlock *big.Int, endBlock *big.Int, events []EventWrapper) ([]interface{}, error) {

	return getLogsInternal(client, startBlock, endBlock, events)
}
