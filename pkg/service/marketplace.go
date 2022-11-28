package service

import (
	"app"
	"app/pkg/repository"
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

func (s *MarketplaceService) GetItemsForSale() ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetItemsForSale()
}

func (s *MarketplaceService) GetMyPurchases(wallet string) ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMyPurchases(wallet)
}

func (s *MarketplaceService) GetMyListings(wallet string) ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMyListings(wallet)
}

func (s *MarketplaceService) BuyItem(itemId int) (app.MarketplaceItemDTO, error) {
	return s.repo.BuyItem(itemId)
}

func (s *MarketplaceService) ValidateSCItems() error {
	items, err := s.repo.GetMarketplaceItemsFromSC()
	if err != nil {
		return err
	}
	return s.repo.ValidateSCItems(items)
}
