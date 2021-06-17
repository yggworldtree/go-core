package main

import (
	"github.com/yggworldtree/go-core/bean"
	"github.com/yggworldtree/go-core/common"
	"testing"
)

func Test1(t *testing.T) {
	s := bean.NewTopicPath("mgr", "cpu_info").String()
	println("msg1:" + s)
	m, e := bean.ParseTopicPath(s)
	if e != nil {
		println("ParseTopicPath err:" + e.Error())
		return
	}
	println("msg2:" + m.String())
}
func Test2(t *testing.T) {
	s := "mgr/cpu_info"
	println("msg1:" + s)
	m, e := bean.ParseTopicPath(s)
	if e != nil {
		println("ParseTopicPath err:" + e.Error())
		return
	}
	println("msg2:" + m.String())
}
func Test3(t *testing.T) {
	s := "123ms"
	ms := common.RegTms.FindAllStringSubmatch(s, -1)[0]
	println(ms)
}
