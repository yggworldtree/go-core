package messages

import (
	"encoding/json"
	"github.com/yggworldtree/go-core/utils"
	"net/url"
)

// MessageBox 长连接 数据包
type MessageBox struct {
	Info   *MessageInfo `json:"info,omitempty"`
	Head   []byte       `json:"head,omitempty"`
	Body   []byte       `json:"body,omitempty"`
	header *utils.Map
}

type MessageInfo struct {
	Id      string     `json:"id,omitempty"`
	Command string     `json:"command,omitempty"`
	Args    url.Values `json:"args,omitempty"`
}

func NewMessageBox(cmd string, args ...url.Values) *MessageBox {
	c := &MessageBox{
		Info: &MessageInfo{
			Command: cmd,
			Args:    url.Values{},
		},
	}
	if len(args) > 0 && args[0] != nil {
		c.Info.Args = args[0]
	}
	return c
}
func (c *MessageBox) PutBody(o interface{}) error {
	if o == nil {
		return nil
	}
	switch o.(type) {
	case []byte:
		c.Body = o.([]byte)
	case string:
		c.Body = []byte(o.(string))
	default:
		bts, err := json.Marshal(o)
		if err != nil {
			return err
		}
		c.Body = bts
	}
	return nil
}
func (c *MessageBox) Header() *utils.Map {
	if c.header == nil {
		c.header = utils.NewMaps(c.Head)
	}
	return c.header
}
