package model_test

import (
	"testing"

	"github.com/resyahrial/go-template/internal/entity"
	"github.com/resyahrial/go-template/internal/repo/postgresql/model"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUserModel(t *testing.T) {
	userEntity := &entity.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	user, err := model.NewUserModel(userEntity)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userEntity.Name, user.Name)
	assert.Equal(t, userEntity.Email, user.Email)
	assert.Equal(t, userEntity.Password, user.Password)
}

func TestConvertToEntityUser(t *testing.T) {
	user := &model.User{
		CommonField: model.CommonField{
			Id: ksuid.New().String(),
		},
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	userEntity, err := user.ConvertToEntity()
	assert.Nil(t, err)
	assert.NotNil(t, userEntity)
	assert.Equal(t, user.Id, userEntity.Id)
	assert.EqualValues(t, user.Name, userEntity.Name)
	assert.EqualValues(t, user.Email, userEntity.Email)
	assert.EqualValues(t, user.Password, userEntity.Password)
}
