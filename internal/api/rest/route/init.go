package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/config"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
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

	initV1Route(e, handler.NewHandler(
		nil,
		nil,
		nil,
	))
}
