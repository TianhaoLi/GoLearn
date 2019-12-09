package proto

import (
	"go_dev/part13/chatbycellnet/cellnet"
	_"go_dev/part13/chatbycellnet/cellnet/codec/json"
	"reflect"
)

type ChatREQ struct {
	Content string  //发送的内容
}


//服务器通知客户端有人发送聊天
type ChatACK struct {
	Content string //发送的内容
	ID int64 //发送者的id
}

func init(){
	cellnet.RegisterMessageMeta("json","proto.ChatREQ",reflect.TypeOf((*ChatREQ)(nil)).Elem(),1)
	cellnet.RegisterMessageMeta("json","proto.ChatACK",reflect.TypeOf((*ChatACK)(nil)).Elem(),2)
}