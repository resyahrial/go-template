package response

type ResponseDecoderFn func(in, out interface{}) error

type Converter struct {
	decoder Decoder
}

func NewConverter(
	decoder Decoder,
) *Converter {
	return &Converter{
		decoder,
	}
}

type DecoderImpl struct {
	Fn ResponseDecoderFn
}

func (v *DecoderImpl) Decode(in, out interface{}) error {
	return v.Fn(in, out)
}
