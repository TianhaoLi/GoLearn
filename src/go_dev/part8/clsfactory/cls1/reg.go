package cls1

import (
	"fmt"
	"go_dev/part8/clsfactory/base"
)

type Class1 struct {

}

func (c *Class1) Do() {
	fmt.Println("class1")
}

func init()  {
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}