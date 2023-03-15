package main

import (
	"fmt"
)

func main() {
	// var sum int =0
	// for i:=1;i<=100;i++{
	// 	sum+=i
	// 	if sum>20{
	// 		fmt.Println("当前数是：",i)
	// 		break
	// 	}
	// }

	var name string
	var password string
	for i:=1;i<=3;i++{
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		if name=="aaa"&&password=="888"{
			fmt.Println("登录成功")
			break
		}else{
			fmt.Printf("你还有%v次机会\n",3-i)
		}
	}
}
