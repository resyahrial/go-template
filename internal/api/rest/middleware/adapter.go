package middleware

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Context
type Context interface {
	Get(key string) (value any, isExist bool)
	JSON(code int, obj any)
	Next()
}
