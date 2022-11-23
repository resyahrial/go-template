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
