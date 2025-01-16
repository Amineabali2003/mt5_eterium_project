package model

type TransactionResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

type Transaction struct {
	Hash      string `json:"hash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	TimeStamp string `json:"timeStamp"`
	GasPrice  string `json:"gasPrice"`
	GasUsed   string `json:"gasUsed"`
}

type WalletDataResponse struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}
