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

type ValidatorImpl struct {
	Fn RequestValidatorFn
}

func (v *ValidatorImpl) Validate(data interface{}) map[string][]string {
	return v.Fn(data)
}

type DecoderImpl struct {
	Fn RequestDecoderFn
}

func (v *DecoderImpl) Decode(in, out interface{}) error {
	return v.Fn(in, out)
}
