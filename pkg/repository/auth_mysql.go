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
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1,$2,$3)", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// func (r *AuthMysql) InsertTest(user rest.User) (int, error) {
// 	var id int
// 	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ('ckala','ckala62rus','123123') RETURNING id", userTable)

// 	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
// 	if err := row.Scan(&id); err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }
