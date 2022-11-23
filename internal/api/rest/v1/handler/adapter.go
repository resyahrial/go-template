package handler

import (
	"context"

	"github.com/resyahrial/go-template/internal/entities"
)

//go:generate mockgen -destination=mocks/mock.go -source=init.go JsonRequestBinderAdapater
type JsonRequestBinderAdapater interface {
	BindJSON(obj any) error
}

//go:generate mockgen -destination=mocks/mock.go -source=init.go BodyConverterAdapter
type BodyConverterAdapter interface {
	GetCreateUserRequest(c JsonRequestBinderAdapater) (user *entities.User, err error)
}

type ContextHandlerAdapter interface {
	context.Context
	JsonRequestBinderAdapater
	Set(key string, obj any)
}
