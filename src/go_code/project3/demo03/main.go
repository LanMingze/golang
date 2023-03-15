package main

import "fmt"

//定义全局变量
var n1 = 100
var name = 200
var n3 = "jack"

//一次性声明全局变量
var (
	n4    = 300
	n5    = 900
	name2 = "mary"
)

func main() {
	//一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	//一次性声明多个变量的方式2
	// var n1, name, n3 = 100, "tom", 888
	// fmt.Println("n1=", n1, "n2=", name, "n3=", n3)
	//一次性声明多个变量的方式3
	//n1, name, n3 := 100, "tom", 888
	fmt.Println("n1=", n1, "n2=", name, "n3=", n3)
	fmt.Println("n4=", n4, "n5=", n5, "name2=", name2)
}
