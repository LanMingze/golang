package main

import "fmt"

func test01(arr [3]int){//[3]int [4]int不是一种数据类型
	arr[0]=88
	fmt.Println(arr)
}
func test02(arr *[3]int){//[3]int [4]int不是一种数据类型
	(*arr)[0]=88
	fmt.Println(*arr)
}

func main() {
	var arr01 [3]int
	 arr01[0]=1
	arr01[1]=30
	// arr01[2]=1.1
	 arr01[2]=3
	// arr01[3]=8
	test01(arr01)
	fmt.Println(arr01)
	test02(&arr01)
	fmt.Println(arr01)
}
