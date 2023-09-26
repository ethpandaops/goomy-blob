package txbuilder

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/holiman/uint256"
)

type TxMetadata struct {
	GasTipCap  *uint256.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *uint256.Int // a.k.a. maxFeePerGas
	BlobFeeCap *uint256.Int // a.k.a. maxFeePerBlobGas
	Gas        uint64
	To         common.Address
	Value      *uint256.Int
	Data       []byte
	AccessList types.AccessList
}

func BuildBlobTx(txData *TxMetadata, blobRefs []string) (*types.BlobTx, error) {
	tx := types.BlobTx{
		GasTipCap:  txData.GasTipCap,
		GasFeeCap:  txData.GasFeeCap,
		BlobFeeCap: txData.BlobFeeCap,
		Gas:        txData.Gas,
		To:         txData.To,
		Value:      txData.Value,
		Data:       txData.Data,
		AccessList: txData.AccessList,
		BlobHashes: make([]common.Hash, 0),
		Sidecar: &types.BlobTxSidecar{
			Blobs:       make([]kzg4844.Blob, 0),
			Commitments: make([]kzg4844.Commitment, 0),
			Proofs:      make([]kzg4844.Proof, 0),
		},
	}

	for _, blobRef := range blobRefs {
		err := parseBlobRef(&tx, blobRef)
		if err != nil {
			return nil, err
		}
	}

	return &tx, nil
}

func parseBlobRef(tx *types.BlobTx, blobRef string) error {
	var err error
	var blobBytes []byte

	if strings.HasPrefix(blobRef, "0x") {
		blobBytes = common.FromHex(blobRef)
	} else if refParts := strings.Split(blobRef, ":"); len(refParts) >= 2 {
		switch refParts[0] {
		case "file":
			blobBytes, err = os.ReadFile(strings.Join(refParts[1:], ":"))
			if err != nil {
				return err
			}
		case "url":
			blobBytes, err = loadUrlRef(strings.Join(refParts[1:], ":"))
			if err != nil {
				return err
			}
		case "repeat":
			if len(refParts) != 3 {
				return fmt.Errorf("invalid repeat ref format: %v", blobRef)
			}
			repeatCount, err := strconv.Atoi(refParts[2])
			if err != nil {
				return fmt.Errorf("invalid repeat count: %v", refParts[2])
			}
			repeatBytes := common.FromHex(refParts[1])
			repeatBytesLen := len(repeatBytes)
			blobBytes = make([]byte, repeatCount*repeatBytesLen)
			for i := 0; i < repeatCount; i++ {
				copy(blobBytes[(i*repeatBytesLen):], repeatBytes)
			}
		}
	}

	if blobBytes == nil {
		return fmt.Errorf("unknown blob ref: %v", blobRef)
	}

	blobCommitment, err := EncodeBlob(blobBytes)
	if err != nil {
		return fmt.Errorf("invalid blob: %w", err)
	}

	tx.BlobHashes = append(tx.BlobHashes, blobCommitment.VersionedHash)
	tx.Sidecar.Blobs = append(tx.Sidecar.Blobs, blobCommitment.Blob)
	tx.Sidecar.Commitments = append(tx.Sidecar.Commitments, blobCommitment.Commitment)
	tx.Sidecar.Proofs = append(tx.Sidecar.Proofs, blobCommitment.Proof)
	return nil
}

func loadUrlRef(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("received http error: %v", response.Status)
	}
	return io.ReadAll(response.Body)
}
