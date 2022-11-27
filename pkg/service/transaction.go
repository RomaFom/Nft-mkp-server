package service

import (
	"app"
	"app/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(tx app.Transaction) (int, error) {
	return s.repo.CreateTransaction(tx)
}

func (s *TransactionService) GetAllTransactions() ([]app.Transaction, error) {
	return s.repo.GetAllTransactions()
}

func (s *TransactionService) GetNftTransactions(id int) ([]app.Transaction, error) {
	return s.repo.GetNftTransactions(id)
}
