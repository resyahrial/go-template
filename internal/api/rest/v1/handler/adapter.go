package handler

import (
	"time"

	"github.com/resyahrial/go-template/internal/entity"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go RequestConverter
type RequestConverter interface {
	GetCreateUserRequest(fn func(obj any) error) (user *entity.User, err error)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go ResponseConverter
type ResponseConverter interface {
	GetCreateUserResponse(user *entity.User) (res interface{}, err error)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Context
type Context interface {
	// context.Context
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
	// RequestBinderFn
	BindJSON(obj any) error
}
