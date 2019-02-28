package main

import "math/big"

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
		block:block,
	}
	// 指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "00001000000000000000000000000000000000"
	// 引入的辅助变量，目的是将上面的难度值转成big.Int
	tmpInt := big.Int{}
	// 将难度值赋值给big.Ing,指定16进制
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt

	return &pow
}

// 提供计算，不断计算hash的值
func (pow *ProofOfWork) Run() ([]byte, uint64)  {

	// TODO
	return []byte("HelloWorld"), 10
}



