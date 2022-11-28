package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinOption func(*gin.Engine)

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
