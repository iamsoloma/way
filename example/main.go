package main

import (
	"fmt"
	"log"
	"time"

	"github.com/TinajXD/way"
)

func main() {
	filename := "./example/blockchain.bc"

	ExpCfg := way.Explorer{Path: filename}

	err := way.Explorer.CreateBlockChain(ExpCfg, "My Way!", time.Now().UTC())
	if err != nil {
		log.Println(err)
	}

	genBlock, err := way.Explorer.GetBlockByID(ExpCfg, 0)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Genesis:\n ID: %d\n Time: %s\n PrevHash: %q\n Hash: %q\n Data: %q\n", genBlock.ID, genBlock.Time_UTC.String(), genBlock.PrevHash, genBlock.Hash, genBlock.Data)
	}

	wowBlock := way.Block.NewBlock(way.Block{}, []byte("WOW! It`s work!"), genBlock, time.Now().UTC())
	

	_, err = way.Explorer.AddBlock(ExpCfg, wowBlock)
	if err != nil {
		log.Println(err)
	}
	wowBlock, err = way.Explorer.GetBlockByID(ExpCfg, 1)
	if err != nil {
		log.Println(err)
	}
	log.Printf("WOW:\n ID: %d\n Time: %s\n PrevHash: %q\n Hash: %q\n Data: %q\n", wowBlock.ID, wowBlock.Time_UTC.String(), wowBlock.PrevHash, wowBlock.Hash, wowBlock.Data)

	lastBlock, err := way.Explorer.GetLastBlock(ExpCfg)
	if err != nil {
		log.Println(err)
	}
	log.Println("Last ID in blockchain is " + fmt.Sprint(lastBlock.ID) + ".")
}
