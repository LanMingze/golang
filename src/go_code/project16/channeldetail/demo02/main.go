package main

import (
	"fmt"

)

func main(){
	inchan:=make(chan int,10)
	for i:=0;i<10;i++{
		inchan<-i
	}
	stringchan:=make(chan string,5)
	for i:=0;i<5;i++{
		stringchan<-"hello"+fmt.Sprintf("%d",i)
	}
	//传统方法不关闭管道会阻塞，形成死锁
	//在某些情况下，无法确定什么时候关闭管道
	//可以使用select
	for {
		select{
		case v:=<-inchan:
			fmt.Printf("数据为%d\n",v)
		case v:=<-stringchan:
			fmt.Printf("数据为%s\n",v)
		default:
			fmt.Println("没有数据")
			return
		}
	}
	// for i:=range inchan {
	// 	fmt.Println(i)
	// }
}