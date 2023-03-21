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
	
	for k,v:=range cities{
		fmt.Printf("key=%v value=%v\n",k,v)
	}	
	studentMap :=make(map[string]map[string]string)
	studentMap["stu01"]=make(map[string]string,2)
	studentMap["stu01"]["name"]="tom"
	studentMap["stu01"]["sex"]="male"

	studentMap["stu02"]=make(map[string]string,2)
	studentMap["stu02"]["name"]="mary"
	studentMap["stu02"]["sex"]="female"

	fmt.Println(studentMap)

	for k1,v1:=range studentMap{
		fmt.Println("k1=",k1)
		for k2,v2:=range v1{
			fmt.Printf("\t k2=%v,v2=%v\n",k2,v2)
		}
	}
}
