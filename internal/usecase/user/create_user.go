package usecase

import (
	"context"

	"github.com/resyahrial/go-template/internal/entity"
)

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, input *entity.User) (user *entity.User, err error) {
	return u.UserRepo.Create(ctx, input)
}
