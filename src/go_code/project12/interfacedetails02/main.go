package main
import (
	"fmt"
)
type A interface{
	B
	C
	test03()
}
type B interface{
	test01()
}
type C interface{
	test02()
}
type Stu struct{

}
func (s *Stu)test01(){
	fmt.Println("test01...")
}
func (s Stu)test02(){

}
func (s Stu)test03(){

}
type T interface{

}
func main(){
	var stu Stu
	var a A=&stu
	a.test01()
	var t T=stu
	fmt.Println(t)
	var t2 interface{}=stu
	fmt.Println(t2)
	var num float64=10.8
	t2=num
	fmt.Println(t2)
}