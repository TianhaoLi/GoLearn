package main

import "fmt"

func main() {
	var house = "Mailbu Point 10880,90265"

	//ptr类型 *string
	ptr := &house

	fmt.Printf("ptr type : %T\n",ptr)

	fmt.Printf("address:%p\n",ptr)

	value := *ptr

	fmt.Printf("value type: %T\n",value)

	fmt.Printf("value: %s\n",value);
}

