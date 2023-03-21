package main

import (
	"fmt"
)

func main() {
	//method 1
	//key不能重复，value可以重复
	var a map[string]string
	//使用map前，需要make给map分配空间
	a=make(map[string]string,10)
	a["no1"]="songjiang"
	a["no3"]="wuyong"
	a["no2"]="linchong"
	a["no1"]="wusong"
	fmt.Println(a)

	//method 2
	cities:=make(map[string]string)
	cities["no1"]="beijing"
	cities["no2"]="tianjin"
	cities["no3"]="shanghai"
	fmt.Println(cities)

	//method 3
	heros:=map[string]string{
		"hero1":"songjiang",
		"hero2":"lujunyi",
	}
	fmt.Println(heros)
}
