package main

import "fmt"

func main() {
	type Skill struct {
		Name string
		Level int
	}

	type Actor struct {
		Name string
		Age int

		Skills []Skill
	}

	//填充基本角色数据
	a := Actor{
		Name:   "cow boy",
		Age:    37,
		Skills: []Skill{
			{Name:"Roll and Roll",Level:1},
			{Name:"Flash your dog eye",Level:2},
			{Name:"Time to have lunch",Level:3},
		},
	}

	if result,err := MarshalJson(a);err == nil{
		fmt.Println(result)
	}else{
		fmt.Println(err)
	}
}
