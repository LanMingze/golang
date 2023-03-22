package main

import "fmt"

type Point struct {
	X int
	Y int
}
type Rect struct {
	leftUp, rightDown Point
}
type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r1.leftUp.x地址是：%p,r1.leftUp.y地址是：%p,r1.rightDown.x地址是：%p,r1.rightDown.y地址是：%p\n", &r1.leftUp.X,
		&r1.leftUp.Y, &r1.rightDown.X, &r1.rightDown.Y)

	fmt.Printf("r2.leftUp地址是：%p,r2.rightDown地址是：%p\n", 
	&r2.leftUp, &r2.rightDown)
}
