package request

type RequestValidatorFn func(data interface{}) map[string][]string

type RequestDecoderFn func(in, out interface{}) error

type Converter struct {
	validator Validator
	decoder   Decoder
}

func NewConverter(
	validator Validator,
	decoder Decoder,
) *Converter {
	return &Converter{
		validator,
		decoder,
	}
}
