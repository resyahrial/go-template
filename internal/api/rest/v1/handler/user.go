package handler

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/resyahrial/go-template/internal/api/rest/v1"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	res, err := h.createUser(ctx)
	h.wrapHandler(ctx, res, err)
}

func (h *Handler) createUser(ctx *gin.Context) (res interface{}, err error) {
	user, err := h.reqConverter.GetCreateUserRequest(ctx.BindJSON)
	if err != nil {
		return
	}

	if user, err = h.userUsecase.CreateUser(ctx, user); err != nil {
		return
	}

	return h.resConverter.GetCreateUserResponse(user)
}

func (h *Handler) GetUserByID(c *gin.Context, userID v1.UserID) {
}
