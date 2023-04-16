package main

import (
	"fmt"
)
func writeData(intchan chan int){
	for i:=1;i<=50;i++{
		intchan<-i
		fmt.Println("write:",i)
	}
	close(intchan)
}
func readData(intchan chan int,exitchan chan bool){
	for {
		v,ok:=<-intchan
		if !ok{
			break
		}
		fmt.Printf("read:%v\n",v)
	}
	exitchan<-true
	close(exitchan)
}
func main() {
	//创建两个管道
	intchan:=make(chan int ,50)
	exitchan:=make(chan bool,1)
	go writeData(intchan)
	go readData(intchan,exitchan)
	for{
		_,ok:=<-exitchan
		if !ok{
			break
		}
	}
}
