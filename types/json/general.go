package jsontypes

//Chat object
type Chat struct {
	Text       string `json:"text"`
	Bold       bool   `json:"bold"`
	Italic     bool   `json:"italic"`
	Underlined bool   `json:"underlined"`
	Obfuscated bool   `json:"obfuscated"`
	Color      string `json:"color"`
	Insertion  string `json:"insertion"`
	//	ClickEvent ClickEvent `json:"clickEvent"`
}

type Configuration struct {
	ServerPort int16
}

type ServerInfo struct {
	//Description can be a string on older minecraft servers or a chat object, returned as a map[string] interface{}
	Description interface{} `json:"description"`
	//Players object
	Players Players `json:"players"`
	//Version object
	Version Version `json:"version"`
	//Favicon, returned as base64
	Favicon string `json:"favicon"`
}

type Players struct {
	//Max server's player capacity
	Max int `json:"max"`
	//Currently players on the server
	Online int `json:"online"`
	//Sample list of the players on the server
	Sample []SamplePlayersList `json:"sample"`
}

type SamplePlayersList struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Version struct {
	//Name of the version for unsupported clients, this string is showed instead of the players count
	Name string `json:"name"`
	//Protocol version of the server
	Protocol int `json:"protocol"`
}

/*
type ClickEvent struct {
	open_url    string
	open_file   string //DISABLED IN THE JSON CHAT
	run_command string
	//twitch_user_info string - Removed in 1.9
	suggest_command string //Only usable in chats
	change_page     string //Only in books
}
*/
