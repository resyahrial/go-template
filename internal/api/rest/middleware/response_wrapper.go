package middleware

import (
	"net/http"

	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/pkg/exception"
)

type success struct {
	Data     interface{} `json:"data"`
	PageInfo interface{} `json:"pageInfo,omitempty"`
}

type pageInfo struct {
	CurrentPage int   `json:"currentPage,omitempty"`
	TotalPage   int   `json:"totalPage,omitempty"`
	Count       int64 `json:"count,omitempty"`
}

type failure struct {
	ErrorMsg interface{} `json:"error"`
}

func (m *Middleware) ResponseWrapper(ctx Context) {
	ctx.Next()

	if val, ok := ctx.Get(response.FailureKey); ok {
		if err, ok := val.(error); ok {
			handleError(ctx, err)
			return
		}
	}

	if data, ok := ctx.Get(response.SuccessKey); ok {
		if paginatedData, ok := ctx.Get(response.PaginatedKey); ok {
			if parsedPaginatedData, ok := paginatedData.(response.PaginatedResultValue); ok {
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

	ctx.AbortWithStatusJSON(code, &failure{
		ErrorMsg: message,
	})
}

func handleSuccess(ctx Context, data interface{}) {
	ctx.JSON(http.StatusOK, &success{
		Data: data,
	})
}

func handleSuccessPaginated(ctx Context, data interface{}, paginatedResultValue response.PaginatedResultValue) {
	totalPage := 1
	if paginatedResultValue.Limit < int(paginatedResultValue.Count) {
		addtional := int(paginatedResultValue.Count) / paginatedResultValue.Limit
		if int(paginatedResultValue.Count)%paginatedResultValue.Limit == 0 {
			addtional -= 1
		}
		totalPage += addtional
	}
	ctx.JSON(http.StatusOK, &success{
		Data: data,
		PageInfo: pageInfo{
			CurrentPage: paginatedResultValue.Page + 1,
			TotalPage:   totalPage,
			Count:       paginatedResultValue.Count,
		},
	})
}
