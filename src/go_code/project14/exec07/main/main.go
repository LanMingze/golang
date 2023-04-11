package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
//定义一个结构体，用于保存统计结果
type CharCount struct{
	ChCount int//英文
	Numount int//数字
	Spaceount int//空格
	Otherount int//其他字符
}

func main(){
	filename:="/home/lan/write.txt"
	file,err:=os.Open(filename)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	var count CharCount
	reader:=bufio.NewReader(file)
	for{
		str,err1:=reader.ReadString('\n')
		if err1==io.EOF{
			break
		}
		for _,v:=range str{
			switch {
			case v>='a'&&v<='z':
				fallthrough
			case v>='A'&&v<='Z':
				count.ChCount++
			case v==' '||v=='\t':
				count.Spaceount++
			case v>='0'&&v<='9':
				count.Numount++
			default:
				count.Otherount++	
			}
		}
	}
	fmt.Println(count.ChCount,count.Numount,count.Spaceount,count.Otherount)
}