package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

//var DifficultyLevel int32

const DifficultyLevel = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DifficultyLevel))

	pofw := &ProofOfWork{
		Block:  block,
		Target: target,
	}
	return pofw
}
func (p *ProofOfWork) ComputeData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(p.Block.PrevHash),
			[]byte(p.Block.Data),
			[]byte(strconv.Itoa(nonce)),
			[]byte(strconv.Itoa(DifficultyLevel)),
		},
		[]byte{},
	)
	return data
}

func (p *ProofOfWork) MineBlock() (int, string) {
	nonce := 0
	var hash [32]byte
	var hashInt big.Int
	for {
		data := p.ComputeData(nonce)

		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		fmt.Printf("\rMining: %x", hash)

		if hashInt.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	return nonce, hex.EncodeToString(hash[:])
}

func (p *ProofOfWork) Validate() bool {
	data := p.ComputeData(p.Block.Nonce)

	hash := sha256.Sum256(data)

	var hashInt big.Int

	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(p.Target) == -1
}
