package exception

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	BaseModule = "BASE"
)

type Base struct {
	Code              int
	Message           string
	Module            string
	CollectionMessage map[string][]string
}

func NewException(statusCode int) *Base {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return &Base{
		Code:   statusCode,
		Module: BaseModule,
	}
}

func (exc *Base) SetMessage(msg string) *Base {
	exc.Message = msg
	return exc
}

func (exc *Base) SetCollectionMessage(msg map[string][]string) *Base {
	exc.CollectionMessage = msg
	return exc
}

func (exc *Base) Error() string {
	if len(exc.CollectionMessage) != 0 {
		return fmt.Sprintf("%v", exc.CollectionMessage)
	}
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
