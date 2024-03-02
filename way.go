package way

import (
	"crypto/sha256"
)

type Block struct {
	PrevHash []byte
	Data []byte
	Hash []byte
}

func (b Block) InitBlock (data []byte) (int, error) {
	hasher := sha256.New()
	dataHash, err := hasher.Write(data)
	if err != nil {
		return 0, err
	}

	return dataHash, err
}

func (b Block) NewBlock (data []byte, prevBlock Block) (Block Block) {
	hasher := sha256.New()
	hasher.Sum(prevBlock.Hash)
	dataHash := hasher.Sum(data)

	b.Hash = dataHash
	b.Data = data
	b.PrevHash = prevBlock.Hash

	return b
}