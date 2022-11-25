package middleware

import (
	"net/http"

	"github.com/resyahrial/go-template/pkg/exception"
)

const (
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

func (m *Middleware) ResponseWrapper(ctx Context) {
	ctx.Next()

	if val, ok := ctx.Get(FailureKey); ok {
		if err, ok := val.(error); ok {
			handleError(ctx, err)
			return
		}
	}

	if data, ok := ctx.Get(SuccessKey); ok {
		if paginatedData, ok := ctx.Get(PaginatedKey); ok {
			if parsedPaginatedData, ok := paginatedData.(PaginatedResultValue); ok {
				handleSuccessPaginated(ctx, data, parsedPaginatedData)
			}
		} else {
			handleSuccess(ctx, data)
		}
		return
	}
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

func handleSuccess(ctx Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Success{
		Data: data,
	})
}

func handleSuccessPaginated(ctx Context, data interface{}, paginatedResultValue PaginatedResultValue) {
	totalPage := 1
	if paginatedResultValue.Limit < int(paginatedResultValue.Count) {
		addtional := int(paginatedResultValue.Count) / paginatedResultValue.Limit
		if int(paginatedResultValue.Count)%paginatedResultValue.Limit == 0 {
			addtional -= 1
		}
		totalPage += addtional
	}
	ctx.JSON(http.StatusOK, &Success{
		Data: data,
		PageInfo: PageInfo{
			CurrentPage: paginatedResultValue.Page + 1,
			TotalPage:   totalPage,
			Count:       paginatedResultValue.Count,
		},
	})
}
