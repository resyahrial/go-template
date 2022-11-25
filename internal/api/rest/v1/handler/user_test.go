package handler_test

import (
	"errors"

	"github.com/resyahrial/go-template/internal/entity"
)

func (s *HandlerTestSuite) TestCreateUser() {
	testCases := []struct {
		name                  string
		mockUserEntity        *entity.User
		mockReqConverterError error
		mockUsecaseError      error
		mockResConverterError error
		expectedError         error
	}{
		{
			name: "should success create user",
			mockUserEntity: &entity.User{
				Id:       "id",
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
		{
			name: "should return error when occur error when convert response",
			mockUserEntity: &entity.User{
				Id:       "id",
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockResConverterError: errors.New("failed to convert to response"),
			expectedError:         errors.New("failed to convert to response"),
		},
		{
			name: "should return error when occur error when create user",
			mockUserEntity: &entity.User{
				Id:       "id",
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockUsecaseError: errors.New("failed to create user"),
			expectedError:    errors.New("failed to create user"),
		},
		{
			name:                  "should return error when occur error when convert request",
			mockReqConverterError: errors.New("failed to convert from request"),
			expectedError:         errors.New("failed to convert from request"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.reqConverter.EXPECT().GetCreateUserRequest(s.ctx).Return(tc.mockUserEntity, tc.mockReqConverterError)
			if tc.mockReqConverterError == nil {
				s.userUsecase.EXPECT().CreateUser(s.ctx, tc.mockUserEntity).Return(tc.mockUserEntity, tc.mockUsecaseError)
				if tc.mockUsecaseError == nil {
					s.resConverter.EXPECT().SetCreateUserResponse(s.ctx, tc.mockUserEntity).Return(tc.mockResConverterError)
				}
			}
			err := s.h.CreateUser(s.ctx)
			s.Equal(tc.expectedError, err)
		})
	}
}
