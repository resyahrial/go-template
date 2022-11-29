package request_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	mock_request "github.com/resyahrial/go-template/internal/api/rest/v1/request/mocks"
	"github.com/stretchr/testify/suite"
)

type RequestConverterTestSuite struct {
	suite.Suite
	validator *mock_request.MockValidator
	decoder   *mock_request.MockDecoder
	converter *request.Converter
}

func TestRequestConverter(t *testing.T) {
	suite.Run(t, new(RequestConverterTestSuite))
}

func (s *RequestConverterTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.validator = mock_request.NewMockValidator(ctrl)
	s.decoder = mock_request.NewMockDecoder(ctrl)
	s.converter = request.NewConverter(
		s.validator,
		s.decoder,
	)
}

func requestBinderFnStub(expectedInput any, expectedResult any, expectedError error) func(obj any) error {
	return func(obj any) error {
		if reflect.TypeOf(expectedInput) != reflect.TypeOf(obj) {
			return errors.New("different input and expected input")
		}
		if expectedError == nil {
			objVal, expVal := reflect.ValueOf(obj), reflect.ValueOf(expectedResult)
			objVal.Elem().Set(expVal.Elem())
		}

		return expectedError
	}
}
