package main

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain struct
type BlockChain struct {
	blocks []*Block
}

// Block struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// InitHash Block struct method, that creates a hash
func (block *Block) initHash() {
	info := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

// createBlock func that creates and returns new block
func createBlock(data string, prevHash []byte) *Block {
	block := &Block{
		[]byte{},
		[]byte(data),
		prevHash,
	}
	block.initHash()

	return block
}

// InitBlockChain func that creates genesis block,
// initializes and returns new instance of BlockChain struct
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{
		createBlock("17700", []byte{}), // Genesis block
	}}
}

// AddBlock is BlockChain struct method, that appends new block to chain
func (blockChain *BlockChain) AddBlock(data string) {
	prevBlock := blockChain.blocks[len(blockChain.blocks)-1]
	newBlock := createBlock(data, prevBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, newBlock)
}
