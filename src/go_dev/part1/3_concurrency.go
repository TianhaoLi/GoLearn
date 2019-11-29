package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生产者
func producer(header string, channel chan <- string)  {
	for{
		channel <- fmt.Sprintf("%s: %v",header,rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <- chan string)  {
	for{
		message := <- channel

		fmt.Println(message)
	}
}

func main() {
	chanel := make(chan string)
	//创建groutinue
	go producer("cat",chanel)
	go producer("dog",chanel)
	//消费者
	customer(chanel)
}
