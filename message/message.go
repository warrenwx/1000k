package message

import (
	"sync"
)

type MessageQueue struct {
	// 写queue时要加锁
	Mutex sync.Mutex
	// 队列中当前消息个数
	MCount uint32
	// 消息队列
	Queue chan *Message
}

type Message struct {
	MType uint8
	MLength uint
	MValue uintptr
}

type MessageC1 struct {
	a int32
	b string
	c map[string]string
	d string
}

func (p *MessageC1) FillMessageC1() {
	p.a = 1
	p.b = "string"
	c := make(map[string]string)
	c["a"] = "av"
	c["b"] = "bv"
	p.c = c
}