package main

import "fmt"

func peach(n int) int {
	if n == 10 {
		return 1
	} else if n >= 1 && n < 10 {
		return 2 * (peach(n+1) + 1)
	} else {
		fmt.Println("输入天数错误")
		return 0
	}
}

func main() {
	fmt.Println("第一天桃子数：", peach(1))
}
