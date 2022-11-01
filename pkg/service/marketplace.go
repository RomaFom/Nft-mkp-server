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

func (s *MarketplaceService) GetMarketplaceItems(count *big.Int) ([]app.MarketplaceItemDTO, error) {
	return s.repo.GetMarketplaceItems(count)
}
