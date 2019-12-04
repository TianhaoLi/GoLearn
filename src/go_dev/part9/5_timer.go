package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	stopper := time.NewTimer(time.Second * 2)

	var i int;
	for  {
		select {
		case <- stopper.C:
			fmt.Println("stop")
			goto StopHere

		case <- ticker.C:
			i++
			fmt.Println("tic",i)

		}
	}

StopHere:
	fmt.Println("done")
}
