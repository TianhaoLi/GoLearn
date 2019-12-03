package main

import (
	"fmt"
)

type Dictionary struct {
	data map[interface{}] interface{}  //键值 都为 interface
}

func (d *Dictionary) Get(key interface{}) interface{}{
	return d.data[key]
}

func (d *Dictionary) Set(key interface{},value interface{}) {
	d.data[key] = value
}

//遍历

func (d *Dictionary) Visit(callback func(k,v interface{}) bool) {
	if callback == nil{
		return
	}

	for k,v := range d.data{
		if !callback(k,v){
			return
		}
	}
}

//清空所有数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}

	d.Clear()
	return d
}

func main() {
	dic := NewDictionary()

	dic.Set("My Factory",60)
	dic.Set("Terra Craft",36)
	dic.Set("Don't Hungry",24)

	favorite := dic.Get("Terra Craft")
	fmt.Println("favorite:",favorite)

	dic.Visit(func(k, v interface{}) bool {


		if v.(int) > 40 {
			fmt.Println(k,"is expensive")
			return true
		}

		fmt.Println(k,"is cheap")

		return true
	})
}