package main

import (
	"fmt"

	"github.com/TinajXD/way"
)

func main() {
	genData := "Egor Solomahin"
	data1 := "rehiopjbhres"
	data2 := "kjvdslzj"

	genesisBlock, err := way.Block.InitBlock(way.Block{}, []byte(genData))
	if err != nil {
		panic(err)
	} else {
		out := fmt.Sprintf("PrevHash: %x\nData: %s\nHash: %x\n--------------------------------\n", 
		genesisBlock.PrevHash, string(genesisBlock.Data), genesisBlock.Hash)
		fmt.Println(out)
	}

	
	block := way.Block.NewBlock(way.Block{}, []byte(data1), genesisBlock)
	out := fmt.Sprintf("PrevHash: %x\nData: %s\nHash: %x\n--------------------------------\n", 
	block.PrevHash, string(block.Data), block.Hash)
	fmt.Println(out)

	block2 := way.Block.NewBlock(way.Block{}, []byte(data2), block)
	out = fmt.Sprintf("PrevHash: %x\nData: %s\nHash: %x\n--------------------------------\n", 
	block2.PrevHash, string(block2.Data), block2.Hash)
	fmt.Println(out)
}