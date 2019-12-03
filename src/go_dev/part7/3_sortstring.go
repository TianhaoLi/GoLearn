package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len()int {
	return len(m)
}

func (m MyStringList) Less(i,j int)	bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i,j int){
	m[i],m[j] = m[j],m[i]
}

func main() {
	names := MyStringList{
		"3. Tripple kill",
		"5. Penta Kill",
		"4. Quadra Kill",
		"1. First Kill",
	}

	sort.Sort(names)

	for _, v := range names{
		fmt.Printf("%s\n",v)
	}
}