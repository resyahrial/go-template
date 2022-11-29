package response_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	mock_response "github.com/resyahrial/go-template/internal/api/rest/v1/response/mocks"
	"github.com/stretchr/testify/suite"
)

type ResponseConverterTestSuite struct {
	suite.Suite
	decoder   *mock_response.MockDecoder
	converter *response.Converter
}

func TestResponseConverter(t *testing.T) {
	suite.Run(t, new(ResponseConverterTestSuite))
}

func (s *ResponseConverterTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.decoder = mock_response.NewMockDecoder(ctrl)
	s.converter = response.NewConverter(
		s.decoder,
	)
}
