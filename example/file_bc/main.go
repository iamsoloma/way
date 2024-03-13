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
	inp := 0
	genesis := ""
	lenght := 0
	fmt.Print("Genesis block`s info data: ")
    fmt.Scanln(&genesis)
	fmt.Print("The desired number of blocks(random data): ")
    fmt.Scanln(&inp)
	fmt.Print("The desired lenght of random data: ")
    fmt.Scanln(&lenght)

//<<<<<<< main
	filename := "./blockchains/ex1.bc"

//=======
//	filename := "./ex1.bc"
//>>>>>>> DEV
	ExpCfg := way.Explorer{Path: filename}

	if _, err := os.Stat(ExpCfg.Path); errors.Is(err, os.ErrNotExist) {
		ExpCfg.File, err = os.Create(ExpCfg.Path)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println(errors.New("BlockChain is Exist! File: " + ExpCfg.Path))
		ExpCfg.File, err = os.Open(ExpCfg.Path)
		if err != nil {
			log.Println(err)
		}
	}

	defer ExpCfg.File.Close()



	err := ExpCfg.CreateBlockChain(genesis, time.Now().UTC())
	if err != nil {
		log.Println(err)
	}

	genBlock, err := ExpCfg.GetBlockByID(0)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Genesis:\n ID: %d\n Time: %s\n PrevHash: %q\n Hash: %q\n Data: %q\n", genBlock.ID, genBlock.Time_UTC.String(), genBlock.PrevHash, genBlock.Hash, genBlock.Data)
	}

	for i := 1; i <= inp; i++ {
		lastblock, _ := ExpCfg.GetLastBlock()
		curblock := way.Block.NewBlock(way.Block{}, []byte(somestr(lenght)), lastblock, time.Now().UTC())
		_, err = ExpCfg.AddBlock(curblock)
		if err != nil {
			log.Println(err)
		}
	}
	for i := 0; i <= inp; i++ {
		curblock, err := ExpCfg.GetBlockByID(i)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Block:\n ID: %d\n Time: %s\n PrevHash: %x\n Hash: %x\n Data: %q\n", curblock.ID, curblock.Time_UTC.String(), curblock.PrevHash, curblock.Hash, curblock.Data)
	}


	lastBlock, err := ExpCfg.GetLastBlock()
	if err != nil {
		log.Println(err)
	}
	log.Println("Last ID in blockchain is " + fmt.Sprint(lastBlock.ID) + ".")
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
