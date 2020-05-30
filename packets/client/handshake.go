package clientpackets

import "github.com/recraft/recraft-lib/types"

// PacketHandshake structure
// PacketID is always 0
type PacketHandshake struct {
	// ProtocolVersion: Protocol version of the client
	ProtocolVersion types.VarInt `rcount:"1"`
	// HostAddress: Address that the client connected with
	HostAddress types.String `rcount:"2"`
	// Port: Port that the client connected with
	Port types.Short `rcount:"3"`
	// NextState: "Enum", can be "1" for status or "2" for login
	NextState types.VarInt `rcount:"4"`
}

//ID of packet
func (packet *PacketHandshake) ID() int32 {
	return 0
}
