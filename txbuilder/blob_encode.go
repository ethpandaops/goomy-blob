package txbuilder

import (
	"crypto/sha256"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
)

type BlobCommitment struct {
	Blob          kzg4844.Blob
	Commitment    kzg4844.Commitment
	Proof         kzg4844.Proof
	VersionedHash common.Hash
}

func EncodeBlob(data []byte) (*BlobCommitment, error) {
	dataLen := len(data)
	if dataLen > params.BlobTxFieldElementsPerBlob*params.BlobTxBytesPerFieldElement {
		return nil, fmt.Errorf("blob data longer than allowed (length: %v, limit: %v)", dataLen, params.BlobTxFieldElementsPerBlob*params.BlobTxBytesPerFieldElement)
	}
	blobCommitment := BlobCommitment{}
	copy(blobCommitment.Blob[:], data)
	var err error

	// generate blob commitment
	blobCommitment.Commitment, err = kzg4844.BlobToCommitment(blobCommitment.Blob)
	if err != nil {
		return nil, fmt.Errorf("failed generating blob commitment: %w", err)
	}

	// generate blob proof
	blobCommitment.Proof, err = kzg4844.ComputeBlobProof(blobCommitment.Blob, blobCommitment.Commitment)
	if err != nil {
		return nil, fmt.Errorf("failed generating blob proof: %w", err)
	}

	// build versioned hash
	blobCommitment.VersionedHash = sha256.Sum256(blobCommitment.Commitment[:])
	blobCommitment.VersionedHash[0] = params.BlobTxHashVersion
	return &blobCommitment, nil
}
