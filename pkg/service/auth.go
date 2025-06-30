package service

import (
	"github.com/kravz2702-byte/todo-app/pkg/entity"
	"github.com/kravz2702-byte/todo-app/pkg/repository"
)

type AuthServise struct {
	repo repository.Repository
}

func NewAuthServise(repo repository.Repository) *AuthServise {
	return &AuthServise{repo: repo}
}

func (s *AuthServise) CreateService(user entity.User) (int, error) {
	return s.repo.CreateUser(user)
}
