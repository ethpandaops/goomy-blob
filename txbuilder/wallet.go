package txbuilder

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

type Wallet struct {
	mutex   sync.Mutex
	privkey *ecdsa.PrivateKey
	address common.Address
	chainid *big.Int
	nonce   uint64
	balance *big.Int
}

func NewWallet(privkey string) (*Wallet, error) {
	wallet := &Wallet{}
	err := wallet.loadPrivateKey(privkey)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (wallet *Wallet) loadPrivateKey(privkey string) error {
	var privateKey *ecdsa.PrivateKey
	if privkey == "" {
		var err error
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			return err
		}
	} else {
		var err error
		privateKey, err = crypto.HexToECDSA(privkey)
		if err != nil {
			return err
		}
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	wallet.privkey = privateKey
	wallet.address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return nil
}

func (wallet *Wallet) GetAddress() common.Address {
	return wallet.address
}

func (wallet *Wallet) GetChainId() *big.Int {
	return wallet.chainid
}

func (wallet *Wallet) GetNonce() uint64 {
	return wallet.nonce
}

func (wallet *Wallet) GetBalance() *big.Int {
	return wallet.balance
}

func (wallet *Wallet) SetChainId(chainid *big.Int) {
	wallet.chainid = chainid
}

func (wallet *Wallet) SetNonce(nonce uint64) {
	wallet.nonce = nonce
}

func (wallet *Wallet) SetBalance(balance *big.Int) {
	wallet.balance = balance
}

func (wallet *Wallet) BuildDynamicFeeTx(txData *types.DynamicFeeTx) (*types.Transaction, error) {
	wallet.mutex.Lock()
	txData.ChainID = wallet.chainid
	txData.Nonce = wallet.nonce
	wallet.nonce++
	wallet.mutex.Unlock()
	return wallet.SignTx(txData)
}

func (wallet *Wallet) BuildBlobTx(txData *types.BlobTx) (*types.Transaction, error) {
	wallet.mutex.Lock()
	txData.ChainID = uint256.MustFromBig(wallet.chainid)
	txData.Nonce = wallet.nonce
	wallet.nonce++
	wallet.mutex.Unlock()
	return wallet.SignTx(txData)
}

func (wallet *Wallet) SignTx(txData types.TxData) (*types.Transaction, error) {
	tx := types.NewTx(txData)
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(wallet.chainid), wallet.privkey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
