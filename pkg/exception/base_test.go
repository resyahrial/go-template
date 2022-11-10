package exception_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/resyahrial/go-template/pkg/exception"
	"github.com/stretchr/testify/suite"
)

type ExceptionTestSuite struct {
	suite.Suite
}

func TestException(t *testing.T) {
	suite.Run(t, new(ExceptionTestSuite))
}

func (s *ExceptionTestSuite) SetupTest() {
}

func (s *ExceptionTestSuite) TestNewException() {
	testCases := []struct {
		name                   string
		inputStatusCode        int
		inputMessage           string
		inputCollectionMessage map[string][]string
		inputModule            string
		expectedOutput         error
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

		{
			name: "should create exception with collection message",
			inputCollectionMessage: map[string][]string{
				"email": {
					"email not valid",
				},
			},
			inputStatusCode: http.StatusBadRequest,
			inputModule:     "USER",
			expectedOutput: &exception.Base{
				Code: http.StatusBadRequest,
				CollectionMessage: map[string][]string{
					"email": {
						"email not valid",
					},
				},
				Module: "USER",
			},
		},
	}

	for _, tc := range testCases {
		exc := exception.NewException(tc.inputStatusCode).
			SetMessage(tc.inputMessage).
			SetCollectionMessage(tc.inputCollectionMessage).
			SetModule(tc.inputModule)

		s.Run(tc.name, func() {
			s.EqualValues(tc.expectedOutput, exc)
			if len(tc.inputCollectionMessage) > 0 {
				s.Equal(fmt.Sprintf("%v", tc.inputCollectionMessage), exc.Error())
			} else {
				s.Equal(exc.Message, exc.Error())
			}
		})
	}
}
