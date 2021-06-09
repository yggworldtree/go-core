package messages

import (
	"context"
	"errors"
	"github.com/yggworldtree/go-core/bean/clientBean"
	"github.com/yggworldtree/go-core/utils"
	"sync"
	"time"
)

type ReplyCallback struct {
	egn  IEngine
	ctx  context.Context
	cncl context.CancelFunc

	msg   *clientBean.MessageBox
	okfn  RpleyCallbackOk
	errfn RpleyCallbackErr

	once  sync.Once
	outms time.Duration
}

func NewReplyCallback(egn IEngine, m *clientBean.MessageBox, outms ...time.Duration) IReply {
	if egn == nil || m == nil || m.Head == nil {
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
func (c *ReplyCallback) Message() *clientBean.MessageBox {
	if c.egn == nil || c.msg == nil {
		return nil
	}
	c.once.Do(c.tmoutCheck)
	return c.msg
}
func (c *ReplyCallback) Ok(fn RpleyCallbackOk) IReply {
	c.okfn = fn
	return c
}
func (c *ReplyCallback) Err(fn RpleyCallbackErr) IReply {
	c.errfn = fn
	return c
}
func (c *ReplyCallback) OkFun() RpleyCallbackOk {
	c.cncl()
	return c.okfn
}
func (c *ReplyCallback) ErrFun() RpleyCallbackErr {
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
					c.errfn(c.egn, errors.New("time out"))
				}
				break
			}
			time.Sleep(time.Millisecond)
		}
	}()
}
