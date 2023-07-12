package packet

type Packet struct {
	Protocol string
	Fields   map[string]any
}

func FromMap(m map[string]any) (Packet, error) {
	temName, exits := m["protocol"]
	pName, valid := temName.(string)
	delete(m, "protocol")
	if !exits || !valid {
		return Packet{}, ErrProtocolNotFound
	}

	return Packet{
		Protocol: pName,
		Fields:   m,
	}, nil
}
