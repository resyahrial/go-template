package server

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/config"
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	"github.com/resyahrial/go-template/internal/api/rest/route"
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
	customMiddleware := middleware.New(middleware.Opts{})

	engine := gin.Default()

	engine.Use(
		func(ctx *gin.Context) {
			customMiddleware.ResponseHandler(ctx)
		},
	)

	engine.Use(gin.CustomRecovery((func(c *gin.Context, recovered interface{}) {
		c.Set(middleware.FailureKey, fmt.Errorf("panic : %v", recovered))
	})))

	engine.NoRoute(func(c *gin.Context) {
		c.Set(middleware.FailureKey, errors.New("route not found"))
	})

	route.InitRoutes(engine, route.RouteOpt{
		Db:  db,
		Cfg: cfg,
	})

	return engine
}
