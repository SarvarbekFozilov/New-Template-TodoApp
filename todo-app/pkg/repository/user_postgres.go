package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUserById(UserId todo.IdRequest) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT name, username, password_hash FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, UserId.Id)

	return user, err
}
