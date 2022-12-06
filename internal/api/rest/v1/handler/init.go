package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	v1 "github.com/resyahrial/go-template/internal/api/rest/v1"
	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/internal/entity"
	"github.com/resyahrial/go-template/internal/factory"
	"github.com/resyahrial/go-template/pkg/validator"
	"gorm.io/gorm"
)

type option struct {
	Db *gorm.DB
}

type Option func(*option)

type Handler struct {
	reqConverter RequestConverter
	resConverter ResponseConverter
	userUsecase  entity.UserUsecase
}

func (h *Handler) wrapHandler(ctx *gin.Context, res interface{}, err error) {
	if err != nil {
		ctx.Set(middleware.ResultKey, err)
	} else {
		ctx.Set(middleware.ResultKey, res)
	}
}

func New(opts ...Option) func(*gin.Engine) {
	return func(engine *gin.Engine) {
		opt := &option{}
		for _, o := range opts {
			o(opt)
		}

		handlerImpl := &Handler{
			reqConverter: request.NewConverter(
				&request.ValidatorImpl{Fn: validator.Validate},
				&request.DecoderImpl{Fn: mapstructure.Decode},
			),
			resConverter: response.NewConverter(
				&response.DecoderImpl{Fn: mapstructure.Decode},
			),
			userUsecase: factory.InitUserUsecase(opt.Db),
		}

		v1.RegisterHandlersWithOptions(
			engine,
			handlerImpl,
			v1.GinServerOptions{
				BaseURL:     "/api/v1",
				Middlewares: []v1.MiddlewareFunc{},
				ErrorHandler: func(ctx *gin.Context, err error, code int) {
					ctx.JSON(code, gin.H{"error": err.Error()})
				},
			},
		)
	}
}

func WithGorm(db *gorm.DB) Option {
	return func(o *option) {
		o.Db = db
	}
}
