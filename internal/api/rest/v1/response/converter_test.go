package response_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_handler "github.com/resyahrial/go-template/internal/api/rest/v1/handler/mocks"
	"github.com/resyahrial/go-template/internal/api/rest/v1/response"
	mock_response "github.com/resyahrial/go-template/internal/api/rest/v1/response/mocks"
	"github.com/stretchr/testify/suite"
)

type ResponseConverterTestSuite struct {
	suite.Suite
	resCtx    *mock_handler.MockResponseContext
	decoder   *mock_response.MockDecoderAdapter
	converter *response.Converter
}

func TestResponseConverter(t *testing.T) {
	suite.Run(t, new(ResponseConverterTestSuite))
}

func (s *ResponseConverterTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.resCtx = mock_handler.NewMockResponseContext(ctrl)
	s.decoder = mock_response.NewMockDecoderAdapter(ctrl)
	s.converter = response.NewConverter(
		s.decoder,
	)
}
