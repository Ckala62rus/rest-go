package service

import (
	"github.com/Ckala62rus/rest-go"
	"github.com/Ckala62rus/rest-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list rest.TodoList) (int, error)
	GetAll(userId int) ([]rest.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
