package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const ROOT_DIR = "/home/llchh/GolandProjects/blockchain/chain/"

func CreateBlock(t Transaction) *Block {
	files, err := FilePathWalkDir(ROOT_DIR)
	if err != nil {
		fmt.Println(err)
	}
	lastFile := files[len(files)-1]
	lastFileContent, err := os.Open(lastFile)
	b := Block{
		Index:        len(files) + 1,
		Timestamp:    time.Now(),
		Transactions: t,

		PrevHash: hash(lastFileContent),
	}
	file, _ := json.MarshalIndent(b, "", " ")
	_ = ioutil.WriteFile(fmt.Sprint(ROOT_DIR, len(files)+1), file, 0644)
	return &b
}

func CheckIntegrity() []string {
	files, _ := FilePathWalkDir(ROOT_DIR)
	res := make([]string, len(files))
	for i := 1; i < len(files); i++ {
		jsonFile, err := os.Open(files[i])
		prevHash := getHash(jsonFile)
		if err != nil {
			fmt.Println(err)
		}
		res[0] = "Genesis block"
		currentFile, _ := os.Open(files[i-1])

		if prevHash != hash(currentFile) {
			res[i] = fmt.Sprintf("block %d was corrupted!", i+1)
		} else {
			res[i] = fmt.Sprintf("block %d is OK!", +i+1)
		}

	}
	return res

}
func Chain() *BlockChain {
	var blockchain BlockChain
	files, _ := FilePathWalkDir(ROOT_DIR)
	//var blocks []Block
	for i := 0; i < len(files); i++ {
		jsonFile, err := os.Open(files[i])
		block, _ := ioutil.ReadAll(jsonFile)
		var result Block
		json.Unmarshal([]byte(block), &result)
		blockchain.Blocks = append(blockchain.Blocks, result)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(blockchain)
	return &blockchain
}
func MakeOutputConfig(t Transaction) *OutputConfig {
	BlockChain := Chain()
	for i := range BlockChain.Blocks {
		fmt.Println(t.User, BlockChain.Blocks[i].Transactions.User)
		if BlockChain.Blocks[i].Transactions.User == t.User && BlockChain.Blocks[i].Transactions.VmId == t.VmId {
			return &OutputConfig{
				IP:      BlockChain.Blocks[i].Transactions.VmIP,
				Account: BlockChain.Blocks[i].Transactions.Account,
				Pass:    BlockChain.Blocks[i].Transactions.Pass,
			}
		}
	}
	return &OutputConfig{}
}
