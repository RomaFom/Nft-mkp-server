package app

type Transaction struct {
	Id      int    `json:"-" db:"id"`
	Wallet  string `json:"wallet" db:"wallet" binding:"required"`
	UserId  int    `json:"user_id" db:"user_id" binding:"required"`
	ItemId  int    `json:"item_id" db:"item_id" binding:"required"`
	TxHash  string `json:"tx_hash" db:"tx_hash" binding:"required"`
	Created string `json:"created" db:"created_at"`
}

type TransactionInput struct {
	TxHash string `json:"tx_hash" binding:"required"`
	Wallet string `json:"wallet" binding:"required"`
	UserId int    `json:"user_id" binding:"required"`
	ItemId int    `json:"item_id" binding:"required"`
}
