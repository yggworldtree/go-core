package bean

import "fmt"

const (
	RegTypeRunner = "runner"
	RegTypeAgent  = "agent"

	RegStatusExist = 1000
)

type ClientRegInfo struct {
	Id    string `json:"id,omitempty"`
	Org   string `json:"org,omitempty"`
	Name  string `json:"name,omitempty"`
	Alias string `json:"alias,omitempty"`
}

func (c *ClientRegInfo) FullPath() string {
	return fmt.Sprintf("%s/%s", c.Org, c.Name)
}
func (c *ClientRegInfo) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}

type ClientRegRes struct {
	Id    string `json:"id,omitempty"`
	Org   string `json:"org,omitempty"`
	Name  string `json:"name,omitempty"`
	Alias string `json:"alias,omitempty"`
	Token string `json:"token,omitempty"`
}

func (c *ClientRegRes) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}

type ClientSubTopic struct {
	Topics []*TopicInfo
}

/*type ClientPushTopic struct {
	Topics map[string]*TopicBody
}*/

type GroupClients struct {
	Id    string `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
	Org   string `json:"org,omitempty"`
	Name  string `json:"name,omitempty"`
	Count int    `json:"count,omitempty"`
}

func (c *GroupClients) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}
