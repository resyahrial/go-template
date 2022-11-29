package middleware_test

import (
	"errors"
	"net/http"

	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	"github.com/resyahrial/go-template/pkg/exception"
)

func (s *MiddlewareTestSuite) TestResponseHandlerError() {
	testCases := []struct {
		name            string
		mockError       any
		expectedCode    int
		expectedMessage *middleware.Failure
	}{
		{
			name:         "success handle ordinary error",
			mockError:    errors.New("message"),
			expectedCode: http.StatusInternalServerError,
			expectedMessage: &middleware.Failure{
				ErrorMsg: map[string]interface{}{
					"message": "message",
				},
			},
		},
		{
			name:         "success handle exception package error - single error",
			mockError:    exception.NewAuthenticationException().SetMessage("message"),
			expectedCode: http.StatusForbidden,
			expectedMessage: &middleware.Failure{
				ErrorMsg: map[string]interface{}{
					"message": "message",
				},
			},
		},
		{
			name: "success handle exception package error - collection error",
			mockError: exception.NewBadRequestException().SetCollectionMessage(map[string][]string{
				"email": {"invalid email"},
			}),
			expectedCode: http.StatusBadRequest,
			expectedMessage: &middleware.Failure{
				ErrorMsg: map[string][]string{
					"email": {"invalid email"},
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.ctx.EXPECT().Next()
			s.ctx.EXPECT().Get(middleware.ResultKey).Return(tc.mockError, true)
			s.ctx.EXPECT().JSON(tc.expectedCode, tc.expectedMessage)
			s.m.ResponseHandler(s.ctx)
		})
	}
}

func (s *MiddlewareTestSuite) TestResponseHandlerSuccess() {
	testCases := []struct {
		name         string
		mockData     any
		expectedData any
	}{
		{
			name: "success handle success response",
			mockData: map[string]interface{}{
				"message": "success",
			},
			expectedData: map[string]interface{}{
				"message": "success",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.ctx.EXPECT().Next()
			s.ctx.EXPECT().Get(middleware.ResultKey).Return(tc.mockData, true)
			s.ctx.EXPECT().JSON(http.StatusOK, tc.expectedData)
			s.m.ResponseHandler(s.ctx)
		})
	}
}
