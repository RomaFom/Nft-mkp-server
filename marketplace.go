package app

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type MarketplaceItemDTO struct {
	ItemId       int64
	Nft          NftDTO
	TokenId      int64
	Price        decimal.Decimal
	ListingPrice decimal.Decimal
	Seller       common.Address
	IsSold       bool
	TotalPrice   decimal.Decimal
}

type NftDTO struct {
	Image       string `json:"image" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
