package repository

import (
	"app"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{
		db: db,
	}
}

func (r *TransactionPostgres) CreateTransaction(wallet string, transaction string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (wallet, tx_hash) VALUES ($1, $2) RETURNING id", transactionsTable)
	row := r.db.QueryRow(query, wallet, transaction)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TransactionPostgres) GetAllTransactions() ([]app.Transaction, error) {
	var txList []app.Transaction
	query := fmt.Sprintf("SELECT * FROM %s", transactionsTable)
	err := r.db.Select(&txList, query)
	return txList, err
}
