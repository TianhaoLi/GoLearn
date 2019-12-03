package main

import "reflect"

type State interface {
	//状态名
	Name() string

	//该状态是否允许转移
	EnableSameTransit()bool

	//响应状态开始时
	OnBegin()

	OnEnd()

	//判断能否转移
	CanTransitTo(name string) bool
}

//根据实例获取状态名字

func StateName(s State) string {
	if s == nil{
		return "none"
	}

	return reflect.TypeOf(s).Elem().Name()
}


