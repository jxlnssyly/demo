package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

// 定义一个工作量证明的结构ProofOfWork

type ProofOfWork struct {
	// a.block
	block *Block

	// b.目标值
	// 一个非常大的数，它有很丰富的方法：比较、赋值方法
	target *big.Int
}

// 提供创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	// 指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	// 引入的辅助变量，目的是将上面的难度值转成big.Int
	tmpInt := big.Int{}
	// 将难度值赋值给big.Ing,指定16进制
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt

	return &pow
}

// 提供计算，不断计算hash的值
func (pow *ProofOfWork) Run() ([]byte, uint64) {



	var nonce uint64
	block := pow.block
	var hash [32]byte
	for {
		// 1.拼装数据（区块的数据、不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		// 2. 做hash运算
		blockInfo := bytes.Join(tmp,[]byte{})
		hash = sha256.Sum256(blockInfo)
		// 3.与pow中的target进行比较
		tmpInt := big.Int{}
		// 将hash数组转换成一个big.int
		tmpInt.SetBytes(hash[:])
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		//func (x *Int) Cmp(y *Int) (r int) {
		// 比较当前的哈希与目标哈希，如果当前的哈希值小于目标哈希值，说明找到了，否则继续找

		if tmpInt.Cmp(pow.target ) == -1 {
			// a.找打了，退出返回
			fmt.Printf("挖矿成功！hash %x, nonce: %d\n", hash, nonce)
			return hash[:], nonce
		} else  {
			// b.没招到，继续找，随机数+1
			nonce++
		}
	}

}
