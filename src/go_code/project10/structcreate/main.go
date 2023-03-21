package main

import "fmt"

type Person struct{
	Name string
	Age int
	
} 
func main(){
	p2:=Person{"mary",20}
	//p2.Name="mary"
	//p2.Age=20
	fmt.Println(p2)

	var p3 *Person=new(Person)
	//(*p3).Name="jack"
	//(*p3).Age=50
	//equal to
	p3.Name="smith"
	p3.Age=40
	//为了程序员使用方便，底层会进行处理
	fmt.Println(*p3)

	var p4 *Person=&Person{}//&Person{"tom",12}
	(*p4).Name="john"
	p4.Age=10//原因同上
	fmt.Println(*p4)
}