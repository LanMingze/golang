package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//获取到reflect.Value
	rVal:=reflect.ValueOf(b)
	//rVal.Elem()可以理解为*b
	rVal.Elem().SetInt(20)
	
}

func main() {
	var num int = 100
	reflectTest01(&num)
	fmt.Println(num)
}
