package main

import (
	"fmt"
	"net"
	"sync"
)

type Acceptor struct {
	//保存侦听器
	l net.Listener

	//侦听器停止同步
	wg sync.WaitGroup

	//连接的数据回调
	OnSessionData func(net.Conn,[]byte)bool
}

//异步开始侦听
func (a *Acceptor) Start(address string) {
	go a.listen(address)
}

func (a *Acceptor) listen(address string) {
	a.wg.Add(1)

	defer a.wg.Done()

	var err error

	a.l,err = net.Listen("tcp",address)

	if err !=nil{
		fmt.Println(err.Error())
		return
	}

	for{
		conn,err := a.l.Accept()

		if err != nil{
			break
		}

		go handleSession(conn,a.OnSessionData)
	}
}

func (a *Acceptor) Stop() {
	a.l.Close()
}

func (a *Acceptor) Wait() {
	a.wg.Wait()
}

func NewAcceptor() *Acceptor {
	return &Acceptor{}
}
