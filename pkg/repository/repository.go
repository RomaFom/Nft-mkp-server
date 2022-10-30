package repository

import (
	"app"
	sc_api "app/pkg/smart-contracts"
	"github.com/jmoiron/sqlx"
	"math/big"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username, password string) (app.User, error)
	GetUserById(id int) (app.UserPublicDTO, error)
}

type Transaction interface {
	CreateTransaction(wallet string, transaction string) (int, error)
	GetAllTransactions() ([]app.Transaction, error)
}

type Marketplace interface {
	GetItemCount() (*big.Int, error)
	GetMarketplaceItems(count *big.Int) ([]app.MarketplaceItemDTO, error)
}

type Repository struct {
	Authorization
	Transaction
	Marketplace
}

func NewRepository(db *sqlx.DB, sc *sc_api.ScApi) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transaction:   NewTransactionPostgres(db),
		Marketplace:   NewMarketplaceSc(sc),
	}
}
