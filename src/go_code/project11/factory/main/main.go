package main

import "fmt"
import "go_code/project11/factory/model"



func main() {
	var stu1=model.NewStudent("tooom",100.2)
	fmt.Println(*stu1)
	fmt.Println(stu1.GetName(),stu1.GetScore())
}
