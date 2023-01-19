package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nounce       int
	prevHash     [32]byte
	timestamp    int64
	transactions []string
}

func main() {
	blockChain := NewBlockChain()
	blockChain.Print()

	prevHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, prevHash)
	blockChain.Print()

	prevHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, prevHash)
	blockChain.Print()
}

func NewBlock(nounce int, prevHash [32]byte) *Block {
	return &Block{
		nounce:    nounce,
		prevHash:  prevHash,
		timestamp: time.Now().UnixNano(),
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp      	 %d\n", b.timestamp)
	fmt.Printf("nounce      	 %d\n", b.nounce)
	fmt.Printf("prevHash     	 %x\n", b.prevHash)
	fmt.Printf("transactions     %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nounce       int      `json:"nounce"`
		PrevHash     [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"x`
	}{
		Timestamp:    b.timestamp,
		Nounce:       b.nounce,
		PrevHash:     b.prevHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nounce int, prevHash [32]byte) *Block {
	b := NewBlock(nounce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}
