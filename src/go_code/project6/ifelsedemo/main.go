package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入您的年龄：")
	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你年龄大于18岁，要对自己的行为负责")
	} else {
		fmt.Println("你年龄不大，这次放过你了")
	}
}
