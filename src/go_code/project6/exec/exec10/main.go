package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("奇数是：", i)
	}
	var positive int=0
	var negative int=0
	var num int
	for {
		fmt.Println("请输入一个整数：")
		fmt.Scanln(&num)
		if num==0{
			break
		}else if num>0{
			positive++
		}else{
			negative++
		}
	}
	fmt.Printf("正数个数是%v，负数个数是%v\n",positive,negative)
}
