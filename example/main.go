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

	genBlock, id, time_utc, err := way.Explorer.GetBlockByID(ExpCfg, 0)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Genesis:\n ID: %d\n Time: %s\n PrevHash: %q\n Hash: %q\n Data: %q\n", id, time_utc.String(), genBlock.PrevHash, genBlock.Hash, genBlock.Data)
	}

	wowBlock := way.Block.NewBlock(way.Block{}, []byte("WOW! It`s work!"), genBlock)
	

	id, err = way.Explorer.AddBlock(ExpCfg, wowBlock, time.Now().UTC())
	if err != nil {
		log.Println(err)
	}
	wowBlock, id, time_utc, err = way.Explorer.GetBlockByID(ExpCfg, 1)
	if err != nil {
		log.Println(err)
	}
	log.Printf("WOW:\n ID: %d\n Time: %s\n PrevHash: %q\n Hash: %q\n Data: %q\n", id, time_utc.String(), wowBlock.PrevHash, wowBlock.Hash, wowBlock.Data)

	lastID, err := way.Explorer.GetLastID(ExpCfg)
	if err != nil {
		log.Println(err)
	}
	log.Println("Last ID in blockchain is " + fmt.Sprint(lastID) + ".")
}
