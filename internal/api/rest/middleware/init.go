package middleware

import (
	"github.com/resyahrial/go-template/config"
	"gorm.io/gorm"
)

type Middleware struct {
}

type Opts struct {
	Db  *gorm.DB
	Cfg config.Config
}

type MiddlewareOptionFn func(*Middleware, Opts)

func New(mOpts Opts, opts ...MiddlewareOptionFn) *Middleware {
	m := &Middleware{}
	for _, opt := range opts {
		opt(m, mOpts)
	}
	return m
}
