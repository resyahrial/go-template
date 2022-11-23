package handler

import (
	"time"

	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go RequestConverter
type RequestConverter interface {
	GetCreateUserRequest(c request.Context) (user *entities.User, err error)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go ResponseConverter
type ResponseConverter interface {
	SetCreateUserResponse(c response.Context, user *entities.User)
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
	// response.Context
	Set(key string, obj any)
}
