package main

import "fmt"

type Person struct {
	Name string
}
func test01(p Person){
	fmt.Println(p.Name)
}

func main() {
	p:=Person{"tom"}
	test01(p)
}
