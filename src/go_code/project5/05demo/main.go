package main

import (
	_ "debug/plan9obj"
	"fmt"
)

func test() bool {
	fmt.Println("test....")
	return true
}
func main() {
	var i int = 10
	//短路与 i<9为false，test()不执行
	if i < 9 && test() {
		fmt.Println("ok...")
	}
	//短路与 i>9为true，test()不执行
	if i > 9 || test() {
		fmt.Println("hello...")
	}
	//逻辑运算符的演示 &&
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}
	//逻辑运算符的演示 ||
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}
	//逻辑运算符的演示 !
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}
}
