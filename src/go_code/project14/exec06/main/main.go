package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//编写一个函数，接收两个文件目录
func copyFile(dst string,src string)(written int64,err error){
	srcfile,err:=os.Open(src)
	if err!=nil{
		fmt.Println(err)
	}
	defer srcfile.Close()
	reader:=bufio.NewReader(srcfile)

	dstfile,err1:=os.OpenFile(dst,os.O_WRONLY|os.O_CREATE,0666)
	if err1!=nil{
		fmt.Println(err)
	}
	defer dstfile.Close()
	writer:=bufio.NewWriter(dstfile)

	return io.Copy(writer,reader)
}
func main(){
	//将一个图片拷贝到另一个文件夹下
	srcfile:="/home/lan/图片/1.jpg"
	dstfile:="/home/lan/abc.jpg"
	_,err:=copyFile(dstfile,srcfile)
	if err==nil{
		fmt.Println("finish")
	}else{
		fmt.Println(err)
	}
}