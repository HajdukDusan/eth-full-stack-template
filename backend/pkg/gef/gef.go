package gef

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
func GetLogs(
	client *ethclient.Client,
	startBlock *big.Int,
	endBlock *big.Int,
	events []EventWrapper,
) ([]interface{}, error) {

	return getLogsInternal(client, startBlock, endBlock, events)
}

// SendTx send the tx to the mempool
//
// Arguments:
//   - client - the ethclient client pointer
//   - privateKey - EOA private key
//   - value - eth value to be sent
//   - initializeTx - method which to call
func SendTx(
	client *ethclient.Client,
	privateKey string,
	value *big.Int,
	initializeTx func(*bind.TransactOpts) (*types.Transaction, error),
) (*types.Transaction, error) {

	return sendTxInternal(client, privateKey, value, initializeTx)
}

// WaitTxReceipt waits for the tx to be mined
//
// Arguments:
//   - client - the ethclient client pointer
//   - tx - the tx that was previously sent to the mempool
func WaitTxReceipt(
	client *ethclient.Client,
	tx *types.Transaction,
) (*types.Receipt, error) {

	return waitTxReceiptInternal(client, tx)
}
