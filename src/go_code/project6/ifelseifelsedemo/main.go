package main

import (
	"fmt"
)

func main() {
	var grade int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&grade)
	if grade == 100 {
		fmt.Println("奖励一台BMW")
	} else if grade > 80 && grade <= 99 {
		fmt.Println("奖励一台iphone")
	} else if grade >= 60 && grade <= 80 {
		fmt.Println("奖励一个ipad")
	} else {
		fmt.Println("什么也没有")
	}
	//使用陷阱，只会输出ok1
	var n int = 10
	if n > 9 {
		fmt.Println("ok1")
	} else if n > 6 {
		fmt.Println("ok2")
	} else if n > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok420")
	}
}
