package main

import (
	"encoding/json"
	"fmt"
)
type Monster struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string
	Sal float64
	Skill string
}
func testStruct(){
	monster:=Monster{
		Name: "niumo" ,
		Age: 500,
		Birthday: "2011-1-1",
		Sal: 200.2,
		Skill: "niumoquan",
	}
	data,err:=json.Marshal(&monster)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("monstor序列化后=%v\n",string(data))
}
func testMap(){
	var a map[string]interface{}=make(map[string]interface{})
	a["name"]="honghaier"
	a["age"]=30
	a["address"]="huoyundong"
	data,err:=json.Marshal(&a)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("a map序列化后=%v\n",string(data))
}
func testSlice(){
	var slice []map[string]interface{}
	a:=make(map[string]interface{})
	a["name"]="bob"
	a["age"]=20
	a["address"]="beijing"
	slice=append(slice, a)
	b:=make(map[string]interface{})
	b["name"]="jack"
	b["age"]=18
	b["address"]="shanghai"
	slice = append(slice, b)
	data,err:=json.Marshal(&slice)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("slice序列化后=%v\n",string(data))

}
func testFloat64(){
	var num1 float64=10.23
	data,err:=json.Marshal(&num1)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("num1序列化后=%v\n",string(data))
}
func main(){
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}