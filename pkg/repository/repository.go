package repository

// import "github.com/jmoiron/sqlx"
import (
	// "database/sql"

	"database/sql"

	"github.com/Ckala62rus/rest-go"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GetUser(username, password string) (rest.User, error)
}

type TodoList interface {
	Create(userId int, list rest.TodoList) (int, error)
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
	return &Repository{
		// Authorization: NewAuthPostgres(db),
		Authorization: NewAuthMysql(db),
		TodoList:      NewTodoListMysql(db),
	}
}
