package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/TinajXD/way"
)

func main() {
	NumOfBlocks := 100
	genesis := "FileToChain"
	lenght := 10

	filename := "./ex3.bc"
	ExpCfg := way.Explorer{Path: filename}

	var file *os.File
	if _, err := os.Stat(ExpCfg.Path); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(ExpCfg.Path)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println(errors.New("BlockChain is Exist! File: " + ExpCfg.Path))
	}

	defer file.Close()

	err := way.Explorer.CreateBlockChain(ExpCfg, genesis, time.Now().UTC())
	if err != nil {
		log.Println(err)
	}

	genBlock, err := way.Explorer.GetBlockByID(ExpCfg, 0)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("Genesis:\n ID: %d\n Time: %s\n PrevHash: %q\n Hash: %q\n Data: %q\n", genBlock.ID, genBlock.Time_UTC.String(), genBlock.PrevHash, genBlock.Hash, genBlock.Data)
	}

	for i := 1; i <= NumOfBlocks; i++ {
		lastblock, _ := way.Explorer.GetLastBlock(ExpCfg)
		curblock := way.Block.NewBlock(way.Block{}, []byte(somestr(lenght)), lastblock, time.Now().UTC())
		_, err = way.Explorer.AddBlock(ExpCfg, curblock)
		if err != nil {
			log.Println(err)
		}
	}

	way.Translate.FileToChain(way.Translate{}, &ExpCfg)

	for i := 0; i <= ExpCfg.Chain.GetLastBlock().ID; i++ {
		curBlock := ExpCfg.Chain.GetBlockByID(i)
		fmt.Printf("ID: %d\nTime: %s\nPrevHash: %x\nHash %x\nData %s\n-------------------------------------------------------\n",
			curBlock.ID, curBlock.Time_UTC.String(), curBlock.PrevHash, curBlock.Hash, string(curBlock.Data))
	}
}

//random
func somestr(lenght int) string {
	letters := []byte("abcdefghijklmnopqrstvwxyzABCDEFGHIGKLMNOPQRSTVWXYZ1234567890!@#$%^&*()_-+=")
	out := []byte{}
	x := len(letters)
	for y := lenght; y > 0; y-- {
		out = append(out, letters[rand.Intn(x)])
	}
	return string(out)
}
