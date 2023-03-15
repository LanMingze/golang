package main

import (
	"fmt"
)

func main() {
	//求两个数的最大值
	var max int
	var n2 int = 10
	var n1 int = 15
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("max=", max)

	//求三个数的最大值
	var n3 int = 20
	if n3 > max {
		max = n3
	}
	fmt.Println("max=", max)
}
