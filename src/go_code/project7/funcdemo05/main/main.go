package main

import (
	"fmt"
)

type myFun1 func(int, int) int

func sum(n1 int, n2 int) int {
	return n1 + n2
}
func myFun(funvar myFun1, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func sumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func sum1(n1 int,args...int)int{
	sum:=n1
	for i:=0;i<len(args);i++{
		sum+=args[i]//args[0]表示args切片的第一个值
	}
	return sum
}
func main() {
	n1 := sum
	fmt.Printf("a的类型是%T，sum的类型是%T\n", n1, sum)
	res := n1(10, 40)
	fmt.Println("res=", res)
	res2 := myFun(sum, 50, 60)
	fmt.Println("res2=", res2)

	type myInt int //myInt和int虽然都是int类型，但go认为他们是两个不同类型
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)
	fmt.Println("num1=", num1, "num2=", num2)

	res3 := myFun(sum, 50, 600)
	fmt.Println("res2=", res3)

	res4, res5 := sumAndSub(60, 40)
	fmt.Println("res4=", res4, "res5=", res5)

	res6:=sum1(10,0,-1,90,10,100)
	fmt.Println("res6=",res6)
}
