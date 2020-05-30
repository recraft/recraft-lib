package packets

const (
	// HandshakePacketID is the first packetID sent
	HandshakePacketID int32 = 0
)

type State int32

const (
	HANDSHAKE State = iota
	STATUS
	LOGIN
	PLAY
)

//Packet interface
type Packet interface {
	ID() int32
}
