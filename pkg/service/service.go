package service

import (
	"github.com/Ckala62rus/rest-go"
	"github.com/Ckala62rus/rest-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
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
	}
}
