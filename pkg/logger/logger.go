package logger

var L LoggerInterface

func init() {
	UseZapLogger()
}

type SetKeyValueFn func() (string, interface{})

type LoggerInterface interface {
	Debug(message string, setKeyValues ...SetKeyValueFn)
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
