package main

import (
	"fmt"
)

func main() {
	var count int = 0
	var sum int=0
	for i:=9;i<=100;i++{
		if i%9==0{
			count++
			sum+=i
		}
	}
	fmt.Println("count=",count,",sum=",sum)

	var n int =6
	for i:=0;i<=n;i++{
		fmt.Printf("%v+%v=%v\n",i,n-i,n)
	}
}
