package main

import (
	"fmt"
	"os"
)
func main(){
	file1Path:="/home/lan/write.txt"
	file2Path:="/home/lan/def.txt"
	content,err:=os.ReadFile(file1Path)
	if err!=nil{
		fmt.Println(err)
		return
	}
	err=os.WriteFile(file2Path,content,0666)
	if err!=nil{
		fmt.Println(err)
	}
}