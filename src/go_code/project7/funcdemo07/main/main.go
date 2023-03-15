package main

import (
	"fmt"
	"strings"
)

func addUpper() func(int) int {
	var n int = 10
	str := "hello"
	return func(i int) int {
		n = n + i
		str += "a"
		fmt.Println("str=", str)
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string)string{
		if strings.HasSuffix(name,suffix)==false{
			return name+suffix
		}
		return name
	}
}
func main() {
	f := addUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	res:=makeSuffix(".jpg")
	fmt.Println(res("winter"))
	fmt.Println(res("spring.avi"))
}
