package main

import (
	"fmt"

	"github.com/fernandoporazzi/blockchain-golang/blockchain"
)

func main() {
	blockchain := blockchain.CreateBlockchain()

	blockchain.CreateGenesisBlock()

	blockchain.AddBlock("second block")
	blockchain.AddBlock("third block")
	blockchain.AddBlock("fourth block")

	for i := 0; i < len(blockchain.Blocks); i++ {
		fmt.Println(blockchain.Blocks[i])
		fmt.Println("")
		fmt.Println("")
	}
}
