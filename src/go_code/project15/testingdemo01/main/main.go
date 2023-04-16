package main
import(
	"fmt"
)
func addUpper(n int)int{
	res:=0
	for i:=1;i<=n;i++{
		res+=i
	}
	return res
}
func main(){
	res:=addUpper(10)
	if res!=55{
		fmt.Println("测试错误！res=",res)
	}else{
		fmt.Println("测试正确，res=",res)
	}

}