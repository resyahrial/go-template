package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinOption func(*gin.Engine)

func defaultSetup() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.Use(gin.CustomRecovery((func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Errorf("panic : %v", recovered),
		})
	})))

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "route not found",
		})
	})
	return engine
}

func InitGinEngine(opts ...GinOption) *gin.Engine {
	engine := defaultSetup()
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

func IsReleaseMode() GinOption {
	return func(e *gin.Engine) {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
}

func WithCustomMiddlewares(middlewares ...gin.HandlerFunc) GinOption {
	return func(e *gin.Engine) {
		for _, middleware := range middlewares {
			e.Use(middleware)
		}
	}
}
