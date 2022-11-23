package handler

import (
	"github.com/resyahrial/go-template/internal/entities"
)

type Handler struct {
	reqConverter RequestConverter
	resConverter ResponseConverter
	userUsecase  entities.UserUsecase
}

func NewHandler(
	reqConverter RequestConverter,
	resConverter ResponseConverter,
	userUsecase entities.UserUsecase,
) *Handler {
	return &Handler{
		reqConverter,
		resConverter,
		userUsecase,
	}
}
