package model

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-template/internal/entity"
)

type User struct {
	CommonField
	Name     string
	Email    string
	Password string
}

func NewUserModel(userEntity *entity.User) (user *User, err error) {
	if err = mapstructure.Decode(userEntity, &user); err != nil {
		return
	}
	return
}

func (u *User) ConvertToEntity() (userEntity *entity.User, err error) {
	if err = mapstructure.Decode(u, &userEntity); err != nil {
		return
	}
	userEntity.Id = u.Id
	return
}
