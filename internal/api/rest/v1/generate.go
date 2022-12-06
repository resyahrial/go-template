package v1

import (
	"fmt"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/pkg/logger"
)

//go:generate oapi-codegen -config types.cfg.yml docs.yml
//go:generate oapi-codegen -config server.cfg.yml docs.yml

func WithSwagger() func(*gin.Engine) {
	return func(e *gin.Engine) {
		swagger, err := GetSwagger()
		if err != nil {
			logger.Fatal(fmt.Sprintf("Error loading swagger spec\n: %s", err))
			return
		}
		swagger.Servers = nil
		e.Use(middleware.OapiRequestValidator(swagger))
	}
}
