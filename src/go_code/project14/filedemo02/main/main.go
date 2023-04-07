package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main(){
	// file	文件对象/文件句柄/文件指针
	file,err:=os.Open("/home/lan/write.txt")
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	defer file.Close()//不关闭会有内存泄漏

	//采用带缓冲区方式
	reader:=bufio.NewReader(file)
	for{
		str,err:=reader.ReadString('\n')//读到一个换行就结束
		if err==io.EOF{
			break
		}
		fmt.Print(str)
	}
}