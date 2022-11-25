package response

import (
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	"github.com/resyahrial/go-template/internal/entity"
)

type CreateUser struct {
	Name  string
	Email string
}

func (e *Converter) SetCreateUserResponse(ctx Context, user *entity.User) (err error) {
	var (
		res *CreateUser
	)

	if err = e.decoder.Decode(user, &res); err != nil {
		return
	}

	ctx.Set(middleware.SuccessKey, res)
	return nil
}
