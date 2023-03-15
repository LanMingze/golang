package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b=false
	fmt.Println("b=",b)
	//占用空间为1个字节
	fmt.Println("b占用空间为：",unsafe.Sizeof(b))
}
