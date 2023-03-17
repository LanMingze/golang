package main

import "fmt"

func main() {
	str:="hello world!!!"
	str1:=str[6:]
	fmt.Println(str1)

	arr1:=[]byte(str)
	arr1[0]='z'
	str=string(arr1)
	fmt.Println(str)

	arr2:=[]rune(str)
	arr2[0]='会'
	str=string(arr2)
	fmt.Println(str)
}
