package service

import (
	"github.com/Ckala62rus/rest-go"
	"github.com/Ckala62rus/rest-go/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list rest.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]rest.TodoList, error) {
	return s.repo.GetAll(userId)
}
