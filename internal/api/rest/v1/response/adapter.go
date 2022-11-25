package response

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Context
type Context interface {
	Set(key string, obj any)
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Decoder
type Decoder interface {
	Decode(in, out interface{}) error
}
