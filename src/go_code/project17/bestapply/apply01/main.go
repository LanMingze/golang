package main

import (
	"fmt"
	"reflect"
)
type Monster struct{
	Name string `json:"name"`
	Age int `json:"monster_age"`
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
	if kd!=reflect.Struct{//如果不是结构体就退出
		fmt.Println("expect sturct")
		return
	}
	//获取到结构体有多少字段
	num:=val.NumField()
	fmt.Printf("有%v个字段\n",num)
	for i:=0;i<num;i++{
		fmt.Printf("field %d 值为%v\n",i,val.Field(i))
		//获取struct标签
		tagVal:=typ.Field(i).Tag.Get("json")
		//没有标签就不显示
		if tagVal!=""{
			fmt.Printf("field %d 值为%v\n",i,tagVal)
		}
	}
	//获取结构体方法数量
	numOfMethod:=val.NumMethod()
	fmt.Printf("有%d方法\n",numOfMethod)
	//按照函数名首字母为顺序，ASCII为标准
	val.Method(1).Call(nil)

	//声明一个reflect.value切片
	var param[]reflect.Value
	param = append(param, reflect.ValueOf(10))
	param = append(param, reflect.ValueOf(20))
	res:=val.Method(0).Call(param)//call需要传入切片，同时返回切片
	fmt.Println("res=",res[0].Int())
}
func main() {
	var m Monster=Monster{
		Name: "hhh",
		Age: 20,
		Score:66.6,
	}
	TestStruct(m)
}
