package request

import (
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-template/internal/entities"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

func (r *CreateUserRequest) CastToUserEntity() (user *entities.User, err error) {
	if err = mapstructure.Decode(r, &user); err != nil {
		return
	}
	return
}
