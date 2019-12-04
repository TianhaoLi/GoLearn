package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//声明一个等待组
	var wg sync.WaitGroup

	var urls = []string{
		"https://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	for _,url := range urls{
		wg.Add(1)

		go func(url string) {
			defer  wg.Done()

			_,err := http.Get(url)

			fmt.Println(url,err)
		}(url)
	}

	wg.Wait()

	fmt.Println("over")
}
