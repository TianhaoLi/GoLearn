package main

import "fmt"

func main() {
	const elementCount = 1000

	srcData := make([]int,elementCount)

	for i :=0;i<elementCount;i++  {
		srcData[i] = i
	}

	refData :=srcData

	copyData := make([]int,elementCount)

	copy(copyData,srcData)

	srcData[0] = 999

	fmt.Println(refData[0])

	fmt.Println(copyData[0],copyData[elementCount -1])

	copy(copyData,srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d",copyData[i])
	}
}
