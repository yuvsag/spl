package packet_test

import (
	"spl/pkg/packet"
	"spl/pkg/packet/protocol"
	"testing"

	"github.com/stretchr/testify/require"
)

type protocolTestParams struct {
	rawProtocol string
	packetMap   map[string]any
	expected    []string
}

func testProtocolOnPacket(t *testing.T, params protocolTestParams) {

	protocol, err := protocol.Parse(params.rawProtocol)
	require.Zero(t, err)

	p, err := packet.FromMap(params.packetMap)
	require.Zero(t, err)

	res, err := protocol.ToHex(p)
	require.Zero(t, err)

	require.Equal(t, params.expected, res)
}

func TestFirstProtocol(t *testing.T) {
	protocol := `{
		"id": "0/2",
		"name": "protocol/example",
		"fields": [
			{
				"name": "num","datatype": "number","length": 2
			},
			{
				"name": "mode","datatype": "enum","length": 1,
				"options": {"high": 420, "luwu": 0}
			}
		]
	}`

	testProtocolOnPacket(t, protocolTestParams{
		rawProtocol: protocol,
		packetMap: map[string]any{
			"protocol": "protocol/example",
			"num":      69,
			"mode":     "high",
		},
		//420 -> 01 A4
		//69 -> 45
		expected: []string{"45", "01", "A4"},
	})
}
