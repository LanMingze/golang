package main
import (
	"fmt"
)
type A interface{
	Say()
}
type B interface{
	Hello()
}
type Stu struct{
	Name string
}
func(stu Stu)Say(){
	fmt.Println("stu say")
}
type intrger int
func(i intrger)Say(){
	fmt.Println("integer say i=",i)
}
type Monster struct{

}
func(m Monster)Hello(){
	fmt.Println("monster hello")
}
func(m Monster)Say(){
	fmt.Println("monster say")
}
func main(){
	var stu Stu
	var a A=stu
	a.Say()
	var i intrger=10
	var b A=i
	b.Say()
	//monster实现a和b接口
	var m Monster
	var c A=m
	var d B=m
	c.Say()
	d.Hello()
}