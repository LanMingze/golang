package main
import(
	"fmt"
)
func main(){
	fmt.Println(10/4)
	var n1 float32=10/4
	fmt.Println(n1)
	//如果有小数需要有浮点数运算
	var n2 float32=10.0/4
	fmt.Println(n2)

	//%演示
	//a%b=a-a/b*b
	fmt.Println("10模3=",10%3)
	fmt.Println("-10模3=",-10%3)
	fmt.Println("10模-3=",10%-3)
	fmt.Println("-10模-3=",-10%-3)
	//++和--演示
	var i int =10
	i++
	fmt.Println("i=",i)
	i--
	fmt.Println("i=",i)
}