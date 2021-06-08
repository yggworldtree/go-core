package clientBean

import (
	"encoding/json"
	"github.com/yggworldtree/go-core/utils"
)

// MessageBox 长连接 数据包
type MessageBox struct {
	Head *MessageHead
	Body []byte
}

type MessageHead struct {
	Id      string
	Type    string
	Command string
	Args    utils.Map
}
type MessageReply struct {
	Id     string
	Status string
}

func NewMessageBox() *MessageBox {
	return &MessageBox{
		Head: &MessageHead{
			Args: utils.Map{},
		},
	}
}
func (c *MessageBox) PutBody(o interface{}) error {
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
