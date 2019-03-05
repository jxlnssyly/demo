package main

import (
	"os"
	"fmt"
)

// 这是一个用来接收命令行参数并且控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage  = `
	addBlock --data DATA "add data to blockchain"
	printChain		"print all blockchain data"
`

// 接收参数的动作，我们放到一个函数中
func (cli *CLI)Run()  {
	args := os.Args // 得到所有的命令

	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}

	// 分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		// 添加区块
		//fmt.Println("添加区块")

		if len(args) == 4 && args[2] == "--data" {
			// a. 获取数据
			data := args[3]

			// b.使用bc添加区块AddBlock
			cli.AddBlock(data)
		} else {
			fmt.Println("添加区块参数使用不当，请检查")
			fmt.Println(Usage)
		}

	case "printChain":
		// 打印区块
		//fmt.Println("无效的命令，请检查")
		//
		//fmt.Println("打印区块")

		cli.PrintBlockChain()

	default:
		fmt.Println(Usage)

	}


	// 执行相应动作
}





