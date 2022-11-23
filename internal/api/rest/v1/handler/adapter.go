package handler

import (
	"time"

	"github.com/resyahrial/go-template/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go RequestContext
type RequestContext interface {
	BindJSON(obj any) error
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go RequestConverter
type RequestConverter interface {
	GetCreateUserRequest(c RequestContext) (user *entities.User, err error)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go ResponseContext
type ResponseContext interface {
	Set(key string, obj any)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go ResponseConverter
type ResponseConverter interface {
	SetCreateUserResponse(c ResponseContext, user *entities.User)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go ContextHandler
type ContextHandler interface {
	// context.Context
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
	// embed method interface
	RequestContext
	ResponseContext
}
