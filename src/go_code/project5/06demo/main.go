package main

import (
	
	"fmt"
)


func main() {
	// var i int 
	// i=10
	//a,b交换变量的值
	a:=9
	b:=2
	fmt.Printf("交换前的情况是：a=%v,b=%v\n",a,b)
	t:=a
	a=b
	b=t
	fmt.Printf("交换后的情况是：a=%v,b=%v\n",a,b)
	a+=17
	fmt.Println("a=",a)
}
