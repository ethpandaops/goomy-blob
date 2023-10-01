package txbuilder

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func DynFeeTx(txData *TxMetadata) (*types.DynamicFeeTx, error) {
	tx := types.DynamicFeeTx{
		GasTipCap:  txData.GasTipCap.ToBig(),
		GasFeeCap:  txData.GasFeeCap.ToBig(),
		Gas:        txData.Gas,
		To:         txData.To,
		Value:      txData.Value.ToBig(),
		Data:       txData.Data,
		AccessList: txData.AccessList,
	}
	return &tx, nil
}
