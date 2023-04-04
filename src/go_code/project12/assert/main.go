package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	fmt.Println(a)
	//类型断言
	var b Point = a.(Point)
	//error
	//b=a
	fmt.Println(b)

	// var x interface{}
	// var C float32=1.1
	// x=C
	// y:=x.(float32)
	// fmt.Printf("%T %v\n",y,y)

	//带检测的断言
	var x interface{}
	var C float32 = 1.1
	x = C

	if y, ok := x.(float64); ok {
		fmt.Printf("%T %v\n", y, y)
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("继续执行")

}
