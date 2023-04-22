package main

import (
	"fmt"
)

func main() {
	var num int=0
	num=90
	//常量必须赋值
	const tax int = 0
	//常量不能修改
	//tax=10
	fmt.Println(num, tax)
	//常量只能修饰bool，数值类型（int float），string类型
	//const tax1=num/3 错误
	const(
		a=iota//a=0
		b//b=1
		c//c=2
	)
	const(
		d=iota//d=0
		e//e=1
		f,g=iota,iota//f=2,g=2
	)
	//可以通过大小写控制访问范围
	fmt.Println(a,b,c)
	fmt.Println(d,e,f,g)
}
