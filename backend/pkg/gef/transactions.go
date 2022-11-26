package gef

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sendTxInternal(client *ethclient.Client, privateKey string, value *big.Int, initializeTx func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {

	txOpts, err := createTransactionOptions(client, privateKey, value)
	if err != nil {
		return nil, err
	}

	tx, err := initializeTx(txOpts)
	if err != nil {
		return nil, tryToExtractErrorCode(err)
	}

	fmt.Println("\nsent tx:")
	fmt.Println("\thash:", tx.Hash().Hex())
	fmt.Println("\tnonce:", txOpts.Nonce)
	fmt.Println("\tvalue:", txOpts.Value)
	fmt.Println("\tgasPrice:", txOpts.GasPrice)
	fmt.Println("\tgasLimit:", tx.Gas())

	return tx, nil
}

func waitTxReceiptInternal(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, tryToExtractErrorCode(err)
	}

	fmt.Println("\ntx receipt:")
	fmt.Println("\tstatus:", receipt.Status)
	fmt.Println("\ttxFee:", receipt.GasUsed*tx.GasPrice().Uint64())

	return receipt, nil
}

func tryToExtractErrorCode(err error) error {
	s_tmp := strings.Split(err.Error(), "'")
	if len(s_tmp) < 2 {
		return err
	}
	return errors.New(s_tmp[1])
}

func createTransactionOptions(client *ethclient.Client, privateKey string, value *big.Int) (*bind.TransactOpts, error) {

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKeyECDSA)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value // in wei
	auth.GasLimit = 0  // 0 = estimate
	auth.GasPrice = gasPrice

	return auth, nil
}
