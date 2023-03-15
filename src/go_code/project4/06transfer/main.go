package main

import (
	"fmt"//_加在前面表示忽略该包
	
)

func main() {
	var i int =100
	//希望将i转为float
	var j float32 =float32(i)
	var k int64=int64(i)
	fmt.Printf("i=%v,j=%v,k=%v\n",i,j,k)
	fmt.Printf("i=%T\n",i)

	var num int64=999999
	var num2 int8=int8(num)
	fmt.Println("num2=",num2)
}
