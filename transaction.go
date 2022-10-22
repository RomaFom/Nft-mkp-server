package app

type Transaction struct {
	Id      int    `json:"-" db:"id"`
	Wallet  int    `json:"wallet" db:"wallet" binding:"required"`
	TxHash  string `json:"tx_hash" db:"tx_hash" binding:"required"`
	Created string `json:"created" db:"created_at"`
}
