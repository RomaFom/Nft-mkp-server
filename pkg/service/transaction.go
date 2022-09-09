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

func (s *TransactionService) CreateTransaction(userId int, transaction string) (int, error) {
	return s.repo.CreateTransaction(userId, transaction)
}

func (s *TransactionService) GetAllTransactions() ([]app.Transaction, error) {
	return s.repo.GetAllTransactions()
}
