package handler

import (
	"github.com/resyahrial/go-template/internal/entities"
)

type Handler struct {
	converter   BodyConverterAdapter
	userUsecase entities.UserUsecase
}

func NewHandler(
	converter BodyConverterAdapter,
	userUsecase entities.UserUsecase,
) *Handler {
	return &Handler{
		converter,
		userUsecase,
	}
}
