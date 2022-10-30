package app

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type MarketplaceItemDTO struct {
	ItemId       *big.Int
	Nft          common.Address
	TokenId      *big.Int
	Price        *big.Int
	ListingPrice *big.Int
	Seller       common.Address
	IsSold       bool
}
