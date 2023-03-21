package main

import "fmt"

type Cat struct{
	Name string
	Age int
	Color string
	Hobby string
} 
func main(){
	// var cat1Name string="xiaobai"
	// var cat1Age int=2
	// var cat1Color string="baise"

	var cat1 Cat
	cat1.Name="xiaobai"
	cat1.Age=3
	cat1.Color="white"
	cat1.Hobby="fish"
	fmt.Println(cat1)

	fmt.Println("name=",cat1.Name)
	fmt.Println("age=",cat1.Age)
	fmt.Println("color=",cat1.Color)
	fmt.Println("hobby=",cat1.Hobby)
}