package entities

import (
	"context"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type UserOption func(*User) error

func NewUser(input CreateUserRequest, opts ...UserOption) (user *User, errs []error) {
	if err := mapstructure.Decode(input, &user); err != nil {
		errs = append(errs, err)
		return
	}

	for _, opt := range opts {
		if err := opt(user); err != nil {
			errs = append(errs, err)
			return
		}
	}

	return
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
