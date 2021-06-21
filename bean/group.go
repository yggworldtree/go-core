package bean

import (
	"errors"
	"fmt"
	"strings"
)

type CliGroupPath struct {
	Org   string `json:"org"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

func NewCliGroupPath(org, nm string, als ...string) *CliGroupPath {
	c := &CliGroupPath{
		Org:   org,
		Name:  nm,
		Alias: "",
	}
	if len(als) > 0 && als[0] != "" {
		c.Alias = als[0]
	}
	/* if c.Alias == "" {
		c.Alias = "main"
	} */
	return c
}
func (c *CliGroupPath) Path() string {
	return fmt.Sprintf("%s/%s", c.Org, c.Name)
}
func (c *CliGroupPath) String() string {
	return fmt.Sprintf("%s/%s:%s", c.Org, c.Name, c.Alias)
}
func ParseCliGroupPath(s string) (*CliGroupPath, error) {
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
	pth := NewCliGroupPath(s[:i1], s[i1+1:i2], s[i2+1:])
	if pth.Org == "" || pth.Name == "" {
		return nil, errors.New("param err")
	}
	return pth, nil
}
