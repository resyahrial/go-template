package route

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/internal/factory"
	"github.com/resyahrial/go-template/pkg/rest"
	"github.com/resyahrial/go-template/pkg/validator"
	"gorm.io/gorm"
)

type option struct {
	Db *gorm.DB
}

type Option func(*option)

var (
	routes []rest.GinRoute
)

type r struct {
	h *handler.Handler
}

func InitRoutes(opts ...Option) []rest.GinRoute {
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

	registerRoute(r{}, handler)
	return routes
}

func WithGorm(db *gorm.DB) Option {
	return func(o *option) {
		o.Db = db
	}
}

func registerRoute(r r, h *handler.Handler) {
	r.h = h
	methodFinder := reflect.TypeOf(&r)
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		method.Func.Call([]reflect.Value{reflect.ValueOf(&r)})
	}
}

func addRoute(method string, path string, fn func(handler.Context) (interface{}, error)) rest.GinRoute {
	return rest.GinRoute{
		Route: rest.Route{
			Method: method,
			Path:   path,
		},
		HandlerFn: func(ctx *gin.Context) (interface{}, error) { return fn(ctx) },
	}
}
