package main

import "fmt"

// Block represents a block in a blockchain
type Block struct {
	Data string
}

// NewBlock creates a new block with the given data
func NewBlock(data string) *Block {
	return &Block{Data: data}
}

// ModifyBlock modifies the data of a block
func ModifyBlock(block *Block, newData string) {
	block.Data = newData
}

// DisplayAllBlocks displays the data of all blocks in the blockchain
func DisplayAllBlocks(blocks []*Block) {
	for i, block := range blocks {
		fmt.Printf("Block %d: %s\n", i+1, block.Data)
	}
}

func main() {
	// Create initial blocks
	block1 := NewBlock("First Block")
	block2 := NewBlock("Second Block")
	block3 := NewBlock("Third Block")

	// Display initial block data
	fmt.Println("Initial Block Data:")
	DisplayAllBlocks([]*Block{block1, block2, block3})

	// Modify the second block
	ModifyBlock(block2, "Modified Second Block")

	// Display modified block data
	fmt.Println("Modified Block Data:")
	DisplayAllBlocks([]*Block{block1, block2, block3})
}
