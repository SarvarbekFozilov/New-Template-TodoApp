package service

import (
	"github.com/zhashkevych/todo-app"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(userId int, list todo.User) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *UserService) GetById(userId, listId int) (todo.User, error) {
	return s.repo.GetById(userId, listId)
}
