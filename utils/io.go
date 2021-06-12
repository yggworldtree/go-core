package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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
	m5 := md5.New()
	m5.Write([]byte(data))
	md5Data := m5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}
