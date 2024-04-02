package main

import (
	"fmt"
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
	
	mem_chain := way.Explorer{}.Chain
	_ = mem_chain.InitChain([]byte(genesis), time.Now())

	//fmt.Println(mem_chain)

	for i := 0; i <= inp; i++ {
		_ = mem_chain.NewBlockInChain([]byte(somestr(lenght)), time.Now().UTC())
	}

	fmt.Println(mem_chain.GetLastBlock().ID)

	for i := 0; i <= mem_chain.GetLastBlock().ID; i++ {
		curBlock := mem_chain.GetBlockByID(i)
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