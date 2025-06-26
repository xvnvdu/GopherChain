package chain

import (
	"fmt"
	"time"
)

// Define the block structure
type Block struct {
	Number     uint64
	Time       time.Time
	Hash       Hash
	ParentHash Hash
	Data       string
}

// Create and return first block in the chain without parent
func GenesisBlock() *Block {
	block := &Block{
		Number:     0,
		Time:       time.Now(),
		ParentHash: Hash{},
		Data:       "Hey there, I am the Genesis Block for GopherChain",
	}
	block.Hash = NewHash(*block)

	return block
}

// Create and return all other blocks in the chain
func NewBlock(parentBlock Block, data string) *Block {
	block := &Block{
		Number:     parentBlock.Number + 1,
		Time:       time.Now(),
		ParentHash: parentBlock.Hash,
		Data:       data,
	}
	block.Hash = NewHash(*block)

	return block
}

// Blockchain is a database that keeps all its blocks
type Blockchain struct {
	Blocks []*Block
}

// Save provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(*lastBlock, data)

	if !VerifyBlock(newBlock, lastBlock) {
		panic(fmt.Sprintf("Failed attempt to AddBlock. "+
			"Block under the number %d is not valid.", newBlock.Number))
	}

	bc.Blocks = append(bc.Blocks, newBlock)
}

// Create a new blockchain, starting with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

// Confirm whether a block is authentic
func VerifyBlock(newBlock, parentBlock *Block) bool {
	if newBlock.Number != parentBlock.Number+1 {
		return false
	}
	if newBlock.ParentHash != parentBlock.Hash {
		return false
	}

	expectedHash := NewHash(struct {
		Number     uint64
		Time       time.Time
		Hash       Hash
		ParentHash Hash
		Data       string
	}{
		Number:     newBlock.Number,
		Time:       newBlock.Time,
		Hash:       Hash{},
		ParentHash: newBlock.ParentHash,
		Data:       newBlock.Data,
	})
	if newBlock.Hash != expectedHash {
		return false
	}

	return true
}
