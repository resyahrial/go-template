package handler

import (
	"github.com/resyahrial/go-template/internal/entity"
)

func (h *Handler) CreateUser(c Context) (err error) {
	var (
		user *entity.User
	)

	if user, err = h.reqConverter.GetCreateUserRequest(c); err != nil {
		return
	}

	if user, err = h.userUsecase.CreateUser(c, user); err != nil {
		return
	}

	if err = h.resConverter.SetCreateUserResponse(c, user); err != nil {
		return
	}

	return
}
