package handler_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-template/internal/api/rest/v1/handler"
	mock_handler "github.com/resyahrial/go-template/internal/api/rest/v1/handler/mocks"
	mock_entities "github.com/resyahrial/go-template/internal/entities/mocks"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	ctx          *mock_handler.MockContext
	reqConverter *mock_handler.MockRequestConverter
	resConverter *mock_handler.MockResponseConverter
	userUsecase  *mock_entities.MockUserUsecase
	h            *handler.Handler
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (s *HandlerTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.ctx = mock_handler.NewMockContext(ctrl)
	s.reqConverter = mock_handler.NewMockRequestConverter(ctrl)
	s.resConverter = mock_handler.NewMockResponseConverter(ctrl)
	s.userUsecase = mock_entities.NewMockUserUsecase(ctrl)
	s.h = handler.NewHandler(
		s.reqConverter,
		s.resConverter,
		s.userUsecase,
	)
}
