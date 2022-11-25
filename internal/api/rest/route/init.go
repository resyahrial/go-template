package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-template/config"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/internal/factory"
	"github.com/resyahrial/go-template/pkg/validator"
	"gorm.io/gorm"
)

type RouteOpt struct {
	Db  *gorm.DB
	Cfg config.Config
}

func InitRoutes(e *gin.Engine, opt RouteOpt) {
	e.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})

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

type HandlerFn func(handler.Context) error

func WrapHandler(fn HandlerFn) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := fn(ctx); err != nil {
			ctx.Set(response.FailureKey, err)
		}
	}
}
