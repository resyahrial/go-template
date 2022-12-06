package response

import (
	v1 "github.com/resyahrial/go-template/internal/api/rest/v1"
	"github.com/resyahrial/go-template/internal/entity"
)

func (e *Converter) GetCreateUserResponse(user *entity.User) (res interface{}, err error) {
	var (
		userModel *v1.User
	)

	if err = e.decoder.Decode(user, &userModel); err != nil {
		return
	}

	return WrapSingleData(userModel), nil
}
