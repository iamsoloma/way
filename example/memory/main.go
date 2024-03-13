package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TinajXD/way"
)

func main() {
	mem_chain := way.Chain{}
	_ = mem_chain.InitChain([]byte("Hi!"), time.Now())

	//fmt.Println(mem_chain)

	for i := 0; i <= 99; i++ {
		_ = mem_chain.NewBlockInChain([]byte(somestr(100)), time.Now().UTC())
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