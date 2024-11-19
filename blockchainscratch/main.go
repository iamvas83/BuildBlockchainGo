package main

import (
	"BUILDBLOCKCHAINGO/blockchainscratch/blockchain"
	"fmt"
)

func main() {
	fmt.Println("Blockchain from scratch")

	block := blockchain.InitBlockChain()

	w := &blockchain.Wallet{}
	AliceWallet, err := w.NewWallet()

	if err != nil {
		fmt.Printf("error creating wallet for Alice")
	}
	fmt.Println("Wallet created successfuly for Alice")
	BobWallet, err := w.NewWallet()
	if err != nil {
		fmt.Printf("error creating wallet for Bob")
	}
	fmt.Println("Wallet created successfuly for Bob")

	CharlieWallet, err := w.NewWallet()
	if err != nil {
		fmt.Printf("error creating wallet for Charlie")
	}
	fmt.Println("Wallet created successfuly for Charlie")

	//1
	t := &blockchain.Transaction{
		Sender:   string(AliceWallet.PublicKey),
		Receiver: string(BobWallet.PublicKey),
		Amount:   12.5,
	}
	fmt.Println("Alice to Bob transaction created successfuly")
	signature, err := AliceWallet.SignTransaction(t)
	if err != nil {
		fmt.Println("error signing the transaction")
	}

	err = w.VerifyTransaction(t, AliceWallet.PublicKey, signature)
	if err != nil {
		fmt.Printf("Signatiure validation failed %s", err)
	}
	fmt.Println("Transaction Verified Successfully")

	block.AddBlock("Block 1", "Alice", []*blockchain.Transaction{t})

	//2
	t1 := &blockchain.Transaction{
		Sender:   string(AliceWallet.PublicKey),
		Receiver: string(CharlieWallet.PublicKey),
		Amount:   21.0,
	}
	fmt.Println("Alice to Charlie transaction created successfuly")
	charlieSign, err := CharlieWallet.SignTransaction(t1)
	if err != nil {
		fmt.Println("error signing the transaction")
	}

	err = w.VerifyTransaction(t1, CharlieWallet.PublicKey, charlieSign)
	if err != nil {
		fmt.Printf("Signatiure validation failed %s", err)
	}
	fmt.Println("Transaction Verified Successfully")
	block.AddBlock("blockchain-hyperledger Block 2", "Alice", []*blockchain.Transaction{t1})
	// block.AddBlock("blockchain-hyperledger Block 2", "Bob", []*blockchain.Transaction{
	// 	{Sender: "Bob", Receiver: "Charlie", Amount: 2.5},
	// 	{Sender: "Bob", Receiver: "Alice", Amount: 3.5},
	// })
	// block.AddBlock("blockchain-hyperledger Block 3", "Charlie", []*blockchain.Transaction{
	// 	{Sender: "Charlie", Receiver: "Alice", Amount: 6.5},
	// 	{Sender: "Charlie", Receiver: "Bob", Amount: 21.5},
	// })

	for _, block := range block.Blocks {

		fmt.Printf("Previous hash %x\n", block.PrevHash)
		fmt.Printf("Data in Block %s\n", block.Data)
		fmt.Printf("Hash of Block %x\n", block.Hash)
		p := blockchain.NewProofOfWork(block)
		flag := p.Validate()
		fmt.Printf("PoW algo ran successfully %v\n", flag)

		fmt.Println()
		fmt.Println("Transactions..")

		for _, tx := range block.Transactions {
			fmt.Printf("Sender %s\n", tx.Sender)
			fmt.Printf("Receiver %s\n", tx.Receiver)
			fmt.Printf("Amount %f\n", tx.Amount)
			fmt.Printf("Coinbase %t\n", tx.Coinbase)
		}
	}

}
