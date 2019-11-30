package main

import (
	"fmt"
	"strings"
)

func StringProccess(list []string,chain []func(string) string)  {
	for index,str := range list{
		result := str
//逐个取出 list切片中的string . 每个string逐一经过chain函数列表的里的函数处理
		for _,proc := range chain{
			result = proc(result)
			}
			list[index] = result
	}
}

//自定义处理函数

func removePrefix(str string) string{
	return strings.TrimPrefix(str,"go")
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	//C处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	StringProccess(list,chain)

	for _,str := range list{
		fmt.Println(str)
	}
}