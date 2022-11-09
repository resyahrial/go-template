package request_test

import (
	"testing"

	request "github.com/resyahrial/go-template/internal/api/handlers/requests"
	"github.com/resyahrial/go-template/internal/entities"
	"github.com/stretchr/testify/suite"
)

type CreateUserRequestTestSuite struct {
	suite.Suite
}

func TestCreateUserRequest(t *testing.T) {
	suite.Run(t, new(CreateUserRequestTestSuite))
}

func (s *CreateUserRequestTestSuite) SetupTest() {
}

func (s *CreateUserRequestTestSuite) TestConvertToUserEntity() {
	createUserReq := &request.CreateUserRequest{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name           string
		expectedOutput *entities.User
		expectedError  error
	}{
		{
			name: "should create basic user",
			expectedOutput: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
	}

	for _, tc := range testCases {
		user, err := createUserReq.CastToUserEntity()
		s.Run(tc.name, func() {
			s.Equal(tc.expectedError, err)
			s.EqualValues(tc.expectedOutput, user)
		})
	}
}
