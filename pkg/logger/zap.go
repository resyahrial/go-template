package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLoggerWrapper struct {
	log *zap.Logger
}

type zapOption struct {
	core               zapcore.Core
	loggerConstuctorFn []zap.Option
}

type ZapLoggerOption func(*zapOption)

func UseZapLogger(opts ...ZapLoggerOption) {
	opt := defaultZapOption()
	for _, o := range opts {
		o(opt)
	}
	L = &zapLoggerWrapper{zap.New(opt.core, opt.loggerConstuctorFn...)}
}

func defaultZapOption() *zapOption {
	return &zapOption{
		core: zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			os.Stdout,
			zap.NewAtomicLevelAt(zapcore.DebugLevel),
		),
		loggerConstuctorFn: []zap.Option{
			zap.AddCaller(),
			zap.AddCallerSkip(2),
			zap.AddStacktrace(zapcore.ErrorLevel),
		},
	}
}

func WithCore(core zapcore.Core) ZapLoggerOption {
	return func(zo *zapOption) {
		zo.core = core
	}
}

func ZapLoggerReleaseModeCore() zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
}

func (l *zapLoggerWrapper) Debug(message string, setKeyValues ...SetKeyValueFn) {
	l.log.Debug(message, l.extractFields(setKeyValues...)...)
}

func (l *zapLoggerWrapper) Info(message string, setKeyValues ...SetKeyValueFn) {
	l.log.Info(message, l.extractFields(setKeyValues...)...)
}

func (l *zapLoggerWrapper) Warn(message string, setKeyValues ...SetKeyValueFn) {
	l.log.Warn(message, l.extractFields(setKeyValues...)...)
}

func (l *zapLoggerWrapper) Error(message string, setKeyValues ...SetKeyValueFn) {
	l.log.Error(message, l.extractFields(setKeyValues...)...)
}

func (l *zapLoggerWrapper) Panic(message string, setKeyValues ...SetKeyValueFn) {
	l.log.Panic(message, l.extractFields(setKeyValues...)...)
}

func (l *zapLoggerWrapper) Fatal(message string, setKeyValues ...SetKeyValueFn) {
	l.log.Fatal(message, l.extractFields(setKeyValues...)...)
}

func (l *zapLoggerWrapper) extractFields(setKeyValues ...SetKeyValueFn) []zap.Field {
	fields := make([]zap.Field, 0)
	for _, setKeyValue := range setKeyValues {
		fields = append(fields, zap.Any(setKeyValue()))
	}
	return fields
}
