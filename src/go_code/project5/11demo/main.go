package main

import (
	"fmt"
)

func main() {
	var i int =5
	fmt.Printf("i=%b\n",i)//二进制输出
	var j int = 011//八进制
	fmt.Println("j=",j)
	var k int =0x11//十六进制
	fmt.Println("k=",k)
}
