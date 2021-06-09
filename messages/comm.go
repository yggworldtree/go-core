package messages

import (
	"context"
	"errors"
	"github.com/yggworldtree/go-core/bean/clientBean"
)

var ReplyTimeoutErr = errors.New("time out")

type ReplyCallbackOk func(c IEngine, m *clientBean.MessageBox)
type ReplyCallbackErr func(c IEngine, errs error)
type IReply interface {
	Message() *clientBean.MessageBox
	Ok(ReplyCallbackOk) IReply
	Err(ReplyCallbackErr) IReply
	OkFun() ReplyCallbackOk
	ErrFun() ReplyCallbackErr
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
