package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ckala62rus/rest-go"
)

type AuthMysql struct {
	db *sql.DB
}

func NewAuthMysql(db *sql.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user rest.User) (int, error) {
	var id int

	// query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1,$2,$3) RETURNING id", userTable)
	// query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1,$2,$3)", userTable)
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values (?, ?, ?)", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	// row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
