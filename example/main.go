package main

import (
	"fmt"

	"github.com/TinajXD/way"
)

func main() {
	genData := "Egor Solomahin"
	data1 := "rehiopjbhres"
	data2 := "kjvdslzj"

	genesis, err := way.Block.InitBlock(way.Block{}, []byte(genData))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Genesis: ", genesis, "\n--------------------------------")
	}

	genesisBlock := way.Block{PrevHash: []byte("0"), Data: []byte(genData), Hash: []byte("genesis")}
	
	block := way.Block.NewBlock(way.Block{}, []byte(data1), genesisBlock)
	fmt.Printf("PrevHash: %s\nData: %s\nHash %s\n--------------------------------\n", 
	string(block.PrevHash), string(block.Data), string(block.Hash))

	block2 := way.Block.NewBlock(way.Block{}, []byte(data2), block)
	fmt.Printf("PrevHash: %s\nData: %s\nHash %s\n--------------------------------\n", 
	string(block2.PrevHash), string(block2.Data), string(block2.Hash))
}