package bean

import "fmt"

const (
	RegTypeRunner = "runner"
	RegTypeAgent  = "agent"

	RegStatusExist = 1000
)

type ClientRegInfo struct {
	Id    string `json:"id"`
	Org   string `json:"org"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Sign  string `json:"sign"`
}

func (c *ClientRegInfo) FullPath() string {
	return fmt.Sprintf("%s/%s", c.Org, c.Name)
}
func (c *ClientRegInfo) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}

type ClientRegRes struct {
	Id    string `json:"id"`
	Org   string `json:"org"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Token string `json:"token"`
}

func (c *ClientRegRes) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}

type ClientSubTopic struct {
	Topics []*TopicInfo
}
type ClientUnSubTopic struct {
	Topics []*TopicPath
}

/*type ClientPushTopic struct {
	Topics map[string]*TopicBody
}*/

type GroupClients struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Org   string `json:"org"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (c *GroupClients) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}
