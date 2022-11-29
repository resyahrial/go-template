package handler

import (
	"github.com/resyahrial/go-template/internal/entity"
)

func (h *Handler) CreateUser(c Context) (res interface{}, err error) {
	var (
		user *entity.User
	)

	if user, err = h.reqConverter.GetCreateUserRequest(c.BindJSON); err != nil {
		return
	}

	if user, err = h.userUsecase.CreateUser(c, user); err != nil {
		return
	}

	return h.resConverter.GetCreateUserResponse(user)
}
