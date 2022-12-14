package usecase

import (
	"context"

	"github.com/resyahrial/go-template/internal/entity"
)

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go UserRepo
type UserRepo interface {
	Create(ctx context.Context, user *entity.User) (res *entity.User, err error)
}
