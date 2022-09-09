package app

type Transaction struct {
	Id      int    `json:"-" db:"id"`
	UserId  int    `json:"user_id" db:"user_id" binding:"required"`
	TxHash  string `json:"tx_hash" db:"tx_hash" binding:"required"`
	Created string `json:"created" db:"created_at"`
}
