package main

import (
	"fmt"
)

func main() {
	var i int = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello", i)
		i++
	}
	fmt.Println("i=", i)
	
	i=1
	for {
		fmt.Println("hello", i)
		i++
		if i > 10 {
			break
		} 
	}
	fmt.Println("i=", i)
}
