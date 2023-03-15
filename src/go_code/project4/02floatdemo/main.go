package main

import (
	"fmt"
)

// 小数类型使用
func main() {
	var price float32 = 26.89
	fmt.Println("price=", price)
	var num1 float32 = -0.009
	var num2 float64 = -7852255.3
	fmt.Println("num1=", num1, "num2=", num2)

	//尾数部分可能丢失，造成精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)
	//默认为float64
	var num5=1.1
	fmt.Printf("num5数据类型是%T\n",num5)
	num6:=5.12
	num7:=.123
	fmt.Println("num6=",num6,"num7=",num7)
	num8:=5.1234e2
	num9:=5.1234e-2
	fmt.Println("num8=",num8,"num9=",num9)
}
