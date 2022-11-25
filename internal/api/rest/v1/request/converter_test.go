package request_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-template/internal/api/rest/v1/request"
	mock_request "github.com/resyahrial/go-template/internal/api/rest/v1/request/mocks"
	"github.com/stretchr/testify/suite"
)

type RequestConverterTestSuite struct {
	suite.Suite
	ctx       *mock_request.MockContext
	validator *mock_request.MockValidator
	decoder   *mock_request.MockDecoder
	converter *request.Converter
}

func TestRequestConverter(t *testing.T) {
	suite.Run(t, new(RequestConverterTestSuite))
}

func (s *RequestConverterTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.ctx = mock_request.NewMockContext(ctrl)
	s.validator = mock_request.NewMockValidator(ctrl)
	s.decoder = mock_request.NewMockDecoder(ctrl)
	s.converter = request.NewConverter(
		s.validator,
		s.decoder,
	)
}
