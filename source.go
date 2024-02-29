package blockchain

import (
    "fmt"
)

// Block represents a block in the blockchain
type Block struct {
    Data     string
    Previous *Block
}

// Blockchain represents a blockchain
type Blockchain struct {
    Head *Block
}

// DisplayAllBlocks prints all blocks in the blockchain
func (bc *Blockchain) DisplayAllBlocks() {
    current := bc.Head
    for current != nil {
        fmt.Printf("Data: %s\n", current.Data)
        current = current.Previous
    }
}

// NewBlock creates a new block in the blockchain
func (bc *Blockchain) NewBlock(data string) {
    newBlock := &Block{
        Data:     data,
        Previous: bc.Head,
    }
    bc.Head = newBlock
}

// ModifyBlock modifies the data of a specific block in the blockchain
func (bc *Blockchain) ModifyBlock(index int, newData string) {
    current := bc.Head
    for i := 0; i < index; i++ {
        if current == nil {
            fmt.Println("Error: Block index out of range")
            return
        }
        current = current.Previous
    }
    if current != nil {
        current.Data = newData
    }
}

