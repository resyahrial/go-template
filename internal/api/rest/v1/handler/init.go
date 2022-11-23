package handler

import (
	"github.com/resyahrial/go-template/internal/entities"
)

type Handler struct {
	converter   RequestConverter
	userUsecase entities.UserUsecase
}

func NewHandler(
	converter RequestConverter,
	userUsecase entities.UserUsecase,
) *Handler {
	return &Handler{
		converter,
		userUsecase,
	}
}
