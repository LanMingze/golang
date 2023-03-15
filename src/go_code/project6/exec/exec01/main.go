package main

import (
	"fmt"
)

func main() {
	var a int32 = 10
	var b int32 = 50
	if a+b > 50 {
		fmt.Println("hello world!")
	}
	var c float64 = 17.0
	var d float64 = 14.0
	if c > 10.0 && d < 20.0 {
		fmt.Println("c+d=", c+d)
	}

	var n1 int32 = 10
	var n2 int32 = 5
	if (n1+n2)%3 == 0 && (n1+n2)%5 == 0 {
		fmt.Println("能被3又能被5整除")
	}
	var year int32 = 2022
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year,"年是闰年")
	}
}
