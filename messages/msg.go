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
	Id     string `json:"id,omitempty"`
	Sender string `json:"sender,omitempty"`

	/*
		Flags:  0  |  0  |  0  |  0  |  0  |  0  |     0     |   0
		Desc:   no    no    no    no    no    no   broadcast   needReply
	*/
	Flags   int8       `json:"flags,omitempty"`
	Command string     `json:"command,omitempty"`
	Args    url.Values `json:"args,omitempty"`
}

func NewMessageBox(cmd string, args ...url.Values) *MessageBox {
	c := &MessageBox{
		Info: &MessageInfo{
			Command: cmd,
		},
	}
	if len(args) > 0 && args[0] != nil {
		c.Info.Args = args[0]
	}
	return c
}
func (c *MessageBox) SetArg(k, v string) {
	if c.Info == nil {
		panic("pleas use new")
	}
	if c.Info.Args == nil {
		c.Info.Args = url.Values{}
	}
	c.Info.Args.Set(k, v)
}
func (c *MessageBox) PutHead(o interface{}) error {
	if o == nil {
		return nil
	}
	switch o.(type) {
	case []byte:
		c.Head = o.([]byte)
	case string:
		c.Head = []byte(o.(string))
	default:
		bts, err := json.Marshal(o)
		if err != nil {
			return err
		}
		c.Head = bts
	}
	return nil
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
func (c *MessageBox) Heads() []byte {
	if c.Head == nil && c.header != nil {
		return c.header.ToBytes()
	}
	return c.Head
}
func (c *MessageBox) Header() *utils.Map {
	if c.header == nil {
		c.header = utils.NewMaps(c.Head)
	}
	return c.header
}

type ReplyInfo struct {
	Status string `json:"status,omitempty"`
	Head   []byte `json:"head,omitempty"`
	Body   []byte `json:"body,omitempty"`
	header *utils.Map
}

func NewReplyInfo(stat string, bts ...[]byte) *ReplyInfo {
	c := &ReplyInfo{Status: stat}
	if len(bts) > 0 {
		c.Body = bts[0]
	}
	if len(bts) > 1 {
		c.Body = bts[1]
	}
	return c
}
func (c *ReplyInfo) Heads() []byte {
	if c.Head == nil && c.header != nil {
		return c.header.ToBytes()
	}
	return c.Head
}
func (c *ReplyInfo) Header() *utils.Map {
	if c.header == nil {
		c.header = utils.NewMaps(c.Head)
	}
	return c.header
}
