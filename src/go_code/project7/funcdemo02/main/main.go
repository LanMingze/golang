package main

import (
	"fmt"
)

func test(n1 int){
	n1+=1
	fmt.Println("test n1=",n1)
}

func cal(n1 int,n2 int)int{
	return n1+n2
}
func sumAndSub(n1 int,n2 int)(int,int){
	sum:=n1+n2
	sub:=n1-n2
	return sum,sub
}
func main() {
	n1:=10
	test(n1)
	fmt.Println("n1=",n1)

	n2:=15
	n3:=cal(n1,n2)
	fmt.Println("n3=",n3)

	res1,res2:=sumAndSub(n1,n2)
	fmt.Println("sum=",res1,"sub=",res2)

	res3,_:=sumAndSub(n1,n2)
	fmt.Println("sum=",res3)

}
