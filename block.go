package main

import (
	"crypto/sha256"
	"time"
	"bytes"
	"encoding/binary"
	"log"
)

type Block struct {
	Version uint64 // 版本号
	MerkelRoot []byte // 梅克尔根
	TimeStamp uint64 // 时间戳
	Difficulty uint64 // 难度值
	Nonce uint64 // 随机数
	PrevHash []byte // 上一个区块的哈希
	// 正常比特币区块中没有当前区块的哈希，我们是为了方便做了简化
	Hash []byte // 当前区块的哈希
	Data []byte  // 数据
}

func (block *Block)SetHash()  {
	var blockInfo []byte
	/*
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...) // ...是把切片里的内容打散
	blockInfo = append(blockInfo, block.MerkelRoot...) // ...是把切片里的内容打散
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.PrevHash...) // ...是把切片里的内容打散
	blockInfo = append(blockInfo, block.Hash...) // ...是把切片里的内容打散
	blockInfo = append(blockInfo, block.Data...) // ...是把切片里的内容打散
	*/

	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.PrevHash,
		block.Hash,
		block.Data,
	}

	blockInfo = bytes.Join(tmp,[]byte{})
	hash := sha256.Sum256(blockInfo)

	block.Hash = hash[:]

}

// 实现一个辅助函数，功能是将uint64转成[]byte

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}


func NewBlock(data string, prevBlockHash []byte) *Block  {
	block := Block {
		Version: 00,
		PrevHash:prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty:0, // 随便填写，无效值
		Nonce:0,

		Hash: []byte{}, // 先填空，后面再计算
		Data: []byte(data),
	}
	//block.SetHash()

	// 根据挖矿结果对区块数据进行更新(补充)
	pow := NewProofOfWork(&block)
	block.Hash, block.Nonce = pow.Run()

	return &block
}
