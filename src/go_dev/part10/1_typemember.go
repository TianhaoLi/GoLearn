package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {
		Name string

		Type int `json:"type" id:"100"`
	}

	ins := cat{
		Name: "mimi",
		Type: 1,
	}

	typeOfCat := reflect.TypeOf(ins)

	for i := 0; i < typeOfCat.NumField(); i++ {
		filedType := typeOfCat.Field(i)

		fmt.Printf("name:%v tag:'%v'\n",filedType.Name,filedType.Tag)
	}

	if catType ,ok := typeOfCat.FieldByName("Type"); ok{
		fmt.Println(catType.Tag.Get("json"),catType.Tag.Get("id"))
	}
}
