package repository

import (
	"app"
	sc_api "app/pkg/smart-contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

type MarketplaceSC struct {
	sc *sc_api.ScApi
}

func NewMarketplaceSc(sc *sc_api.ScApi) *MarketplaceSC {
	return &MarketplaceSC{
		sc: sc,
	}
}

func (r *MarketplaceSC) GetItemCount() (*big.Int, error) {
	count, err := r.sc.ItemCount(&bind.CallOpts{})
	return count, err
}

func (r *MarketplaceSC) GetMarketplaceItems(count *big.Int) ([]app.MarketplaceItemDTO, error) {
	var items []app.MarketplaceItemDTO

	for i := 1; i <= int(count.Int64()); i++ {
		item, err := r.sc.Items(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}
		items = append(items, app.MarketplaceItemDTO{
			ItemId:       big.NewInt(int64(i)),
			Nft:          item.Nft,
			TokenId:      item.TokenId,
			Price:        item.Price,
			ListingPrice: item.ListingPrice,
			Seller:       item.Seller,
			IsSold:       item.IsSold,
		})
	}
	return items, nil
}
