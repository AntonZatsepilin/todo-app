package service

import (
	"github.com/AntonZatsepilin/todo-app/internal/models"
	"github.com/AntonZatsepilin/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
