package main

import (
	"github.com/pkg/profile"
	"time"
)

func joinSlice() []string {
	var arr []string

	for i := 0; i < 100000; i++ {
		arr = append(arr,"arr")
	}

	return arr
}

func main() {
	stopper := profile.Start(profile.CPUProfile,profile.ProfilePath("."))

	defer stopper.Stop()

	joinSlice()

	time.Sleep(time.Second)
}
