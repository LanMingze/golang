package main

import (
	"fmt"
)

func main() {
	//不用中间变量交换两个数的值
	var a int = 3
	var b int = 5
	fmt.Println("a=", a, "b=", b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, "b=", b)
}
