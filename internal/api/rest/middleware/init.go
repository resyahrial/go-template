package middleware

import (
	"github.com/resyahrial/go-template/config"
	"gorm.io/gorm"
)

type Middleware struct {
	db  *gorm.DB
	cfg config.Config
}

type Option func(*Middleware)

func New(opts ...Option) *Middleware {
	m := &Middleware{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func WithDbInstance(db *gorm.DB) Option {
	return func(m *Middleware) {
		m.db = db
	}
}

func WithConfig(cfg config.Config) Option {
	return func(m *Middleware) {
		m.cfg = cfg
	}
}
