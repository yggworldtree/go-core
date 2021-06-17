package bean

import (
	"errors"
	"fmt"
	"strings"
)

type TopicPath struct {
	NameSpace string `json:"nameSpace"`
	Key       string `json:"key"`
	Tag       string `json:"tag"`
}

func NewTopicPath(namespace, key string, tag ...string) *TopicPath {
	c := &TopicPath{
		NameSpace: namespace,
		Key:       key,
		Tag:       "",
	}
	if len(tag) > 0 && tag[0] != "" {
		c.Tag = tag[0]
	}
	if c.Tag == "" {
		c.Tag = "main"
	}
	return c
}
func (c *TopicPath) String() string {
	return fmt.Sprintf("%s/%s:%s", c.NameSpace, c.Key, c.Tag)
}
func ParseTopicPath(s string) (*TopicPath, error) {
	if s == "" {
		return nil, errors.New("param blank")
	}
	i1 := strings.LastIndexByte(s, '/')
	i2 := strings.LastIndexByte(s, ':')
	if i1 <= 0 {
		return nil, errors.New("path err")
	}
	if i2 < 0 {
		i2 = len(s) - 1
	}
	pth := NewTopicPath(s[:i1], s[i1+1:i2], s[i2+1:])
	if pth.NameSpace == "" || pth.Key == "" {
		return nil, errors.New("param err")
	}
	if pth.Tag == "" {
		pth.Tag = "main"
	}
	return pth, nil
}

type TopicInfo struct {
	Path  string
	Safed bool //是否确保发送必到
}
