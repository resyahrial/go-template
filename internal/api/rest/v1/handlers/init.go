package v1

import (
	"github.com/resyahrial/go-template/config"
	"github.com/resyahrial/go-template/internal/entities"
	"github.com/resyahrial/go-template/internal/factory"
	"gorm.io/gorm"
)

type Handler struct {
	userUsecase entities.UserUsecase
}

func NewHandler(db *gorm.DB, cfg config.Config) *Handler {
	return &Handler{
		userUsecase: factory.InitUserUsecase(db),
	}
}
