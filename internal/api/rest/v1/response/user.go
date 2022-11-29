package response

import (
	"github.com/resyahrial/go-template/internal/entity"
)

type CreateUser struct {
	Name  string
	Email string
}

func (e *Converter) GetCreateUserResponse(user *entity.User) (res interface{}, err error) {
	var (
		createUserRes *CreateUser
	)

	if err = e.decoder.Decode(user, &createUserRes); err != nil {
		return
	}

	return WrapSingleData(createUserRes), err
}
