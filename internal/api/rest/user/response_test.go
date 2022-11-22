package user_test

import (
	"testing"

	"github.com/resyahrial/go-template/internal/api/rest/user"
	"github.com/resyahrial/go-template/internal/entities"
	"github.com/stretchr/testify/suite"
)

type CreateUserResponseTestSuite struct {
	suite.Suite
}

func TestCreateUserResponse(t *testing.T) {
	suite.Run(t, new(CreateUserResponseTestSuite))
}

func (s *CreateUserResponseTestSuite) SetupTest() {
}

func (s *CreateUserResponseTestSuite) TestConvertToUserEntity() {
	entity := &entities.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name           string
		expectedOutput *user.CreateUserResponse
		expectedError  error
	}{
		{
			name: "should create basic user",
			expectedOutput: &user.CreateUserResponse{
				Name:  "user",
				Email: "user@mail.com",
			},
		},
	}

	for _, tc := range testCases {
		res, err := user.NewCreateUserResponse(entity)
		s.Run(tc.name, func() {
			s.Equal(tc.expectedError, err)
			s.EqualValues(tc.expectedOutput, res)
		})
	}
}
