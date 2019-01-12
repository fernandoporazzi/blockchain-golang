package block

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int
	PreviousHash string
	Data         string
	Timestamp    time.Time
	Difficulty   int
	Hash         string
	Nonce        int
}

func CreateBlock() Block {
	var block Block

	block.Timestamp = time.Now()
	block.Nonce = 0

	return block
}

func (b *Block) GenerateHash() {

	index := strconv.Itoa(b.Index)
	nonce := strconv.Itoa(b.Nonce)

	b.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(index+b.PreviousHash+b.Data+b.Timestamp.String()+nonce)))
}

func (b *Block) Mine() {
	prefix := getPrefix(b.Difficulty)

	for {
		b.GenerateHash()

		if strings.HasPrefix(b.Hash, prefix) {
			break
		} else {
			b.Nonce = b.Nonce + 1
			b.GenerateHash()
		}
	}
}

func getPrefix(length int) string {
	letterBytes := "0"
	b := make([]byte, length)

	for i := range b {
		b[i] = letterBytes[0]
	}

	return string(b)
}
