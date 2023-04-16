package main

import (
	"fmt"
)
func main() {
	intchan:=make(chan int,2)
	intchan<-2
	close(intchan)
	n1:=<-intchan
	fmt.Println(n1)

	//遍历管道
	intchan2:=make(chan int,100)
	for i:=0;i<100;i++{
		intchan2<-i*2//管道中放入100个数据
	}
	close(intchan2)//不关闭管道会出现死锁
	for v:=range intchan2{
		fmt.Println(v)
	}

}
