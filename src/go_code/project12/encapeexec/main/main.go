package main
import (
	"fmt"
	"go_code/project12/encapeexec/model"
)
func main(){
	account:=model.NewAccount("sefsfsf","888888",50.2)
	if account!=nil{
		fmt.Println(*account)
	}else{
		fmt.Println("fail")
	}
}