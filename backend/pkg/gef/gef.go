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
//   - contractAddresses - list of contracts to scan for logs
//   - startBlock - if nil start from block 0
//   - endBlock - if nil end at the most recent block
//   - events - the list of EventWrapper objects that contains the name of the event, its arg types and
//     a ParseMethod which determains how the fetched log should be parsed
func GetLogs(
	client *ethclient.Client,
	contractAddresses []string,
	startBlock *big.Int,
	endBlock *big.Int,
	events []EventWrapper,
) ([]interface{}, error) {

	return getLogsInternal(client, contractAddresses, startBlock, endBlock, events)
}

// SubscribeToEvent lets you listen for new events (need to use wss).
//
// Arguments:
//   - client - the ethclient client pointer
//   - contractAddresses - list of contracts to scan for logs
//   - events - the list of EventWrapper objects
func SubscribeToEvent(
	client *ethclient.Client,
	contractAddresses []string,
	events []EventWrapper,
) {
	subscribeToEventsInternal(client, contractAddresses, events)
}

// SendContractTx sends a contract tx to the mempool
//
// Arguments:
//   - client - the ethclient client pointer
//   - privateKey - EOA private key
//   - value - eth value to be sent
//   - initializeTx - method which to call
func SendContractTx(
	client *ethclient.Client,
	privateKey string,
	value *big.Int,
	initializeTx func(*bind.TransactOpts) (*types.Transaction, error),
) (*types.Transaction, error) {

	return sendContractTxInternal(client, privateKey, value, initializeTx)
}

// SendContractTx sends a normal tx to the mempool
//
// Arguments:
//   - client - the ethclient client pointer
//   - privateKey - EOA private key
//   - value - eth value to be sent
//   - receiverAddress - address of the tx receiver
func SendNormalTx(
	client *ethclient.Client,
	privateKey string,
	value *big.Int,
	receiverAddress string,
) (*types.Transaction, error) {

	return sendNormalTxInternal(client, privateKey, value, receiverAddress)
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
