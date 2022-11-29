package middleware_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-template/internal/api/rest/middleware"
	mock_middleware "github.com/resyahrial/go-template/internal/api/rest/middleware/mocks"
	"github.com/stretchr/testify/suite"
)

type MiddlewareTestSuite struct {
	suite.Suite
	ctx *mock_middleware.MockContext
	m   *middleware.Middleware
}

func TestMiddleware(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

func (s *MiddlewareTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.ctx = mock_middleware.NewMockContext(ctrl)
	s.m = middleware.New()
}
