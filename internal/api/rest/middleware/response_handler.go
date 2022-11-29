package middleware

import (
	"net/http"

	"github.com/resyahrial/go-template/pkg/exception"
)

const (
	ResultKey    = "ResultKey"
	SuccessKey   = "SuccessKey"
	FailureKey   = "FailureKey"
	PaginatedKey = "PaginatedKey"
)

type PaginatedResultValue struct {
	Page  int
	Limit int
	Count int64
}

type Success struct {
	Data     interface{} `json:"data"`
	PageInfo interface{} `json:"pageInfo,omitempty"`
}

type PageInfo struct {
	CurrentPage int   `json:"currentPage,omitempty"`
	TotalPage   int   `json:"totalPage,omitempty"`
	Count       int64 `json:"count,omitempty"`
}

type Failure struct {
	ErrorMsg interface{} `json:"error"`
}

func (m *Middleware) ResponseHandler(ctx Context) {
	ctx.Next()

	val, ok := ctx.Get(ResultKey)
	if !ok {
		return
	}

	if err, ok := val.(error); ok {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, val)
}

func handleError(ctx Context, err error) {
	var (
		code                = http.StatusInternalServerError
		message interface{} = map[string]interface{}{
			"message": err.Error(),
		}
	)

	switch typeErr := err.(type) {
	case *exception.Base:
		typeErr.LogError()
		code = typeErr.Code
		if typeErr.CollectionMessage != nil && len(typeErr.CollectionMessage) > 0 {
			message = typeErr.CollectionMessage
		}
	}

	ctx.JSON(code, &Failure{
		ErrorMsg: message,
	})
}
