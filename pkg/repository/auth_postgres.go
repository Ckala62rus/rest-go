package repository

import (
	"fmt"

	"github.com/Ckala62rus/rest-go"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user rest.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1,$2,$3) RETURNING id", userTable)
	// query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1,$2,$3)", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
	// return 1, nil
}
