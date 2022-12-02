package logger_test

import (
	"os"

	"bou.ke/monkey"
	"github.com/resyahrial/go-template/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func (s *LoggerTestSuite) initZapLogger() *observer.ObservedLogs {
	observedZapCore, observedLogs := observer.New(zapcore.DebugLevel)
	baseCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		os.Stdout,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	)
	logger.UseZapLogger(
		logger.WithCore(zapcore.NewTee(baseCore, observedZapCore)),
	)
	return observedLogs
}

func (s *LoggerTestSuite) TestZapLogger() {
	observedLogs := s.initZapLogger()

	logger.Debug("debug")
	logger.Debug("debug with field", logger.WithKeyValue("foo", "bar"))
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	s.Equal(observedLogs.Len(), 5)
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
	s.Equal("info", logs[2].Message)
	s.Equal("warn", logs[3].Message)
	s.Equal("error", logs[4].Message)
}

func (s *LoggerTestSuite) TestZapLoggerPanic() {
	observedLogs := s.initZapLogger()
	defer func() {
		if r := recover(); r != nil {
			s.Equal(observedLogs.Len(), 1)
			logs := observedLogs.All()
			s.Equal("panic", logs[0].Message)
			s.Equal(zapcore.PanicLevel, logs[0].Level)
		}
	}()

	logger.Panic("panic")
}

func (s *LoggerTestSuite) TestZapLoggerFatal() {
	observedLogs := s.initZapLogger()
	patch := monkey.Patch(os.Exit, func(int) {
		panic("os.Exit called")
	})

	defer func() {
		if r := recover(); r != nil {
			s.Equal(observedLogs.Len(), 1)
			logs := observedLogs.All()
			s.Equal("fatal", logs[0].Message)
			s.Equal(zapcore.FatalLevel, logs[0].Level)
		}
		patch.Unpatch()
	}()

	logger.Fatal("fatal")
}
