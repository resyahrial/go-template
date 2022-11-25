package middleware

type Context interface {
	AbortWithStatusJSON(code int, obj any)
	Get(key string) (value any, isExist bool)
	JSON(code int, obj any)
	Next()
}
