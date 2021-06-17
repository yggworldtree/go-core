package bean

import "fmt"

const (
	RegTypeRunner = "runner"
	RegTypeAgent  = "agent"

	RegStatusExist = 1000
)

type ClientRegInfo struct {
	Id     string `json:"id"`
	Org    string `json:"org"`
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Subs   string `json:"subs"`
	Sign   string `json:"sign"`
	Secret string `json:"secret"`
}

func (c *ClientRegInfo) GroupPath() string {
	return fmt.Sprintf("%s/%s", c.Org, c.Name)
}
func (c *ClientRegInfo) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}
func (c *ClientRegInfo) Sources() []byte {
	s := fmt.Sprintf("%s/%s|%s", c.Org, c.Name, c.Subs)
	return []byte(s)
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
