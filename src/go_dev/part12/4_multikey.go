package main

import (
	"fmt"
	"go/ast"
)

type queryKey struct {
	Name string
	Age int
}

var mapper  = make(map[queryKey]*Profile)

func buildIndex(list []*Profile){
	for _,profile:= range list{
		key := queryKey{
			Name: profile.name,
			Age:  profile.name,
		}

		mapper[key] = profile
	}
}

//查询逻辑
func queryData(name string, age int) {
	key := queryKey{name,age}

	result,ok := mapper[key]

	if ok{
		fmt.Println(result)
	}else {
		fmt.Println("not found")
	}
}

