package converters

import (
	"fmt"
	"spl/pkg/packet"
	"spl/pkg/utils"
)

type numericConverter struct {
}

// Numeric is used to convert simple numbers
func Numeric() *numericConverter {
	return &numericConverter{}
}

func (f *numericConverter) ToHex(v int) ([]string, error) {
	hex := fmt.Sprintf("%02X", v)
	if len(hex)%2 != 0 {
		hex = "0" + hex
	}
	arr := utils.SliceString(hex, 2)
	return arr, nil
}
func (f *numericConverter) FromHex(hex []string) (int, error) {
	return 0, packet.ErrNotImplemented
}
