package app

import "github.com/shopspring/decimal"

type MarketplaceItem struct {
	Id           int             `json:"-" db:"id"`
	ItemId       int             `json:"item_id" db:"item_id" binding:"required"`
	NftId        int             `json:"nft_id" db:"nft_id" binding:"required"`
	Price        decimal.Decimal `json:"price" db:"price" binding:"required"`
	ListingPrice decimal.Decimal `json:"listing_price" db:"listing_price" binding:"required"`
	TotalPrice   decimal.Decimal `json:"total_price" db:"total_price" binding:"required"`
	SellerWallet string          `json:"seller_wallet" db:"seller_wallet" binding:"required"`
	//SellerId     int             `json:"seller_id" db:"seller_id" binding:"required"`
	IsSold bool `json:"is_sold" db:"is_sold" binding:"required"`
}

//id serial not null unique,
//item_id varchar(255) not null unique,
//nft_id int not null references nfts(nft_id),
//price decimal not null,
//listing_price decimal not null,
//total_price decimal not null,
//seller_wallet varchar(255) not null,
//seller_id int not null references users(id),
//is_sold boolean not null default false
