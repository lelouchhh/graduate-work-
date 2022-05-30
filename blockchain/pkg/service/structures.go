package service

import (
	"time"
)

type Block struct {
	Index        int         `json:"index"`
	Timestamp    time.Time   `json:"Timestamp"`
	Transactions Transaction `json:"Transactions"`
	PrevHash     string      `json:"PrevHash"`
}

type Transaction struct {
	User   string
	Action string

	VmId    string
	VmIP    string
	Account string
	Ssh     string
	Ram     string
	Disk    string
	Cpu     string
	Pass    string
}

type BlockChain struct {
	Blocks []Block
}
type OutputConfig struct {
	IP      string
	Account string
	Pass    string
}
