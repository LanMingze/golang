package main

import (
	"fmt"
)

func main() {
	//控制台获取用户信息,包括姓名，年龄，薪水，是否通过考试
	var name string
	var age byte
	var sal float32
	var isPass bool
	// fmt.Println("请输入姓名:")
	// //等待用户输入并回车
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水:")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入是否通过:")
	// fmt.Scanln(&isPass)
	// fmt.Printf("名字是：%v\n年龄是：%v\n薪水是：%v\n是否通过:%v\n",
	// name,age,sal,isPass)

	//按指定格式输入
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，使用空格隔开:")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
	fmt.Printf("名字是：%v\n年龄是：%v\n薪水是：%v\n是否通过:%v\n",
	 name,age,sal,isPass)
}
