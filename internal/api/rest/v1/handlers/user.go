package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/internal/api/rest/middlewares"
	request "github.com/resyahrial/go-template/internal/api/rest/v1/requests"

	// response "github.com/resyahrial/go-template/internal/api/rest/v1/responses"
	"github.com/resyahrial/go-template/internal/entities"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var (
		err  error
		req  *request.CreateUserRequest
		user *entities.User
	)

	if err = c.BindJSON(&req); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if user, err = req.CastToEntity(); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	if user, err = h.userUsecase.CreateUser(c.Request.Context(), user); err != nil {
		c.Set(middlewares.FailureKey, err)
		return
	}

	c.Set(middlewares.SuccessKey, user)
}
