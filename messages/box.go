package messages

import (
	"context"
	"errors"
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"net"
	"net/url"
)

var lenMsgInfo=hbtp.SizeOf(new(msgInfo))
type msgInfo struct {
	LenId   uint8
	LenSndr uint16
	LenCmd  uint16
	LenArg  uint16
	LenHead uint32
	LenBody uint32
}

func ReadMessageBox(ctx context.Context, conn net.Conn, cfg *hbtp.Config) (*MessageBox, error) {
	info := &msgInfo{}
	ctx, _ = context.WithTimeout(ctx, cfg.TmsInfo)
	bts, err := hbtp.TcpRead(ctx, conn, uint(lenMsgInfo))
	if err != nil {
		return nil, err
	}
	err = hbtp.Byte2Struct(bts, info)
	if err != nil {
		return nil, err
	}
	rt := &MessageBox{
		Info: &MessageInfo{},
	}
	bts, err = hbtp.TcpRead(ctx, conn, uint(info.LenId))
	if err != nil {
		return nil, err
	}
	rt.Info.Id = string(bts)
	bts, err = hbtp.TcpRead(ctx, conn, 1)
	if err != nil {
		return nil, err
	}
	rt.Info.Flags = int8(bts[0])
	ctx, _ = context.WithTimeout(ctx, cfg.TmsHead)
	if info.LenSndr > 0 {
		bts, err = hbtp.TcpRead(ctx, conn, uint(info.LenSndr))
		if err != nil {
			return nil, err
		}
		rt.Info.Sender = string(bts)
	}
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
	hds := msg.Heads()
	var args string
	if msg.Info.Args != nil {
		args = msg.Info.Args.Encode()
	}
	info := &msgInfo{
		LenId:   uint8(len(msg.Info.Id)),
		LenSndr: uint16(len(msg.Info.Sender)),
		LenCmd:  uint16(len(msg.Info.Command)),
		LenArg:  uint16(len(args)),
		LenHead: uint32(len(hds)),
		LenBody: uint32(len(msg.Body)),
	}

	bts, err := hbtp.Struct2ByteLen(info,lenMsgInfo)
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
	_, err = conn.Write([]byte{byte(msg.Info.Flags)})
	if err != nil {
		return err
	}
	if info.LenSndr > 0 {
		_, err = conn.Write([]byte(msg.Info.Sender))
		if err != nil {
			return err
		}
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
		_, err = conn.Write(hds)
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
