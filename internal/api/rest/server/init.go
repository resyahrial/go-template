package server

import (
	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	"github.com/resyahrial/go-template/pkg/rest"
)

func InitServer(isDebugMode bool, routes []rest.GinRoute) *gin.Engine {
	var mode string
	if isDebugMode {
		mode = gin.ReleaseMode
	}
	customMiddleware := middleware.New()
	return rest.InitGinEngine(
		mode,
		rest.WithCustomMiddlewares(
			func(ctx *gin.Context) {
				customMiddleware.ResponseHandler(ctx)
			},
		),
		rest.WithRoutes(routes...),
	)
}
