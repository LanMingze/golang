package main
import(
	"fmt"
	"go_code/project12/encapsulate/model"
)
func main(){
	p:=model.NewPerson("tom")
	fmt.Println(*p)
	p.SetAge(20)
	fmt.Println(p.GetAge())
	p.SetSalary(3555.2)
	fmt.Println(p.GetSalary())
	fmt.Println(*p)
}