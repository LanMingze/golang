package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}
type B struct {
	Num int
}
type Monster struct{
	Name string `json:"name"`//tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)

	monster:=Monster{"Niumi",500,"niuquan"}
	//将monster序列化为json格式字符串
	jsonStr,err:=json.Marshal(monster)//使用到反射
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
}
