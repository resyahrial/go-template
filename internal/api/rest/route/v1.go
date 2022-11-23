package route

import (
	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
)

func initV1Route(e *gin.Engine, h *handler.Handler) {
	r := e.Group("v1")
	r.POST("", func(ctx *gin.Context) { h.CreateUser(ctx) })
}
