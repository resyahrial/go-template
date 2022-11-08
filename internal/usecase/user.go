package usecase

import (
	"context"

	"github.com/resyahrial/go-template/internal/entities"
)

type UserUsecaseImpl struct{}

func NewUserUsecase() entities.UserUsecase {
	return &UserUsecaseImpl{}
}

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, req entities.CreateUserRequest) (user *entities.User, err error) {
	return
}

func (u *UserUsecaseImpl) GetByEmail(ctx context.Context, email string) (user *entities.User, err error) {
	return
}
