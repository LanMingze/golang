package main

import "fmt"

type integer int
func (i integer)print(){
	fmt.Println(i)
}
func (i *integer)change(){
	*i=*i+1
}
type Student struct{
	Name string
	Age int
}
func(stu *Student)String()string{
	str:=fmt.Sprintf("Name=[%v],Age=[%v]",stu.Name,stu.Age)
	return str
}
func main() {
	var i integer=10
	i.print()
	i.change()
	i.print()
	stu:=Student{
		Name: "tom",
		Age: 15,
	}
	//str:=(&stu).String()
	fmt.Println(&stu)
}
