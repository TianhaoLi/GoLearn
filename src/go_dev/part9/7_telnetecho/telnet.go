package main

import (
	"fmt"
	"strings"
)

func processTelnetCommand(str string, exitChan chan int) bool {
	//@close 前缀
	if strings.HasPrefix(str,"@close"){
		fmt.Println("Session closed")

		return false
	}else if strings.HasPrefix(str,"@shutdown"){
		fmt.Println("Server shutdown")

		exitChan <- 0

		return false
	}

	fmt.Println(str)

	return true
}
