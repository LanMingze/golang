package main

import (
	"fmt"
)
func putnum(intchan chan int){
	for i:=1;i<=8000;i++{
		intchan<-i
	}
	close(intchan)
}
func primenum(intchan chan int,primechan chan int,exitchan chan bool){
	//var num int
	var flag bool
	for {
		num,ok:=<-intchan
		if !ok{
			break
		}
		flag=true
		for i:=2;i<num;i++{
			if num%i==0{
				flag=false
				break
			}
		}
		if flag{
			primechan<-num
		}
	}
	fmt.Println("有一个协程取不到数据，退出！！！")
	exitchan<-true
}
func main() {
	//创建两个管道
	intchan:=make(chan int ,1000)
	primechan:=make(chan int ,2000)
	exitchan:=make(chan bool,4)

	go putnum(intchan)

	//从intchan取数据，判断是否为素数
	for i:=0;i<4;i++{
		go primenum(intchan,primechan,exitchan)
	}
	go func(){
		for i:=0;i<4;i++{
			<-exitchan
		}
		close(primechan)
	}()
	for{
		res,ok:=<-primechan
		if !ok{
			break
		}
		fmt.Println(res)
	}
}
