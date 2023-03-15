package main

import (
	"fmt"
	"go_code/project7/funcInitdemo/utils"
)
var age =test()
func test()int{
	fmt.Println("test...")
	return 90
}
//init通常完成初始化工作
func init(){
	fmt.Println("init...")
}
func main() {
	fmt.Println("main....")
	fmt.Println(age)
	fmt.Println("Age=",utils.Age,"Name=",utils.Name)
}
