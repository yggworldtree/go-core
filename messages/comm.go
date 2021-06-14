package messages

import (
	"context"
	"errors"
)

var ReplyTimeoutErr = errors.New("time out")

type ReplyCallbackOk func(c IEngine, m *ReplyInfo)
type ReplyCallbackErr func(c IEngine, errs error)
type IReply interface {
	Message() *MessageBox
	Ok(ReplyCallbackOk) IReply
	Err(ReplyCallbackErr) IReply
	OkFun() ReplyCallbackOk
	ErrFun() ReplyCallbackErr
}
type ISender interface {
	Sends(msg *MessageBox) error
}
type IEngine interface {
	ISender
	Ctx() context.Context
	SendForReply(e IReply) error
	RmReply(e IReply) error
}
