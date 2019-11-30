package main

import "fmt"

//闭包的记忆效应

func Accumulate(value int) func() int {
	return func() int {
		value ++

		return value
	}
}

func main() {
	accumulator := Accumulate(1)

	fmt.Println(accumulator())

	fmt.Println(accumulator())

	accumulator2 := Accumulate(10)

	fmt.Println(accumulator2())
}
