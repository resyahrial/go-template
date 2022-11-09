package v1

import (
	"github.com/gin-gonic/gin"
	handler "github.com/resyahrial/go-template/internal/api/handlers/v1"
)

func HealthCheckRoute(router *gin.RouterGroup) {
	router.GET("/health-check", func(c *gin.Context) {
		handler.HealthCheckHandler(c)
	})
}
