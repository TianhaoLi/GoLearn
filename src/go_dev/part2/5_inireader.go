package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValue(filename, expectSection, expectKey string) string {
	//1.读取文件
	//2.读取行文本
	//3.读取段
	//4.读取键值

	file, err := os.Open(filename)

	if err != nil {
		return ""
	}

	defer file.Close() //延迟到返回时退出

	reader := bufio.NewReader(file)

	var sectionName string

	for {
		linestr, err := reader.ReadString('\n') //读到换行
		if err != nil {
			break
		}

		linestr = strings.TrimSpace(linestr) //去除两边空格

		if linestr == "" { //忽略空白行
			continue
		}

		if linestr[0] == ';' { //忽略注释
			continue
		}

		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sectionName = linestr[1 : len(linestr)-1]
		} else if sectionName == expectSection {
			pair := strings.Split(linestr, "=")

			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])

				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}

func main() {

	fmt.Println(getValue("example.ini","remote \"origin\"","fetch"))
	fmt.Println(getValue("example.ini","core","hideDotFiles"))
}
