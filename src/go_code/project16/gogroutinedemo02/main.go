package main

import (
	"fmt"
	"sync"
	"time"
)

//计算各个数的阶乘，并放入到一个map中
//启动多个协程，将结果放入到map中
//map应该为全局变量

var (
	mymap =make(map[int]int,10)
	lock sync.Mutex//全局互斥锁
)
func test(n int){
	res:=1
	for i:=1;i<=n;i++{
		res*=i
	}
	lock.Lock()
	mymap[n]=res
	lock.Unlock()
}
func main(){
	for i:=1;i<=20;i++{
		go test(i)
	}
	time.Sleep(5*time.Second)
	lock.Lock()
	for i,v:=range mymap{
		fmt.Printf("map[%d]=%d",i,v)
	}
	lock.Unlock()
}