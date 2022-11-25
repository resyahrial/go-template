package response_test

import (
	"errors"

	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	"github.com/resyahrial/go-template/internal/entities"
)

func (s *ResponseConverterTestSuite) TestConvertCreateUser() {
	testCases := []struct {
		name               string
		userEntity         *entities.User
		mockDecodeError    error
		mockDecodeResponse *response.CreateUser
		expectedError      error
	}{
		{
			name: "should success get create user response",
			userEntity: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockDecodeResponse: &response.CreateUser{
				Name:  "user",
				Email: "user@mail.com",
			},
		},
		{
			name: "should return error when occure error on decode request",
			userEntity: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockDecodeError: errors.New("failed to decode request"),
			expectedError:   errors.New("failed to decode request"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			var (
				res *response.CreateUser
			)
			s.decoder.EXPECT().Decode(tc.userEntity, &res).SetArg(1, tc.mockDecodeResponse).Return(tc.expectedError)
			if tc.expectedError == nil {
				s.ctx.EXPECT().Set(response.SuccessKey, tc.mockDecodeResponse)
			}
			err := s.converter.SetCreateUserResponse(s.ctx, tc.userEntity)
			s.Equal(tc.expectedError, err)
		})
	}
}
