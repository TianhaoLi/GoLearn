package main

import "fmt"

func createLogger() *Logger {
	l := NewLogger()

	cw := newConsoleWriter()

	l.RegisterWriter(cw)

	fw := newFileWriter()

	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}

	l.RegisterWriter(fw)

	return l
}

func main() {
	l := createLogger()

	l.Log("hello")
}
