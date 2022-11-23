package response

import (
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
	"github.com/resyahrial/go-template/internal/entities"
)

type CreateUser struct {
	Name  string
	Email string
}

func (e *Converter) SetCreateUserResponse(c handler.ResponseContext, user *entities.User) (err error) {
	var (
		res *CreateUser
	)

	if err = e.decoder.Decode(user, &res); err != nil {
		return
	}

	c.Set(middleware.SuccessKey, res)
	return nil
}
