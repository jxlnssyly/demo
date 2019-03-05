package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	//blocks []* Block
	db   *bolt.DB
	tail []byte // 存储最后一个区块的哈希
}

func (bc *BlockChain) AddBlock(data string) {
	/*
	lastBlock := bc.blocks[len(bc.blocks)-1]

	block := NewBlock(data, lastBlock.Hash)

	bc.blocks = append(bc.blocks, block)
	*/

	db := bc.db // 数据库
	lastHash := bc.tail // 最后一个区块的哈希


	db.Update(func(tx *bolt.Tx) error {

		// 完成数据添加
		bucket := tx.Bucket([]byte(blockBucket))

		if bucket ==nil {
			 log.Panic("bucket 不应该为空，请检查!")
		}

		block := NewBlock(data, lastHash)

		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte(LastHashKey),block.Hash)
		bc.tail = block.Hash

		return nil
	})

}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"
const LastHashKey = "LastHashKey"

// 5.定义一个区块链
func NewBlockChain() *BlockChain {



	var lastHash []byte

	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败")
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}

			// 创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock()

			// 3.写数据
			// hash作为key，block的字节流作为value
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte(LastHashKey), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte(LastHashKey))

		}

		return nil
	})

	return &BlockChain{
		db,
		lastHash,
	}

}

// 创世快
func GenesisBlock() *Block {
	return NewBlock("Go创世块，老牛逼了", []byte{})
}
