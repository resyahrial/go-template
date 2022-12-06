package handler

import (
	"github.com/resyahrial/go-template/internal/entity"
)

type RequestConverter interface {
	GetCreateUserRequest(fn func(obj any) error) (user *entity.User, err error)
}

type ResponseConverter interface {
	GetCreateUserResponse(user *entity.User) (res interface{}, err error)
}
