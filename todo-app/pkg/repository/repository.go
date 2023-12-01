package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type User interface {
	CreateUser(user todo.User) (int, error)
	GetUserById(UserId todo.IdRequest) (todo.User, error)
}
type Repository struct {
	Authorization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
	}
}
