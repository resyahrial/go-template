package server

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/config"
	"github.com/resyahrial/go-template/internal/api/rest/middlewares"
	route "github.com/resyahrial/go-template/internal/api/rest/routes"
	"gorm.io/gorm"
)

func InitServer(db *gorm.DB, cfg config.Config) *gin.Engine {
	var (
		ginMode string
	)

	if cfg.App.DebugMode {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
		gin.DisableConsoleColor()
	}

	gin.SetMode(ginMode)
	customMiddleware := middlewares.New(middlewares.Opts{})

	engine := gin.Default()

	engine.Use(customMiddleware.ResponseWrapper)

	engine.Use(gin.CustomRecovery((func(c *gin.Context, recovered interface{}) {
		c.Set(middlewares.FailureKey, fmt.Errorf("panic : %v", recovered))
	})))

	engine.NoRoute(func(c *gin.Context) {
		c.Set(middlewares.FailureKey, errors.New("route not found"))
	})

	route.InitRoutes(engine, route.RouteOpt{
		Db:  db,
		Cfg: cfg,
	})

	return engine
}
