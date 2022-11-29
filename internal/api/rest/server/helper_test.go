package server_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type responseOption struct {
	method string
	path   string
	body   io.Reader
}

type responseOptionFn func(*responseOption)

func defaultResponseOption() *responseOption {
	return &responseOption{
		method: http.MethodGet,
	}
}

func getResponse(handler http.Handler, opts ...responseOptionFn) (statusCode int, responseBody map[string]interface{}) {
	opt := defaultResponseOption()
	for _, o := range opts {
		o(opt)
	}
	host := "localhost:3000"
	request := httptest.NewRequest(opt.method, fmt.Sprintf("http://%s%s", host, opt.path), opt.body)
	recorder := httptest.NewRecorder()
	server := &http.Server{
		Addr:    host,
		Handler: handler,
	}
	server.Handler.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	return response.StatusCode, responseBody
}

func withMethod(method string) responseOptionFn {
	return func(ro *responseOption) {
		ro.method = method
	}
}

func withPath(path string) responseOptionFn {
	return func(ro *responseOption) {
		ro.path = path
	}
}

func withBodyReader(body io.Reader) responseOptionFn {
	return func(ro *responseOption) {
		ro.body = body
	}
}
