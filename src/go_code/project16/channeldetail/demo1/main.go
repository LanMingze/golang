package main

import (
	"fmt"

)

func main(){
	//var chan1 chan int//可读可写
	var chan2 chan<-int=make(chan int,3)//只写
	chan2<-2
	//num:=<-chan2 报错
	fmt.Println(chan2)
	var chan3 <-chan int//只读
	num:=<-chan3
	//chan3<-2
	fmt.Println(num)
}