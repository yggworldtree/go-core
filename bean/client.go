package bean

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yggworldtree/go-core/common"
)

const (
	RegTypeRunner = "runner"
	RegTypeAgent  = "agent"

	RegStatusExist = 1000
)

type ClientRegInfo struct {
	Id        string `json:"id"`
	Org       string `json:"org"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	Frequency string `json:"frequency"`
	Subs      []TopicSubInfo
	Pushs     []TopicPushInfo
	Sign      string `json:"sign"`
	Secret    string `json:"secret"`
}

func (c *ClientRegInfo) GroupPath() string {
	return fmt.Sprintf("%s/%s", c.Org, c.Name)
}
func (c *ClientRegInfo) CliGroupPath() *CliGroupPath {
	return NewCliGroupPath(c.Org, c.Name, c.Alias)
}
func (c *ClientRegInfo) MaxFreqTm() time.Duration {
	bs := time.Second
	if !common.RegTms.MatchString(c.Frequency) {
		return bs
	}
	s := common.RegTms.FindAllStringSubmatch(c.Frequency, -1)[0]
	switch s[2] {
	case "h":
		bs = time.Hour
	case "m":
		bs = time.Minute
	case "ms":
		bs = time.Millisecond
	}
	n, _ := strconv.ParseInt(s[1], 10, 64)
	if n <= 0 {
		return bs
	}
	return bs * time.Duration(n)
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
	Topics []*TopicPath
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
