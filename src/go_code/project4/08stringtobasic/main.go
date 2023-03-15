package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	//返回值有两个，第二个为err，不想使用可用 _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T b=%v\n", b, b)
	var str2 string = "123456789"
	var n2 int64
	n2, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n2 type is %T n2=%v\n", n2, n2)
	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type is %T f1=%v\n", f1, f1)

	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type is %T n3=%v\n", n3, n3)
}
