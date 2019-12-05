package main

import (
	"bufio"
	"net"
)

//连接会话的逻辑
func handleSession(conn net.Conn, callback func(net.Conn, []byte)bool)  {
	dataReader := bufio.NewReader(conn)

	for{
		pkt,err := readPacket(dataReader)

		if err !=nil || !callback(conn,pkt.Body){
			conn.Close()
			break
		}
	}
}
