package middleware

import (
	"github.com/resyahrial/go-template/config"
	"gorm.io/gorm"
)

type Middleware struct {
}

type MiddlewareOpts struct {
	Db  *gorm.DB
	Cfg config.Config
}

type MiddlewareOptionFn func(*Middleware, MiddlewareOpts)

func NewMiddleware(mOpts MiddlewareOpts, opts ...MiddlewareOptionFn) *Middleware {
	m := &Middleware{}
	for _, opt := range opts {
		opt(m, mOpts)
	}
	return m
}
