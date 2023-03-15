package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println("i=", i)

	var j int8 = -128
	fmt.Println("j=", j)
	var k uint8 = 255
	fmt.Println("k=", k)

	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 89
	fmt.Println("b=", b)
	var c byte = 255
	fmt.Println("c=", c)
	//使用细节
	var n1 = 100
	fmt.Printf("n1的类型是：%T\n", n1)
	//查看变量占用多少字节和数据类型
	var n2 int64 = 10
	fmt.Printf("n2的类型是：%T,n2占用的字节数是：%d\n", n2, unsafe.Sizeof(n2))
	//尽量使用占用空间小的数据类型
	var age byte=22
	fmt.Println("age=",age)
	
}
