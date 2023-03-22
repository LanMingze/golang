package main

import (
	"fmt"
)
func modify(users map[string]map[string]string,name string){
	if users[name]!=nil{
		users[name]["psw"]="88888888"
	}else{
		users[name]=make(map[string]string,2)
		users[name]["psw"]="88888888"
		users[name]["nickname"]="nick"
	}
}
func main() {
	users:=make(map[string]map[string]string,10)
	users["smith"]=make(map[string]string,2)
	users["smith"]["psw"]="aaa"
	users["smith"]["nickname"]="smi"
	modify(users,"smith")
	modify(users,"tom")
	modify(users,"mary")
	fmt.Println(users)
}
