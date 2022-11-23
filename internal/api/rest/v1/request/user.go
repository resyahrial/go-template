package request

import (
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
	"github.com/resyahrial/go-template/internal/entities"
	"github.com/resyahrial/go-template/pkg/exception"
)

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

func (e *Converter) GetCreateUserRequest(binder handler.JsonRequestBinderAdapater) (user *entities.User, err error) {
	var (
		req *CreateUser
	)

	if err = binder.BindJSON(&req); err != nil {
		return
	}

	if mapErr := e.validator.Validate(req); len(mapErr) > 0 {
		err = exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(mapErr)
		return
	}

	if err = e.decoder.Decode(req, &user); err != nil {
		return
	}

	return
}
