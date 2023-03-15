package main

import "fmt"

func main() {
	// for i:=1; i<=3;i++{
	// 	fmt.Println("hello",i)
	// }
	// j:=1
	// for j<=10{
	// 	fmt.Println("nihao",j)
	// 	j++
	// }
	// var str string = "hello,world"
	// for i:=0;i<len(str);i++{
	// 	fmt.Printf("%c\n",str[i])
	// }
	var str string = "hello,world北京"
	str1:=[]rune(str)
	for i:=0;i<len(str1);i++{
		fmt.Printf("%c\n",str1[i])
	}
	for index,val:=range str{
		fmt.Printf("index=%d,val=%c\n",index,val)
	}
}
