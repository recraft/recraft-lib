package utils

import "github.com/recraft/recraft-lib/types"

const (
	//HandshakeState is the first state, used when handshake has not yet been done.
	HandshakeState types.VarInt = 0
	//StatusState is when handshake has been done and NextState was defined as "1"
	StatusState types.VarInt = 1
	//LoginState is when handshake has been done and NextState was defined as "2"
	LoginState types.VarInt = 2
	//ProtocolVersion of the client/server
	ProtocolVersion int = 578
)
