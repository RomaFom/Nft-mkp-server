package app

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Wallet   string `json:"wallet"`
	Password string `json:"password" binding:"required"`
}
