package request_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_handler "github.com/resyahrial/go-template/internal/api/rest/v1/handler/mocks"
	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	mock_request "github.com/resyahrial/go-template/internal/api/rest/v1/request/mocks"
	"github.com/stretchr/testify/suite"
)

type RequestConverterTestSuite struct {
	suite.Suite
	binder    *mock_handler.MockJsonRequestBinderAdapater
	validator *mock_request.MockValidatorAdapter
	decoder   *mock_request.MockDecoderAdapter
	converter *request.Converter
}

func TestRequestConverter(t *testing.T) {
	suite.Run(t, new(RequestConverterTestSuite))
}

func (s *RequestConverterTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.binder = mock_handler.NewMockJsonRequestBinderAdapater(ctrl)
	s.validator = mock_request.NewMockValidatorAdapter(ctrl)
	s.decoder = mock_request.NewMockDecoderAdapter(ctrl)
	s.converter = request.NewConverter(
		s.validator,
		s.decoder,
	)
}
