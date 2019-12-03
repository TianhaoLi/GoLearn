package main

import "fmt"

type Actor struct {

}

func (a *Actor) OnEvent(param interface{})  {
	fmt.Println("actor event:",param)
}

//全局事件
func GlobalEvent(param interface{})  {
	fmt.Println("global event:",param)
}

func main() {
	a := new(Actor)

	//注册一个onskill的回调
	RegisterEvent("OnSkill",a.OnEvent)

	//再次注册 全局事件
	RegisterEvent("OnSkill",GlobalEvent)

	CallEvent("OnSkill",100)
}
