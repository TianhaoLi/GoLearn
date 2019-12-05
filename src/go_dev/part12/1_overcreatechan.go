package main

import (
	"fmt"
	"runtime"
)

func consumer(ch chan int) {
	for{
		data := <- ch

		//需要设置gorounte 的退出条件，否则会无限制创建
		if data == 0{
			break
		}

		fmt.Println(data)
	}

}

func main() {
	ch := make(chan int)

	for{
		var dummy string

		fmt.Scan(&dummy)

		//需要设置gorounte 的退出条件，否则会无限制创建
		if dummy == "quit"{
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- 0
			}
			continue
		}

		//并发执行
		go consumer(ch)

		fmt.Println("goroutines:",runtime.NumGoroutine())
	}
}
