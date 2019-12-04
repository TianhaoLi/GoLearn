package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string

	//嵌入字段
	float32
	bool

	next *dummy
}

func main() {
	d := reflect.ValueOf(
		dummy{next:&dummy{},
		})

	fmt.Println("NumFiled",d.NumField())

	//获得索引为2的字段
	floatField := d.Field(2)

	fmt.Println("Filed",floatField.Type())

	fmt.Println("FieldByName(\"b\").Type",d.FieldByName("b").Type())

	fmt.Println("FieldByIndex([]int{4,0}.Type()",d.FieldByIndex([]int{4,0}).Type())
}
