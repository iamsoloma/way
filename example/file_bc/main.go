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
	lenght := 0
	fmt.Print("Genesis block`s info data: ")
    fmt.Scanln(&genesis)
	fmt.Print("The desired number of blocks(random data): ")
    fmt.Scanln(&inp)
	fmt.Print("The desired lenght of random data: ")
    fmt.Scanln(&lenght)

	path := "./blockchains"
	name := "ex1"
	partition := 5

	ExpCfg := way.Explorer{Path: path, Name: name, Partition: partition}

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

	startWrite := time.Now()
	for i := 1; i <= inp; i++ {
		_, err = ExpCfg.AddBlock([]byte(somestr(lenght)), time.Now().UTC())
		if err != nil {
			log.Println(err)
		}
	}
	endWrite := time.Since(startWrite)


	lastBlock, err := way.Explorer.GetLastBlock(ExpCfg)
	if err != nil {
		log.Println(err)
	}

	readed := []string{}
	startRead := time.Now()
	for i := 0; i <= lastBlock.ID; i++ {
		curblock, err := way.Explorer.GetBlockByID(ExpCfg, i)
		if err != nil {
			log.Println(err)
		}
		curstring := fmt.Sprintf("Block:\n ID: %d\n Time: %s\n PrevHash: %x\n Hash: %x\n Data: %q\n", curblock.ID, curblock.Time_UTC.String(), curblock.PrevHash, curblock.Hash, curblock.Data)
		readed = append(readed, curstring)
	}
	endRead := time.Since(startRead)

	fmt.Println("-------------------------------------------------------------\nAll blocks:")
	for i := 0; i < len(readed); i++ {
		log.Print(readed[i])
	}
	fmt.Println("-------------------------------------------------------------")


	log.Println("Last ID in blockchain is " + fmt.Sprint(lastBlock.ID) + ".")
	log.Println("Recording time per block: ", endWrite / time.Duration(inp))
	log.Println("Reading time per block: ", endRead / time.Duration(inp))

	var del bool
	fmt.Print("Enter true or false if you want or do not want to delete the blockchain: ")
    fmt.Scanln(&del)

	if del {
		found, err := ExpCfg.DeleteBlockChain()
		if found && err == nil {
			log.Println("Blockchain is deleted.")
		} else {
			log.Printf("Blockchain is deleted: %t.\nError: %s.", found, err.Error())
		}
		
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