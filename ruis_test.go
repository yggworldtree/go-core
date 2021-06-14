package go_core

import (
	"fmt"
	"github.com/yggworldtree/go-core/bean"
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

type Person struct {
}

func (this *Person) Eat() {
	fmt.Println("Person Eat")
}

func (this *Person) Run() {
	fmt.Println("Person Run")
}

func (this *Person) Sleep() {
	fmt.Println("Person Sleep")
}

type Man struct {
	Person
}

func (this *Man) Eat() {
	// 类似于Java的 super.Eat()
	fmt.Println("Man Eat")
	this.Person.Eat()
}

func (this *Man) Run() {
	fmt.Println("Man Run")
}
func Test3(t *testing.T) {
	m := &Man{}
	m.Eat()
	m.Run()
	m.Sleep()
}
