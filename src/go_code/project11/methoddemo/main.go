package main

import "fmt"

type A struct {
	Num int
}

//b绑定一个方法
func (a A) test() {
	fmt.Println("test()", a.Num)
}
type Person struct{
	Name string
}
func(person Person)speak(){
	fmt.Println(person.Name,"is a good man")
}
func (person Person)calculate(){
	res:=0
	for i:=1;i<=1000;i++{
		res+=i
	}
	fmt.Println(person.Name,res)
}
func (person Person)calculate2(n int){
	res:=0
	for i:=1;i<=n;i++{
		res+=i
	}
	fmt.Println(person.Name,res)
}
func(person Person)add(m int,n int)int{
	return m+n
}
func main() {
	var a A
	a.Num = 5
	a.test()
	var p Person
	p.Name="tom"
	p.speak()
	p.calculate()
	p.calculate2(100)
	n1:=10
	n2:=20
	res:=p.add(n1,n2)
	fmt.Println(res)
}
