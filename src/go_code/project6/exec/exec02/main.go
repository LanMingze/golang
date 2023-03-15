package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64=3.0
	var b float64 =100.0
	var c float64=6.0
	var m float64 =b*b-4*a*c
	var x1 float64
	var x2 float64
	if m>0{
		x1=(-b+math.Sqrt(m))/(2*a)
		x2=(-b-math.Sqrt(m))/(2*a)
		fmt.Printf("x1=%v,x2=%v",x1,x2)
	}else if m==0{
		x1=(-b+math.Sqrt(m))/(2*a)
		fmt.Printf("x1=x2=%v",x1)
	}else{
		fmt.Println("无解")
	}

}
