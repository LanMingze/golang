package main

import (
	"fmt"
	
)
func test(slice[]int){
	slice[0]=100//会改变实参
}
func main(){
	var slice[]int
	var arr[5]int =[...]int{1,2,3,4,5}
	slice=arr[:]
	var slice2=slice
	slice2[0]=10
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(arr)

	test(slice)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(arr)
}