package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ResultKey = "ResultKey"
)

type GinOption func(*gin.Engine)

type HandlerFn func(*gin.Context) (interface{}, error)
type GinRoute struct {
	Route
	HandlerFn HandlerFn
}

func defaultSetup() *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.CustomRecovery((func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("panic : %v", recovered),
		})
	})))

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "route not found",
		})
	})

	engine.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})

	return engine
}

func InitGinEngine(mode string, opts ...GinOption) *gin.Engine {
	if mode != "" {
		gin.SetMode(mode)
		if mode == gin.ReleaseMode {
			gin.DisableConsoleColor()
		}
	}
	engine := defaultSetup()
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

func WithCustomMiddlewares(middlewares ...gin.HandlerFunc) GinOption {
	return func(e *gin.Engine) {
		for _, middleware := range middlewares {
			e.Use(middleware)
		}
	}
}

func WithRoutes(routes ...GinRoute) GinOption {
	return func(e *gin.Engine) {
		fn := func(handler HandlerFn) gin.HandlerFunc {
			return func(ctx *gin.Context) {
				res, err := handler(ctx)
				if err != nil {
					ctx.Set(ResultKey, err)
				} else {
					ctx.Set(ResultKey, res)
				}
			}
		}

		for _, route := range routes {
			handler := fn(route.HandlerFn)
			switch route.Method {
			case http.MethodPost:
				e.POST(route.Path, handler)
			case http.MethodDelete:
				e.DELETE(route.Path, handler)
			case http.MethodPatch:
				e.PATCH(route.Path, handler)
			case http.MethodPut:
				e.PUT(route.Path, handler)
			default:
				e.GET(route.Path, handler)
			}
		}
	}
}

func WithDefaultResponseWrapper(routes ...GinRoute) GinOption {
	return func(e *gin.Engine) {
		e.Use(func(ctx *gin.Context) {
			ctx.Next()

			val, ok := ctx.Get(ResultKey)
			if !ok {
				return
			}
			if err, ok := val.(error); ok {
				ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"data": val,
			})
		})
	}
}
