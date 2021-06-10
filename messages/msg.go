package messages

import (
	"encoding/json"
	"github.com/yggworldtree/go-core/utils"
)

// MessageBox 长连接 数据包
type MessageBox struct {
	Head *MessageHead `json:"head,omitempty"`
	Body []byte       `json:"body,omitempty"`
}

type MessageHead struct {
	Id        string    `json:"id,omitempty"`
	Type      string    `json:"type,omitempty"`
	NeedReply bool      `json:"needReply,omitempty"`
	Command   string    `json:"command,omitempty"`
	Args      utils.Map `json:"args,omitempty"`
}

func NewMessageBox(cmd string, args ...utils.Map) *MessageBox {
	c := &MessageBox{
		Head: &MessageHead{
			Command: cmd,
			Args:    utils.Map{},
		},
	}
	if len(args) > 0 && args[0] != nil {
		c.Head.Args = args[0]
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
