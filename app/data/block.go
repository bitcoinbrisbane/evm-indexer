package data

import (
	"encoding/json"
	"log"
)

// Block - Block related info to be delivered to client in this format
type Block struct {
	Hash                string  `json:"hash" gorm:"column:hash"`
	Number              uint64  `json:"number" gorm:"column:number"`
	Time                uint64  `json:"time" gorm:"column:time"`
	ParentHash          string  `json:"parentHash" gorm:"column:parenthash"`
	Difficulty          string  `json:"difficulty" gorm:"column:difficulty"`
	GasUsed             uint64  `json:"gasUsed" gorm:"column:gasused"`
	GasLimit            uint64  `json:"gasLimit" gorm:"column:gaslimit"`
	Nonce               string  `json:"nonce" gorm:"column:nonce"`
	Miner               string  `json:"miner" gorm:"column:miner"`
	Size                float64 `json:"size" gorm:"column:size"`
	StateRootHash       string  `json:"stateRootHash" gorm:"column:stateroothash"`
	UncleHash           string  `json:"uncleHash" gorm:"column:unclehash"`
	TransactionRootHash string  `json:"txRootHash" gorm:"column:txroothash"`
	ReceiptRootHash     string  `json:"receiptRootHash" gorm:"column:receiptroothash"`
	ExtraData           []byte  `json:"extraData" gorm:"column:extradata"`
}

// ToJSON - Encodes into JSON, to be supplied when queried for block data
func (b *Block) ToJSON() []byte {
	data, err := json.Marshal(b)
	if err != nil {
		log.Printf("[!] Failed to encode block data to JSON : %s\n", err.Error())
		return nil
	}

	return data
}

// Blocks - A set of blocks to be held, extracted from DB query result
// also to be supplied to client in JSON encoded form
type Blocks struct {
	Blocks []*Block `json:"blocks"`
}

// ToJSON - Encoding into JSON, to be invoked when delivering query result to client
func (b *Blocks) ToJSON() []byte {
	data, err := json.Marshal(b)
	if err != nil {
		log.Printf("[!] Failed to encode block data to JSON : %s\n", err.Error())
		return nil
	}

	return data
}
