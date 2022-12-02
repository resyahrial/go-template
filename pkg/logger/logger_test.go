package logger_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerTestSuite struct {
	suite.Suite
}

func TestLogger(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

var (
	fieldTester = struct {
		Foo string `json:"foo"`
	}{"bar"}
)
