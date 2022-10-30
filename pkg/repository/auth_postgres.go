package repository

import (
	"app"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int

	user_check_query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", userTable)
	err := r.db.Get(&id, user_check_query, user.Username)

	if err == nil {
		return 0, errors.New("User already exists")
	}

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetUserById(id int) (app.UserPublicDTO, error) {
	var user app.UserPublicDTO
	query := fmt.Sprintf(`SELECT id, name, username, wallet FROM %s WHERE id=$1`, userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}
