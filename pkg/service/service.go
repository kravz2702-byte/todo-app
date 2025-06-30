package service

import (
	"github.com/kravz2702-byte/todo-app/pkg/entity"
	"github.com/kravz2702-byte/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
