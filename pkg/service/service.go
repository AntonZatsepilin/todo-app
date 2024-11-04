package service

import (
	"github.com/AntonZatsepilin/todo-app/internal/models"
	"github.com/AntonZatsepilin/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
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
