package main

import (
	"fmt"
	"go_dev/part13/chatbycellnet/cellnet"
	"go_dev/part13/chatbycellnet/cellnet/packet"
	"go_dev/part13/chatbycellnet/cellnet/socket"
	"go_dev/part13/chatbycellnet/chat/proto"
	"go_dev/part13/chatbycellnet/golog"
)

var log = golog.New("main")

func main() {
	queue := cellnet.NewEventQueue()

	peer := socket.NewAcceptor(packet.NewMessageCallback(onMessage),queue)

	peer.Start("127.0.0.1:8801")

	peer.SetName("server")

	queue.StartLoop()

	queue.Wait()

	peer.Stop()
}

func onMessage(ses cellnet.Session,raw interface{}){
	switch ev := raw.(type) {
	case socket.AcceptedEvent:
		log.Infof("client accepted:%d",ses.ID())
	case packet.MsgEvent:
		msg := ev.Msg.(*proto.ChatREQ)
		ack := proto.ChatACK{
			Content: msg.Content,
			ID:      ses.ID(),
		}

		ses.Peer().VisitSession(func(ses cellnet.Session) bool {
			ses.Send(&ack)

			return true
		})

	case socket.SessionClosedEvent:
		fmt.Println("client disconnected:",ses.ID())

	}
}
