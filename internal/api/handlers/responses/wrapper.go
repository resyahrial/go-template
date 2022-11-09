package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/pkg/exception"
	// "github.com/kargotech/ltms-shipper-integration/internal/pkg/errors"
)

type Success struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
}

func HandleSuccess(data interface{}) *Success {
	return &Success{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

type Failure struct {
	StatusCode int         `json:"-"`
	Error      interface{} `json:"error"`
}

func HandleError(c *gin.Context, err error) *Failure {
	ginErr := c.Error(err)

	switch typeErr := ginErr.Err.(type) {
	case *exception.Base:
		typeErr.LogError()
		return generateFailure(typeErr.Code, typeErr)
	default:
		return generateFailure(http.StatusInternalServerError, typeErr)
	}
}

func generateFailure(statusCode int, err error) *Failure {
	return &Failure{
		StatusCode: statusCode,
		Error: map[string]interface{}{
			"message": err.Error(),
		},
	}
}
