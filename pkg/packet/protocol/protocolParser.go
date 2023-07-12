package protocol

import (
	"encoding/json"
	"spl/pkg/packet"
	"spl/pkg/packet/converters"
)

func Parse(p string) (*protocol, error) {
	res := rawProtocol{}
	err := json.Unmarshal([]byte(p), &res)
	if err != nil {
		return nil, err
	}
	return res.Protocol()
}
func ParseCollection(c string) (*Collection, error) {
	rawPs := rawProtocolCollection{}
	err := json.Unmarshal([]byte(c), &rawPs)
	if err != nil {
		return nil, err
	}

	res := &Collection{}
	for _, p := range rawPs.Protocols {
		proto, err := p.Protocol()
		if err != nil {
			return nil, err
		}
		res.protocols = append(res.protocols, *proto)
	}
	return res, nil
}

type rawProtocolCollection struct {
	Protocols []rawProtocol
}
type rawProtocol struct {
	Id     string
	Name   string
	Fields []rawField
}
type rawField struct {
	Name     string
	Datatype string
	Length   int
	Options  map[string]int
}

func (rawP rawProtocol) Protocol() (*protocol, error) {
	proto := protocol{}
	pos := 0
	for _, f := range rawP.Fields {
		fc, err := f.Converter(pos)
		if err != nil {
			return nil, err
		}
		proto.fieldNames = append(proto.fieldNames, f.Name)
		proto.fieldConverters = append(proto.fieldConverters, fc)
		pos += f.Length
	}
	return &proto, nil
}

func (r *rawField) Converter(start int) (fieldConverter[any], error) {
	var fc fieldConverter[any]

	switch r.Datatype {
	case "number":
		fc = &protocolField[int]{
			converters.Numeric(),
			r.Name, start, r.Length,
		}

	case "enum":
		fc = &protocolField[string]{
			converters.Enumic(r.Options),
			r.Name, start, r.Length,
		}
	default:
		return nil, packet.ErrNotAnOption
	}

	return fc, nil
}
