package protocol

import (
	"spl/pkg/packet"
)

type protocol struct {
	name            string
	fieldNames      []string
	fieldConverters []fieldConverter[any]
}

func (proto *protocol) ToHex(pack packet.Packet) ([]string, error) {
	res := []string{}
	for i, fc := range proto.fieldConverters {

		field, ok := pack.Fields[proto.fieldNames[i]]
		if !ok {
			return nil, packet.ErrMissingField
		}

		hex, err := fc.ToHex(field)
		if err != nil {
			return nil, err
		}
		res = append(res, hex...)
	}
	return res, nil
}
