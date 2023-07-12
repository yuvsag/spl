package protocol

import (
	"spl/pkg/packet"
)

type fieldConverter[T any] interface {
	ToHex(T) ([]string, error)
	FromHex([]string) (T, error)
}

type protocolField[T any] struct {
	converter fieldConverter[T]
	name      string
	start     int
	length    int
}

func (f *protocolField[T]) ToHex(v any) ([]string, error) {
	value, ok := v.(T)
	if !ok {
		return nil, packet.ErrIncorrectType
	}
	return f.converter.ToHex(value)
}
func (f *protocolField[T]) FromHex(hex []string) (any, error) {
	return f.converter.FromHex(hex[f.start : f.start+f.length])
}
