package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const(
	None HeroKind = iota
	Tank
	Assasin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

func (s Heros) Len() int{
	return len(s)
}

func (s Heros) Less(i, j int)bool {
	if s[i].Kind != s[j].Kind{
		return s[i].Kind < s[j].Kind
	}

	return s[i].Name < s[j].Name
}

func (s Heros) Swap(i,j int)  {
	s[i],s[j] = s[j],s[i]
}

func main() {
	heros := Heros{
		&Hero{"吕布",Tank},
		&Hero{"李白",Assasin},
		&Hero{"妲己",Mage},
		&Hero{"貂蝉",Assasin},
		&Hero{"关羽",Tank},
		&Hero{"诸葛亮",Mage},
	}

	//sort.Sort(heros)

	//go 1.8新姿势
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind{
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _,v:= range heros{
		fmt.Printf("%+v\n",v)
	}
}