package main

type BlockChain struct {
	blocks []* Block
}

func (bc *BlockChain)AddBlock(data string)  {
	lastBlock := bc.blocks[len(bc.blocks) - 1]

	block := NewBlock(data, lastBlock.Hash)

	bc.blocks = append(bc.blocks, block)
}

// 5.定义一个区块链
func NewBlockChain() *BlockChain  {
	// 创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// 创世快
func GenesisBlock() *Block {
	return NewBlock("Go创世块，老牛逼了",[]byte{})
}