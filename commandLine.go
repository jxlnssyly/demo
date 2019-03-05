package main

import "fmt"

func (cli *CLI)PrintBlockChain()  {
	bc := cli.bc
	it := bc.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("===================================\n")
		fmt.Printf("前区块哈希值 %x\n",block.PrevHash)
		fmt.Printf("当前区块哈希值 %x\n",block.Hash)
		fmt.Printf("区块数据 %s\n",block.Data)
		fmt.Printf("===================================\n")

		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束")
			break
		}

	}
}

func (cli *CLI) AddBlock(data string)  {
	cli.bc.AddBlock(data)
	fmt.Println("添加区块成功")
}

