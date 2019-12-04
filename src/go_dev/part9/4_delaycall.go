package main

import (
	"fmt"
	"time"
)

func main() {
	//声明一个退出的通道
	exit := make(chan int)

	fmt.Println("start")

	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")

		exit <- 0
	})

	//等待结束
	<-exit
}
