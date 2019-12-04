package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 3; i >= 0; i -- {
			ch <- i

			time.Sleep(time.Second)
		}
	}()

	for data := range ch{
		fmt.Println(data)

		if data == 0{
			break
		}
	}
}
