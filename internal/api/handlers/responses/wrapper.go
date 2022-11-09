package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	// 	case errors.InternalError:
	// 		return generateFailure(http.StatusInternalServerError, typeErr.BaseError, typeErr.Cause.Error())
	// 	case errors.BadRequestError:
	// 		return generateFailure(http.StatusBadRequest, typeErr.BaseError, typeErr.Cause.Error())
	// 	case errors.UnprocessableEntityError:
	// 		return generateFailure(http.StatusUnprocessableEntity, typeErr.BaseError, typeErr.Cause.Error())
	// 	case errors.UnauthorizedError:
	// 		return generateFailure(http.StatusUnauthorized, typeErr.BaseError, typeErr.Cause.Error())
	// 	case errors.TooManyRequestError:
	// 		return generateFailure(http.StatusTooManyRequests, typeErr.BaseError, typeErr.Cause.Error())
	// 	case errors.RequestTimeoutError:
	// 		return generateFailure(http.StatusRequestTimeout, typeErr.BaseError, typeErr.Cause.Error())
	// 	case errors.NotFoundError:
	// 		return generateFailure(http.StatusNotFound, typeErr.BaseError, typeErr.Cause.Error())
	default:
		// unexpectedErr := errors.NewInternalError(ginErr.Err, "unexpected error")
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
