package jsontypes

//ServerInfo contains informations about the server, obtained by the status request
type ServerInfo struct {
	// Description can be a string on older minecraft servers or a chat object, returned as a map[string] interface{}
	Description interface{} `json:"description"`
	// Players object
	Players Players `json:"players"`
	// Version object
	Version Version `json:"version"`
	// Favicon, returned as base64
	Favicon string `json:"favicon"`
}

type Players struct {
	// Max server's player capacity
	Max int `json:"max"`
	// Currently players on the server
	Online int `json:"online"`
	// Sample list of the players on the server
	Sample []SamplePlayersList `json:"sample"`
}

type SamplePlayersList struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Version struct {
	// Name of the version for unsupported clients, this string is showed instead of the players count
	Name string `json:"name"`
	// Protocol version of the server
	Protocol int `json:"protocol"`
}
