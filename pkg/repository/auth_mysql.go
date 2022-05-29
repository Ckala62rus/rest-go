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
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values (?, ?, ?)", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthMysql) GetUser(username, password string) (rest.User, error) {
	var user rest.User
	logrus.Info(username, password)
	query := fmt.Sprintf("select id from %s where username=? and password_hash=?", userTable)
	res := r.db.QueryRow(query, username, password)
	res.Scan(&user)
	logrus.Info("wwwwwwwwwwwwwwwwwwwwww")

	// if err != nil {
	// 	panic(err)
	// // }

	// for res.Next() {
	// 	// err = res.Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	// 	err = res.Scan(&user.Username, &user.Password)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// fmt.Println(user)
	// }

	return user, nil
}
