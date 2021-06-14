package messages

import (
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/yggworldtree/go-core/utils"
	"sync"
	"time"
)

type RetryCallback struct {
	ReplyCallback
	once      sync.Once
	retryTime time.Duration
}

func NewRetryCallback(egn IEngine, m *MessageBox, retryTime time.Duration, outms ...time.Duration) *RetryCallback {
	if egn == nil || m == nil || m.Info == nil {
		panic("param err")
	}
	c := &RetryCallback{
		retryTime: retryTime,
	}
	c.egn = egn
	c.msg = m
	if len(outms) > 0 {
		c.outms = outms[0]
	}
	return c
}
func (c *RetryCallback) Exec() error {
	err := c.ReplyCallback.Exec()
	if err != nil {
		return err
	}
	if c.retryTime < time.Second*3 {
		c.retryTime = time.Second * 3
	}
	c.once.Do(c.retryCheck)
	return nil
}
func (c *RetryCallback) retryCheck() {
	go func() {
		tm := time.Now()
		for !utils.EndContext(c.ctx) {
			if time.Since(tm) > c.retryTime {
				tm = time.Now()
				c.egn.SendForReply(c)
				hbtp.Debugf("retryCheck reSend:%s", c.Message().Info.Id)
			}
			time.Sleep(time.Millisecond)
		}
	}()
}