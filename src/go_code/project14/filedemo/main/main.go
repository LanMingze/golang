package main
import (
	"fmt"
	"os"
)
func main(){
	// file	文件对象/文件句柄/文件指针
	file,err:=os.Open("/home/lan/write.txt")
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	fmt.Printf("file=%v\n",file)
	err=file.Close()
	if err!=nil{
		fmt.Println("error=",err)
	}
}