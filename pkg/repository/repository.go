package repository

// import "github.com/jmoiron/sqlx"
import (
	// "database/sql"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// func NewRepository(db *sqlx.DB) *Repository {
func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
