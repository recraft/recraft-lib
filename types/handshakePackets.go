package types

//ClientHandshake - Structure of the first packet sended by the client
//PacketID is always 0
type ClientHandshake struct {
	//ProtocolVersion: Protocol version of the client
	ProtocolVersion VarInt `rcount:"1"`
	//HostAddress: Address that the client connected with
	HostAddress String `rcount:"2"`
	//Port: Port that the client connected with
	Port Short `rcount:"3"`
	//NextState: "Enum", can be "1" for status or "2" for login
	NextState VarInt `rcount:"4"`
}

//ServerListPingResponse - Response with server info
type ServerListPingResponse struct {
	//Json containing the server info
	JSON String `rcount:"1"`
}

//Pong packet
type Pong struct {
	Payload Long `rcount:"1"`
}
