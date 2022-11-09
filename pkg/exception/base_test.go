package exception_test

import (
	"net/http"
	"testing"

	"github.com/resyahrial/go-template/pkg/exception"
	"github.com/stretchr/testify/suite"
)

type BaseExceptionTestSuite struct {
	suite.Suite
}

func TestBaseException(t *testing.T) {
	suite.Run(t, new(BaseExceptionTestSuite))
}

func (s *BaseExceptionTestSuite) SetupTest() {
}

func (s *BaseExceptionTestSuite) TestNewBaseException() {
	testCases := []struct {
		name            string
		inputStatusCode int
		inputMessage    string
		inputModule     string
		expectedOutput  error
	}{
		{
			name:         "should create base exception",
			inputMessage: "internal server error",
			expectedOutput: &exception.Base{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
				Module:  exception.BaseModule,
			},
		},
		{
			name:            "should create exception with module",
			inputMessage:    "bad request",
			inputStatusCode: http.StatusBadRequest,
			inputModule:     "USER",
			expectedOutput: &exception.Base{
				Code:    http.StatusBadRequest,
				Message: "bad request",
				Module:  "USER",
			},
		},
	}

	for _, tc := range testCases {
		exc := exception.NewBaseException(tc.inputStatusCode, tc.inputMessage).SetModule(tc.inputModule)
		s.Run(tc.name, func() {
			s.EqualValues(tc.expectedOutput, exc)
			s.Equal(exc.Message, exc.Error())
		})
	}
}
