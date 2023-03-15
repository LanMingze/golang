package main

import (
	"fmt"
)

func main() {
	var grade float64
	var sex string
	fmt.Println("请输入成绩：")
	fmt.Scanln(&grade)

	if grade <= 8 {
		fmt.Println("请输入性别：")
		fmt.Scanln(&sex)
		if sex == "男" {
			fmt.Println("进入男子组决赛")
		} else {
			fmt.Println("进入女子组决赛")
		}

	} else {
		fmt.Println("out...")
	}

	var month byte
	var age byte
	var price float64=60
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	if month>=4&&month<=10{
		if age>60{
			fmt.Printf("%v月 %v年龄 票价是：%v",month,age,price/3)
		}else if age>=18{
			fmt.Printf("%v月 %v年龄 票价是：%v",month,age,price)
		}else{
			fmt.Printf("%v月 %v年龄 票价是：%v",month,age,price/2)
		}
	}else{
		if age>=18&&age<=60{
			fmt.Println("淡季成人 票价是：",40)
		}else{
			fmt.Println("淡季老人和小孩 票价是：",20)
		}
	}
}
