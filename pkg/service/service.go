package service

import (
	"TODO-APP/pkg/repository"

	"github.com/AntonZatsepilin/todo-app/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
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
