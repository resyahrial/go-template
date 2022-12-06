package request_test

import (
	"errors"

	v1 "github.com/resyahrial/go-template/internal/api/rest/v1"
	"github.com/resyahrial/go-template/internal/entity"
	"github.com/resyahrial/go-template/pkg/exception"
)

func (s *RequestConverterTestSuite) TestConvertCreateUser() {
	user := &entity.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}
	testCases := []struct {
		name               string
		createUserRequest  *v1.UserCreate
		mockBindJSONError  error
		mockValidateResult map[string][]string
		mockDecodeError    error
		expectedResult     *entity.User
		expectedError      error
	}{
		{
			name: "should success get create user request",
			createUserRequest: &v1.UserCreate{
				Email:    &user.Email,
				Name:     &user.Name,
				Password: &user.Password,
			},
			expectedResult: user,
		},
		{
			name: "should return error when occure error on decode request",
			createUserRequest: &v1.UserCreate{
				Email:    &user.Email,
				Name:     &user.Name,
				Password: &user.Password,
			},
			mockDecodeError: errors.New("failed to decode request"),
			expectedError:   errors.New("failed to decode request"),
		},
		{
			name: "should return error when occure error on validate request",
			createUserRequest: &v1.UserCreate{
				Email:    &user.Email,
				Name:     &user.Name,
				Password: &user.Password,
			},
			mockValidateResult: map[string][]string{
				"email": {"invalid email"},
			},
			expectedError: exception.NewBadRequestException().SetModule(entity.UserModule).SetCollectionMessage(map[string][]string{
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
				req    *v1.UserCreate
				entity *entity.User
			)
			reqBinderFn := requestBinderFnStub(&req, &tc.createUserRequest, tc.mockBindJSONError)
			if tc.mockBindJSONError == nil {
				s.validator.EXPECT().Validate(tc.createUserRequest).Return(tc.mockValidateResult)
				if tc.mockValidateResult == nil {
					s.decoder.EXPECT().Decode(tc.createUserRequest, &entity).SetArg(1, tc.expectedResult).Return(tc.expectedError)
				}
			}
			user, err := s.converter.GetCreateUserRequest(reqBinderFn)
			s.Equal(tc.expectedError, err)
			s.Equal(tc.expectedResult, user)
		})
	}
}
