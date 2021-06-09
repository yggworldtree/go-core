package messages

import (
	"context"
	"github.com/yggworldtree/go-core/bean/clientBean"
)

type RpleyCallbackOk func(c IEngine, m *clientBean.MessageBox)
type RpleyCallbackErr func(c IEngine, errs ...interface{})
type IReply interface {
	Message() *clientBean.MessageBox
	Ok(RpleyCallbackOk) IReply
	Err(RpleyCallbackErr) IReply
	OkFun() RpleyCallbackOk
	ErrFun() RpleyCallbackErr
	Exec() error
}
type ISender interface {
	Sends(msg *clientBean.MessageBox) error
}
type IEngine interface {
	ISender
	Ctx() context.Context
	SendForReply(e IReply) error
	RmReply(e IReply) error
}
