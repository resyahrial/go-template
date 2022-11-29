package route

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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

type R struct {
	engine  *gin.Engine
	handler *handler.Handler
}

func InitRoutes(opts ...Option) func(*gin.Engine) {
	return func(engine *gin.Engine) {
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

		handler := handler.NewHandler(
			reqConverter,
			resConverter,
			factory.InitUserUsecase(opt.Db),
		)

		registerRoute(engine, handler)
	}
}

func WithGorm(db *gorm.DB) Option {
	return func(o *option) {
		o.Db = db
	}
}

func registerRoute(e *gin.Engine, h *handler.Handler) {
	r := R{e, h}
	methodFinder := reflect.TypeOf(&r)
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		method.Func.Call([]reflect.Value{reflect.ValueOf(&r)})
	}
}

func wrapHandler(fn func(handler.Context) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := fn(ctx)
		if err != nil {
			ctx.Set("Result", err)
		} else {
			ctx.Set("Result", res)
		}
	}
}
