package main

import (
	"fmt"
	"go_code/project4/10identifierdemo/model"
	//为了使用utils.go变量或函数，先引用该包
)

//标识符
func main() {
	var num int =10
	var Num int =20
	fmt.Printf("num=%v,Num=%v\n",num,Num)
	fmt.Println("hero的名字是：",model.HeroName)
}
