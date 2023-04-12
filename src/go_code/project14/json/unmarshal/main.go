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
func unmarshalStruct(){
	str:="{\"name\":\"niumo\",\"age\":500,\"Birthday\":\"2011-1-1\",\"Sal\":200.2,\"Skill\":\"niumoquan\"}"
	var monster Monster
	err:=json.Unmarshal([]byte(str),&monster)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%v\n",monster)
}
func unmarshalMap(){
	str:="{\"address\":\"huoyundong\",\"age\":30,\"name\":\"honghaier\"}"
	var a map[string]interface{}
	err:=json.Unmarshal([]byte(str),&a)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%v\n",a)
}
func unmarshalSlice(){
	str:="[{\"address\":\"beijing\",\"age\":20,\"name\":\"bob\"},"+
	"{\"address\":\"shanghai\",\"age\":18,\"name\":\"jack\"}]"
	var slice []map[string]interface{}
	err:=json.Unmarshal([]byte(str),&slice)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%v\n",slice)
}
func main(){
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}