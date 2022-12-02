package logger

var L LoggerInterface

func init() {
	UseZapLogger()
}

type SetKeyValueFn func() (string, interface{})

type LoggerInterface interface {
	Debug(message string, setKeyValues ...SetKeyValueFn)
	Info(message string, setKeyValues ...SetKeyValueFn)
	Warn(message string, setKeyValues ...SetKeyValueFn)
	Error(message string, setKeyValues ...SetKeyValueFn)
	Panic(message string, setKeyValues ...SetKeyValueFn)
	Fatal(message string, setKeyValues ...SetKeyValueFn)
}

func WithKeyValue(key string, value interface{}) SetKeyValueFn {
	return func() (string, interface{}) {
		return key, value
	}
}

// Basic Debug Level logging
func Debug(message string, setKeyValues ...SetKeyValueFn) {
	L.Debug(message, setKeyValues...)
}

// Basic Info Level logging
func Info(message string, setKeyValues ...SetKeyValueFn) {
	L.Info(message, setKeyValues...)
}

// Basic Warn Level logging
func Warn(message string, setKeyValues ...SetKeyValueFn) {
	L.Warn(message, setKeyValues...)
}

// Basic Error Level logging
func Error(message string, setKeyValues ...SetKeyValueFn) {
	L.Error(message, setKeyValues...)
}

// Basic Panic Level logging
func Panic(message string, setKeyValues ...SetKeyValueFn) {
	L.Panic(message, setKeyValues...)
}

// Basic Fatal Level logging
func Fatal(message string, setKeyValues ...SetKeyValueFn) {
	L.Fatal(message, setKeyValues...)
}
