package main

import (
	"flag"
	"fmt"
)
func main(){
	//定义几个变量用于接收命令行参数值
	var user string
	var pwd string
	var host string
	var port int
	//user用于接收-u后面参数，“”为默认值
	flag.StringVar(&user,"u","","用户名默认为空")
	flag.StringVar(&pwd,"pwd","","密码默认为空")
	flag.StringVar(&host,"h","","主机名默认为空")
	flag.IntVar(&port,"port",3306,"端口号默认为3306")
	//必须执行
	flag.Parse()
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n",user,pwd,host,port)
}