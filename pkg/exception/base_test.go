package exception_test

import (
	"fmt"
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
		inputMessages   []string
		expectedOutput  error
	}{
		{
			name: "should create base exception",
			inputMessages: []string{
				"internal server error",
			},
			expectedOutput: &exception.Base{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
			},
		},
		{
			name:            "should create exception with joined multiple message",
			inputStatusCode: http.StatusBadRequest,
			inputMessages: []string{
				"validate 1",
				"validate 2",
			},
			expectedOutput: &exception.Base{
				Code:    http.StatusBadRequest,
				Message: "validate 1, validate 2",
			},
		},
	}

	for _, tc := range testCases {
		exc := exception.NewBaseException(tc.inputStatusCode, tc.inputMessages...)
		s.Run(tc.name, func() {
			s.EqualValues(tc.expectedOutput, exc)
			s.Equal(fmt.Sprintf("[%v]%s\n", exc.Code, exc.Message), exc.Error())
		})
	}
}
