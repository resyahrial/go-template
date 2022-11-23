package handler

import (
	middlewares "github.com/resyahrial/go-template/internal/api/rest/middleware"

	// response "github.com/resyahrial/go-template/internal/api/rest/v1/responses"
	"github.com/resyahrial/go-template/internal/entities"
)

func (h *Handler) CreateUser(c ContextHandler) {
	var (
		err  error
		user *entities.User
	)

	if user, err = h.converter.GetCreateUserRequest(c); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if user, err = h.userUsecase.CreateUser(c, user); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	c.Set(middlewares.SuccessKey, user)
}
