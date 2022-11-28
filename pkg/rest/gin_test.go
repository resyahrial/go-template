package rest_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/health-check", nil)
	recorder := httptest.NewRecorder()
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: rest.InitGinEngine(gin.TestMode),
	}
	server.Handler.ServeHTTP(recorder, request)
	response := recorder.Result()
	s.Equal(http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	s.Nil(err)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	s.Equal("OK", responseBody["message"])
}
