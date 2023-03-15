package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str=%T\n", str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str=%T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str=%T str=%v\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str=%T str=%v\n", str, str)

	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str=%T str=%v\n", str, str)
	//'f'是一种格式，10指小数位保存10位，64表示float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str=%T str=%v\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str=%T str=%v\n", str, str)

	var num5 int64=4567
	str=strconv.Itoa(int(num5))
	fmt.Printf("str=%T str=%v\n", str, str)
}
