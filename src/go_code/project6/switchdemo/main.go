package main

import (
	"fmt"
)

func main() {
/*	var a byte
	fmt.Println("请输入一个字符：")
	fmt.Scanf("%c", &a)
	switch a {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("输入有误")
	}

	var n1 int32 = 5
	var n2 int32 = 20
	var n3 int32 = 5
	switch n1 {
	case n2, 10, 5:
		fmt.Println("ok1")
	case n3:
		fmt.Println("ok2")
	default:
		fmt.Println("没有匹配到")

	}
*/
	var age int =10
	switch{
	case age==10:
		fmt.Println("age==10")
	case age==20:
		fmt.Println("age==20")
	default:
		fmt.Println("没有匹配到")
	}

	var grade int =80
	switch {
	case grade>90:
		fmt.Println("优秀")
	case grade>70&&grade<=90:
		fmt.Println("良好")
	default:
		fmt.Println("没有匹配到")
	}
	
	switch score:=95;{
	case score>90:
		fmt.Println("优秀")
	case score>70&&score<=90:
		fmt.Println("良好")
	default:
		fmt.Println("没有匹配到")
	}

	var num int32=10
	switch num{
		case 10:
			fmt.Println("ok1")
			fallthrough
		case 20:
			fmt.Println("ok2")
	}

}
