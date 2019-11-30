package main

import "fmt"

type Data struct {
	complax []int

	instance InnerData

	ptr *InnerData
}

type InnerData struct {
	a int
}

//值传递

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value: %+v\n",inFunc)

	fmt.Printf("inFunc ptr:%p\n",&inFunc)

	return inFunc
}

func main() {
	in := Data{
		complax:  []int{1,2,3},
		instance: InnerData{5,},
		ptr:      &InnerData{1},
	}

	fmt.Printf("in value:%+v\n",in)
	fmt.Printf("in ptr:%p\n",&in)

	out := passByValue(in)

	fmt.Printf("out value:%+v\n",out)
	fmt.Printf("in ptr:%p\n",&out)
}
