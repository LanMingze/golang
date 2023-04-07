package main

import (
	"fmt"
	"os"
)
func main(){
	// file	文件对象/文件句柄/文件指针
	file:="/home/lan/write.txt"
	content,err:=os.ReadFile(file)
	if err!=nil{
		fmt.Println("error=",err)
	}
	fmt.Printf("%v\n",string(content))
}