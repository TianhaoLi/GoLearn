package main

import "fmt"

type Alipay struct {

}

func (a *Alipay) CanUseFaceID() {

}

type Cash struct {

}

func (a *Cash) Stolen() {
}

type CantainCanUseFaceID interface {
	CanUseFaceID()
}

type CantainCanStolen interface {
	Stolen()
}

func print(payMethod interface{})  {
	switch payMethod.(type) {
		case CantainCanUseFaceID:
			fmt.Printf("%T can use faceid\n",payMethod)
		case CantainCanStolen:
			fmt.Printf("%T may be stolen\n",payMethod)
	}
}

func main() {
	print(new(Alipay))

	print(new(Cash))
}