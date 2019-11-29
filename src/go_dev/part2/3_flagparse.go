package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode","","process mode")

func main() {
	flag.Parse()

	fmt.Println(*mode)
}