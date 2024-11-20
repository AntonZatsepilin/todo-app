package service

import (
	"github.com/AntonZatsepilin/todo-app/internal/models"
	"github.com/AntonZatsepilin/todo-app/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item models.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]models.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return nil, err
	}

	return s.repo.GetAll(listId)
}
