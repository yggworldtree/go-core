package main

import (
	"github.com/yggworldtree/go-core/bean"
	"github.com/yggworldtree/go-core/common"
	"testing"
	"time"
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
func Test4(t *testing.T) {
	n := 0
	tmr := time.NewTicker(time.Second)
	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			if n > 5 {
				tmr.Stop()
				break
			}
		}
	}()
	for {
		<-tmr.C
		n++
		println("tick:", n)
		if n > 10 {
			break
		}
	}
}
