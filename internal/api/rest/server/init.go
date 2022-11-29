package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
)

type Option func(*gin.Engine)

func InitServerDebugMode(opts ...Option) *gin.Engine {
	return initServer(opts...)
}

func InitServerReleaseMode(opts ...Option) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	return initServer(opts...)
}

func InitServerTestMode(opts ...Option) *gin.Engine {
	gin.SetMode(gin.TestMode)
	return initServer(opts...)
}

func initServer(opts ...Option) *gin.Engine {
	engine := gin.Default()
	customRecovery := gin.CustomRecovery((func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("panic : %v", recovered),
		})
	}))

	customMiddleware := middleware.New()
	engine.Use(
		customRecovery,
		func(ctx *gin.Context) { customMiddleware.ResponseHandler(ctx) },
	)

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "route not found",
		})
	})

	engine.GET("/health-check", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})

	for _, opt := range opts {
		opt(engine)
	}

	return engine
}
