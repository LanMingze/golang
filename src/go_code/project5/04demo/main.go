package main

import (
	_ "debug/plan9obj"
	"fmt"
)

func main() {
	//关系运算符的演示
	var n1 int =9
	var n2 int =8
	fmt.Println(n1==n2)//false
	fmt.Println(n1!=n2)//true
	fmt.Println(n1>=n2)//true
	fmt.Println(n1<=n2)//false
	fmt.Println(n1>n2)//true
	fmt.Println(n1<n2)//false
	flag:=n1>n2
	fmt.Printf("flag=%v\n",flag)
}
