package way

import "time"


type Chain struct {
	blocks []Block
}

func (c Chain) InitChain (genesis []byte, time_utc time.Time) (Chain, error){
	InBl, err := Block.InitBlock(Block{}, genesis, time_utc)
	if err != nil {
		return Chain{}, err
	}
	c.blocks = append(c.blocks, InBl)

	return c, nil
}

func (c Chain) AddToChain (data []byte, time_utc time.Time) (Chain) {
	NewBlock := Block.NewBlock(Block{}, data, c.blocks[len(c.blocks) - 1], time_utc)
	c.blocks = append(c.blocks, NewBlock)

	return c
}

func (c Chain) GetLastID() (int) {
	lastID := len(c.blocks)
	return lastID
}