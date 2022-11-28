package rest_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/pkg/rest"
	"github.com/stretchr/testify/suite"
)

type GinEngineTestSuite struct {
	suite.Suite
}

func TestGinEngine(t *testing.T) {
	suite.Run(t, new(GinEngineTestSuite))
}

func (s *GinEngineTestSuite) SetupTest() {
}

func (s *GinEngineTestSuite) TestHealthCheck() {
	code, resBody := getResponse(http.MethodGet, "/health-check", nil, func(host string) *http.Server {
		return &http.Server{
			Addr:    host,
			Handler: rest.InitGinEngine(gin.TestMode),
		}
	})
	s.Equal(http.StatusOK, code)
	s.NotNil(resBody)
	s.Equal("OK", resBody["message"])
}

func (s *GinEngineTestSuite) TestNoRoute() {
	code, resBody := getResponse(http.MethodGet, "/unregistered_route", nil, func(host string) *http.Server {
		return &http.Server{
			Addr:    host,
			Handler: rest.InitGinEngine(gin.TestMode),
		}
	})
	s.Equal(http.StatusInternalServerError, code)
	s.NotNil(resBody)
	s.Equal("route not found", resBody["error"])
}
