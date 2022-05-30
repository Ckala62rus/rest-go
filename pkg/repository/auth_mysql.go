package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ckala62rus/rest-go"
	"github.com/sirupsen/logrus"
)

type AuthMysql struct {
	db *sql.DB
}

func NewAuthMysql(db *sql.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user rest.User) (int, error) {
	// var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values (?, ?, ?)", userTable)

	// row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	// if err := row.Scan(&id); err != nil {
	// 	return 0, err
	// }

	// return id, nil

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	returnId, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(returnId), nil
}

func (r *AuthMysql) GetUser(username, password string) (rest.User, error) {
	var user rest.User
	query := fmt.Sprintf("select id from %s where username=? and password_hash=?", userTable)
	// res := r.db.QueryRow(query, username, password)
	res, err := r.db.Query(query, username, password)

	if err != nil {
		logrus.Info(err)
	}

	for res.Next() {
		err = res.Scan(&user.Id)

		if err != nil {
			panic(err)
		}
	}

	return user, nil
}
