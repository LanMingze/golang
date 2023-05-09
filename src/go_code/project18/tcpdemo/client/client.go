package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
func main(){
	conn,err:=net.Dial("tcp","127.0.0.1:8888")
	if err!=nil{
		fmt.Println(err)
		return
	}
	//fmt.Println("conn=",conn)
	//发送单行数据，然后退出
	reader:=bufio.NewReader(os.Stdin)//标准输入即终端
	for{
		line,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println(err)
		}
		line=strings.Trim(line," \n\r")
		if line=="exit"{
			fmt.Println("客户端退出")
			break
		}
		_,err=conn.Write([]byte(line+"\n"))
		if err!=nil{
			fmt.Println(err)
		}
	}
}