package main

import (
	"fmt"
)

func main() {
	//在golang中++和--只能独立使用
	var i int = 8
	var a int
	//a=i++  错误
	i++
	a = i
	fmt.Println("a=", a)

	var b int =1
	//++b 没有前置++和--
	fmt.Println("b=",b)
	
}
