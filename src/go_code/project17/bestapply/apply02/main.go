package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)
type Monster struct{
	Name string `json:"name"`
	Age int 
	Score float32
	Sex string
}
func (m Monster)Print(){
	fmt.Println("----start----")
	fmt.Println(m)
	fmt.Println("----end----")
}
func(m Monster)GetSum(n1,n2 int)int{
	return n1+n2
}
func (m *Monster)Set(name string,age int,score float32,sex string){
	m.Name=name
	m.Age=age
	m.Score=score
	m.Sex=sex
}
func TestStruct(a interface{}){
	typ:=reflect.TypeOf(a)//获取type类型
	val:=reflect.ValueOf(a)//获取value类型
	kd:=val.Kind()//获取a的类别
	if kd!=reflect.Ptr&&val.Elem().Kind()==reflect.Struct{//如果不是结构体就退出
		fmt.Println("expect sturct")
		return
	}
	//获取到结构体有多少字段
	num:=val.Elem().NumField()
	val.Elem().Field(0).SetString("baixiang")
	fmt.Printf("有%v个字段\n",num)
	for i:=0;i<num;i++{
		fmt.Printf("field %d 值为%v\n",i,val.Elem().Field(i).Kind())
		
	}
	//获取struct标签
	tag:=typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n",tag)
	//获取结构体方法数量
	numOfMethod:=val.Elem().NumMethod()
	fmt.Printf("有%d方法\n",numOfMethod)
	//按照函数名首字母为顺序，ASCII为标准
	val.Method(1).Call(nil)
}
func main() {
	var m Monster=Monster{
		Name: "hhh",
		Age: 20,
		Score:66.6,
	}
	result,_:=json.Marshal(m)
	fmt.Println(string(result))
	TestStruct(&m)
	fmt.Println(m)
}
