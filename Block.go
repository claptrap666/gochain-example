package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block : a hashblock
type Block struct {
	Index     int
	Data      string
	Hash      string
	PrevHash  string
	Timestamp time.Time
}

// make sure block is valid by checking index, and comparing the hash of the previous block
func (block *Block) isBlockValid(oldBlock Block) bool {
	if oldBlock.Index+1 != block.Index {
		return false
	}

	if oldBlock.Hash != block.PrevHash {
		return false
	}

	if block.hash() != block.Hash {
		return false
	}

	return true
}

// SHA256 hasing
func (block *Block) hash() string {
	record := strconv.Itoa(block.Index) + block.Timestamp.String() + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// create a new block using previous block's hash
func (block *Block) generate(Data string) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = block.Index + 1
	newBlock.Timestamp = t
	newBlock.Data = Data
	newBlock.PrevHash = block.Hash
	newBlock.Hash = newBlock.hash()
	return newBlock
}
