package main

import (
	"errors"
	"fmt"
	"os"
)

type fileWriter struct {
	file *os.File
}


//设置写入的文件名
func (f *fileWriter) SetFile(filename string) (err error) {

	if f.file != nil{
		f.file.Close()
	}

	//创建一个文件并保存文件句柄
	f.file,err = os.Create(filename)

	return err
}

func (f *fileWriter) Write(data interface{}) error{
	if f.file == nil {
		//创建失败
		return errors.New("file not create")
	}

	//数据序列成字符串
	str := fmt.Sprintf("%v\n",data)

	_,err := f.file.Write([]byte(str))

	return err
}

//创建文件写入器的实例
func newFileWriter() *fileWriter{
	return &fileWriter{}
}
