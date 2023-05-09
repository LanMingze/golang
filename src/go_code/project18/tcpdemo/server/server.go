package main

import (
	"fmt"
	"net"
)
func process(conn net.Conn){
	defer conn.Close()
	
	for{
		//创建一个切片
		buf:=make([]byte,1024)
		//fmt.Printf("服务器等待客户端%s发送消息\n",conn.RemoteAddr().String())
		n,err:=conn.Read(buf)//不发信息会一直阻塞
		if err!=nil{
			fmt.Println(err)
			return
		}
		//显示内容到服务器终端
		fmt.Print(string(buf[:n]))
	}

}
func main(){
	fmt.Println("服务器开始监听")
	//协议使用tcp
	//监听8888端口
	listen,err:=net.Listen("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for{
		//等待客户端连接
		fmt.Println("等待客户端连接。。。")
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println(err)
			
		}else{
			fmt.Printf("accept success con=%v 客户端ip=%v\n",
			conn,conn.RemoteAddr().String())
		}
		//准备启动一个协程，为客户端服务
		go process(conn)
	}
	//fmt.Printf("listen=%v\n",listen) 
}