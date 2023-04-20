package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取传入值的type，kind，值
	//先获取到reflect.Type
	rType:=reflect.TypeOf(b)
	fmt.Println("rtype=",rType)
	//获取到reflect.Value
	rVal:=reflect.ValueOf(b)
	fmt.Println("rval=",rVal,rVal.Int()+2)
	fmt.Printf("rval type=%T\n",rVal)
	//转为interface
	iV:=rVal.Interface()
	//将interface通过断言转为需要的类型
	num2:=iV.(int)
	fmt.Println("num2=",num2)
}
type Students struct{
	Name string
	Age int
}
func reflectTest02(b interface{}) {
	//通过反射获取传入值的type，kind，值
	//先获取到reflect.Type
	rType:=reflect.TypeOf(b)
	fmt.Println("rtype=",rType)
	//获取到reflect.Value
	rVal:=reflect.ValueOf(b)

	//转为interface
	iV:=rVal.Interface()
	//将interface通过断言转为需要的类型
	fmt.Printf("iv=%v iv=%T\n",iV,iV)
	stu,ok:=iV.(Students)
	if ok{
		fmt.Printf("stu name=%v\n",stu.Name)
	}
}
func main() {
	var num int = 100
	reflectTest01(num)
	stu:=Students{"tom",20}
	reflectTest02(stu)
}
