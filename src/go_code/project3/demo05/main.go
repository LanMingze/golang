package main

import "fmt"

func main() {
	//+的使用
	var j = 1
	var i = 2
	var k = i + j
	fmt.Println("k=", k)

	var str1 = "hello "
	var str2 = "world"
	var str3 = str1 + str2
	fmt.Println(str3)
}
