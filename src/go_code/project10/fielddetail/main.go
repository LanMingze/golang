package main

import "fmt"

type Person struct{
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
} 
type monster struct{
	Name string
	Age int
}
func main(){
	var mons monster
	mons.Name="niumo"
	mons.Age=500
	mons2:=mons
	mons2.Name="qingniu"
	fmt.Println(mons,mons2)
	var p1 Person
	fmt.Println(p1)
	if p1.ptr==nil{
		fmt.Println("ok1")
	}
	if p1.slice==nil{
		fmt.Println("ok2")
	}
	if p1.map1==nil{
		fmt.Println("ok3")
	}
	p1.slice=make([]int, 10)
	p1.slice[0]=10
	fmt.Println(p1)

	p1.map1=make(map[string]string, 2)
	p1.map1["name"]="tom"
	fmt.Println(p1)
}