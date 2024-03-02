package way

import (
	"crypto/sha256"
	"strconv"
)

type Block struct {
	PrevHash []byte
	Data []byte
	Hash []byte
}

func (b Block) InitBlock (genesis []byte) (Block, error) {
	hasher := sha256.New()
	genesisHash, err := hasher.Write(genesis)
	if err != nil {
		return Block{}, err
	}

	b.Hash = []byte(strconv.Itoa(genesisHash))
	b.Data = genesis
	b.PrevHash = []byte{}

	return b, err
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