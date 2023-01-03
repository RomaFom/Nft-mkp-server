package service

import (
	"app"
	"app/pkg/repository"
	"fmt"
	"math/big"
)

type MarketplaceService struct {
	repo repository.Marketplace
}

func NewMarketplaceService(repo repository.Marketplace) *MarketplaceService {
	return &MarketplaceService{repo: repo}
}

func (s *MarketplaceService) GetItemCount() (*big.Int, error) {
	return s.repo.GetItemCount()
}

func (s *MarketplaceService) GetMarketplaceItemsFromSC() ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMarketplaceItemsFromSC()
}

func (s *MarketplaceService) GetItemsForSale(page int, size int) ([]app.CombinedItemDTO, error) {
	return s.repo.GetItemsForSale(page, size)
}

func (s *MarketplaceService) GetMyPurchases(wallet string, page int, size int) ([]app.CombinedItemDTO, error) {
	return s.repo.GetMyPurchases(wallet, page, size)
}

func (s *MarketplaceService) GetMyListings(wallet string) ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMyListings(wallet)
}

func (s *MarketplaceService) BuyItem(itemId int) (app.MarketplaceItemDTO, error) {
	return s.repo.BuyItem(itemId)
}

func (s *MarketplaceService) ValidateSCItems() {
	items, err := s.repo.GetMarketplaceItemsFromSC()
	fmt.Println("items", len(items))
	if err != nil {
		fmt.Printf("Error getting items from SC: %v", err)
	}
	err = s.repo.ValidateSCItems(items)
	if err != nil {
		fmt.Printf("Error validating items: %v", err)
	}
}
