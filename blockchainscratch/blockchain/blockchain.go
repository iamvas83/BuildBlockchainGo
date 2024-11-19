package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
	Coinbase bool
}

func InitBlockChain() *BlockChain {
	genesis := Genesis()
	bchain := &BlockChain{
		Blocks: []*Block{genesis},
	}

	return bchain
}

func (chain *BlockChain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	coinbaseTranx := &Transaction{
		Sender:   "Coinbase",
		Receiver: coinbaseRcpt,
		Amount:   10.0,
		Coinbase: true,
	}

	coinbaseTransaction := append([]*Transaction{coinbaseTranx}, transactions...)

	newBlock := CreateBlock(data, prevBlock.Hash, coinbaseTransaction)
	chain.Blocks = append(chain.Blocks, newBlock)
}
