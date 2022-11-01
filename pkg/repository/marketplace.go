package repository

import (
	"app"
	mkp_api "app/pkg/MkpSc"
	nft_api "app/pkg/NftSc"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"io"
	"math/big"
	"net/http"
	"sync"
)

type MarketplaceSC struct {
	MkpSc *mkp_api.MkpApi
	NftSc *nft_api.NftApi
}

func NewMarketplaceSc(mkp *mkp_api.MkpApi, nft *nft_api.NftApi) *MarketplaceSC {
	return &MarketplaceSC{
		MkpSc: mkp,
		NftSc: nft,
	}
}

func (r *MarketplaceSC) GetItemCount() (*big.Int, error) {
	count, err := r.MkpSc.ItemCount(&bind.CallOpts{})
	return count, err
}

func (r *MarketplaceSC) GetMarketplaceItems(count *big.Int) ([]app.MarketplaceItemDTO, error) {
	wg := &sync.WaitGroup{}
	wg.Add(int(count.Int64()))
	var items []app.MarketplaceItemDTO

	ch := make(chan app.MarketplaceItemDTO)

	for i := 1; i <= int(count.Int64()); i++ {
		//wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Get item from marketplace Contract
			item, err := r.MkpSc.Items(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				return
			}

			if item.IsSold == false {
				// Call NFT contract to get tokenURI
				nftItem, err := r.GetNftMetadata(item.TokenId)
				if err != nil {
					return
				}

				finalPrice, err := r.MkpSc.GetFinalPrice(&bind.CallOpts{}, item.ItemId)
				if err != nil {
					return
				}

				ch <- app.MarketplaceItemDTO{
					ItemId:       big.NewInt(int64(i)),
					Nft:          nftItem,
					TokenId:      item.TokenId,
					Price:        item.Price,
					ListingPrice: item.ListingPrice,
					Seller:       item.Seller,
					IsSold:       item.IsSold,
					TotalPrice:   finalPrice,
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for item := range ch {
		items = append(items, item)
	}

	fmt.Println(items)
	return items, nil
}

func (r *MarketplaceSC) GetNftMetadata(id *big.Int) (app.NftDTO, error) {
	var nftItem app.NftDTO

	// Get token URI from NFT contract
	tokenURI, _ := r.NftSc.TokenURI(&bind.CallOpts{}, id)
	// Call tokenURI to get metadata
	resp, err := http.Get(tokenURI)
	if err != nil {
		return nftItem, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nftItem, err
	}

	err = json.Unmarshal(body, &nftItem)
	if err != nil {
		return nftItem, err
	}

	return nftItem, nil
}
