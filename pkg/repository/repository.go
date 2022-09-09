package repository

import (
	"app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username, password string) (app.User, error)
	GetUserById(id int) (app.User, error)
}

type Transaction interface {
	CreateTransaction(userId int, transaction string) (int, error)
	GetAllTransactions() ([]app.Transaction, error)
}

type Repository struct {
	Authorization
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transaction:   NewTransactionPostgres(db),
	}
}
