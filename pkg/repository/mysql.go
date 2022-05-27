package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ConfigMySQL struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewMysqlDB(cfg Config) (*sql.DB, error) {
	// db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	// if err != nil {
	// 	return nil, err
	// }

	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	// mysql-education 3306 root 000000 go
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
