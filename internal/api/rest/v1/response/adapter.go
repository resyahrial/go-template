package response

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go Decoder
type Decoder interface {
	Decode(in, out interface{}) error
}
