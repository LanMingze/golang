package main

import (
	"fmt"
	_"math"
)

func main() {
	var height int32
	var money float32
	var handsome bool
	fmt.Println("请输入身高：")
	fmt.Scanln(&height)
	fmt.Println("请输入财富(千万)：")
	fmt.Scanln(&money)
	fmt.Println("请输入是否帅气：")
	fmt.Scanln(&handsome)
	if height>180&&money>1.0&&handsome==true{
		fmt.Println("我一定嫁给他")
	}else if height<180&&money<1.0&&handsome==false{
		fmt.Println("不嫁。。。")
	}else{
		fmt.Println("比上不足比下有余")
	}
}
