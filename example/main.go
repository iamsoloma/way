package main

import (
	"fmt"
	"time"

	"github.com/TinajXD/way"
)

func main() {
	genData := "Egor Solomahin"
	data1 := "rehiopjbhres"
	data2 := "kjvdslzj"


	exp, err := way.Explorer.OpenBlockChain(way.Explorer{}, "./example/test.bc")
	if err != nil {
		panic(err)
	}

	genesisBlock, err := way.Block.InitBlock(way.Block{}, []byte(genData))
	if err != nil {
		panic(err)
	} else {
		id, err := way.Explorer.SaveBlock(*exp, genesisBlock, time.Now().UTC())
		if err != nil {
			panic(err)
		}
		out := fmt.Sprintf("ID: %d\nPrevHash: %x\nData: %s\nHash: %x\n--------------------------------\n", 
		id, genesisBlock.PrevHash, string(genesisBlock.Data), genesisBlock.Hash)
		fmt.Println(out)
	}

	
	
	block := way.Block.NewBlock(way.Block{}, []byte(data1), genesisBlock)
	id, err := way.Explorer.SaveBlock(*exp, block, time.Now().UTC())
	if err != nil {
		panic(err)
	}
	out := fmt.Sprintf("ID: %d\nPrevHash: %x\nData: %s\nHash: %x\n--------------------------------\n", 
	id, block.PrevHash, string(block.Data), block.Hash)
	fmt.Println(out)

	block2 := way.Block.NewBlock(way.Block{}, []byte(data2), block)
	id, err = way.Explorer.SaveBlock(way.Explorer{}, block2, time.Now().UTC())
	if err != nil {
		panic(err)
	}
	out = fmt.Sprintf("ID: %d\nPrevHash: %x\nData: %s\nHash: %x\n--------------------------------\n", 
	id, block2.PrevHash, string(block2.Data), block2.Hash)
	fmt.Println(out)
}