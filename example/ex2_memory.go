package main

import (
	"fmt"
	"log"
	"time"

	"github.com/TinajXD/way"
)

func main() {
	mem_chain := way.Chain{}
	_ = mem_chain.InitChain([]byte("Hi!"), time.Now())

	//fmt.Println(mem_chain)

	for i := 0; i <= 99; i++ {
		_ = mem_chain.NewBlockInChain([]byte(somelet(i)), time.Now().UTC())
	}

	fmt.Println(mem_chain.GetLastBlock().ID)

	for i := 0; i <= mem_chain.GetLastBlock().ID; i++ {
		curBlock := mem_chain.GetBlockByID(i)
		log.Printf("ID: %d\n Time: %s\n PrevHash: %x\n Hash %x\n Data %s\n-------------------------------------------------------",
			curBlock.ID, curBlock.Time_UTC.String(), curBlock.PrevHash, curBlock.Hash, string(curBlock.Data))
	}
}

func somelet(i int) string {
	letters := "abcdefghijklmnopqrstvwxyz"
	x := len(letters)
	return string(letters[i%x])
}
