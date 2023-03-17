package main

import "fmt"

func main() {
	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	//slice是切片名
	//intArr[1:3]表示slice引用到intArr这个数组的第2（下标为1）
	//个元素到第3（下标为3-1）个元素，不包含下标为3的元素
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice的元素是：", slice)
	fmt.Println("slice的个数是：", len(slice))
	fmt.Println("slice的容量是：", cap(slice))
	
}
