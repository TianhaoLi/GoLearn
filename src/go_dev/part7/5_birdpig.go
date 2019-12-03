package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {

}

func (b *bird) Fly() {
	fmt.Println("bird:fly")
}

func (b *bird) Walk()  {
	fmt.Println("bird:walk")
}

type pig struct {

}

func (p *pig) Walk()  {
	fmt.Println("pig:walk")
}

func main()  {
	animals := map[string]interface{}{
		"birds":new(bird),
		"pigs":new(pig),
	}

	for name,obj := range animals{


		//断言
		f,isFlyer := obj.(Flyer)

		w,isWalker := obj.(Walker)

		fmt.Printf("name:%s isFlayer:%v isWalker:%v\n",name,isFlyer,isWalker)

		if isFlyer{
			f.Fly()
		}

		if isWalker{
			w.Walk()
		}
	}
}