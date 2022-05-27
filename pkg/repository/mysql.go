package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userTable       = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemTable  = "lists_items"
)

type ConfigMySQL struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewMysqlDB(cfg ConfigMySQL) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:3306@tcp(mysql-education:3306)/go")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
