package entity

import (
	"context"
)

const (
	UserModule = "USER"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

//go:generate mockgen -destination=mocks/user_mock.go -source=user.go UserUsecase
type UserUsecase interface {
	CreateUser(ctx context.Context, input *User) (user *User, err error)
}
