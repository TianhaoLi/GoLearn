package main

import (
	"fmt"
	"net"
)

func server(address string, exitChan chan int)  {
	l,err := net.Listen("tcp",address)

	if err != nil{
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("listen:"+address)

	defer l.Close()


	//侦听循环
	for{
		//新连接没有来是阻塞的
		conn,err := l.Accept()

		if err !=nil{
			fmt.Println(err.Error())
			continue
		}

		go handleSession(conn,exitChan)
	}
}
