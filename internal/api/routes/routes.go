package route

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/resyahrial/go-template/internal/api/routes/v1"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	v1Path := r.Group("/api/v1")
	{
		v1.HealthCheckRoute(v1Path)
	}

	return r
}
