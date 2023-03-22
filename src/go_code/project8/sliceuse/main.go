package main

import "fmt"

func main() {
	var slice []float64=make([]float64, 5,10)//切片必须make后才能使用
	slice[1]=20
	slice[3]=40
	fmt.Println(slice)
	fmt.Println("slice size=",len(slice))
	fmt.Println("slice cap=",cap(slice))

	var slice1 []string=[]string{"tom","jack","mary"}
	fmt.Println(slice1)
	fmt.Println("slice1 size=",len(slice1))
	fmt.Println("slice1 cap=",cap(slice1))
}
