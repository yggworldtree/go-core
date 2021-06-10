package bean

import "fmt"

const (
	RegTypeRunner = "runner"
	RegTypeAgent  = "agent"

	RegStatusExist = 1000
)

type ClientRegInfo struct {
	Id    string
	Org   string
	Name  string
	Alias string
}

func (c *ClientRegInfo) FullName() string {
	return fmt.Sprintf("%s/%s", c.Org, c.Name)
}

type ClientRegRes struct {
	Id    string
	Alias string
	Token string
}

type ClientSubTopic struct {
	Topics []*TopicInfo
}

/*type ClientPushTopic struct {
	Topics map[string]*TopicBody
}*/
