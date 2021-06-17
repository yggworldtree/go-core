package bean

type GroupListRes struct {
	Active    bool     `json:"active"`
	FullPath  string   `json:"fullPath"`
	ClientLen int      `json:"clientLen"`
	Clients   []string `json:"clients"`
}
type ClientListRes struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	GroupPath   string `json:"groupPath"`
	GroupActive bool   `json:"groupActive"`
}
