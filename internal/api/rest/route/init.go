package route

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/internal/factory"
	"github.com/resyahrial/go-template/pkg/validator"
	"gorm.io/gorm"
)

type option struct {
	Db *gorm.DB
}

type Option func(*option)

func InitRoutes(e *gin.Engine, opts ...Option) {
	opt := &option{}
	for _, o := range opts {
		o(opt)
	}

	reqConverter := request.NewConverter(
		&request.ValidatorImpl{
			Fn: validator.Validate,
		},
		&request.DecoderImpl{
			Fn: mapstructure.Decode,
		},
	)

	resConverter := response.NewConverter(
		&response.DecoderImpl{
			Fn: mapstructure.Decode,
		},
	)

	initV1Route(e, handler.NewHandler(
		reqConverter,
		resConverter,
		factory.InitUserUsecase(opt.Db),
	))
}

func WithGorm(db *gorm.DB) Option {
	return func(o *option) {
		o.Db = db
	}
}

type handlerFn func(handler.Context) error

func WrapHandler(fn handlerFn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := fn(ctx); err != nil {
			ctx.Set(middleware.FailureKey, err)
		}
	}
}
