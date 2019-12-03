package main

import "fmt"

type ParseError struct {
	Filename string
	Line int
}

func (e *ParseError)  Error() string{
	return fmt.Sprintf("%s:%d",e.Filename,e.Line)
}

func newParseError(filename string, line int) error {
	return &ParseError{filename,line}
}

func main() {
	var e error

	e = newParseError("3_main.go",1)
	fmt.Println(e.Error())

	switch detail := e.(type){
	case *ParseError:
		fmt.Printf("filename:%s Line:%d\n",detail.Filename,detail.Line)
	default:
		fmt.Println("otherError")
	}
}