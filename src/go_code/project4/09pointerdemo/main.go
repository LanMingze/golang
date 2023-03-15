package main

import (
	"fmt"
)
//指针类型
func main() {
	//基本数据类型在内存布局
	var i int =10
	//i的地址,&i
	fmt.Println("i的地址是=",&i)
	//ptr是一个指针变量，类型是*int，ptr本身的值是&i
	var ptr *int=&i
	fmt.Printf("ptr的值是=%v\n",ptr)
	fmt.Println("ptr的地址是=",&ptr)
	fmt.Printf("ptr指向的值是=%v\n",*ptr)

	var num int =10
	var ptr1 *int=&num
	fmt.Printf("num的地址是=%v\n",ptr1)
	*ptr1=100
	fmt.Printf("num修改后的值是=%v\n",num)
}
