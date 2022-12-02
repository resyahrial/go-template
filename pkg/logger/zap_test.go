package logger_test

import (
	"os"

	"github.com/resyahrial/go-template/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func (s *LoggerTestSuite) TestZapLogger() {
	observedZapCore, observedLogs := observer.New(zapcore.DebugLevel)
	baseCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		os.Stdout,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	)
	logger.UseZapLogger(
		logger.WithCore(zapcore.NewTee(baseCore, observedZapCore)),
	)

	logger.Debug("debug")
	logger.Debug("debug with field", logger.WithKeyValue("foo", "bar"))
	s.Equal(observedLogs.Len(), 2)
	logs := observedLogs.All()
	s.Equal("debug", logs[0].Message)
	s.Equal(zapcore.DebugLevel, logs[0].Level)
	s.Equal("debug with field", logs[1].Message)
	s.ElementsMatch([]zap.Field{
		{
			Key:    "foo",
			Type:   zapcore.StringType,
			String: "bar",
		},
	}, logs[1].Context)
}
