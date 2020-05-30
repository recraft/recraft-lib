package jsontypes

//Chat object
type Chat struct {
	Text       string     `json:"text"`
	Bold       bool       `json:"bold"`
	Italic     bool       `json:"italic"`
	Underlined bool       `json:"underlined"`
	Obfuscated bool       `json:"obfuscated"`
	Color      string     `json:"color"`
	Insertion  string     `json:"insertion"`
	ClickEvent ClickEvent `json:"clickEvent"`
	Extra      []Chat     `json:"extra"`
}

//ClickEvent object
type ClickEvent struct {
	OpenUrl        string `json:"open_url"`
	RunCommand     string `json:"run_command"`
	SuggestCommand string `json:"suggest_command"`
	ChangePage     string `json:"change_page"`
}
