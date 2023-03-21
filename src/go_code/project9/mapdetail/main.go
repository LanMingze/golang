package main

import (
	"fmt"
	
)

func modify(map1 map[int]int){
	map1[1]=100
}

type stu struct{
	Name string
	Age int
	Add string
}
func main() {
	

	stus:=make(map[string]stu,10)
	stu1:=stu{"tom",25,"beijing"}
	stu2:=stu{"mary",15,"shanghai"}
	stus["no1"]=stu1
	stus["no2"]=stu2
	fmt.Println(stus)
	//遍历学生信息
	for k,v:=range stus{
		fmt.Printf("id=%v\n",k)
		fmt.Printf("name=%v\n",v.Name)
		fmt.Printf("age=%v\n",v.Age)
		fmt.Printf("address=%v\n",v.Add)
	}
	map1:=make(map[int]int)
	map1[1]=1
	map1[2]=50
	map1[5]=100
	fmt.Println(map1)
	modify(map1)//map是引用类型
	fmt.Println(map1)
}
