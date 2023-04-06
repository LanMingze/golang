package main

import (
	"fmt"
)

func main() {
	key := ""    //保存输入
	loop := true //是否退出循环
	//收支详情，使用字符串记录
	detail := "收支\t账户余额\t收支金额\t说明"
	//定义变量是否有收支
	flag:=false
	//定义账户余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支说明
	note := ""
	for {
		fmt.Println("\n--------家庭收支记账软件--------")
		fmt.Println("        1.收支明细        ")
		fmt.Println("        2.登记收入        ")
		fmt.Println("        3.登记支出        ")
		fmt.Println("        4.退出软件        ")
		fmt.Print("请选择1-4：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------当前收支明细--------")
			if flag{
				fmt.Println(detail)
			}else{
				fmt.Println("当前没有收支，请来一笔吧！")
			}
			
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money //修改余额
			flag=true
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//将收入情况记录到detail中
			detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
		case "3":
			fmt.Println("--------登记支出--------")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money>balance{
				fmt.Println("余额不足")
				break
			}
			balance -= money //修改余额
			flag=true
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			detail += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
		case "4":
			fmt.Println("你确定要退出吗？ y/n")
			choice:=""
			for{
				fmt.Scanln(&choice)
				if choice=="y"||choice=="n"{
					break
				}
				fmt.Println("输入有误，重新输入")
			}
			if choice=="y"{
				loop = false
			}
			
		default:
			fmt.Println("输入错误...")
		}
		if !loop {
			break
		}
	}
	fmt.Println("已退出！")
}
