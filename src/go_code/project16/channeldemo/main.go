package main

import (
	"fmt"

)

func main(){
	var intchan chan int=make(chan int,3)
	fmt.Printf("intchan=%v\n",intchan)
	//写入数据
	intchan<-10
	num:=200
	intchan<-num
	fmt.Printf("intchan len=%v cap=%v\n",len(intchan),cap(intchan))
	//写入数据时不能超过其容量
	//读取数据
	var num2 int=<-intchan
	fmt.Println("num2=",num2)
	//没有协程的情况下，若管道数据全部取出，再去会报错
}