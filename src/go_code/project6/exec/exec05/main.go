package main

import (
	"fmt"
)

func main() {
	// var x byte
	// fmt.Println("请输入一个字符：")
	// fmt.Scanf("%c", &x)
	// switch x {
	// 	case 'a':
	// 		fmt.Println("A")
	// 	case 'b':
	// 		fmt.Println("B")
	// 	case 'c':
	// 		fmt.Println("C")
	// 	case 'd':
	// 		fmt.Println("D")
	// 	default:
	// 		fmt.Println("other")
	// }

	// var grade float64
	// fmt.Println("请输入成绩：")
	// fmt.Scanln(&grade)
	// switch int(grade/60){
	// 	case 1:
	// 		fmt.Println("及格了")
	// 	case 0:
	// 		fmt.Println("不及格")
	// 	default:
	// 		fmt.Println("输入有误")
	// }

	var month byte
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	switch month{
		case 3,4,5:
			fmt.Println("spring")
		case 6,7,8:
			fmt.Println("summer")
		case 9,10,11:
			fmt.Println("fall")
		case 12,1,2:
			fmt.Println("winter")
		default:
			fmt.Println("输入有误")
	}
}
