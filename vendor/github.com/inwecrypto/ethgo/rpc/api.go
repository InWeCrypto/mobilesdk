package rpc

// Block eth block object
type Block struct {
	Number           string         `json:"number"`
	Hash             string         `json:"hash"`
	Parent           string         `json:"parentHash"`
	Nonce            string         `json:"nonce"`
	SHA3Uncles       string         `json:"sha3Uncles"`
	LogsBloom        string         `json:"logsBloom"`
	TransactionsRoot string         `json:"transactionsRoot"`
	StateRoot        string         `json:"stateRoot"`
	ReceiptsRoot     string         `json:"receiptsRoot"`
	Miner            string         `json:"miner"`
	Difficulty       string         `json:"difficulty"`
	TotalDifficulty  string         `json:"totalDifficulty"`
	ExtraData        string         `json:"extraData"`
	Size             string         `json:"size"`
	GasLimit         string         `json:"gasLimit"`
	GasUsed          string         `json:"gasUsed"`
	Timestamp        string         `json:"timestamp"`
	Transactions     []*Transaction `json:"transactions"`
	Uncles           []string       `json:"uncles"`
}

// Transaction .
type Transaction struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	GasPrice         string `json:"gasPrice"`
	Gas              string `json:"gas"`
	Input            string `json:"input"`
}

// CallSite .
type CallSite struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Value    string `json:"value,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Gas      string `json:"gas,omitempty"`
	Data     string `json:"data,omitempty"`
}
