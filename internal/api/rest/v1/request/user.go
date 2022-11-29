package request

import (
	"github.com/resyahrial/go-template/internal/entity"
	"github.com/resyahrial/go-template/pkg/exception"
)

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

func (e *Converter) GetCreateUserRequest(fn func(obj any) error) (user *entity.User, err error) {
	var (
		req *CreateUser
	)

	if err = fn(&req); err != nil {
		return
	}

	if mapErr := e.validator.Validate(req); len(mapErr) > 0 {
		err = exception.NewBadRequestException().SetModule(entity.UserModule).SetCollectionMessage(mapErr)
		return
	}

	if err = e.decoder.Decode(req, &user); err != nil {
		return
	}

	return
}
