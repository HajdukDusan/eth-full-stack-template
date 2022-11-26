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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sendNormalTxInternal(client *ethclient.Client, privateKey string, value *big.Int, receiverAddress string) (*types.Transaction, error) {

	privateKeyECDSA, nonce, err := getPendingNonceAndECDSAPrivKey(client, privateKey)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(receiverAddress),
		value,
		uint64(21000), // limit for standard txs
		gasPrice,
		nil,
	)

	err = signAndSendTx(client, tx, privateKeyECDSA)

	printSentTx(tx.Hash().Hex(), nonce, value, gasPrice, tx.Gas())

	return tx, nil
}

func signAndSendTx(client *ethclient.Client, tx *types.Transaction, privateKeyECDSA *ecdsa.PrivateKey) error {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		return err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	return nil
}

func getPendingNonceAndECDSAPrivKey(client *ethclient.Client, privateKey string) (*ecdsa.PrivateKey, uint64, error) {

	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, 0, err
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, 0, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	return privateKeyECDSA, nonce, nil
}

func sendContractTxInternal(client *ethclient.Client, privateKey string, value *big.Int, initializeTx func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {

	txOpts, err := createTransactionOptions(client, privateKey, value)
	if err != nil {
		return nil, err
	}

	tx, err := initializeTx(txOpts)
	if err != nil {
		return nil, tryToExtractErrorCode(err)
	}

	printSentTx(tx.Hash().Hex(), txOpts.Nonce.Uint64(), txOpts.Value, txOpts.GasPrice, tx.Gas())

	return tx, nil
}

func printSentTx(hash string, nonce uint64, value *big.Int, gasPrice *big.Int, gasLimit uint64) {
	fmt.Println("\nsent tx:")
	fmt.Println("\thash:", hash)
	fmt.Println("\tnonce:", nonce)
	fmt.Println("\tvalue:", value)
	fmt.Println("\tgasPrice:", gasPrice)
	fmt.Println("\tgasLimit:", gasLimit)
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

	privateKeyECDSA, nonce, err := getPendingNonceAndECDSAPrivKey(client, privateKey)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKeyECDSA)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value // in wei
	auth.GasLimit = 0  // 0 = estimate
	auth.GasPrice = gasPrice

	return auth, nil
}
