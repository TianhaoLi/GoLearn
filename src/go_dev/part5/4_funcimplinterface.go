package main

import "fmt"

type Invoker interface {
	Call(interface{})
}


type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from sturct",p)
}

func main() {
	var invoker Invoker

	s := new(Struct)

	invoker = s

	invoker.Call("hello")
}