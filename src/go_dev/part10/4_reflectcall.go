package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

func main() {
	funcValue := reflect.ValueOf(add)

	paramList := []reflect.Value{reflect.ValueOf(10),reflect.ValueOf(20)}

	retList := funcValue.Call(paramList)

	fmt.Println(retList[0].Int())
}
