package server_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/resyahrial/go-template/internal/api/rest/server"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	engine *gin.Engine
}

func TestGinEngine(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (s *ServerTestSuite) SetupTest() {
	s.engine = server.InitServerTestMode()
}

func (s *ServerTestSuite) TestHealthCheck() {
	code, resBody := getResponse(s.engine, withPath("/health-check"))
	s.Equal(http.StatusOK, code)
	s.NotNil(resBody)
	s.Equal("OK", resBody["message"])
}

func (s *ServerTestSuite) TestNoRoute() {
	code, resBody := getResponse(
		s.engine,
		withPath("/unregistered_route"),
		withMethod(http.MethodPost),
		withBodyReader(strings.NewReader(`{"name" : "Gadget"}`)),
	)
	s.Equal(http.StatusInternalServerError, code)
	s.NotNil(resBody)
	s.Equal("route not found", resBody["error"])
}

func (s *ServerTestSuite) TestPanicRecovery() {
	panicRoutePath := "/panic-route"
	s.engine.GET(panicRoutePath, func(ctx *gin.Context) {
		emptySlice := make([]string, 0)
		fmt.Println(emptySlice[0])
	})
	code, resBody := getResponse(s.engine, withPath(panicRoutePath))
	s.Equal(http.StatusInternalServerError, code)
	s.NotNil(resBody)
	s.Contains(resBody["error"], "panic")
}
