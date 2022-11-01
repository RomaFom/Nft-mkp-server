package app

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type MarketplaceItemDTO struct {
	ItemId       *big.Int
	Nft          NftDTO
	TokenId      *big.Int
	Price        *big.Int
	ListingPrice *big.Int
	Seller       common.Address
	IsSold       bool
	TotalPrice   *big.Int
}

type NftDTO struct {
	Image       string `json:"image" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
