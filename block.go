package way

import (
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	ID       int       `json:"id"`
	Time_UTC time.Time `json:"time_utc"`
	PrevHash []byte    `json:"prev_hash"`
	Hash     []byte    `json:"hash"`
	Data     []byte    `json:"data"`
}

func (b *Block) InitBlock (genesis []byte, time_utc time.Time) (error) {
	hasher := sha256.New()
	genesisHash, err := hasher.Write(genesis)
	if err != nil {
		return err
	}

	b.Hash = []byte(strconv.Itoa(genesisHash))
	b.Data = genesis
	b.PrevHash = []byte{'0'}
	b.ID = 0
	b.Time_UTC = time_utc

	return err
}

func (b *Block) NewBlock (data []byte, prevBlock Block, time_utc time.Time) {
	hasher := sha256.New()
	hasher.Sum(prevBlock.Hash)
	dataHash := hasher.Sum(data)

	b.Hash = dataHash
	b.Data = data
	b.PrevHash = prevBlock.Hash
	b.ID = prevBlock.ID + 1
	b.Time_UTC = time_utc
}