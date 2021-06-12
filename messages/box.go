package messages

import (
	"context"
	"errors"
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"net"
	"net/url"
)

type msgInfo struct {
	LenId   uint8
	LenCmd  uint16
	LenArg  uint16
	LenHead uint32
	LenBody uint32
}

func ReadMessageBox(ctx context.Context, conn net.Conn, cfg *hbtp.Config) (*MessageBox, error) {
	info := &msgInfo{}
	infoln, _ := hbtp.Size4Struct(info)
	ctx, _ = context.WithTimeout(ctx, cfg.TmsInfo)
	bts, err := hbtp.TcpRead(ctx, conn, uint(infoln))
	if err != nil {
		return nil, err
	}
	err = hbtp.Byte2Struct(bts, info)
	if err != nil {
		return nil, err
	}
	rt := &MessageBox{
		Info: &MessageInfo{
			Args: url.Values{},
		},
	}
	bts, err = hbtp.TcpRead(ctx, conn, uint(info.LenId))
	if err != nil {
		return nil, err
	}
	rt.Info.Id = string(bts)
	ctx, _ = context.WithTimeout(ctx, cfg.TmsHead)
	if info.LenCmd > 0 {
		bts, err = hbtp.TcpRead(ctx, conn, uint(info.LenCmd))
		if err != nil {
			return nil, err
		}
		rt.Info.Command = string(bts)
	}
	if info.LenArg > 0 {
		bts, err = hbtp.TcpRead(ctx, conn, uint(info.LenArg))
		if err != nil {
			return nil, err
		}
		args, err := url.ParseQuery(string(bts))
		if err == nil {
			rt.Info.Args = args
		}
	}
	if info.LenHead > 0 {
		rt.Head, err = hbtp.TcpRead(ctx, conn, uint(info.LenHead))
		if err != nil {
			return nil, err
		}
	}

	ctx, _ = context.WithTimeout(ctx, cfg.TmsBody)
	if info.LenBody > 0 {
		rt.Body, err = hbtp.TcpRead(ctx, conn, uint(info.LenBody))
		if err != nil {
			return nil, err
		}
	}
	return rt, nil
}
func WriteMessageBox(conn net.Conn, msg *MessageBox) error {
	if conn == nil {
		return errors.New("param err1")
	}
	if msg.Info == nil {
		return errors.New("param err2")
	}
	if msg.Head == nil && msg.header != nil {
		msg.Head = msg.header.ToBytes()
	}
	var args string
	if msg.Info.Args != nil {
		args = msg.Info.Args.Encode()
	}
	info := &msgInfo{
		LenId:   uint8(len(msg.Info.Id)),
		LenCmd:  uint16(len(msg.Info.Command)),
		LenArg:  uint16(len(args)),
		LenHead: uint32(len(msg.Head)),
		LenBody: uint32(len(msg.Body)),
	}

	bts, err := hbtp.Struct2Byte(info)
	if err != nil {
		return err
	}
	_, err = conn.Write(bts)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg.Info.Id))
	if err != nil {
		return err
	}
	if info.LenCmd > 0 {
		_, err = conn.Write([]byte(msg.Info.Command))
		if err != nil {
			return err
		}
	}
	if info.LenArg > 0 {
		_, err = conn.Write([]byte(args))
		if err != nil {
			return err
		}
	}
	if info.LenHead > 0 {
		_, err = conn.Write(msg.Head)
		if err != nil {
			return err
		}
	}
	if info.LenBody > 0 {
		_, err = conn.Write(msg.Body)
		if err != nil {
			return err
		}
	}
	return nil
}
