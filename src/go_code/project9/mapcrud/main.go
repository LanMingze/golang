package main

import (
	"fmt"
)

func main() {
	cities:=make(map[string]string)
	cities["no1"]="beijing"
	cities["no2"]="tianjin"
	cities["no3"]="shanghai"
	fmt.Println(cities)
	//有no3就更改，无就增加
	cities["no3"]="hangzhou"
	fmt.Println(cities)
	//delete,当指定key不存在时，也不会报错
	delete(cities,"no2")
	fmt.Println(cities)
	//
	val,ok:=cities["no1"]
	if ok{
		fmt.Printf("value=%v\n",val)
	}else{
		fmt.Println("no found")
	}
	//full delete
	cities=make(map[string]string)
	fmt.Println(cities)

	
}
