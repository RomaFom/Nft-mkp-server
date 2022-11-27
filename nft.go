package app

import "github.com/ethereum/go-ethereum/common"

type Nft struct {
	Id          int            `json:"-" db:"id"`
	NftId       int            `json:"nft_id" db:"nft_id" binding:"required"`
	OwnerId     common.Address `json:"owner_id" db:"owner_id" binding:"required"`
	Image       string         `json:"image" db:"image" binding:"required"`
	Name        string         `json:"name" db:"name" binding:"required"`
	Description string         `json:"description" db:"description" binding:"required"`
}

//owner_id varchar(255) not null,
//image varchar(255) not null,
//name varchar(255) not null,
//description varchar(255) not null
