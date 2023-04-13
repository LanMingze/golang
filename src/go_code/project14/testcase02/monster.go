package monster

import (
	"encoding/json"
	"fmt"
	"os"
)
type Monster struct{
	Name string
	Age int
	Skill string
}
func(m *Monster)Store()bool{
	data,err:=json.Marshal(m)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	filePath:="/home/lan/a.txt"
	err=os.WriteFile(filePath,data,0666)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}
func(m *Monster)ReStore()bool{
	filePath:="/home/lan/a.txt"
	data,err:=os.ReadFile(filePath)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	err=json.Unmarshal(data,&m)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}