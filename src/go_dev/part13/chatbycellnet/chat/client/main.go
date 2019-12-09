package main

import (
	"bufio"
	"fmt"
	"go_dev/part13/chatbycellnet/cellnet"
	"go_dev/part13/chatbycellnet/cellnet/packet"
	"go_dev/part13/chatbycellnet/cellnet/socket"
	"go_dev/part13/chatbycellnet/chat/proto"
	"go_dev/part13/chatbycellnet/golog"
	"os"
	"strings"
)

var log = golog.New("main")

func main() {
	//传入一个事件队列
	queue := cellnet.NewEventQueue()

	//创建一个连接器，传入消息处理的响应函数和事件队列
	peer := socket.NewConnector(packet.NewMessageCallback(onMessage),queue)

	peer.Start("127.0.0.1:8801")

	peer.SetName("client")

	queue.StartLoop()

	ReadConsole(func(str string) {
		//使用peer获取会话
		ses := peer.(interface{
			Session() cellnet.Session
		}).Session()


		//发送使用回车输入的字符串
		ses.Send(&proto.ChatREQ{
			Content:str,
		})
	})




}

func ReadConsole(callback func(string))  {
	for{
		text,err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil{
			break
		}

		text = strings.TrimSpace(text)

		//将数据返回调用者
		callback(text)
	}
}

func onMessage(ses cellnet.Session,raw interface{})  {
	switch ev := raw.(type){
	case socket.ConnectedEvent:
		fmt.Println("connected")
	case packet.MsgEvent://解码后消息
		msg := ev.Msg.(*proto.ChatACK)

		log.Infof("sid:%d say:%s",msg.ID,msg.Content)

	case socket.SessionClosedEvent:
		fmt.Println("disconnected")

	}
}


