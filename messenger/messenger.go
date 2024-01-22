package messenger

type Transaction struct {
	sender  string
	message string
}

type Block struct {
	header []byte
	txs    []Transaction
}

func NewBlock(txs []Transaction) Block {
	// TODO: hash the txs
	return Block{
		header: []byte{0x0},
		txs:    txs,
	}
}

type Messenger struct {
	Blocks map[uint32]*Block
}

func NewMessenger() *Messenger {
	return &Messenger{
		Blocks: map[uint32]*Block{},
	}
}
