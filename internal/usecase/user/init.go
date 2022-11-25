package usecase

import "github.com/resyahrial/go-template/internal/entity"

type UserUsecaseImpl struct {
	UserRepo
}

func NewUserUsecase(
	userRepo UserRepo,
) entity.UserUsecase {
	return &UserUsecaseImpl{
		UserRepo: userRepo,
	}
}
