package main

import "fmt"

type Circle struct{
	radius float64
	
}
func (c Circle)area()float64{
	return 3.14*c.radius*c.radius
} 
func (c *Circle)area1()float64{
	return 3.14*(*c).radius*(*c).radius
} 
func main(){
	var circle Circle
	circle.radius=4.0
	res:=circle.area()
	fmt.Println(res)
	circle.radius=5.0
	//res=(&circle).area1() equal to 编译器自动添加
	res=circle.area1()
	fmt.Println(res)
}