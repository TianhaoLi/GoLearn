package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//避免不必要的地方使用通道
/*
func socketRecv(conn net.Conn,exitChan chan string)  {
	buff := make([]byte,1024)

	for{
		//不停地接收数据

		_,err := conn.Read(buff)

		if err !=nil{
			 break
		}
	}
	exitChan <- "recv exit"
}
*/

//优化 通道内部也是锁实现的。用等待锁来优化锁操作的消耗
func socketRecv(conn net.Conn,wg *sync.WaitGroup)  {
	buff := make([]byte,1024)

	for{
		//不停地接收数据

		_,err := conn.Read(buff)

		if err !=nil{
			break
		}
	}
	wg.Done()
}

/*
func main() {
	//连接一个地址
	conn,err := net.Dial("tcp","www.163.com:80")

	if err != nil{
		fmt.Println(err)
		return
	}

	exit := make(chan  string)
	go socketRecv(conn,exit)

	time.Sleep(time.Second)

	conn.Close()

	fmt.Println(<-exit)
}
 */

//优化 通道内部也是锁实现的。用等待锁来优化锁操作的消耗
func main() {
	//连接一个地址
	conn,err := net.Dial("tcp","www.163.com:80")

	if err != nil{
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go socketRecv(conn,&wg)

	time.Sleep(time.Second)

	conn.Close()

	wg.Wait()
	fmt.Println("recv done")
}
