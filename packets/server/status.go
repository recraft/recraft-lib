package serverpackets

import "github.com/recraft/recraft-lib/types"

// PacketPong structure
type PacketPong struct {
	Payload types.Long `rcount:"1"`
}

//ID of packet
func (packet *PacketPong) ID() int32 {
	return 1
}

//PacketStatus structure
type PacketStatus struct {

	//JSON containing payload
	JSON types.String `rcount:"1"`
}

//ID of packet
func (packet *PacketStatus) ID() int32 {
	return 0
}
