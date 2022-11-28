package repository

import (
	"app"
	mkp_api "app/pkg/MkpSc"
	nft_api "app/pkg/NftSc"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
)

type MarketplaceSC struct {
	MkpSc *mkp_api.MkpApi
	NftSc *nft_api.NftApi
	db    *sqlx.DB
}

func NewMarketplaceSc(mkp *mkp_api.MkpApi, nft *nft_api.NftApi, db *sqlx.DB) *MarketplaceSC {
	return &MarketplaceSC{
		MkpSc: mkp,
		NftSc: nft,
		db:    db,
	}
}

func (r *MarketplaceSC) GetItemCount() (*big.Int, error) {
	count, err := r.MkpSc.ItemCount(&bind.CallOpts{})
	return count, err
}

func (r *MarketplaceSC) GetItemsForSale() ([]app.MarketplaceItemDTO, error) {
	var items []app.MarketplaceItemDTO
	query := fmt.Sprintf(`SELECT * FROM %s it INNER JOIN %s nt ON it.nft_id = nt.nft_id WHERE it.is_sold=false ORDER BY it.item_id`, itemsTable, nftsTable)

	if err := r.db.Select(&items, query); err != nil {
		fmt.Printf("Error: %v", err.Error())
		return nil, err
	}
	return items, nil
}

func (r *MarketplaceSC) GetMarketplaceItemFromSCById(itemId int) (app.MarketplaceItemDTO, error) {
	var mkpItem app.MarketplaceItemDTO
	item, err := r.MkpSc.Items(&bind.CallOpts{}, big.NewInt(int64(itemId)))
	if err != nil {
		return mkpItem, err
	}
	nftItem, err := r.GetNftMetadata(item.TokenId)
	if err != nil {
		return mkpItem, err
	}
	finalPrice, err := r.MkpSc.GetFinalPrice(&bind.CallOpts{}, item.ItemId)
	if err != nil {
		return mkpItem, err
	}

	mkpItem = app.MarketplaceItemDTO{
		ItemId: int64(itemId),
		//Nft:          nftItem,
		TokenId:      item.TokenId.Int64(),
		Price:        app.ToDecimal(item.Price, 18),
		ListingPrice: app.ToDecimal(item.ListingPrice, 18),
		Seller:       item.Seller,
		IsSold:       item.IsSold,
		TotalPrice:   app.ToDecimal(finalPrice, 18),
		Image:        nftItem.Image,
		Name:         nftItem.Name,
		Description:  nftItem.Description,
		Owner:        nftItem.Owner,
	}
	return mkpItem, nil
}

func (r *MarketplaceSC) GetMarketplaceItemsFromSC() ([]app.MarketplaceItemDTO, error) {
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

			//if item.IsSold == false {
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
				ItemId: int64(i),
				//Nft:          nftItem,
				TokenId:      item.TokenId.Int64(),
				Price:        app.ToDecimal(item.Price, 18),
				ListingPrice: app.ToDecimal(item.ListingPrice, 18),
				Seller:       item.Seller,
				IsSold:       item.IsSold,
				TotalPrice:   app.ToDecimal(finalPrice, 18),
				Image:        nftItem.Image,
				Name:         nftItem.Name,
				Description:  nftItem.Description,
				Owner:        nftItem.Owner,
			}
			//}
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
					ItemId: int64(i),
					//Nft:          nftItem,
					TokenId:      item.TokenId.Int64(),
					Price:        app.ToDecimal(item.Price, 18),
					ListingPrice: app.ToDecimal(item.ListingPrice, 18),
					Seller:       item.Seller,
					IsSold:       item.IsSold,
					TotalPrice:   app.ToDecimal(finalPrice, 18),
					Image:        nftItem.Image,
					Name:         nftItem.Name,
					Description:  nftItem.Description,
					Owner:        nftItem.Owner,
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
			nftItem.Owner = owner

			finalPrice, err := r.MkpSc.GetFinalPrice(&bind.CallOpts{}, item.ItemId)
			if err != nil {
				return nil, err
			}

			items = append(items, app.MarketplaceItemDTO{
				ItemId: item.ItemId.Int64(),
				//Nft:          nftItem,
				TokenId:      item.TokenId.Int64(),
				Price:        app.ToDecimal(item.Price, 18),
				ListingPrice: app.ToDecimal(item.ListingPrice, 18),
				Seller:       item.Seller,
				IsSold:       item.IsSold,
				TotalPrice:   app.ToDecimal(finalPrice, 18),
				Image:        nftItem.Image,
				Name:         nftItem.Name,
				Description:  nftItem.Description,
				Owner:        nftItem.Owner,
			})
		}
	}

	return items, nil
}

func (r *MarketplaceSC) GetNftMetadata(id *big.Int) (app.NftDTO, error) {
	var nftItem app.NftDTO

	//Get Owner
	owner, err := r.NftSc.GetOwner(&bind.CallOpts{}, id)
	if err != nil {
		return nftItem, err
	}
	nftItem.Owner = owner

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

func (r *MarketplaceSC) CheckItemInDB(itemId int64) (bool, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE item_id=$1", itemsTable)
	err := r.db.Get(&id, query, itemId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *MarketplaceSC) CheckNftInDB(nftId int64) (bool, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE nft_id=$1", nftsTable)
	err := r.db.Get(&id, query, nftId)
	if err != nil {

		return false, err
	}
	return true, nil
}

func (r *MarketplaceSC) UpdateItemInDB(item app.MarketplaceItemDTO) error {
	query := fmt.Sprintf("UPDATE %s SET owner=$1, image=$2, name=$3, description=$4 WHERE nft_id=$5", nftsTable)
	_, err := r.db.Exec(query, item.Owner, item.Image, item.Name, item.Description, item.TokenId)
	if err != nil {
		logrus.Fatalf("ERROR UPDATE NFT %d", item.TokenId)
		return err
	}
	query = fmt.Sprintf("UPDATE %s SET price=$1, listing_price=$2, total_price=$3, seller_wallet=$4, is_sold=$5 WHERE item_id=$6", itemsTable)
	_, err = r.db.Exec(query, item.Price, item.ListingPrice, item.TotalPrice, item.Seller, item.IsSold, item.ItemId)
	if err != nil {
		logrus.Fatalf("ERROR UPDATE ITEM %d", item.ItemId)
		return err
	}
	return nil
}

func (r *MarketplaceSC) ValidateSCItems(items []app.MarketplaceItemDTO) error {
	var wg sync.WaitGroup
	wg.Add(len(items))
	logrus.Printf("Starting validate items go routines num %d", len(items))

	for i := 0; i < len(items); i++ {
		go func(i int) {
			defer wg.Done()
			item := items[i]
			var id int
			isExists, err := r.CheckNftInDB(item.TokenId)
			if isExists {
				err = r.UpdateItemInDB(item)
				if err != nil {
					logrus.Fatalf("ERROR UPDATE ITEM %d", item.ItemId)
					return
				}
				//query := fmt.Sprintf("UPDATE %s SET owner=$1, image=$2, name=$3, description=$4 WHERE nft_id=$5", nftsTable)
				//_, err = r.db.Exec(query, item.Owner, item.Image, item.Name, item.Description, item.TokenId)
				//if err != nil {
				//	logrus.Fatalf("ERROR UPDATE NFT %d", item.TokenId)
				//	return
				//}
				//query = fmt.Sprintf("UPDATE %s SET price=$1, listing_price=$2, total_price=$3, seller_wallet=$4, is_sold=$5 WHERE item_id=$6", itemsTable)
				//_, err = r.db.Exec(query, item.Price, item.ListingPrice, item.TotalPrice, item.Seller, item.IsSold, item.ItemId)
				//if err != nil {
				//	logrus.Fatalf("ERROR UPDATE ITEM %d", item.ItemId)
				//	return
				//}
				return
			}

			query := fmt.Sprintf("INSERT INTO %s (nft_id, owner ,image, name, description) VALUES ($1, $2, $3, $4, $5) RETURNING id", nftsTable)
			row := r.db.QueryRow(query, item.TokenId, item.Owner, item.Image, item.Name, item.Description)
			if err := row.Scan(&id); err != nil {
				logrus.Fatalf("ERROR ADD NFT %d", item.TokenId)
				return
			}
			query = fmt.Sprintf("INSERT INTO %s (item_id, nft_id ,price, listing_price, total_price,seller_wallet,is_sold) VALUES ($1, $2, $3, $4, $5, $6 ,$7) RETURNING id", itemsTable)
			row = r.db.QueryRow(query, item.ItemId, item.TokenId, item.Price, item.ListingPrice, item.TotalPrice, item.Seller, item.IsSold)
			if err := row.Scan(&id); err != nil {
				logrus.Fatalf("ERROR ADD ITEM %d", item.TokenId)
				return
			}
		}(i)
	}
	wg.Wait()
	logrus.Printf("Successfully validated %d items", len(items))
	return nil
}

func (r *MarketplaceSC) BuyItem(itemId int) (app.MarketplaceItemDTO, error) {
	item, err := r.GetMarketplaceItemFromSCById(itemId)
	fmt.Printf("ITEM %v", item)
	//mkpItem := app.MarketplaceItemDTO{}
	if err != nil {
		return item, err
	}
	err = r.UpdateItemInDB(item)
	if err != nil {
		return item, err
	}
	query := fmt.Sprintf(`SELECT * FROM %s it INNER JOIN %s nt ON it.nft_id = nt.nft_id WHERE it.item_id=$1`, itemsTable, nftsTable)

	if err := r.db.Get(&item, query, itemId); err != nil {
		fmt.Printf("Error: %v", err.Error())
		return item, err
	}
	return item, nil
}
