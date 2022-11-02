package service

import (
	"app"
	"app/pkg/repository"
	"math/big"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUserById(id int) (app.UserPublicDTO, error)
}

type Transaction interface {
	CreateTransaction(wallet string, transaction string) (int, error)
	GetAllTransactions() ([]app.Transaction, error)
}

type Marketplace interface {
	GetItemCount() (*big.Int, error)
	GetMarketplaceItems() ([]app.MarketplaceItemDTO, error)
	GetMyListings(wallet string) ([]app.MarketplaceItemDTO, error)
	GetMyPurchases(wallet string) ([]app.MarketplaceItemDTO, error)
	GetItemsNoThreads() ([]app.MarketplaceItemDTO, error)
}

type Service struct {
	Authorization
	Transaction
	Marketplace
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Transaction:   NewTransactionService(repos.Transaction),
		Marketplace:   NewMarketplaceService(repos.Marketplace),
	}
}
