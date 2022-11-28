package handler

import (
	"time"

	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	"github.com/resyahrial/go-template/internal/entity"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go RequestConverter
type RequestConverter interface {
	GetCreateUserRequest(c request.Context) (user *entity.User, err error)
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
	// request.Context
	BindJSON(obj any) error
}
