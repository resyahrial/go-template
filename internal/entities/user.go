package entities

import "context"

type User struct {
	Id    string
	Name  string
	Email string
	Account
}

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type UserUsecase interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (user *User, err error)
	GetByEmail(ctx context.Context, email string) (user *User, err error)
}
