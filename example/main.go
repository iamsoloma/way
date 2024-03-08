package main

import (
	"log"

	"github.com/TinajXD/way"
)

func main() {
	filename := "./example/blockchain.bc"


	err := way.Explorer.CreateBlockChain(way.Explorer{Path: filename}, "My Way!")
	if err != nil {
		log.Println(err)
	}
}