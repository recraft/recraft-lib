package clientpackets

import "github.com/recraft/recraft-lib/types"

// PacketStatusRequest structure
// Client's request to get status
type PacketStatusRequest struct {
}

//ID of packet
func (packet *PacketStatusRequest) ID() int32 {
	return 0
}

// PacketPing structure
type PacketPing struct {
	Payload types.Long `rcount:"1"`
}

//ID of packet
func (packet *PacketPing) ID() int32 {
	return 1
}
