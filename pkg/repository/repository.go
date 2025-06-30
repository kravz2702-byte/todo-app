package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kravz2702-byte/todo-app/pkg/entity"
)

type Authorization interface {
	createUser(user entity.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
