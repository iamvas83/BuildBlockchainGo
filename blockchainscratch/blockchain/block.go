package blockchain

import (
	"math/rand"
	"time"
)

type Block struct {
	Hash         string
	Data         string
	PrevHash     string
	Nonce        int
	Transactions []*Transaction
}

// func (b *Block) ComputeHash() {
// 	// info := b.PrevHash + b.Data
// 	// hash := sha256.Sum256([]byte(info))
// 	// b.Hash=hex.EncodeToString(hash[:])
// 	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
// 	computedHash := md5.Sum(concatenatedData)
// 	b.Hash = string(computedHash[:])
// }

func CreateBlock(data string, prevhash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano())
	initialNonce := rand.Intn(10000)
	block := &Block{
		Hash:         "",
		Data:         data,
		PrevHash:     prevhash,
		Nonce:        initialNonce,
		Transactions: transactions,
	}
	pofw := NewProofOfWork(block)

	nonce, hash := pofw.MineBlock()

	block.Nonce = nonce
	block.Hash = string(hash[:])
	return block
}

func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	}

	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
