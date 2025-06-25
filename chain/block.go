package chain

import (
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
	blocks []*Block
}

// Save provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(*lastBlock, data)
	bc.blocks = append(bc.blocks, newBlock)
}

// Create a new blockchain, starting with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
