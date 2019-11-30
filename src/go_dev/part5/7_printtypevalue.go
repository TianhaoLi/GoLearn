package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer

	for _,s := range slist{
		str := fmt.Sprintf("%v",s)

		var typeString string

		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"

		}

		b.WriteString("value: ")

		b.WriteString(str)

		b.WriteString(" type: ")

		b.WriteString(typeString)

		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	fmt.Println(printTypeValue(100,"str",true))
}