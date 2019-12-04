package main

import (
	"bytes"
	"reflect"
)

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")
	for s := 0 ; s < value.Len();s++{
		sliceValue := value.Index(s)
		writeAny(buff,sliceValue)

		if s < value.Len() -1{
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return nil
}
