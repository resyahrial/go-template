package handler

import (
	"github.com/resyahrial/go-template/internal/entity"
)

type Handler struct {
	reqConverter RequestConverter
	resConverter ResponseConverter
	userUsecase  entity.UserUsecase
}

func NewHandler(
	reqConverter RequestConverter,
	resConverter ResponseConverter,
	userUsecase entity.UserUsecase,
) *Handler {
	return &Handler{
		reqConverter,
		resConverter,
		userUsecase,
	}
}
