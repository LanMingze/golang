package main

import (
	"fmt"
)

var (
	Fun1=func(n1 int,n2 int)int{
		return n1*n2
	}
)
func main() {
	res1:=func(n1 int,n2 int)int{
		return n1+n2
	}(10,20)
	fmt.Println("res1=",res1)

	a:=func(n1 int,n2 int)int{
		return n1-n2
	}
	res2:=a(20,10)
	fmt.Println("res2=",res2)
	res3:=a(20,90)
	fmt.Println("res3=",res3)

	res4:=Fun1(4,9)
	fmt.Println("res4=",res4)
}
