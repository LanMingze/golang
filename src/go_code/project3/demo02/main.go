package main

import "fmt"

func main() {
	//一、指定变量类型，不赋值，使用默认值
	//int 默认值为0
	var i int
	fmt.Println("i=", i)
	//二、根据值自行判断变量类型
	var num = 10.11
	fmt.Println("num=", num)
	//三、省略var，=左侧的变量不应该是已经声明过的。否则编译错误
	name := "tom"
	fmt.Println("name=", name)
}
