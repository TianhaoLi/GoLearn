package main

import "fmt"

type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return other + int(m)
}

func main() {
	var b MyInt

	fmt.Println(b.IsZero())

	b = 1

	fmt.Println(b.Add(2))
}
