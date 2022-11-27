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

func (r *TransactionPostgres) CreateTransaction(tx app.Transaction) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (wallet, tx_hash,user_id,item_id) VALUES ($1, $2,$3,$4) RETURNING id", transactionsTable)
	row := r.db.QueryRow(query, tx.Wallet, tx.TxHash, tx.UserId, tx.ItemId)
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

func (r *TransactionPostgres) GetNftTransactions(id int) ([]app.Transaction, error) {
	var txList []app.Transaction
	query := fmt.Sprintf("SELECT * FROM %s WHERE item_id=$1", transactionsTable)
	err := r.db.Select(&txList, query, id)
	return txList, err
}
