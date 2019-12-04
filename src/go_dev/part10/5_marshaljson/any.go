package main

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(),10))
	case reflect.Slice:
		return writeSlice(buff,value)
	case reflect.Struct:
		return writeStruct(buff,value)
	default:
		return errors.New("unsupport kind :"+ value.Kind().String())

	}
	return nil
}