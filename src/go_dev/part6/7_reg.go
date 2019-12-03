package main

//字符串 映射的函数切片
var eventByName = make(map[string][]func(interface{}))

//注册事件 事件的名字和回调
func RegisterEvent(name string,callback func(interface{}))  {
	list := eventByName[name]

	//在列表中添加函数
	list = append(list,callback)

	eventByName[name] = list
}

//调用事件

func CallEvent(name string, param interface{}) {
	list := eventByName[name]

	for _, callback := range list{
		callback(param)
	}
}