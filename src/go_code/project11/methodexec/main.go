package main

import "fmt"

type Circle struct{
	radius float64
	
}
func (c Circle)area()float64{
	return 3.14*c.radius*c.radius
} 
func main(){
	var circle Circle
	circle.radius=4.0
	res:=circle.area()
	fmt.Println(res)
}