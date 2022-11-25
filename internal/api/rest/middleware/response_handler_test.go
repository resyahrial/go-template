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
			s.ctx.EXPECT().Get(middleware.FailureKey).Return(tc.mockError, true)
			s.ctx.EXPECT().JSON(tc.expectedCode, tc.expectedMessage)
			s.m.ResponseHandler(s.ctx)
		})
	}
}

func (s *MiddlewareTestSuite) TestResponseHandlerSuccess() {
	testCases := []struct {
		name         string
		mockData     any
		expectedData *middleware.Success
	}{
		{
			name: "success handle success response",
			mockData: map[string]interface{}{
				"message": "success",
			},
			expectedData: &middleware.Success{
				Data: map[string]interface{}{
					"message": "success",
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.ctx.EXPECT().Next()
			s.ctx.EXPECT().Get(middleware.FailureKey).Return(nil, false)
			s.ctx.EXPECT().Get(middleware.SuccessKey).Return(tc.mockData, true)
			s.ctx.EXPECT().Get(middleware.PaginatedKey).Return(nil, false)
			s.ctx.EXPECT().JSON(http.StatusOK, tc.expectedData)
			s.m.ResponseHandler(s.ctx)
		})
	}
}

func (s *MiddlewareTestSuite) TestResponseHandlerPaginated() {
	testCases := []struct {
		name              string
		mockData          any
		mockPaginatedData middleware.PaginatedResultValue
		expectedData      *middleware.Success
	}{
		{
			name: "success handle paginated response",
			mockData: []struct{ ID string }{
				{ID: "id"},
			},
			mockPaginatedData: middleware.PaginatedResultValue{
				Page:  0,
				Limit: 10,
				Count: 1,
			},
			expectedData: &middleware.Success{
				Data: []struct{ ID string }{
					{ID: "id"},
				},
				PageInfo: middleware.PageInfo{
					CurrentPage: 1,
					TotalPage:   1,
					Count:       1,
				}},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.ctx.EXPECT().Next()
			s.ctx.EXPECT().Get(middleware.FailureKey).Return(nil, false)
			s.ctx.EXPECT().Get(middleware.SuccessKey).Return(tc.mockData, true)
			s.ctx.EXPECT().Get(middleware.PaginatedKey).Return(tc.mockPaginatedData, true)
			s.ctx.EXPECT().JSON(http.StatusOK, tc.expectedData)
			s.m.ResponseHandler(s.ctx)
		})
	}
}
