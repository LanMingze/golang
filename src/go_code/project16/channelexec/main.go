package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allchan chan interface{}= make(chan interface{}, 3)
	allchan<-10
	allchan<-"tom jack"
	cat:=Cat{
		Name: "huamao",
		Age: 4,
	}
	allchan<-cat
	//若需要第三个，需要先推出前两个
	<-allchan
	<-allchan
	newCat:=<-allchan
	fmt.Printf("%T %v\n",newCat,newCat)
	a:=newCat.(Cat)//需要类型断言转换
	fmt.Printf("%v\n",a.Name)
}
