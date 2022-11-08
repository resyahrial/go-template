package usecase

import (
	"context"

	"github.com/resyahrial/go-template/internal/entities"
)

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, req entities.CreateUserRequest) (user *entities.User, err error) {
	if user, err = entities.NewUser(req); err != nil {
		return
	}

	return u.UserRepo.Create(ctx, user)
}

func (u *UserUsecaseImpl) GetByEmail(ctx context.Context, email string) (user *entities.User, err error) {
	return
}
