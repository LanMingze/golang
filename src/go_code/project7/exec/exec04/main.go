package main

import (
	"fmt"
	"time"
	"strconv"
)

func test(){
	str:=""
	for i:=0;i<100000;i++{
		str+="hello"+strconv.Itoa(i)
	}

}

func main() {
	start:=time.Now().Unix()
	test()
	end:=time.Now().Unix()
	fmt.Printf("耗费时间是：%v\n",end-start)
}
