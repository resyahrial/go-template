package request

type RequestValidatorFn func(data interface{}) map[string][]string

//go:generate mockgen -destination=mocks/mock.go -source=converter.go ValidatorAdapter
type ValidatorAdapter interface {
	Validate(data interface{}) map[string][]string
}

type RequestDecoderFn func(in, out interface{}) error

//go:generate mockgen -destination=mocks/mock.go -source=converter.go DecoderAdapter
type DecoderAdapter interface {
	Decode(in, out interface{}) error
}

type Converter struct {
	validator ValidatorAdapter
	decoder   DecoderAdapter
}

func NewConverter(
	validator ValidatorAdapter,
	decoder DecoderAdapter,
) *Converter {
	return &Converter{
		validator,
		decoder,
	}
}
