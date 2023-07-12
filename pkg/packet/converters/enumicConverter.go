package converters

import "spl/pkg/packet"

// Enumic is used to convert numbers represented by strings
func Enumic(options map[string]int) *enumicConverter {
	return &enumicConverter{
		Options: options,
	}
}

type enumicConverter struct {
	Options map[string]int
}

func (f *enumicConverter) ToHex(v string) ([]string, error) {
	item, exits := f.Options[v]
	if !exits {
		return nil, packet.ErrNotAnOption
	}
	return Numeric().ToHex(item)
}
func (f *enumicConverter) FromHex(hex []string) (string, error) {
	return "", packet.ErrNotImplemented
}
