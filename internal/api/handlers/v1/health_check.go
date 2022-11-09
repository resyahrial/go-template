package v1

import (
	"github.com/gin-gonic/gin"
	response "github.com/resyahrial/go-template/internal/api/handlers/responses"
)

func HealthCheckHandler(c *gin.Context) {
	res := response.HandleSuccess(map[string]interface{}{
		"message": "OK",
	})
	c.JSON(res.StatusCode, res)
}
