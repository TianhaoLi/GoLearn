package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y==2 {
				goto breakHere
			}
		}
	}

	return

	breakHere:
		fmt.Println("done")
}
