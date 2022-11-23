package request_test

import (
	"errors"

	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	"github.com/resyahrial/go-template/internal/entities"
	"github.com/resyahrial/go-template/pkg/exception"
)

func (s *RequestConverterTestSuite) TestConvertCreateUser() {
	testCases := []struct {
		name               string
		createUserRequest  *request.CreateUser
		mockBindJSONError  error
		mockValidateResult map[string][]string
		mockDecodeError    error
		expectedResult     *entities.User
		expectedError      error
	}{
		{
			name: "should success get create user request",
			createUserRequest: &request.CreateUser{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			expectedResult: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
		{
			name: "should return error when occure error on decode request",
			createUserRequest: &request.CreateUser{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockDecodeError: errors.New("failed to decode request"),
			expectedError:   errors.New("failed to decode request"),
		},
		{
			name: "should return error when occure error on validate request",
			createUserRequest: &request.CreateUser{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockValidateResult: map[string][]string{
				"email": {"invalid email"},
			},
			expectedError: exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(map[string][]string{
				"email": {"invalid email"},
			}),
		},
		{
			name:              "should return error when occure error on binding json request",
			mockBindJSONError: errors.New("failed bind json request"),
			expectedError:     errors.New("failed bind json request"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			var (
				req    *request.CreateUser
				entity *entities.User
			)
			s.reqCtx.EXPECT().BindJSON(&req).SetArg(0, tc.createUserRequest).Return(tc.mockBindJSONError)
			if tc.mockBindJSONError == nil {
				s.validator.EXPECT().Validate(tc.createUserRequest).Return(tc.mockValidateResult)
				if tc.mockValidateResult == nil {
					s.decoder.EXPECT().Decode(tc.createUserRequest, &entity).SetArg(1, tc.expectedResult).Return(tc.expectedError)
				}
			}
			user, err := s.converter.GetCreateUserRequest(s.reqCtx)
			s.Equal(tc.expectedError, err)
			s.Equal(tc.expectedResult, user)
		})
	}
}
