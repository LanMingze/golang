package main
import (
	"fmt"
)
type Student struct{

} 
func TypeJudge(items... interface{}){
	for i,v:=range items{
		switch v.(type){
		case bool:
			fmt.Printf("第%v个参数是bool类型,值是:%v\n",i,v)
		case float32:
			fmt.Printf("第%v个参数是float32类型,值是:%v\n",i,v)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是：%v\n",i,v)
		case int ,int32,int64:
			fmt.Printf("第%v个参数是整数类型，值是：%v\n",i,v)
		case string:
			fmt.Printf("第%v个参数是string类型，值是：%v\n",i,v)
		case Student:
			fmt.Printf("第%v个参数是student类型，值是：%v\n",i,v)
		case *Student:
			fmt.Printf("第%v个参数是student类型，值是：%v\n",i,v)
		default:
			fmt.Printf("第%v个参数是不确定类型，值是：%v\n",i,v)
		}
	}
}


func main(){
	var n1 float32=1.1
	var n2 float64=2.3
	var n3 int32=30
	var name string="tom"
	address:="北京"
	n4:=300
	stu:=Student{}
	stu1:=&Student{}
	TypeJudge(n1,n2,n3,name,address,n4,stu,*stu1)
}