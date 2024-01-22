package messenger

import "time"

type Transaction struct {
	sender  string
	message string
}

type Block struct {
	hash      []byte
	timestamp time.Time
	txs       []Transaction
}

func NewBlock(txs []Transaction) Block {
	// TODO: hash the txs
	return Block{
		hash: []byte{0x0},
		txs:  txs,
	}
}

type Messenger struct {
	Blocks []Block
}

func NewMessenger() *Messenger {
	return &Messenger{
		Blocks: []Block{},
	}
}
