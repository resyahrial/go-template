package repo

import (
	"context"

	"github.com/resyahrial/go-template/internal/entity"
	"github.com/resyahrial/go-template/internal/repo/postgresql/model"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(
	db *gorm.DB,
) *UserRepoImpl {
	return &UserRepoImpl{
		db,
	}
}

func (u *UserRepoImpl) Create(ctx context.Context, user *entity.User) (res *entity.User, err error) {
	var (
		userModel *model.User
	)

	if userModel, err = model.NewUserModel(user); err != nil {
		return
	}

	userModel.Id = ksuid.New().String()
	if err = u.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return
	}

	return userModel.ConvertToEntity()
}
