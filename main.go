package main

import (
	"fmt"
)

// 0.定义结构
//type Block struct {
//	// 1.前区块哈希
//	PrevHash []byte
//
//	// 2. 当前区块哈希
//	Hash []byte
//
//	// 3.数据
//	Data []byte
//
//}

// 2.创建区块
//func NewBlock(data string, prevBlockHash []byte) *Block  {
//	block := Block {
//		PrevHash:prevBlockHash,
//		Hash: []byte{}, // 先填空，后面再计算
//		Data: []byte(data),
//	}
//	block.SetHash()
//
//	return &block
//}

// 3. 生成哈希
//func (block *Block) SetHash() {
//	// 1. 拼装数据
//	blockInfo := append(block.PrevHash, block.Data...)
//	// 2.sha256
//	hash := sha256.Sum256(blockInfo)
//	block.Hash = hash[:]
//}

// 4.引入区块链
//type BlockChain struct {
//	// 1.定义一个区块链数据
//	blocks []*Block
//}

// 5.定义一个区块链
//func NewBlockChain() *BlockChain  {
//	// 创建一个创世块，并作为第一个区块添加到区块链中
//	genesisBlock := GenesisBlock()
//	return &BlockChain{
//		blocks: []*Block{genesisBlock},
//	}
//}

// 创世快
//func GenesisBlock() *Block {
//	return NewBlock("Go创世块，老牛逼了",[]byte{})
//}

// 6.添加区块
//func (bc *BlockChain)AddBlock(data string)  {
//
//	// 获取前区块的哈希
//	lastBlock := bc.blocks[len(bc.blocks) - 1]
//
//	// a.创建新的区块
//	block := NewBlock(data,lastBlock.Hash)
//
//	// b.添加到区块链数组中
//	bc.blocks = append(bc.blocks,block)
//}


func main() {
	//block := NewBlock("老师转班长一枚比特币!",[]byte{})

	bc := NewBlockChain()
	bc.AddBlock("班长向班花转了50枚比特币")
	bc.AddBlock("班长又向班花转了50枚比特币")
	for i, block := range bc.blocks {
		fmt.Printf("=======当前区块高度 %d\n", i + 1)
		fmt.Printf("前区块哈希值 %x\n",block.PrevHash)
		fmt.Printf("当前区块哈希值 %x\n",block.Hash)
		fmt.Printf("区块数据 %s\n",block.Data)
	}


}
