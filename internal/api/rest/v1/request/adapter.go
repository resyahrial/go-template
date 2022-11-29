package request

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go ValidatorAdapter
type Validator interface {
	Validate(data interface{}) map[string][]string
}

//go:generate mockgen -destination=mocks/mock.go -source=adapter.go DecoderAdapter
type Decoder interface {
	Decode(in, out interface{}) error
}
