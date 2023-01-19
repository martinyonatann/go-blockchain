package main

import "github.com/martinyonatann/go-blockchain/internal/blockchain"

func main() {
	blockChain := blockchain.NewBlockChain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	prevHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, prevHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("D", "E", 3.0)
	prevHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, prevHash)
	blockChain.Print()
}
