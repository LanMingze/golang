package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)
func main(){
	//通过go向redis写入和读取数据
	conn,err:=redis.Dial("tcp","127.0.0.1:6379")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("conn success",conn)
}