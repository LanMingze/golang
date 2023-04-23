package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	line,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println(err)
	}
	n,err:=conn.Write([]byte(line))
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("客户端发送了%d字节数据，并退出\n",n)
}