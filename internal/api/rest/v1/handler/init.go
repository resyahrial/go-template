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
	userUsecase entities.UserUsecase,
) *Handler {
	return &Handler{
		userUsecase: userUsecase,
	}
}
