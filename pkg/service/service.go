package service

import (
	"app"
	"app/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUserById(id int) (app.User, error)
}

type Transaction interface {
	CreateTransaction(userId int, transaction string) (int, error)
	GetAllTransactions() ([]app.Transaction, error)
}

type Service struct {
	Authorization
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Transaction:   NewTransactionService(repos.Transaction),
	}
}
