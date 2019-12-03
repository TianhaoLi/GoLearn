package main

import "fmt"

func main() {
	p := NewPlayer(0.5)

	p.MoveTo(Vec2{3,1	})

	for !p.IsArrived(){
		p.Update()

		fmt.Println(p.Pos())
	}
}
