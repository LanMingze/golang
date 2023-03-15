package main

import (
	"errors"
	"fmt"
)

func test(){
	defer func ()  {
		err:=recover()
		if err!=nil{
			fmt.Println("err=",err)
			//发送错误信息

		}
	}()
	num1:=10
	num2:=0
	res:=num1/num2
	fmt.Println("res=",res)
}

func readconf(name string)(err error){
	if name=="config.ini"{
		return nil
	}else{
		return errors.New("读取文件错误..")
	}
}

func test02(){
	err:=readconf("config.in")
	if err!=nil{
		panic(err)
	}
	fmt.Println("test02()继续执行..")
}
func main() {
	//test()
	fmt.Println("main...")

	test02()
	
}
