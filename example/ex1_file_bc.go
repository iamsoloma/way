package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TinajXD/way"
)

func main() {
	inp := 0
	genesis := ""
	fmt.Print("Genesis block`s info: ")
    fmt.Scanln(&genesis)
	fmt.Print("The desired number of blocks: ")
    fmt.Scanln(&inp)

	filename := "./example/ex1.bc"

	ExpCfg := way.Explorer{Path: filename}

	err := way.Explorer.CreateBlockChain(ExpCfg, genesis, time.Now().UTC())
	if err != nil {
		log.Println(err)
	}

	genBlock, err := way.Explorer.GetBlockByID(ExpCfg, 0)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Genesis:\n ID: %d\n Time: %s\n PrevHash: %q\n Hash: %q\n Data: %q\n", genBlock.ID, genBlock.Time_UTC.String(), genBlock.PrevHash, genBlock.Hash, genBlock.Data)
	}

	for i := 1; i <= inp; i++ {
		lastblock, _ := way.Explorer.GetLastBlock(ExpCfg)
		curblock := way.Block.NewBlock(way.Block{}, []byte(somelet(i)), lastblock, time.Now().UTC())
		_, err = way.Explorer.AddBlock(ExpCfg, curblock)
		if err != nil {
			log.Println(err)
		}
	}
	for i := 0; i <= inp; i++ {
		curblock, err := way.Explorer.GetBlockByID(ExpCfg, i)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Block:\n ID: %d\n Time: %s\n PrevHash: %x\n Hash: %x\n Data: %q\n", curblock.ID, curblock.Time_UTC.String(), curblock.PrevHash, curblock.Hash, curblock.Data)
	}


	lastBlock, err := way.Explorer.GetLastBlock(ExpCfg)
	if err != nil {
		log.Println(err)
	}
	log.Println("Last ID in blockchain is " + fmt.Sprint(lastBlock.ID) + ".")
}

func somelet(i int) string {
	letters := []byte("abcdefghijklmnopqrstvwxyzABCDEFGHIGKLMNOPQRSTVWXYZ1234567890!@#$%^&*()_-+=")
	out := []byte{}
	x := len(letters)
	for y := 1000; y > 0; y-- {
		out = append(out, letters[rand.Intn(x)])
	}
	return string(out)
}
