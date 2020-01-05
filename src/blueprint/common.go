package blueprint

type StaffData struct {
	Type    string `json:"type"`
	Login   string `json:"login"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
}

type ProjectData struct {
	Id           int    `json:"id"`
	Scolaryear   string `json:"scolaryear"`
	Codemodule   string `json:"codemodule"`
	Codeinstance string `json:"codeinstance"`
	Title        string `json:"title"`
}
