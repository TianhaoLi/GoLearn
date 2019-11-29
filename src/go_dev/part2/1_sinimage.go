package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	const size = 300

	//新建的一个方框 x1,y1   x2,y2
	pic := image.NewGray(image.Rect(0,0,size,size))

	//遍历每个像素点
	for x := 0; x < size; x++ {
		for y:=0;y<size;y++{
			pic.SetGray(x,y,color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x)*2*math.Pi/size
		//sin 值范围在 0~2pi  计算定义域
		y:=size/2 - math.Sin(s)*size/2
		//sin幅度为一半向下偏移一半并翻转
		pic.SetGray(x,int(y),color.Gray{0})
	}


	//写入文件
	file,err := os.Create("sin.png")

	if err!=nil{
		log.Fatal(err)
	}
	//png格式写入文件
	png.Encode(file,pic)

	file.Close()
}
