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

func (s *MarketplaceService) GetMarketplaceItems() ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMarketplaceItems()
}

func (s *MarketplaceService) GetItemsNoThreads() ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetItemsNoThreads()
}

func (s *MarketplaceService) GetMyPurchases(wallet string) ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMyPurchases(wallet)
}

func (s *MarketplaceService) GetMyListings(wallet string) ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMyListings(wallet)
}
