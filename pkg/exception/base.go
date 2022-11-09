package exception

import (
	"fmt"
	"net/http"
	"strings"
)

type Base struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewBaseException(statusCode int, msg ...string) *Base {
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	return &Base{
		Code:    statusCode,
		Message: strings.TrimSpace(strings.Join(msg, ", ")),
	}
}

func (exc Base) Error() string {
	return fmt.Sprintf("[%v]%s\n", exc.Code, exc.Message)
}
