package main

import "os"

func main() {
	exitChan := make(chan int)

	go server("127.0.0.1:3824",exitChan)

	code := <- exitChan


	//标记程序返回值退出
	os.Exit(code)
}
