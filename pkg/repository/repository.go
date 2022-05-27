package repository

// import "github.com/jmoiron/sqlx"
import (
	// "database/sql"

	"github.com/Ckala62rus/rest-go"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
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

func NewRepository(db *sqlx.DB) *Repository {
	// func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
