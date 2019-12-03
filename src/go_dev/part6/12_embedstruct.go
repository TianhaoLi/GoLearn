package main

import "fmt"

type Wheels struct {
	Size int
}


type Cars struct {
	Wheels
	Engine struct{
		Type string
		Power int
	}
}

func main() {
	c:= Cars{
		Wheels:  Wheels{
			Size:18,
		},
		Engine: struct {
			Type  string
			Power int
		}{Type:"1.4T" , Power:143 },
	}

	fmt.Println("%+v\n",c)
}
