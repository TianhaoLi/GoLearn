package main

import (
	"fmt"
	"net"
	"strconv"
)

func Connector(address string,sendTimes int	)  {
	conn,err := net.Dial("tcp",address)

	if err != nil{
		fmt.Println(err)
		return
	}

	for i:= 0; i<sendTimes;i++{
		str := strconv.Itoa(i)

		if err := writePacket(conn,[]byte(str));err != nil{
			fmt.Println(err)
			break
		}
	}
}
