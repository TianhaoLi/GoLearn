package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Packet struct {
	Size uint16 //包体大小
	Body []byte //包体数据
}


func writePacket(dataWriter io.Writer,data []byte)error{
	var buffer bytes.Buffer


	//size 写入缓冲
	if err := binary.Write(&buffer,binary.LittleEndian,uint16(len(data)));err !=nil{
		return err
	}

	//写入包体数据
	if _,err := buffer.Write(data);err!=nil{
		return err
	}

	//获得写入数据
	out:=buffer.Bytes()

	//写入完整数据
	if _,err := dataWriter.Write(out);err !=nil{
		return err
	}

	return nil
}

//从dataReader读取封包
func readPacket(dataReader io.Reader)(pkt Packet,err error){
	var sizeBuffer = make([]byte,2)

	_,err = io.ReadFull(dataReader,sizeBuffer)

	if err != nil{
		return
	}

	sizeReader := bytes.NewReader(sizeBuffer)

	err = binary.Read(sizeReader,binary.LittleEndian,&pkt.Size)

	if err !=nil{
		return
	}

	//分配包体大小
	pkt.Body = make([]byte,pkt.Size)

	_,err = io.ReadFull(dataReader,pkt.Body)
	return
}