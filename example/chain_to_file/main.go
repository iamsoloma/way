package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/iamsoloma/way"
)

func main() {
	fmt.Println("From memory to file demo.")
	inp := 0
	genesis := ""
	lenght := 0
	partition := 1000
	fmt.Print("Genesis block`s info data: ")
	fmt.Scanln(&genesis)
	fmt.Print("The desired number of blocks(random data): ")
	fmt.Scanln(&inp)
	fmt.Print("The desired number of blocks in one part: ")
	fmt.Scanln(&partition)
	fmt.Print("The desired lenght of random data: ")
	fmt.Scanln(&lenght)

	path := "./blockchains"
	name := "ex4"

	ExpCfg := way.Explorer{Path: path, Name: name, Partition: partition}

	log.Println("Generating...")
	err := ExpCfg.Chain.InitChain([]byte(genesis), time.Now().UTC())
	if err != nil {
		panic(err)
	}

	for i := 1; i <= inp; i++ {
		ExpCfg.Chain.NewBlockInChain([]byte(somestr(lenght)), time.Now().UTC())
	}

	fmt.Println("Translating...")
	startTime := time.Now()
	way.Translate.ChainToFile(way.Translate{}, &ExpCfg)
	endTime := time.Since(startTime)

	fmt.Println("-------------------------------------------------------------\nAll blocks:")
	lastBlock, err := ExpCfg.GetLastBlock()
	if err != nil {
		panic(err)
	}
	for i := 0; i < lastBlock.ID; i++ {
		curblock, err := ExpCfg.GetBlockByID(i)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Block:\n ID: %d\n Time: %s\n PrevHash: %x\n Hash: %x\n Data: %q\n", curblock.ID, curblock.Time_UTC.String(), curblock.PrevHash, curblock.Hash, curblock.Data)
	}
	fmt.Println("-------------------------------------------------------------")

	log.Println("Translating time per block: ", endTime/time.Duration(inp))

}

// random
func somestr(lenght int) string {
	letters := []byte("abcdefghijklmnopqrstvwxyzABCDEFGHIGKLMNOPQRSTVWXYZ1234567890!@#$%^&*()_-+=")
	out := []byte{}
	x := len(letters)
	for y := lenght; y > 0; y-- {
		out = append(out, letters[rand.Intn(x)])
	}
	return string(out)
}
