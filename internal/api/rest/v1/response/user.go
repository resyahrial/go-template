package response

import (
	"github.com/resyahrial/go-template/internal/entities"
)

type CreateUser struct {
	Name  string
	Email string
}

func (e *Converter) SetCreateUserResponse(ctx Context, user *entities.User) (err error) {
	var (
		res *CreateUser
	)

	if err = e.decoder.Decode(user, &res); err != nil {
		return
	}

	ctx.Set(SuccessKey, res)
	return nil
}
