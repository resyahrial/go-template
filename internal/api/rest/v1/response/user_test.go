package response_test

import (
	"errors"

	v1 "github.com/resyahrial/go-template/internal/api/rest/v1"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/internal/entity"
)

func (s *ResponseConverterTestSuite) TestConvertCreateUser() {
	user := &entity.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}
	testCases := []struct {
		name            string
		userEntity      *entity.User
		mockDecodeError error
		expectedResult  *v1.User
		expectedError   error
	}{
		{
			name:       "should success get create user response",
			userEntity: user,
			expectedResult: &v1.User{
				Name:  &user.Name,
				Email: &user.Email,
			},
		},
		{
			name:            "should return error when occure error on decode request",
			userEntity:      user,
			mockDecodeError: errors.New("failed to decode request"),
			expectedError:   errors.New("failed to decode request"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			var (
				createUserRes *v1.User
			)
			s.decoder.EXPECT().Decode(tc.userEntity, &createUserRes).SetArg(1, tc.expectedResult).Return(tc.expectedError)
			res, err := s.converter.GetCreateUserResponse(tc.userEntity)
			s.Equal(tc.expectedError, err)
			if tc.expectedError == nil {
				s.Equal(&response.Success{
					Data: tc.expectedResult,
				}, res)
			} else {
				s.Nil(res)
			}
		})
	}
}
