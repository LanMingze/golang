package main

import (
	"fmt"
)

func main() {
	//还有97天放假，问还有多少星期零多少天
	var day int =97
	var week int =day/7
	var day1 int =day%7
	fmt.Printf("还有%d星期%d天\n",week,day1)
	//定义一个华氏温度变量，5/9*（华氏温度-100）=摄氏温度，求摄氏温度
	var huashi float32=134.2
	var sheshi float32=5.0/9*(huashi-100)
	fmt.Printf("%v对应的摄氏温度是：%v",huashi,sheshi)
}
