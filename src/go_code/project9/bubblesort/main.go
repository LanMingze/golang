package main

import "fmt"

func bubble(arr *[5]int){
	fmt.Println(*arr)
	temp:=0
	target:=0
	for i:=0;i<len(*arr)-1;i++{
		target=0
		for j:=0;j<len(*arr)-i-1;j++{
			if (*arr)[j]>(*arr)[j+1]{
				temp=(*arr)[j]
				(*arr)[j]=(*arr)[j+1]
				(*arr)[j+1]=temp
				target=1
			}
		}
		if target==0{
			break
		}
	}
	fmt.Println(*arr)
}
func main() {
	var array [5]int=[...]int{24,69,80,57,13}
	bubble(&array)
}
