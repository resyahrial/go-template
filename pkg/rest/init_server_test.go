package rest_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func getResponse(method string, path string, bodyReader io.Reader, fn func(host string) *http.Server) (statusCode int, responseBody map[string]interface{}) {
	host := "localhost:3000"
	request := httptest.NewRequest(method, fmt.Sprintf("http://%s%s", host, path), bodyReader)
	recorder := httptest.NewRecorder()
	server := fn(host)
	server.Handler.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	return response.StatusCode, responseBody
}
