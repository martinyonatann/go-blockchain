package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/martinyonatann/go-blockchain/internal/transaction"
)

type Block struct {
	nounce       int
	prevHash     [32]byte
	timestamp    int64
	transactions []*transaction.Transaction
}

func NewBlock(nounce int, prevHash [32]byte, transactions []*transaction.Transaction) *Block {
	return &Block{
		nounce:       nounce,
		prevHash:     prevHash,
		timestamp:    time.Now().UnixNano(),
		transactions: transactions,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp      	 %d\n", b.timestamp)
	fmt.Printf("nounce      	 %d\n", b.nounce)
	fmt.Printf("prevHash     	 %x\n", b.prevHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64                      `json:"timestamp"`
		Nounce       int                        `json:"nounce"`
		PrevHash     [32]byte                   `json:"previous_hash"`
		Transactions []*transaction.Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nounce:       b.nounce,
		PrevHash:     b.prevHash,
		Transactions: b.transactions,
	})
}
