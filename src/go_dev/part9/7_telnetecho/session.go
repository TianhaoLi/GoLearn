package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleSession(conn net.Conn,exitChan chan int)  {
	fmt.Println("Session started:")

	reader := bufio.NewReader(conn)

	for{
		str,err := reader.ReadString('\n')

		if err == nil{
			str = strings.TrimSpace(str)

			//处理telnet指令
			if !processTelnetCommand(str,exitChan){
				conn.Close()
				break
			}

			//echo 逻辑 发什么数据 原样返回
			conn.Write([]byte(str+"\r\n"))
		}else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}
