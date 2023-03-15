package main

import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	//直接输出byte，就是输出对应字符的ASCII值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//希望输出对应的字符，需要格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2)

	//var c3 byte='你'//overflow溢出
	var c4 int='你'
	fmt.Printf("c4=%c 码值=%d\n",c4,c4)

	var c5 int =120
	fmt.Printf("c5=%c\n",c5)

	var n1  =10+'a'
	fmt.Println("n1=",n1)
}
