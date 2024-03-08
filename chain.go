package way


type Chain struct {
	blocks []Block
}

func (c Chain) InitChain (genesis []byte) (Chain, error){
	InBl, err := Block.InitBlock(Block{}, genesis)
	if err != nil {
		return Chain{}, err
	}
	c.blocks = append(c.blocks, InBl)

	return c, nil
}

func (c Chain) AddToChain (data []byte) (Chain) {
	NewBlock := Block.NewBlock(Block{}, data, c.blocks[len(c.blocks) - 1])
	c.blocks = append(c.blocks, NewBlock)

	return c
}