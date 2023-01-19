package blockchain

import (
	"fmt"
	"strings"

	"github.com/martinyonatann/go-blockchain/internal/block"
	"github.com/martinyonatann/go-blockchain/internal/transaction"
)

type Blockchain struct {
	transactionPool []*transaction.Transaction
	chain           []*block.Block
}

func NewBlockChain() *Blockchain {
	b := &block.Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nounce int, prevHash [32]byte) *block.Block {
	b := block.NewBlock(nounce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*transaction.Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := transaction.NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}
