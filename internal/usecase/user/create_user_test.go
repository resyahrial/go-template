package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-template/internal/entity"
	usecase "github.com/resyahrial/go-template/internal/usecase/user"
	adapter_mock "github.com/resyahrial/go-template/internal/usecase/user/mocks"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type CreateUserUsecaseTestSuite struct {
	suite.Suite
	userRepo *adapter_mock.MockUserRepo
	ucase    entity.UserUsecase
}

func TestCreateUserUsecase(t *testing.T) {
	suite.Run(t, new(CreateUserUsecaseTestSuite))
}

func (s *CreateUserUsecaseTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = adapter_mock.NewMockUserRepo(ctrl)
	s.ucase = usecase.NewUserUsecase(
		s.userRepo,
	)
}

func (s *CreateUserUsecaseTestSuite) TestCreateUser() {
	userId := ksuid.New().String()
	input := &entity.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name                 string
		resultMockCreateUser *entity.User
		errorMockCreateUser  error
		expectedOutput       *entity.User
		expectedError        error
	}{
		{
			name: "should create user",
			resultMockCreateUser: &entity.User{
				Id:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: "password",
			},
			expectedOutput: &entity.User{
				Id:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: "password",
			},
		},
		{
			name:                "should raise error when failed persist user",
			errorMockCreateUser: errors.New("failed persist user"),
			expectedError:       errors.New("failed persist user"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.userRepo.EXPECT().Create(gomock.Any(), input).Return(tc.resultMockCreateUser, tc.errorMockCreateUser)

			res, err := s.ucase.CreateUser(context.Background(), input)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.EqualValues(tc.expectedOutput, res)
			} else {
				s.Nil(res)
			}
		})
	}
}
