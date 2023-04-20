package main

import (
	"fmt"
	"time"
)
func sayhello(){
	for i:=0;i<10;i++{
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}
func test(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	//var mymap map[int]string
	//mymap[0]="go"
}
func main(){
	sayhello()
	test()
	for i:=0;i<10;i++{
		fmt.Println("ok",i)
	}
}