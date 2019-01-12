package blockchain

import (
	"github.com/fernandoporazzi/blockchain-golang/block"
)

type Blockchain struct {
	Blocks     []block.Block
	Index      int
	Difficulty int
}

func CreateBlockchain() Blockchain {
	var b Blockchain

	b.Index = 0
	b.Difficulty = 3

	return b
}

func (b *Blockchain) CreateGenesisBlock() {
	block := block.CreateBlock()

	block.Index = 0
	block.PreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"
	block.Data = "Genesis Block"
	block.Difficulty = b.Difficulty

	block.Mine()

	b.Append(block)
}

func (b *Blockchain) AddBlock(data string) {
	block := block.CreateBlock()

	block.Index = b.Index
	block.PreviousHash = b.GetLastBlock().Hash
	block.Data = data
	block.Difficulty = b.Difficulty

	block.Mine()

	b.Append(block)
}

func (b *Blockchain) Append(block block.Block) {
	b.Index = b.Index + 1
	b.Blocks = append(b.Blocks, block)

}

func (b *Blockchain) GetLastBlock() block.Block {
	return b.Blocks[len(b.Blocks)-1]
}
