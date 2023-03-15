package main

import "fmt"

func main() {
	var i int = 10
	i = 30
	//i = 50.05 //不能改变数据类型
	fmt.Println("i=", i)
	//变量在同一作用域不能重命名
	//var i int = 59
	//i:=99
}
