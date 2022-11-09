package exception

import (
	"log"
	"net/http"
	"strings"
)

const (
	BaseModule = "BASE"
)

type Base struct {
	Code    int
	Message string
	Module  string
}

func NewBaseException(statusCode int, msg string) *Base {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return &Base{
		Code:    statusCode,
		Message: msg,
		Module:  BaseModule,
	}
}

func (exc *Base) Error() string {
	return exc.Message
}

func (exc *Base) LogError() {
	log.Printf("[%s] %s", exc.Module, exc.Message)
}

func (exc *Base) SetModule(moduleName string) *Base {
	if strings.TrimSpace(moduleName) != "" {
		exc.Module = moduleName
	}
	return exc
}
