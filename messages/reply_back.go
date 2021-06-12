package messages

import (
	"context"
	"errors"
	"github.com/yggworldtree/go-core/utils"
	"sync"
	"time"
)

type ReplyCallback struct {
	egn  IEngine
	ctx  context.Context
	cncl context.CancelFunc

	msg   *MessageBox
	okfn  ReplyCallbackOk
	errfn ReplyCallbackErr

	once  sync.Once
	outms time.Duration
}

func NewReplyCallback(egn IEngine, m *MessageBox, outms ...time.Duration) IReply {
	if egn == nil || m == nil || m.Info == nil {
		panic("param err")
	}
	c := &ReplyCallback{
		egn: egn,
		msg: m,
	}
	if len(outms) > 0 {
		c.outms = outms[0]
	}
	if c.outms < time.Millisecond*100 {
		c.outms = time.Second * 20
	}
	c.ctx, c.cncl = context.WithCancel(egn.Ctx())
	return c
}
func (c *ReplyCallback) Message() *MessageBox {
	if c.egn == nil || c.msg == nil {
		return nil
	}
	c.once.Do(c.tmoutCheck)
	return c.msg
}
func (c *ReplyCallback) Ok(fn ReplyCallbackOk) IReply {
	c.okfn = fn
	return c
}
func (c *ReplyCallback) Err(fn ReplyCallbackErr) IReply {
	c.errfn = fn
	return c
}
func (c *ReplyCallback) OkFun() ReplyCallbackOk {
	c.cncl()
	return c.okfn
}
func (c *ReplyCallback) ErrFun() ReplyCallbackErr {
	c.cncl()
	return c.errfn
}
func (c *ReplyCallback) Exec() error {
	if c.egn == nil || c.msg == nil {
		return errors.New("param err")
	}
	c.once.Do(c.tmoutCheck)
	return c.egn.SendForReply(c)
}
func (c *ReplyCallback) tmoutCheck() {
	tm := time.Now()
	go func() {
		for !utils.EndContext(c.ctx) {
			if time.Since(tm) > c.outms {
				c.egn.RmReply(c)
				if c.errfn != nil {
					c.errfn(c.egn, ReplyTimeoutErr)
				}
				break
			}
			time.Sleep(time.Millisecond)
		}
	}()
}
