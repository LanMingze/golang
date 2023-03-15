package main

import "fmt"




func main() {
	num1:=100
	fmt.Printf("num1类型是：%T，值是：%v，地址是：%v\n",num1,num1,&num1)

	num2:=new(int)
	//num2是指针
	//值是地址 x系统分配
	fmt.Printf("num2类型是：%T，值是：%v，地址是：%v，指向的值是：%v\n",num2,
	num2,&num2,*num2)


}
