package response

type ResponseDecoderFn func(in, out interface{}) error

//go:generate mockgen -destination=mocks/mock.go -source=converter.go DecoderAdapter
type DecoderAdapter interface {
	Decode(in, out interface{}) error
}

type Converter struct {
	decoder DecoderAdapter
}

func NewConverter(
	decoder DecoderAdapter,
) *Converter {
	return &Converter{
		decoder,
	}
}
