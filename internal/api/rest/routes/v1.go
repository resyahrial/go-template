package route

import (
	"github.com/gin-gonic/gin"
	handler "github.com/resyahrial/go-template/internal/api/rest/v1/handlers"
)

func initV1Route(e *gin.Engine, h *handler.Handler) {
	r := e.Group("v1")
	r.POST("", h.CreateUser)
}
