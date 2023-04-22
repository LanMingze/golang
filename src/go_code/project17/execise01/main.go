package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "100"
	fs:=reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Println(str)
}
