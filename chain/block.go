package chain

import (
	"time"
)

type Block struct {
	Number     uint64
	Time       time.Time
	Hash       Hash
	ParentHash Hash
	Data       string
}

func GenesisBlock() Block {
	block := Block{
		Number:     0,
		Time:       time.Now(),
		ParentHash: Hash{},
		Data:       "Hey there, I am the Genesis Block for GopherChain",
	}
	block.Hash = NewHash(block)

	return block
}

func NewBlock(parentBlock Block, data string) Block {
	block := Block{
		Number:     parentBlock.Number + 1,
		Time:       time.Now(),
		ParentHash: parentBlock.Hash,
		Data:       data,
	}
	block.Hash = NewHash(block)

	return block
}
