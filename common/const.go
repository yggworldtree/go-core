package common

import "regexp"

const (
	TimeFmt  = "2006-01-02 15:04:05"
	TimeFmts = "2006-01-02"
)

const (
	MaxTopicLen   = 1024 * 1024 * 10
	MaxCliHeadLen = 1024 * 1024 * 100
	MaxCliBodyLen = 1024 * 1024 * 1024
)

var (
	RegName      = regexp.MustCompile("[`~!@#$%^&*()\\+=<>?:\"{}|,\\.\\/;'\\\\[\\]·~！@#￥%……&*（）——={}|《》？：“”【】、；‘'，。、]")
	RegNameSpace = regexp.MustCompile("[`~!@#$%^&*()\\+=<>?:\"{}|,\\.;'\\\\[\\]·~！@#￥%……&*（）——={}|《》？：“”【】、；‘'，。、]")
)
