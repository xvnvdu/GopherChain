package main

import (
	"fmt"

	"github.com/xvnvdu/GopherChain/chain"
)

// Quick check for the way it all works
func main() {
	bc := chain.NewBlockchain()

	bc.AddBlock("Second block")
	bc.AddBlock("This is the 3rd block")

	for _, block := range bc.Blocks {
		fmt.Println(*block)
	}
}
