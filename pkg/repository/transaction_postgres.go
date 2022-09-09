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

func (r *TransactionPostgres) CreateTransaction(userId int, transaction string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, tx_hash) VALUES ($1, $2) RETURNING id", transactionsTable)
	row := r.db.QueryRow(query, userId, transaction)
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
