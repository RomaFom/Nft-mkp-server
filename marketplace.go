package app

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type NftDTO struct {
	Id          int            `json:"-" db:"id"`
	NftId       int            `json:"nft_id" db:"nft_id" binding:"required"`
	Image       string         `json:"image" db:"image" binding:"required"`
	Name        string         `json:"name" db:"name" binding:"required"`
	Description string         `json:"description" db:"description" binding:"required"`
	Owner       common.Address `json:"owner" db:"owner" binding:"required"`
}

type MarketplaceItemDTO struct {
	Id           int             `json:"-" db:"id"`
	ItemId       int64           `json:"item_id" db:"item_id" binding:"required"`
	TokenId      int64           `json:"token_id" db:"nft_id" binding:"required"`
	Price        decimal.Decimal `json:"price" db:"price" binding:"required"`
	ListingPrice decimal.Decimal `json:"listing_price" db:"listing_price" binding:"required"`
	Seller       common.Address  `json:"seller" db:"seller_wallet" binding:"required"`
	IsSold       bool            `json:"is_sold" db:"is_sold" binding:"required"`
	TotalPrice   decimal.Decimal `json:"total_price" db:"total_price" binding:"required"`
	Image        string          `json:"image" db:"image" binding:"required"`
	Name         string          `json:"name" db:"name" binding:"required"`
	Description  string          `json:"description" db:"description" binding:"required"`
	Owner        common.Address  `json:"owner" db:"owner" binding:"required"`
}

type BuyItemInput struct {
	TxHash string `json:"tx_hash" binding:"required"`
	ItemId int    `json:"item_id" binding:"required"`
	Wallet string `json:"wallet" binding:"required"`
	NftId  int    `json:"nft_id" binding:"required"`
}
