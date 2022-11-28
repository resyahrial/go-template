package rest_test

import (
	"fmt"
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
	code, resBody := getResponse(http.MethodGet, "/health-check", nil, rest.InitGinEngine(gin.TestMode))
	s.Equal(http.StatusOK, code)
	s.NotNil(resBody)
	s.Equal("OK", resBody["message"])
}

func (s *GinEngineTestSuite) TestNoRoute() {
	code, resBody := getResponse(http.MethodGet, "/unregistered_route", nil, rest.InitGinEngine(gin.TestMode))
	s.Equal(http.StatusInternalServerError, code)
	s.NotNil(resBody)
	s.Equal("route not found", resBody["error"])
}

func (s *GinEngineTestSuite) TestPanicRecovery() {
	panicRoutePath := "/panic-route"
	engine := rest.InitGinEngine(gin.TestMode)
	engine.GET(panicRoutePath, func(ctx *gin.Context) {
		emptySlice := make([]string, 0)
		fmt.Println(emptySlice[0])
	})
	code, resBody := getResponse(http.MethodGet, panicRoutePath, nil, engine)
	s.Equal(http.StatusInternalServerError, code)
	s.NotNil(resBody)
	s.Contains(resBody["error"], "panic")
}
