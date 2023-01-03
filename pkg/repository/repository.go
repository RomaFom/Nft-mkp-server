package repository

import (
	"app"
	mkp_api "app/pkg/MkpSc"
	nft_api "app/pkg/NftSc"
	"github.com/jmoiron/sqlx"
	"math/big"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username, password string) (app.User, error)
	GetUserById(id int) (app.UserPublicDTO, error)
}

type Transaction interface {
	CreateTransaction(tx app.Transaction) (int, error)
	GetAllTransactions() ([]app.Transaction, error)
	GetNftTransactions(id int) ([]app.Transaction, error)
}

type Marketplace interface {
	GetItemCount() (*big.Int, error)
	GetMarketplaceItemsFromSC() ([]app.MarketplaceItemDTO, error)
	GetItemsForSale(page int, size int) ([]app.CombinedItemDTO, error)
	GetMyListings(wallet string) ([]app.MarketplaceItemDTO, error)
	GetMyPurchases(wallet string, page int, size int) ([]app.CombinedItemDTO, error)
	BuyItem(itemId int) (app.MarketplaceItemDTO, error)
	ValidateSCItems(items []app.MarketplaceItemDTO) error
}

type Repository struct {
	Authorization
	Transaction
	Marketplace
}

func NewRepository(db *sqlx.DB, mkp *mkp_api.MkpApi, nft *nft_api.NftApi) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transaction:   NewTransactionPostgres(db),
		Marketplace:   NewMarketplaceSc(mkp, nft, db),
	}
}
