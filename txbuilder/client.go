package txbuilder

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client *ethclient.Client
}

func NewClient(rpchost string) (*Client, error) {
	client, err := ethclient.Dial(rpchost)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: client,
	}, nil
}

func (client *Client) UpdateWallet(wallet *Wallet) error {
	chainId, err := client.GetChainId()
	if err != nil {
		return err
	}
	wallet.SetChainId(chainId)

	nonce, err := client.GetPendingNonceAt(wallet.GetAddress())
	if err != nil {
		return err
	}
	wallet.SetNonce(nonce)

	balance, err := client.GetPendingBalanceAt(wallet.GetAddress())
	if err != nil {
		return err
	}
	wallet.SetBalance(balance)

	return nil
}

func (client *Client) GetChainId() (*big.Int, error) {
	return client.client.ChainID(context.Background())
}

func (client *Client) GetPendingNonceAt(wallet common.Address) (uint64, error) {
	return client.client.PendingNonceAt(context.Background(), wallet)
}

func (client *Client) GetPendingBalanceAt(wallet common.Address) (*big.Int, error) {
	return client.client.PendingBalanceAt(context.Background(), wallet)
}

func (client *Client) SubmitTransaction(txBytes []byte) *common.Hash {
	tx := new(types.Transaction)
	err := tx.UnmarshalBinary(txBytes)
	if err != nil {
		fmt.Printf("Error while decoding transaction: %v (%v)\n", err, len(txBytes))
		return nil
	}

	err = client.client.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Printf("Error while sending transaction: %v\n", err)
		return nil
	}

	txHash := tx.Hash()
	fmt.Printf("    submitted transaction %v\n", txHash.String())

	return &txHash
}

func (client *Client) GetTransactionReceipt(txHash []byte) *types.Receipt {
	hash := common.Hash{}
	hash.SetBytes(txHash)
	receipt, err := client.client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil
	}
	return receipt
}
