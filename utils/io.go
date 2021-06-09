package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net"
)

func EndContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func Md5String(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

func TcpRead(ctx context.Context, conn net.Conn, ln uint) ([]byte, error) {
	if conn == nil || ln <= 0 {
		return nil, errors.New("handleRead ln<0")
	}
	rn := uint(0)
	rt := make([]byte, ln)
	for {
		if EndContext(ctx) {
			return nil, errors.New("context dead")
		}
		n, err := conn.Read(rt[rn:])
		if n > 0 {
			rn += uint(n)
		}
		if rn >= ln {
			break
		}
		if err != nil {
			return nil, err
		}
		if n <= 0 {
			return nil, errors.New("conn abort")
		}
	}
	return rt, nil
}
