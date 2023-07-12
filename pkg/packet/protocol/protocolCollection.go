package protocol

import "spl/pkg/packet"

type Collection struct {
	protocols []protocol
}

func (c *Collection) ProtocolForPacket(p packet.Packet) (*protocol, error) {
	for _, proto := range c.protocols {
		if proto.name == p.Protocol {
			return &proto, nil
		}
	}
	return nil, packet.ErrProtocolNotFound
}
