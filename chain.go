package way

import "time"


type Chain struct {
	blocks []Block
}

func (c *Chain) InitChain (genesis []byte, time_utc time.Time) (error){
	InBl, err := Block.InitBlock(Block{}, genesis, time_utc)
	if err != nil {
		return err
	}
	c.blocks = append(c.blocks, InBl)

	return nil
}

func (c *Chain) NewBlockInChain (data []byte, time_utc time.Time) (id int) {
	NewBlock := Block.NewBlock(Block{}, data, c.blocks[len(c.blocks) - 1], time_utc)
	c.blocks = append(c.blocks, NewBlock)

	return NewBlock.ID
}

func (c Chain) GetLastBlock() (Block) {
	lastID := len(c.blocks) - 1
	return c.blocks[lastID]
}

func (c Chain) GetBlockByID(id int) (Block) {
	return c.blocks[id]
}