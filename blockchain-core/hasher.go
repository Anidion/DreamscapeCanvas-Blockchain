package core

import (
	"crypto/sha256"

	types "github.com/AzlanAmjad/DreamscapeCanvas-Blockchain/data-types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

// BlockHeaderHasher is a hasher for the block header. Implementation of the Hasher interface.
type BlockHeaderHasher struct{}

func NewBlockHeaderHasher() *BlockHeaderHasher {
	return &BlockHeaderHasher{}
}

func (h *BlockHeaderHasher) Hash(bh *BlockHeader) types.Hash {
	blockHeaderHash := sha256.Sum256(bh.GetBytes())
	return types.Hash(blockHeaderHash[:])
}
