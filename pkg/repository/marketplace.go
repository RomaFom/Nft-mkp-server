package repository

import (
	"app"
	mkp_api "app/pkg/MkpSc"
	nft_api "app/pkg/NftSc"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"math/big"
	"net/http"
	"strings"
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

func (r *MarketplaceSC) GetItemsNoThreads() ([]app.MarketplaceItemDTO, error) {
	count, err := r.GetItemCount()
	if err != nil {
		return nil, err
	}
	var items []app.MarketplaceItemDTO

	for i := 1; i <= int(count.Int64()); i++ {

		// Get item from marketplace Contract
		item, err := r.MkpSc.Items(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		if item.IsSold == false {
			// Call NFT contract to get tokenURI
			nftItem, err := r.GetNftMetadata(item.TokenId)
			if err != nil {
				return nil, err
			}

			finalPrice, err := r.MkpSc.GetFinalPrice(&bind.CallOpts{}, item.ItemId)
			if err != nil {
				return nil, err
			}

			items = append(items, app.MarketplaceItemDTO{
				ItemId:       int64(i),
				Nft:          nftItem,
				TokenId:      item.TokenId.Int64(),
				Price:        app.ToDecimal(item.Price, 18),
				ListingPrice: app.ToDecimal(item.ListingPrice, 18),
				Seller:       item.Seller,
				IsSold:       item.IsSold,
				TotalPrice:   app.ToDecimal(finalPrice, 18),
			})
		}

	}
	return items, nil
}

func (r *MarketplaceSC) GetMarketplaceItems() ([]app.MarketplaceItemDTO, error) {
	wg := &sync.WaitGroup{}
	count, err := r.GetItemCount()
	if err != nil {
		return nil, err
	}
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
					ItemId:       int64(i),
					Nft:          nftItem,
					TokenId:      item.TokenId.Int64(),
					Price:        app.ToDecimal(item.Price, 18),
					ListingPrice: app.ToDecimal(item.ListingPrice, 18),
					Seller:       item.Seller,
					IsSold:       item.IsSold,
					TotalPrice:   app.ToDecimal(finalPrice, 18),
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

	return items, nil
}

func (r *MarketplaceSC) GetMyListings(wallet string) ([]app.MarketplaceItemDTO, error) {
	wg := &sync.WaitGroup{}
	count, err := r.GetItemCount()
	if err != nil {
		return nil, err
	}
	wg.Add(int(count.Int64()))
	var items []app.MarketplaceItemDTO

	ch := make(chan app.MarketplaceItemDTO)

	for i := 1; i <= int(count.Int64()); i++ {
		go func(i int) {
			defer wg.Done()
			item, err := r.MkpSc.Items(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				return
			}

			if item.IsSold == false && strings.ToLower(item.Seller.String()) == strings.ToLower(wallet) {
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
					ItemId:       int64(i),
					Nft:          nftItem,
					TokenId:      item.TokenId.Int64(),
					Price:        app.ToDecimal(item.Price, 18),
					ListingPrice: app.ToDecimal(item.ListingPrice, 18),
					Seller:       item.Seller,
					IsSold:       item.IsSold,
					TotalPrice:   app.ToDecimal(finalPrice, 18),
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

	return items, nil
}

func (r *MarketplaceSC) GetMyPurchases(wallet string) ([]app.MarketplaceItemDTO, error) {
	var items []app.MarketplaceItemDTO
	address := common.HexToAddress(wallet)

	bought, err := r.MkpSc.FilterBought(&bind.FilterOpts{}, nil, nil, []common.Address{address})
	if err != nil {
		return nil, err
	}

	for bought.Next() {
		item, err := r.MkpSc.Items(&bind.CallOpts{}, bought.Event.ItemId)
		if err != nil {
			return nil, err
		}
		owner, err := r.NftSc.GetOwner(&bind.CallOpts{}, item.TokenId)
		if err != nil {
			return nil, err
		}

		if strings.ToLower(owner.String()) == strings.ToLower(wallet) {

			// Call NFT contract to get tokenURI
			nftItem, err := r.GetNftMetadata(item.TokenId)
			if err != nil {
				return nil, err
			}

			finalPrice, err := r.MkpSc.GetFinalPrice(&bind.CallOpts{}, item.ItemId)
			if err != nil {
				return nil, err
			}

			items = append(items, app.MarketplaceItemDTO{
				ItemId:       item.ItemId.Int64(),
				Nft:          nftItem,
				TokenId:      item.TokenId.Int64(),
				Price:        app.ToDecimal(item.Price, 18),
				ListingPrice: app.ToDecimal(item.ListingPrice, 18),
				Seller:       item.Seller,
				IsSold:       item.IsSold,
				TotalPrice:   app.ToDecimal(finalPrice, 18),
			})
		}
	}

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
