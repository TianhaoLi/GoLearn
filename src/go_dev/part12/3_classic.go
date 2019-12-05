package main

import (
	"fmt"
)

func simpleHash(str string) (ret int) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		ret +=int(c)
	}

	return
}

type classicQueryKey struct {
	Name string
	Age int
}

func (c *classicQueryKey) hash() int{
	return simpleHash(c.Name) + c.Age*1000000
}

var mapper = make(map[int][]*Profile)

func buildIndex(list []*Profile)  {
	for _,profile := range list{
		key := classicQueryKey{profile.Name,profile.Age}

		existValue := mapper[key.hash()]

		existValue = append(existValue,profile)

		mapper[key.hash()] = existValue
	}
}

//查询逻辑
func queryData(name string, age int)  {
	KeyToQuery := classicQueryKey{name,age	}

	resultList := mapper[KeyToQuery.hash()]

	for _, result := range resultList{
		if result.Name == name && result.Age == age{
			fmt.Println(result)
			return
		}
	}

	fmt.Println("not found")
}