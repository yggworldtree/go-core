package hbtpBean

import "fmt"

const (
	RegTypeRunner = "runner"
	RegTypeAgent  = "agent"

	RegStatusExist = 1000
)

type ClientRegInfo struct {
	Org  string
	Name string
}

func (c *ClientRegInfo) FullName() string {
	return fmt.Sprintf("%s/%s", c.Org, c.Name)
}

type ClientRegRes struct {
	Id    string
	Token string
}

type ClientSubTopic struct {
	Topics []string
}
