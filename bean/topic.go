package bean

import "fmt"

type TopicPath struct {
	NameSpace string `json:"nameSpace,omitempty"`
	Key       string `json:"key,omitempty"`
	Tag       string `json:"tag,omitempty"`
}

func NewTopicPath(namespace, key string, tag ...string) *TopicPath {
	c := &TopicPath{
		NameSpace: namespace,
		Key:       key,
		Tag:       "main",
	}
	if len(tag) > 0 && tag[0] != "" {
		c.Tag = tag[0]
	}
	return c
}
func (c *TopicPath) String() string {
	return fmt.Sprintf("%s/%s:%s", c.NameSpace, c.Key, c.Tag)
}

type TopicInfo struct {
	Path  *TopicPath
	Safed bool //是否确保发送必到
}

type TopicParam struct {
	Type string
}
