package main

import "fmt"

func main() {
	var arr [4][6]int
	arr[1][2]=1
	arr[2][1]=2
	arr[2][3]=3
	fmt.Println(arr)
	for i:=0;i<4;i++{
		for j:=0;j<6;j++{
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}

	var arr3 [2][3]int=[2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr3)
}
