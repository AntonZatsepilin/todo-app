package repository

import (
	"github.com/AntonZatsepilin/todo-app/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
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
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
